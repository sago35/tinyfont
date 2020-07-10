package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/sago35/go-bdf"
	"tinygo.org/x/tinyfont"
)

func bdf2glyphOs(f bdf.Character, r rune, font *bdf.Font) (tinyfont.Glyph, error) {
	// 4x4 oversampling
	// DWIDTH の 30 が進むべき幅 → 4 の倍数に切り上げる
	// FONTBOUNDINGBOX の 40 - (-8) = 48 が高さ
	img := f.Alpha

	fmt.Printf("-- %c --\n", r)
	//fmt.Printf("%#v\n", f)
	//fmt.Printf("%#v\n", img)

	w := (f.Advance[0] + 7) / 8 * 8
	h := (font.FontBoundingBox[1] + 7) / 8 * 8
	//yofs := font.FontBoundingBox[1]
	yofs := 48

	bmp := image.NewAlpha(image.Rectangle{Max: image.Point{X: w, Y: h}})

	fmt.Printf("-- -- maxy %d, low[1] %d, yofs %d\n", img.Rect.Max.Y, f.LowerPoint[1], yofs)
	for y := img.Rect.Min.Y; y < img.Rect.Max.Y; y++ {
		for x := img.Rect.Min.X; x < img.Rect.Max.X; x++ {
			r, _, _, _ := img.At(x, y).RGBA()
			if r != 0 {
				bmp.Set(x+f.LowerPoint[0], y-img.Rect.Max.Y-f.LowerPoint[1]+yofs, color.RGBA{1, 0, 0, 255})
			}
		}
	}
	//fmt.Printf("%#v\n", bmp)
	for x := 0; x < w; x++ {
		fmt.Printf("--")
	}
	fmt.Println()
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, _, _, _ := bmp.At(x, y).RGBA()
			if r != 0 {
				fmt.Printf("■")
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	for x := 0; x < w; x++ {
		fmt.Printf("--")
	}
	fmt.Println()

	b := byte(0)
	ret := make([]byte, w/4*h/4/2)
	fmt.Printf("len %d, w = %d, h = %d\n", len(ret), bmp.Rect.Max.X, bmp.Rect.Max.Y)
	if true {
		bitpos := 0
		for y := bmp.Rect.Min.Y; y < bmp.Rect.Max.Y; y += 4 {
			for x := bmp.Rect.Min.X; x < bmp.Rect.Max.X; x += 4 {
				//fmt.Printf("* (%d, %d)\n", x, y)
				cnt := 0

				for y1 := 0; y1 < 4; y1++ {
					for x1 := 0; x1 < 4; x1++ {
						r, _, _, _ := bmp.At(x+x1, y+y1).RGBA()
						if r != 0 {
							cnt++
						}
					}
				}
				if cnt > 0 {
					cnt--
				}
				fmt.Printf("%X", byte(cnt))
				b = (b << 4) + byte(cnt)

				bitpos += 4
				if bitpos == 8 {
					bitpos = 0
					//fmt.Printf(": %d %d %d : %d\n", x, y, w, ((x-4)/4+(y/4)*(w/4))/2)
					ret[(((x-4)/4)+(y/4)*(w/4))/2] = b
					b = 0
				}
			}
			fmt.Println()
		}
		if bitpos != 0 {
			bitpos = 0
			//bmp = append(bmp, b)
			// TODO:
		}
	}
	//fmt.Println("--")
	//for x := 0; x < w/4*h/4/2; x += 4 {
	//	fmt.Printf("%02X%02X%02X%02X\n", ret[x+0], ret[x+1], ret[x+2], ret[x+3])
	//}

	fmt.Println("--")
	for y := 0; y < h/4; y++ {
		for x := 0; x < w/4/2; x++ {
			fmt.Printf("%02X", ret[x+y*(w/4/2)])
		}
		fmt.Println()
	}

	g := tinyfont.Glyph{
		Rune:     r,
		Width:    uint8(w / 4),
		Height:   uint8(h / 4),
		XAdvance: uint8(f.Advance[0] / 4),
		XOffset:  int8(f.LowerPoint[0] / 4),
		YOffset:  -1 * int8(font.FontBoundingBox[1]+font.FontBoundingBox[3]) / 4,
		Bitmaps:  ret,
		Mode:     tinyfont.Grayscale16,
	}

	return g, nil
}
