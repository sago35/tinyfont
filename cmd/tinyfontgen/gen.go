package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"unsafe"

	"github.com/hajimehoshi/go-jisx0208"
	"github.com/sago35/go-bdf"
	"tinygo.org/x/tinyfont"
)

type fontgen struct {
	pkgname  string
	fontname string
	fonts    []string
}

type option func(*options)

type options struct {
	ascii     bool
	all       bool
	verbose   bool
	yadvance  int
	withEmoji bool
}

var defaultOption = options{
	ascii:     true,
	all:       false,
	verbose:   false,
	yadvance:  0,
	withEmoji: true,
}

func withAscii(b bool) option {
	return func(o *options) {
		o.ascii = b
	}
}

func withAll(b bool) option {
	return func(o *options) {
		o.all = b
	}
}

func withVerbose(b bool) option {
	return func(o *options) {
		o.verbose = b
	}
}

func withYAdvance(a int) option {
	return func(o *options) {
		o.yadvance = a
	}
}

func withEmoji(b bool) option {
	return func(o *options) {
		o.withEmoji = b
	}
}

const bom = 0xFEFF // byte order mark, only permitted as very first character
var re = regexp.MustCompile(`\b(Rune:)(\d+)`)

func (f *fontgen) generate(w io.Writer, runes []rune, opt ...option) error {
	opts := defaultOption
	for _, o := range opt {
		o(&opts)
	}

	if opts.ascii {
		runes = append(runes, allAscii()...)
	}
	runes = sortAndUniq(runes)

	fonts := []*bdf.Font{}
	for _, font := range f.fonts {
		fn, err := readFont(font)
		if err != nil {
			return err
		}
		fonts = append(fonts, fn)
	}

	ufont := tinyfont.Font{}
	exists := map[rune]bool{}
	for _, font := range fonts {
		ufont.YAdvance = uint8(float64(font.Size) * float64(font.DPI[1]) / 75)
		if opts.yadvance > 0 {
			ufont.YAdvance = uint8(opts.yadvance)
		}

		code2rune := func(c int) (rune, error) { return rune(c), nil }
		switch strings.ToLower(font.CharsetRegistry) {
		case "iso08859":
		case "jisx0208.1990":
			code2rune = jisx0208.Rune
		}

		for _, f := range font.Characters {
			fr, err := code2rune(int(f.Encoding))
			if err != nil {
				// skip
				continue
			}

			if fr == bom {
				continue
			} else if opts.all {
				if exists[fr] {
					continue
				}
				exists[fr] = true

				g, err := bdf2glyph(f, fr)
				if err != nil {
					continue
				}
				ufont.Glyphs = append(ufont.Glyphs, g)
			} else {
				for _, r := range runes {
					if r == fr {
						if exists[fr] {
							continue
						}
						exists[fr] = true

						g, err := bdf2glyph(f, fr)
						if err != nil {
							continue
						}
						ufont.Glyphs = append(ufont.Glyphs, g)
					}
				}
			}
		}
	}

	if opts.withEmoji {
		g, err := makeEmoji(`.\cmd\emoji\png`, exists)
		if err != nil {
			return err
		}
		ufont.Glyphs = append(ufont.Glyphs, g...)
	}

	ufont.Glyphs = sortGlyphs(ufont.Glyphs)

	tmp, err := ioutil.TempFile(``, `tinyfontgen`)
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name())

	fmt.Fprintln(tmp, `//`, filepath.Base(os.Args[0]), strings.Join(os.Args[1:], ` `))
	fmt.Fprintln(tmp)
	fmt.Fprintln(tmp, `package `+f.pkgname)
	fmt.Fprintln(tmp)
	fmt.Fprintln(tmp, `import (`)
	fmt.Fprintln(tmp, `	"tinygo.org/x/tinyfont"`)
	fmt.Fprintln(tmp, `)`)
	fmt.Fprintln(tmp)

	fontname := strings.ToUpper(f.fontname[0:1]) + f.fontname[1:]

	fmt.Fprintf(tmp, "\n")
	fmt.Fprintf(tmp, "var %s = s%s{}\n", fontname, fontname)
	fmt.Fprintf(tmp, "\n")
	fmt.Fprintf(tmp, "type s%s struct {\n", fontname)
	fmt.Fprintf(tmp, "}\n")
	fmt.Fprintf(tmp, "\n")

	fmt.Fprintf(tmp, `const c%s = "" +`, fontname)
	fmt.Fprintln(tmp)

	idxMap := map[rune]int{}
	current := 0
	for i, g := range ufont.Glyphs {
		idxMap[g.Rune] = current
		c := fmt.Sprintf("%c", ufont.Glyphs[i].Rune)
		if ufont.Glyphs[i].Rune == 0 {
			c = ""
		}
		length := 4 + 1 + 1 + 1 + 1 + 1 + 1 + len(g.Bitmaps)
		fmt.Fprintf(tmp, `	/* %s */ `, c)
		fmt.Fprintf(tmp, `"\x%02X\x%02X" +`, (uint16(length) >> 8), uint8(length))
		fmt.Fprintf(tmp, `"\x%02X\x%02X\x%02X\x%02X" +`, ((uint32(g.Rune) & 0xFF000000) >> 24), ((uint32(g.Rune) & 0x00FF0000) >> 16), ((uint32(g.Rune) & 0x0000FF00) >> 8), (uint32(g.Rune) & 0x000000FF))
		fmt.Fprintf(tmp, `"\x%02X\x%02X" +`, uint8(g.Width), uint8(g.Height))
		fmt.Fprintf(tmp, `"\x%02X\x%02X" +`, uint8(g.XAdvance), uint8(g.XOffset))
		fmt.Fprintf(tmp, `"\x%02X" +`, uint8(g.YOffset))
		fmt.Fprintf(tmp, `"\x%02X" +`, uint8(g.Mode))
		fmt.Fprintf(tmp, `"`)
		for _, b := range g.Bitmaps {
			fmt.Fprintf(tmp, `\x%02X`, uint8(b))
		}
		fmt.Fprintf(tmp, `" +`)
		fmt.Fprintln(tmp)
		current += 2 + length
	}
	fmt.Fprintln(tmp, `""`)
	fmt.Fprintln(tmp)

	fmt.Fprintf(tmp, "func (f *s%s) GetGlyph(r rune) tinyfont.Glyph {\n", fontname)
	fmt.Fprintf(tmp, "	idx := -1\n")
	fmt.Fprintf(tmp, "\n")

	// switch
	fmt.Fprintf(tmp, "switch r {\n")
	keys := []int{}
	for k := range idxMap {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	for _, k := range keys {
		v := idxMap[rune(k)]
		fmt.Fprintf(tmp, "case 0x%04X: idx = %d\n", k, v)
	}
	fmt.Fprintf(tmp, "}\n")
	fmt.Fprintf(tmp, "\n")

	fmt.Fprintf(tmp, "if idx == -1 {\n")
	fmt.Fprintf(tmp, "	return tinyfont.Glyph{\n")
	fmt.Fprintf(tmp, "		Rune:     r,\n")
	fmt.Fprintf(tmp, "		Width:    0,\n")
	fmt.Fprintf(tmp, "		Height:   0,\n")
	fmt.Fprintf(tmp, "		XAdvance: uint8(c%s[6]),\n", fontname)
	fmt.Fprintf(tmp, "		XOffset:  int8(c%s[7]),\n", fontname)
	fmt.Fprintf(tmp, "		YOffset:  int8(c%s[8]),\n", fontname)
	fmt.Fprintf(tmp, "		Bitmaps:  []uint8{},\n")
	fmt.Fprintf(tmp, "	}\n")
	fmt.Fprintf(tmp, "}\n")
	fmt.Fprintf(tmp, "\n")

	fmt.Fprintf(tmp, "	length := int((uint16(c%s[idx+0]) << 8) + uint16(c%s[idx+1]))\n", fontname, fontname)
	fmt.Fprintf(tmp, "	idx += 2\n")
	fmt.Fprintf(tmp, "	ret := tinyfont.Glyph{\n")
	fmt.Fprintf(tmp, "		Rune:     rune((uint32(c%s[idx+0]) << 24) + (uint32(c%s[idx+1]) << 16) + (uint32(c%s[idx+2]) << 8) + uint32(c%s[idx+3])),\n", fontname, fontname, fontname, fontname)
	fmt.Fprintf(tmp, "		Width:    uint8(c%s[idx+4]),\n", fontname)
	fmt.Fprintf(tmp, "		Height:   uint8(c%s[idx+5]),\n", fontname)
	fmt.Fprintf(tmp, "		XAdvance: uint8(c%s[idx+6]),\n", fontname)
	fmt.Fprintf(tmp, "		XOffset:  int8(c%s[idx+7]),\n", fontname)
	fmt.Fprintf(tmp, "		YOffset:  int8(c%s[idx+8]),\n", fontname)
	fmt.Fprintf(tmp, "		Mode   :  tinyfont.Mode(c%s[idx+9]),\n", fontname)
	fmt.Fprintf(tmp, "		Bitmaps:  []uint8(c%s[idx+10 : idx+length]),\n", fontname)
	fmt.Fprintf(tmp, "	}\n")
	fmt.Fprintf(tmp, "\n")
	fmt.Fprintf(tmp, "	return ret\n")
	fmt.Fprintf(tmp, "}\n")
	fmt.Fprintf(tmp, "\n")
	fmt.Fprintf(tmp, "func (f *s%s) GetYAdvance() uint8 {\n", fontname)
	fmt.Fprintf(tmp, "	return 0x14\n")
	fmt.Fprintf(tmp, "}\n")

	//fmt.Fprintf(tmp, "	YAdvance:%#v,\n", ufont.YAdvance)
	//fmt.Fprintf(tmp, "}\n")

	if opts.verbose {
		fmt.Printf("Approx. %d bytes\n", calcSize(ufont))
	}

	// gofmt
	buf := bytes.Buffer{}
	cmd := exec.Command(`gofmt`, tmp.Name())
	cmd.Stdout = w
	cmd.Stderr = &buf
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("%s : %s", err.Error(), strings.TrimSpace(buf.String()))
	}

	return nil
}

func readFont(p string) (*bdf.Font, error) {
	buf, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}

	font, err := bdf.Parse(buf)
	if err != nil {
		return nil, err
	}

	return font, nil

}

func calcSize(f tinyfont.Font) int {
	sz := 0
	for _, g := range f.Glyphs {
		sz += int(unsafe.Sizeof(g.Rune))
		sz += int(unsafe.Sizeof(g.Width))
		sz += int(unsafe.Sizeof(g.Height))
		sz += int(unsafe.Sizeof(g.XAdvance))
		sz += int(unsafe.Sizeof(g.XOffset))
		sz += int(unsafe.Sizeof(g.YOffset))
		sz += len(*(*[]byte)(unsafe.Pointer(&g.Bitmaps)))
	}
	sz += int(unsafe.Sizeof(f.YAdvance))
	return sz

}

func bdf2glyph(f bdf.Character, r rune) (tinyfont.Glyph, error) {
	img := f.Alpha

	bmp := []byte{}
	bitpos := 0
	b := byte(0)
	for y := img.Rect.Min.Y; y < img.Rect.Max.Y; y++ {
		for x := img.Rect.Min.X; x < img.Rect.Max.X; x++ {
			r, _, _, _ := img.At(x, y).RGBA()
			if r != 0 {
				b += (byte(1) << (8 - byte(bitpos) - 1))
			}

			bitpos++
			if bitpos == 8 {
				bitpos = 0
				bmp = append(bmp, b)
				b = 0
			}
		}
	}
	if bitpos != 0 {
		bitpos = 0
		bmp = append(bmp, b)
	}

	g := tinyfont.Glyph{
		Rune:     r,
		Width:    uint8(img.Rect.Max.X),
		Height:   uint8(img.Rect.Max.Y),
		XAdvance: uint8(f.Advance[0]),
		XOffset:  int8(f.LowerPoint[0]),
		YOffset:  -1 * int8(f.Alpha.Rect.Max.Y+f.LowerPoint[1]),
		Bitmaps:  bmp,
	}

	return g, nil
}

func allAscii() []rune {
	ret := []rune{}
	for i := 0; i <= 255; i++ {
		ret = append(ret, rune(i))
	}
	return ret
}

func sortAndUniq(runes []rune) []rune {
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })

	ret := []rune{}
	var prev rune
	for _, r := range runes {
		if prev != r {
			ret = append(ret, r)
		}
		prev = r
	}
	return ret
}

func sortGlyphs(glyphs []tinyfont.Glyph) []tinyfont.Glyph {
	sort.Slice(glyphs, func(i, j int) bool { return glyphs[i].Rune < glyphs[j].Rune })
	return glyphs
}
