package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"tinygo.org/x/tinyfont"
)

func png2glyph(r rune) (tinyfont.Glyph, error) {
	g := tinyfont.Glyph{
		Rune: r,
	}
	return g, nil
}

func makeEmoji(dir string, exists map[rune]bool) ([]tinyfont.Glyph, error) {
	ret := []tinyfont.Glyph{}

	matches, err := filepath.Glob(filepath.Join(dir, `*`))

	if err != nil {
		return nil, err
	}

	emojis := []string{}
	for _, match := range matches {
		if !strings.Contains(match, `-`) {
			emojis = append(emojis, match)
		}
	}

	for _, match := range emojis {
		//fmt.Println(match)
		code := strings.TrimRight(filepath.Base(match), filepath.Ext(match))
		codeh, err := strconv.ParseUint(code, 16, 0)
		if err != nil {
			return nil, err
		}
		r := rune(codeh)

		if exists[r] {
			continue
		}
		exists[r] = true

		g, err := getEmojiGlyph(r, match)
		if err != nil {
			return nil, err
		}
		ret = append(ret, g)
	}
	return ret, nil
}

func getEmojiData(png string, addr int) (string, string, error) {
	data := &bytes.Buffer{}
	sc := &bytes.Buffer{}

	r, err := os.Open(png)
	if err != nil {
		return "", "", err
	}
	defer r.Close()

	img, _, err := image.Decode(r)
	if err != nil {
		return "", "", err
	}

	ru := filepath.Base(png)
	ru = strings.TrimRight(ru, filepath.Ext(ru))
	rr := uint32(0)
	_, err = fmt.Sscan("0x"+ru, &rr)
	if err != nil {
		return "", "", err
	}

	// data
	length := 4 + 2 + 2 + 1 + 12*12*2
	fmt.Fprintf(data, `	/* %s */ `, string(rr))
	fmt.Fprintf(data, `"\x%02X\x%02X" + `, uint8(length>>8), uint8(length))
	fmt.Fprintf(data, `"\x%02X\x%02X\x%02X\x%02X" + `, uint8((rr&0xFF000000)>>24), uint8((rr&0x00FF0000)>>16), uint8((rr&0x0000FF00)>>8), uint8((rr & 0x000000FF)))
	fmt.Fprintf(data, `"\x0C\x0C" + "\x0C\x00" + "\xF5" + `)
	//fmt.Fprintf(data, "\n")

	fmt.Fprintf(data, `"`)
	for y := 0; y < img.Bounds().Dy(); y++ {

		for x := 0; x < img.Bounds().Dx(); x++ {
			d := RGBATo55a5(img.At(x, y))
			fmt.Fprintf(data, `\x%02X\x%02X`, (d&0xFF00)>>8, d&0xFF)
		}
	}
	fmt.Fprintf(data, `" +`)
	fmt.Fprintf(data, "\n")

	// switch case
	fmt.Fprintf(sc, "	case 0x%X:\n", rr)
	fmt.Fprintf(sc, "		idx = %d\n", addr)
	fmt.Fprintf(sc, "		emoji = true\n")

	//case 0x1F430:
	//	idx = 216909 + 2 + 0x1D + 2 + 0x129
	//	emoji = true

	return data.String(), sc.String(), nil
}

func getEmojiGlyph(rr rune, png string) (tinyfont.Glyph, error) {
	ret := tinyfont.Glyph{
		Rune:     rr,
		Width:    0,
		Height:   0,
		XAdvance: 0,
		XOffset:  0,
		YOffset:  0,
		Bitmaps:  []byte{},
		Emoji:    false,
		Mode:     tinyfont.Rgb55a5,
	}

	r, err := os.Open(png)
	if err != nil {
		return ret, err
	}
	defer r.Close()

	img, _, err := image.Decode(r)
	if err != nil {
		return ret, err
	}

	// data
	ret.Width = 0x0C    // TODO
	ret.Height = 0x0C   // TODO
	ret.XAdvance = 0x0C // TODO
	ret.XOffset = 0
	//ret.YOffset = -11 // TODO: -? // for mplus
	ret.YOffset = -15 // TODO: -? // for noto sans 48 to 12pt

	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			d := RGBATo55a5(img.At(x, y))
			ret.Bitmaps = append(ret.Bitmaps, uint8((d&0xFF00)>>8), uint8(d&0xFF))
		}
	}
	return ret, nil
}

// RGBATo565 converts a color.RGBA to uint16 used in the display
func RGBATo565(c color.Color) uint16 {
	r, g, b, _ := c.RGBA()
	return uint16((r & 0xF800) +
		((g & 0xFC00) >> 5) +
		((b & 0xF800) >> 11))
}

// RGBATo55a5 converts a color.RGBA to uint16 used in the display
func RGBATo55a5(c color.Color) uint16 {
	r, g, b, a := c.RGBA()
	if a > 0 {
		a = 1
	}
	return uint16((r & 0xF800) +
		((g & 0xF800) >> 5) +
		(a << 5) +
		((b & 0xF800) >> 11))
}
