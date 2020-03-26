package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"unsafe"

	ntinyfont "github.com/sago35/tinyfont"
	"github.com/tinygo-org/tinyfont/freemono"
	"github.com/tinygo-org/tinyfont/freeserif"
	"github.com/tinygo-org/tinyfont/gophers"
	"github.com/tinygo-org/tinyfont/proggy"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

func main() {
	targets := []struct {
		pkgname, path, fontname string
		font                    tinyfont.Font
	}{
		{`tinyfont`, `./tinyfont/org_01.go`, `Org01`, tinyfont.Org01},
		{`tinyfont`, `./tinyfont/picopixel.go`, `Picopixel`, tinyfont.Picopixel},
		{`tinyfont`, `./tinyfont/tiny3x3a2pt7b.go`, `Tiny3x3a2pt7b`, tinyfont.Tiny3x3a2pt7b},
		{`tinyfont`, `./tinyfont/tomthumb.go`, `TomThumb`, tinyfont.TomThumb},
		{`freemono`, `./freemono/freemono12pt7b.go`, `Regular12pt7b`, freemono.Regular12pt7b},
		{`freemono`, `./freemono/freemono18pt7b.go`, `Regular18pt7b`, freemono.Regular18pt7b},
		{`freemono`, `./freemono/freemono24pt7b.go`, `Regular24pt7b`, freemono.Regular24pt7b},
		{`freemono`, `./freemono/freemono9pt7b.go`, `Regular9pt7b`, freemono.Regular9pt7b},
		{`freemono`, `./freemono/freemonobold12pt7b.go`, `Bold12pt7b`, freemono.Bold12pt7b},
		{`freemono`, `./freemono/freemonobold18pt7b.go`, `Bold18pt7b`, freemono.Bold18pt7b},
		{`freemono`, `./freemono/freemonobold24pt7b.go`, `Bold24pt7b`, freemono.Bold24pt7b},
		{`freemono`, `./freemono/freemonobold9pt7b.go`, `Bold9pt7b`, freemono.Bold9pt7b},
		{`freemono`, `./freemono/freemonoboldoblique12pt7b.go`, `BoldOblique12pt7b`, freemono.BoldOblique12pt7b},
		{`freemono`, `./freemono/freemonoboldoblique18pt7b.go`, `BoldOblique18pt7b`, freemono.BoldOblique18pt7b},
		{`freemono`, `./freemono/freemonoboldoblique24pt7b.go`, `BoldOblique24pt7b`, freemono.BoldOblique24pt7b},
		{`freemono`, `./freemono/freemonoboldoblique9pt7b.go`, `BoldOblique9pt7b`, freemono.BoldOblique9pt7b},
		{`freemono`, `./freemono/freemonooblique12pt7b.go`, `Oblique12pt7b`, freemono.Oblique12pt7b},
		{`freemono`, `./freemono/freemonooblique18pt7b.go`, `Oblique18pt7b`, freemono.Oblique18pt7b},
		{`freemono`, `./freemono/freemonooblique24pt7b.go`, `Oblique24pt7b`, freemono.Oblique24pt7b},
		{`freemono`, `./freemono/freemonooblique9pt7b.go`, `Oblique9pt7b`, freemono.Oblique9pt7b},
		{`freesans`, `./freesans/freesans12pt7b.go`, `Regular12pt7b`, freesans.Regular12pt7b},
		{`freesans`, `./freesans/freesans18pt7b.go`, `Regular18pt7b`, freesans.Regular18pt7b},
		{`freesans`, `./freesans/freesans24pt7b.go`, `Regular24pt7b`, freesans.Regular24pt7b},
		{`freesans`, `./freesans/freesans9pt7b.go`, `Regular9pt7b`, freesans.Regular9pt7b},
		{`freesans`, `./freesans/freesansbold12pt7b.go`, `Bold12pt7b`, freesans.Bold12pt7b},
		{`freesans`, `./freesans/freesansbold18pt7b.go`, `Bold18pt7b`, freesans.Bold18pt7b},
		{`freesans`, `./freesans/freesansbold24pt7b.go`, `Bold24pt7b`, freesans.Bold24pt7b},
		{`freesans`, `./freesans/freesansbold9pt7b.go`, `Bold9pt7b`, freesans.Bold9pt7b},
		{`freesans`, `./freesans/freesansboldoblique12pt7b.go`, `BoldOblique12pt7b`, freesans.BoldOblique12pt7b},
		{`freesans`, `./freesans/freesansboldoblique18pt7b.go`, `BoldOblique18pt7b`, freesans.BoldOblique18pt7b},
		{`freesans`, `./freesans/freesansboldoblique24pt7b.go`, `BoldOblique24pt7b`, freesans.BoldOblique24pt7b},
		{`freesans`, `./freesans/freesansboldoblique9pt7b.go`, `BoldOblique9pt7b`, freesans.BoldOblique9pt7b},
		{`freesans`, `./freesans/freesansoblique12pt7b.go`, `Oblique12pt7b`, freesans.Oblique12pt7b},
		{`freesans`, `./freesans/freesansoblique18pt7b.go`, `Oblique18pt7b`, freesans.Oblique18pt7b},
		{`freesans`, `./freesans/freesansoblique24pt7b.go`, `Oblique24pt7b`, freesans.Oblique24pt7b},
		{`freesans`, `./freesans/freesansoblique9pt7b.go`, `Oblique9pt7b`, freesans.Oblique9pt7b},
		{`freeserif`, `./freeserif/freeserif12pt7b.go`, `Regular12pt7b`, freeserif.Regular12pt7b},
		{`freeserif`, `./freeserif/freeserif18pt7b.go`, `Regular18pt7b`, freeserif.Regular18pt7b},
		{`freeserif`, `./freeserif/freeserif24pt7b.go`, `Regular24pt7b`, freeserif.Regular24pt7b},
		{`freeserif`, `./freeserif/freeserif9pt7b.go`, `Regular9pt7b`, freeserif.Regular9pt7b},
		{`freeserif`, `./freeserif/freeserifbold12pt7b.go`, `Bold12pt7b`, freeserif.Bold12pt7b},
		{`freeserif`, `./freeserif/freeserifbold18pt7b.go`, `Bold18pt7b`, freeserif.Bold18pt7b},
		{`freeserif`, `./freeserif/freeserifbold24pt7b.go`, `Bold24pt7b`, freeserif.Bold24pt7b},
		{`freeserif`, `./freeserif/freeserifbold9pt7b.go`, `Bold9pt7b`, freeserif.Bold9pt7b},
		{`freeserif`, `./freeserif/freeserifbolditalic12pt7b.go`, `BoldItalic12pt7b`, freeserif.BoldItalic12pt7b},
		{`freeserif`, `./freeserif/freeserifbolditalic18pt7b.go`, `BoldItalic18pt7b`, freeserif.BoldItalic18pt7b},
		{`freeserif`, `./freeserif/freeserifbolditalic24pt7b.go`, `BoldItalic24pt7b`, freeserif.BoldItalic24pt7b},
		{`freeserif`, `./freeserif/freeserifbolditalic9pt7b.go`, `BoldItalic9pt7b`, freeserif.BoldItalic9pt7b},
		{`freeserif`, `./freeserif/freeserifitalic12pt7b.go`, `Italic12pt7b`, freeserif.Italic12pt7b},
		{`freeserif`, `./freeserif/freeserifitalic18pt7b.go`, `Italic18pt7b`, freeserif.Italic18pt7b},
		{`freeserif`, `./freeserif/freeserifitalic24pt7b.go`, `Italic24pt7b`, freeserif.Italic24pt7b},
		{`freeserif`, `./freeserif/freeserifitalic9pt7b.go`, `Italic9pt7b`, freeserif.Italic9pt7b},
		{`gophers`, `./gophers/gophers121pt.go`, `Regular121pt`, gophers.Regular121pt},
		{`gophers`, `./gophers/gophers14pt.go`, `Regular14pt`, gophers.Regular14pt},
		{`gophers`, `./gophers/gophers18pt.go`, `Regular18pt`, gophers.Regular18pt},
		{`gophers`, `./gophers/gophers22pt.go`, `Regular22pt`, gophers.Regular22pt},
		{`gophers`, `./gophers/gophers32pt.go`, `Regular32pt`, gophers.Regular32pt},
		{`gophers`, `./gophers/gophers58pt.go`, `Regular58pt`, gophers.Regular58pt},
		{`proggy`, `./proggy/tinysz8pt7b.go`, `TinySZ8pt7b`, proggy.TinySZ8pt7b},
	}

	fmt.Printf("| fontname | org | new | diff | %% |\n")
	fmt.Printf("| -------- | --- | --- | ---- | -- |\n")
	for _, tgt := range targets {
		nfont, err := convertFont(tgt.pkgname, tgt.fontname, tgt.font)
		if err != nil {
			log.Fatal(err)
		}

		orgLen := len(*(*[]byte)(unsafe.Pointer(&tgt.font.Bitmaps)))
		for _, x := range tgt.font.Glyphs {
			orgLen += int(unsafe.Sizeof(x.BitmapIndex))
			orgLen += int(unsafe.Sizeof(x.Width))
			orgLen += int(unsafe.Sizeof(x.Height))
			orgLen += int(unsafe.Sizeof(x.XAdvance))
			orgLen += int(unsafe.Sizeof(x.XOffset))
			orgLen += int(unsafe.Sizeof(x.YOffset))
		}
		orgLen += int(unsafe.Sizeof(tgt.font.First))
		orgLen += int(unsafe.Sizeof(tgt.font.Last))
		orgLen += int(unsafe.Sizeof(tgt.font.YAdvance))

		newLen := 0
		for _, x := range nfont.Glyphs {
			newLen += int(unsafe.Sizeof(x.Rune))
			newLen += int(unsafe.Sizeof(x.Width))
			newLen += int(unsafe.Sizeof(x.Height))
			newLen += int(unsafe.Sizeof(x.XAdvance))
			newLen += int(unsafe.Sizeof(x.XOffset))
			newLen += int(unsafe.Sizeof(x.YOffset))
			newLen += len(*(*[]byte)(unsafe.Pointer(&x.Bitmaps)))
		}
		newLen += int(unsafe.Sizeof(nfont.YAdvance))
		fmt.Printf("| %s.%s | %d byte | %d byte | %+d byte | %.2f %% |\n", tgt.pkgname, tgt.fontname, orgLen, newLen, newLen-orgLen, float64(newLen)/float64(orgLen)*100)

		os.Mkdir(filepath.Dir(tgt.path), 0666)

		w, err := os.Create(tgt.path)
		if err != nil {
			log.Fatal(err)
		}
		write(w, tgt.pkgname, tgt.fontname, nfont)
		err = w.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	//fmt.Printf("%d x 255 = %d\n", int(unsafe.Sizeof(byte('A'))), int(unsafe.Sizeof(byte('A')))*255)
	//fmt.Printf("%d x 255 = %d\n", int(unsafe.Sizeof(rune('A'))), int(unsafe.Sizeof(rune('A')))*255)

}

func convertFont(packagename, fontname string, orgfont tinyfont.Font) (ntinyfont.Font, error) {
	ret := ntinyfont.Font{}
	ret.YAdvance = orgfont.YAdvance

	for i := range orgfont.Glyphs {
		og := orgfont.Glyphs[i]

		length := (uint16(og.Width)*uint16(og.Height) + 7) / 8

		g := ntinyfont.Glyph{
			Rune:     rune(int(orgfont.First) + i),
			Width:    og.Width,
			Height:   og.Height,
			XAdvance: og.XAdvance,
			XOffset:  og.XOffset,
			YOffset:  og.YOffset,
			Bitmaps:  orgfont.Bitmaps[og.BitmapIndex : og.BitmapIndex+length],
		}

		ret.Glyphs = append(ret.Glyphs, g)
	}
	return ret, nil
}

func write(w io.Writer, pkgname, fontname string, ufont ntinyfont.Font) {
	fmt.Fprintln(w, `package `+pkgname)
	fmt.Fprintln(w, ``)
	fmt.Fprintln(w, `import (`)
	fmt.Fprintln(w, `	"tinygo.org/x/tinyfont"`)
	fmt.Fprintln(w, `)`)
	fmt.Fprintln(w)

	fmt.Fprintf(w, "var %s = %T{\n", fontname, ufont)
	fmt.Fprintf(w, "	Glyphs:%T{\n", ufont.Glyphs)
	for i, g := range ufont.Glyphs {
		c := fmt.Sprintf("%c", ufont.Glyphs[i].Rune)
		if ufont.Glyphs[i].Rune == 0 {
			c = ""
		}
		fmt.Fprintf(w, "		/* %s */ %#v,\n", c, g)
	}
	fmt.Fprintf(w, "	},\n")
	fmt.Fprintln(w)

	fmt.Fprintf(w, "	YAdvance:%#v,\n", ufont.YAdvance)
	fmt.Fprintf(w, "}\n")
}
