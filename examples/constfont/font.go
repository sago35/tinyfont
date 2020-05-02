package main

import (
	"tinygo.org/x/tinyfont"
)

// 0 only
const fontdata = "\x00\x30\x08\x0C\x10\x04\xF6\x3C\x66\xC2\xC3\xC3\xC3\xC3\xC3\xC3\xC3\x66\x3C"

type Font struct {
}

var ConstFont = Font{}

func (f *Font) GetGlyph(r rune) tinyfont.Glyph {
	idx := 0
	ret := tinyfont.Glyph{
		Rune:     rune((uint16(fontdata[idx+0]) << 8) + uint16(fontdata[idx+1])),
		Width:    uint8(fontdata[idx+2]),
		Height:   uint8(fontdata[idx+3]),
		XAdvance: uint8(fontdata[idx+4]),
		XOffset:  int8(fontdata[idx+5]),
		YOffset:  int8(fontdata[idx+6]),
		Bitmaps:  []uint8(fontdata[idx+7:]),
	}

	return ret
}

func (f *Font) GetYAdvance() uint8 {
	return 0x0C
}
