package freeserif

import (
	"tinygo.org/x/tinyfont"
)

var Italic12pt7b = tinyfont.Font{
	Glyphs:[]tinyfont.Glyph{
		/*   */ tinyfont.Glyph{Width:0x0, Height:0x0, XAdvance:0x6, XOffset:0, YOffset:1, Bitmaps:[]uint8{}},
		/* ! */ tinyfont.Glyph{Width:0x6, Height:0x10, XAdvance:0x8, XOffset:1, YOffset:-15, Bitmaps:[]uint8{0xc, 0x31, 0xc6, 0x18, 0x43, 0xc, 0x20, 0x84, 0x10, 0x3, 0xc, 0x30}},
		/* " */ tinyfont.Glyph{Width:0x7, Height:0x6, XAdvance:0x8, XOffset:3, YOffset:-15, Bitmaps:[]uint8{0x66, 0xcd, 0x12, 0x24, 0x51, 0x0}},
		/* # */ tinyfont.Glyph{Width:0xd, Height:0x10, XAdvance:0xc, XOffset:0, YOffset:-15, Bitmaps:[]uint8{0x3, 0x10, 0x11, 0x80, 0x8c, 0xc, 0x40, 0x46, 0x1f, 0xfc, 0x21, 0x1, 0x18, 0x18, 0x80, 0x84, 0x3f, 0xf8, 0x62, 0x2, 0x30, 0x31, 0x1, 0x8, 0x8, 0xc0}},
		/* $ */ tinyfont.Glyph{Width:0xc, Height:0x14, XAdvance:0xc, XOffset:0, YOffset:-17, Bitmaps:[]uint8{0x0, 0x40, 0x8, 0x7, 0xc0, 0xca, 0x18, 0xa1, 0x92, 0x19, 0x1, 0xd0, 0xf, 0x0, 0x78, 0x3, 0xc0, 0x2e, 0x2, 0x64, 0x46, 0x44, 0x64, 0x46, 0x64, 0xc1, 0xf0, 0x8, 0x0, 0x80}},
		/* % */ tinyfont.Glyph{Width:0x11, Height:0x11, XAdvance:0x14, XOffset:2, YOffset:-16, Bitmaps:[]uint8{0x0, 0x8, 0xf, 0xc, 0xc, 0x7c, 0xc, 0x22, 0x6, 0x12, 0x6, 0x9, 0x3, 0x9, 0x1, 0x84, 0x80, 0xc4, 0x8f, 0x3c, 0x4c, 0x40, 0x4c, 0x20, 0x4e, 0x10, 0x26, 0x8, 0x23, 0x8, 0x11, 0x84, 0x10, 0xc4, 0x8, 0x3c, 0x0}},
		/* & */ tinyfont.Glyph{Width:0xf, Height:0x10, XAdvance:0x13, XOffset:2, YOffset:-15, Bitmaps:[]uint8{0x0, 0xe0, 0x2, 0x60, 0xc, 0xc0, 0x19, 0x80, 0x36, 0x0, 0x70, 0x0, 0xc0, 0x7, 0x9f, 0x33, 0x8, 0xc3, 0x13, 0x6, 0x46, 0xd, 0xc, 0xc, 0x18, 0x1c, 0x1c, 0x5c, 0x9f, 0x1e}},
		/* ' */ tinyfont.Glyph{Width:0x2, Height:0x6, XAdvance:0x5, XOffset:4, YOffset:-15, Bitmaps:[]uint8{0xfa, 0xa0}},
		/* ( */ tinyfont.Glyph{Width:0x7, Height:0x14, XAdvance:0x8, XOffset:1, YOffset:-15, Bitmaps:[]uint8{0x2, 0x8, 0x20, 0xc3, 0x6, 0x18, 0x30, 0xe1, 0x83, 0x6, 0xc, 0x18, 0x30, 0x60, 0x40, 0x80, 0x81, 0x0}},
		/* ) */ tinyfont.Glyph{Width:0x7, Height:0x14, XAdvance:0x8, XOffset:0, YOffset:-15, Bitmaps:[]uint8{0x8, 0x10, 0x10, 0x20, 0x40, 0xc1, 0x83, 0x6, 0xc, 0x18, 0x70, 0xc1, 0x83, 0xc, 0x10, 0x41, 0x4, 0x0}},
		/* * */ tinyfont.Glyph{Width:0x8, Height:0xa, XAdvance:0xc, XOffset:4, YOffset:-15, Bitmaps:[]uint8{0x18, 0x18, 0x18, 0x93, 0x74, 0x38, 0xd7, 0x93, 0x18, 0x18}},
		/* + */ tinyfont.Glyph{Width:0xb, Height:0xb, XAdvance:0x10, XOffset:2, YOffset:-10, Bitmaps:[]uint8{0x4, 0x0, 0x80, 0x10, 0x2, 0x0, 0x41, 0xff, 0xc1, 0x0, 0x20, 0x4, 0x0, 0x80, 0x10, 0x0}},
		/* , */ tinyfont.Glyph{Width:0x3, Height:0x6, XAdvance:0x6, XOffset:0, YOffset:-2, Bitmaps:[]uint8{0x6c, 0x95, 0x0}},
		/* - */ tinyfont.Glyph{Width:0x5, Height:0x1, XAdvance:0x8, XOffset:1, YOffset:-5, Bitmaps:[]uint8{0xf8}},
		/* . */ tinyfont.Glyph{Width:0x2, Height:0x3, XAdvance:0x6, XOffset:1, YOffset:-2, Bitmaps:[]uint8{0xfc}},
		/* / */ tinyfont.Glyph{Width:0xb, Height:0x10, XAdvance:0x7, XOffset:0, YOffset:-15, Bitmaps:[]uint8{0x0, 0x40, 0x18, 0x2, 0x0, 0xc0, 0x30, 0x6, 0x1, 0x80, 0x20, 0xc, 0x1, 0x0, 0x60, 0x18, 0x3, 0x0, 0xc0, 0x10, 0x6, 0x0}},
		/* 0 */ tinyfont.Glyph{Width:0xb, Height:0x11, XAdvance:0xc, XOffset:1, YOffset:-16, Bitmaps:[]uint8{0x7, 0x81, 0x98, 0x61, 0x18, 0x33, 0x6, 0xc0, 0xd8, 0x1b, 0x3, 0xe0, 0xf8, 0x1f, 0x3, 0x60, 0x6c, 0x19, 0x83, 0x10, 0xc3, 0x30, 0x3c, 0x0}},
		/* 1 */ tinyfont.Glyph{Width:0x9, Height:0x11, XAdvance:0xc, XOffset:1, YOffset:-16, Bitmaps:[]uint8{0x1, 0x87, 0xc0, 0xc0, 0x60, 0x30, 0x18, 0x18, 0xc, 0x6, 0x7, 0x3, 0x1, 0x80, 0xc0, 0xc0, 0x60, 0x30, 0xfe, 0x0}},
		/* 2 */ tinyfont.Glyph{Width:0xa, Height:0xf, XAdvance:0xc, XOffset:1, YOffset:-14, Bitmaps:[]uint8{0xf, 0xc, 0x64, 0xc, 0x3, 0x0, 0xc0, 0x20, 0x18, 0xc, 0x2, 0x1, 0x0, 0x80, 0x40, 0x20, 0x10, 0x2f, 0xf0}},
		/* 3 */ tinyfont.Glyph{Width:0xa, Height:0x10, XAdvance:0xc, XOffset:1, YOffset:-15, Bitmaps:[]uint8{0x7, 0x86, 0x30, 0xc, 0x3, 0x1, 0x81, 0x81, 0xf0, 0x1e, 0x3, 0x80, 0x60, 0x18, 0x6, 0x1, 0x0, 0xcc, 0x63, 0xe0}},
		/* 4 */ tinyfont.Glyph{Width:0xb, Height:0x10, XAdvance:0xc, XOffset:0, YOffset:-15, Bitmaps:[]uint8{0x0, 0x20, 0xc, 0x3, 0x80, 0xa0, 0x2c, 0x9, 0x82, 0x30, 0x84, 0x31, 0x8c, 0x33, 0x6, 0x7f, 0xe0, 0x30, 0x6, 0x0, 0x80, 0x30}},
		/* 5 */ tinyfont.Glyph{Width:0xb, Height:0x10, XAdvance:0xc, XOffset:0, YOffset:-15, Bitmaps:[]uint8{0x3, 0xe1, 0x80, 0x20, 0x6, 0x0, 0xf0, 0xf, 0x0, 0x60, 0x6, 0x0, 0xc0, 0x18, 0x3, 0x0, 0x40, 0x18, 0x2, 0x30, 0x87, 0xe0}},
		/* 6 */ tinyfont.Glyph{Width:0xc, Height:0x11, XAdvance:0xc, XOffset:1, YOffset:-16, Bitmaps:[]uint8{0x0, 0x70, 0x3c, 0x7, 0x0, 0xe0, 0x1c, 0x3, 0x80, 0x7f, 0x7, 0x18, 0x60, 0xce, 0xc, 0xc0, 0xcc, 0xc, 0xc0, 0xcc, 0x18, 0x41, 0x86, 0x30, 0x3e, 0x0}},
		/* 7 */ tinyfont.Glyph{Width:0xb, Height:0x10, XAdvance:0xc, XOffset:2, YOffset:-15, Bitmaps:[]uint8{0x7f, 0xf0, 0x18, 0x3, 0x0, 0xc0, 0x10, 0x6, 0x1, 0x80, 0x30, 0xc, 0x1, 0x0, 0x60, 0x8, 0x3, 0x0, 0xc0, 0x10, 0x6, 0x0}},
		/* 8 */ tinyfont.Glyph{Width:0xb, Height:0x11, XAdvance:0xc, XOffset:1, YOffset:-16, Bitmaps:[]uint8{0xf, 0x83, 0x18, 0xc1, 0x98, 0x33, 0x6, 0x71, 0x87, 0x60, 0x70, 0x17, 0xc, 0x71, 0x7, 0x60, 0x6c, 0xd, 0x81, 0xb0, 0x63, 0x1c, 0x3e, 0x0}},
		/* 9 */ tinyfont.Glyph{Width:0xb, Height:0x11, XAdvance:0xc, XOffset:1, YOffset:-16, Bitmaps:[]uint8{0x7, 0x83, 0x18, 0xc1, 0x18, 0x36, 0x6, 0xc0, 0xd8, 0x1b, 0x7, 0x60, 0xe6, 0x38, 0x7f, 0x0, 0xc0, 0x30, 0xc, 0x7, 0x3, 0xc0, 0xc0, 0x0}},
		/* : */ tinyfont.Glyph{Width:0x4, Height:0xb, XAdvance:0x6, XOffset:1, YOffset:-10, Bitmaps:[]uint8{0x33, 0x30, 0x0, 0x0, 0xcc, 0xc0}},
		/* ; */ tinyfont.Glyph{Width:0x5, Height:0xe, XAdvance:0x6, XOffset:0, YOffset:-10, Bitmaps:[]uint8{0x18, 0xc6, 0x0, 0x0, 0x0, 0x3, 0x18, 0x44, 0x40}},
		/* < */ tinyfont.Glyph{Width:0xc, Height:0xd, XAdvance:0xe, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0x0, 0x0, 0x3, 0x0, 0xf0, 0x38, 0x1e, 0x7, 0x80, 0xe0, 0xf, 0x0, 0x1c, 0x0, 0x78, 0x1, 0xe0, 0x7, 0x0, 0x10}},
		/* = */ tinyfont.Glyph{Width:0xc, Height:0x6, XAdvance:0x10, XOffset:2, YOffset:-8, Bitmaps:[]uint8{0xff, 0xf0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xf, 0xff}},
		/* > */ tinyfont.Glyph{Width:0xc, Height:0xd, XAdvance:0xe, XOffset:2, YOffset:-12, Bitmaps:[]uint8{0x0, 0xc, 0x0, 0xf0, 0x1, 0xc0, 0x7, 0x80, 0x1e, 0x0, 0x70, 0xf, 0x3, 0xc1, 0xe0, 0x78, 0xe, 0x0, 0x80, 0x0}},
		/* ? */ tinyfont.Glyph{Width:0x9, Height:0x10, XAdvance:0xb, XOffset:3, YOffset:-15, Bitmaps:[]uint8{0x3e, 0x21, 0x90, 0x60, 0x30, 0x38, 0x38, 0x30, 0x30, 0x20, 0x20, 0x10, 0x0, 0x0, 0x6, 0x3, 0x1, 0x80}},
		/* @ */ tinyfont.Glyph{Width:0x10, Height:0x10, XAdvance:0x13, XOffset:2, YOffset:-15, Bitmaps:[]uint8{0x7, 0xe0, 0x1c, 0x18, 0x30, 0x4, 0x60, 0x2, 0x61, 0xda, 0xc3, 0x31, 0xc6, 0x31, 0xc4, 0x31, 0xcc, 0x31, 0xcc, 0x21, 0xcc, 0x62, 0x6c, 0xe4, 0x67, 0x38, 0x30, 0x0, 0x1c, 0x8, 0x7, 0xf0}},
		/* A */ tinyfont.Glyph{Width:0xf, Height:0xf, XAdvance:0x10, XOffset:0, YOffset:-14, Bitmaps:[]uint8{0x0, 0x20, 0x0, 0xc0, 0x3, 0x80, 0xb, 0x0, 0x16, 0x0, 0x4e, 0x0, 0x9c, 0x2, 0x18, 0x8, 0x30, 0x1f, 0xe0, 0x40, 0xc1, 0x81, 0xc2, 0x3, 0x8c, 0x7, 0x3c, 0x1f, 0x80}},
		/* B */ tinyfont.Glyph{Width:0xe, Height:0x10, XAdvance:0xe, XOffset:0, YOffset:-15, Bitmaps:[]uint8{0x1f, 0xf0, 0x1c, 0x60, 0x60, 0xc1, 0x83, 0x6, 0xc, 0x38, 0x60, 0xc3, 0x3, 0xf0, 0x1c, 0x30, 0x60, 0x61, 0x81, 0x86, 0x6, 0x38, 0x18, 0xc0, 0xc3, 0x6, 0x3f, 0xf0}},
		/* C */ tinyfont.Glyph{Width:0x10, Height:0x10, XAdvance:0xf, XOffset:1, YOffset:-15, Bitmaps:[]uint8{0x1, 0xf9, 0x6, 0xf, 0x1c, 0x6, 0x38, 0x2, 0x30, 0x2, 0x60, 0x0, 0x60, 0x0, 0xc0, 0x0, 0xc0, 0x0, 0xc0, 0x0, 0xc0, 0x0, 0xc0, 0x0, 0xc0, 0x8, 0x60, 0x10, 0x30, 0x60, 0x1f, 0x80}},
		/* D */ tinyfont.Glyph{Width:0x10, Height:0x10, XAdvance:0x11, XOffset:0, YOffset:-15, Bitmaps:[]uint8{0x1f, 0xf0, 0x7, 0xc, 0x6, 0x6, 0x6, 0x6, 0x6, 0x3, 0xe, 0x3, 0xc, 0x3, 0xc, 0x3, 0x1c, 0x3, 0x1c, 0x7, 0x18, 0x6, 0x18, 0x6, 0x38, 0xc, 0x30, 0x18, 0x30, 0x70, 0xff, 0x80}},
		/* E */ tinyfont.Glyph{Width:0x10, Height:0x10, XAdvance:0xe, XOffset:0, YOffset:-15, Bitmaps:[]uint8{0x1f, 0xff, 0x7, 0x7, 0x6, 0x2, 0x6, 0x2, 0x6, 0x0, 0xe, 0x10, 0xc, 0x30, 0xf, 0xf0, 0x1c, 0x20, 0x18, 0x20, 0x18, 0x0, 0x18, 0x0, 0x38, 0x4, 0x30, 0x8, 0x30, 0x38, 0xff, 0xf8}},
		/* F */ tinyfont.Glyph{Width:0x10, Height:0x10, XAdvance:0xe, XOffset:0, YOffset:-15, Bitmaps:[]uint8{0x1f, 0xff, 0x7, 0x7, 0x7, 0x2, 0x6, 0x2, 0x6, 0x0, 0xe, 0x10, 0xc, 0x30, 0xf, 0xf0, 0x1c, 0x20, 0x1c, 0x20, 0x18, 0x0, 0x18, 0x0, 0x38, 0x0, 0x30, 0x0, 0x30, 0x0, 0xfc, 0x0}},
		/* G */ tinyfont.Glyph{Width:0x10, Height:0x10, XAdvance:0x11, XOffset:1, YOffset:-15, Bitmaps:[]uint8{0x1, 0xf1, 0x6, 0xf, 0x18, 0x7, 0x38, 0x2, 0x30, 0x2, 0x60, 0x0, 0x60, 0x0, 0xe0, 0x0, 0xc0, 0x7f, 0xc0, 0x1c, 0xc0, 0x1c, 0xc0, 0x18, 0xc0, 0x18, 0x60, 0x18, 0x30, 0x38, 0xf, 0xc0}},
		/* H */ tinyfont.Glyph{Width:0x13, Height:0x10, XAdvance:0x11, XOffset:0, YOffset:-15, Bitmaps:[]uint8{0x1f, 0xc7, 0xe0, 0xe0, 0x70, 0x18, 0xe, 0x3, 0x1, 0x80, 0x60, 0x30, 0x1c, 0xe, 0x3, 0x1, 0x80, 0x7f, 0xf0, 0x1c, 0x6, 0x3, 0x1, 0xc0, 0x60, 0x30, 0xc, 0x6, 0x3, 0x81, 0xc0, 0x60, 0x38, 0xc, 0x6, 0x7, 0xe3, 0xf0}},
		/* I */ tinyfont.Glyph{Width:0x9, Height:0x10, XAdvance:0x8, XOffset:0, YOffset:-15, Bitmaps:[]uint8{0x1f, 0x83, 0x81, 0x80, 0xc0, 0x60, 0x70, 0x30, 0x18, 0x1c, 0xc, 0x6, 0x3, 0x3, 0x81, 0x80, 0xc1, 0xf8}},
		/* J */ tinyfont.Glyph{Width:0xc, Height:0x10, XAdvance:0xa, XOffset:0, YOffset:-15, Bitmaps:[]uint8{0x3, 0xf0, 0xc, 0x0, 0xc0, 0x1c, 0x1, 0x80, 0x18, 0x3, 0x80, 0x30, 0x3, 0x0, 0x30, 0x7, 0x0, 0x60, 0x6, 0xc, 0xe0, 0xcc, 0x7, 0x80}},
		/* K */ tinyfont.Glyph{Width:0x11, Height:0x10, XAdvance:0xf, XOffset:0, YOffset:-15, Bitmaps:[]uint8{0x1f, 0xcf, 0x83, 0x83, 0x81, 0x81, 0x0, 0xc3, 0x0, 0x62, 0x0, 0x72, 0x0, 0x36, 0x0, 0x1e, 0x0, 0x1d, 0x80, 0xc, 0xe0, 0x6, 0x30, 0x3, 0x1c, 0x3, 0x87, 0x1, 0x81, 0x80, 0xc0, 0xe1, 0xf9, 0xfc}},
		/* L */ tinyfont.Glyph{Width:0xe, Height:0x10, XAdvance:0xe, XOffset:0, YOffset:-15, Bitmaps:[]uint8{0x1f, 0xc0, 0x1c, 0x0, 0x60, 0x1, 0x80, 0x6, 0x0, 0x38, 0x0, 0xc0, 0x3, 0x0, 0x1c, 0x0, 0x60, 0x1, 0x80, 0x6, 0x0, 0x38, 0xc, 0xc0, 0x23, 0x3, 0xbf, 0xfe}},
		/* M */ tinyfont.Glyph{Width:0x15, Height:0x10, XAdvance:0x14, XOffset:0, YOffset:-15, Bitmaps:[]uint8{0xf, 0x0, 0x78, 0x38, 0x7, 0x81, 0xc0, 0x38, 0xe, 0x2, 0xc0, 0x70, 0x3e, 0x5, 0xc1, 0x70, 0x2e, 0x13, 0x1, 0x31, 0x98, 0x11, 0x89, 0xc0, 0x8c, 0x8c, 0x4, 0x6c, 0x60, 0x23, 0x43, 0x2, 0x1c, 0x38, 0x10, 0xe1, 0x81, 0x86, 0x1c, 0x1f, 0x23, 0xf8}},
		/* N */ tinyfont.Glyph{Width:0x12, Height:0x10, XAdvance:0x10, XOffset:0, YOffset:-15, Bitmaps:[]uint8{0x1e, 0x7, 0xc1, 0xc0, 0x60, 0x70, 0x10, 0x1c, 0xc, 0x5, 0x82, 0x2, 0x60, 0x80, 0x9c, 0x60, 0x23, 0x10, 0x10, 0xc4, 0x4, 0x19, 0x1, 0x6, 0xc0, 0x40, 0xe0, 0x20, 0x38, 0x8, 0xe, 0x6, 0x1, 0x3, 0xe0, 0x40}},
		/* O */ tinyfont.Glyph{Width:0xf, Height:0x10, XAdvance:0x10, XOffset:1, YOffset:-15, Bitmaps:[]uint8{0x1, 0xf0, 0xc, 0x10, 0x30, 0x10, 0xc0, 0x33, 0x0, 0x6e, 0x0, 0xd8, 0x1, 0xf0, 0x3, 0xc0, 0xd, 0x80, 0x1b, 0x0, 0x76, 0x0, 0xcc, 0x3, 0x8, 0xc, 0x18, 0x70, 0xf, 0x80}},
		/* P */ tinyfont.Glyph{Width:0xe, Height:0x10, XAdvance:0xe, XOffset:0, YOffset:-15, Bitmaps:[]uint8{0x1f, 0xf0, 0x1c, 0x60, 0x60, 0xc1, 0x83, 0x6, 0xc, 0x38, 0x30, 0xc1, 0x83, 0xe, 0x1f, 0xe0, 0x60, 0x1, 0x80, 0x6, 0x0, 0x38, 0x0, 0xc0, 0x3, 0x0, 0x3f, 0x0}},
		/* Q */ tinyfont.Glyph{Width:0xf, Height:0x14, XAdvance:0x10, XOffset:1, YOffset:-15, Bitmaps:[]uint8{0x1, 0xf0, 0x6, 0x10, 0x30, 0x30, 0xc0, 0x33, 0x0, 0x66, 0x0, 0xd8, 0x1, 0xb0, 0x3, 0xe0, 0xf, 0x80, 0x1b, 0x0, 0x36, 0x0, 0xcc, 0x3, 0x98, 0x6, 0x18, 0x18, 0x18, 0xc0, 0xe, 0x0, 0x20, 0x1, 0xf8, 0x36, 0x7f, 0x80}},
		/* R */ tinyfont.Glyph{Width:0xe, Height:0x10, XAdvance:0xf, XOffset:0, YOffset:-15, Bitmaps:[]uint8{0x1f, 0xf0, 0x1c, 0x60, 0x60, 0xc1, 0x83, 0x6, 0xc, 0x38, 0x70, 0xc3, 0x83, 0xf8, 0x1d, 0xc0, 0x63, 0x1, 0x8c, 0x6, 0x18, 0x38, 0x60, 0xc1, 0xc3, 0x3, 0x3f, 0xf}},
		/* S */ tinyfont.Glyph{Width:0xc, Height:0x10, XAdvance:0xb, XOffset:0, YOffset:-15, Bitmaps:[]uint8{0x7, 0x90, 0xc7, 0x18, 0x21, 0x82, 0x18, 0x1, 0xc0, 0xe, 0x0, 0x70, 0x3, 0x80, 0x1c, 0x0, 0xc4, 0xc, 0x40, 0xc6, 0x8, 0xe1, 0x89, 0xe0}},
		/* T */ tinyfont.Glyph{Width:0xf, Height:0x10, XAdvance:0xe, XOffset:2, YOffset:-15, Bitmaps:[]uint8{0x7f, 0xfe, 0xc7, 0x1d, 0xc, 0x14, 0x18, 0x20, 0x70, 0x0, 0xe0, 0x1, 0x80, 0x3, 0x0, 0xe, 0x0, 0x18, 0x0, 0x30, 0x0, 0x60, 0x1, 0xc0, 0x3, 0x0, 0xe, 0x0, 0x7f, 0x80}},
		/* U */ tinyfont.Glyph{Width:0x10, Height:0x10, XAdvance:0x11, XOffset:3, YOffset:-15, Bitmaps:[]uint8{0x7e, 0x1f, 0x38, 0xc, 0x38, 0xc, 0x30, 0x8, 0x30, 0x8, 0x70, 0x8, 0x70, 0x10, 0x60, 0x10, 0x60, 0x10, 0xe0, 0x10, 0xc0, 0x20, 0xc0, 0x20, 0xc0, 0x60, 0xc0, 0x40, 0x61, 0x80, 0x3f, 0x0}},
		/* V */ tinyfont.Glyph{Width:0xf, Height:0x10, XAdvance:0x10, XOffset:3, YOffset:-15, Bitmaps:[]uint8{0xfc, 0x3e, 0xe0, 0x18, 0xc0, 0x21, 0x80, 0xc3, 0x81, 0x7, 0x4, 0xe, 0x8, 0xc, 0x20, 0x18, 0x80, 0x31, 0x0, 0x64, 0x0, 0xf0, 0x1, 0xe0, 0x1, 0x80, 0x2, 0x0, 0x4, 0x0}},
		/* W */ tinyfont.Glyph{Width:0x14, Height:0x10, XAdvance:0x15, XOffset:3, YOffset:-15, Bitmaps:[]uint8{0xfd, 0xf8, 0xf7, 0x7, 0x6, 0x30, 0x60, 0x63, 0x7, 0x4, 0x30, 0x70, 0x83, 0x8f, 0x8, 0x38, 0xb1, 0x3, 0x93, 0x10, 0x19, 0x32, 0x1, 0xa3, 0x20, 0x1a, 0x34, 0x1, 0xc3, 0x40, 0x1c, 0x38, 0x1, 0x83, 0x0, 0x18, 0x30, 0x1, 0x2, 0x0}},
		/* X */ tinyfont.Glyph{Width:0x10, Height:0x10, XAdvance:0x10, XOffset:0, YOffset:-15, Bitmaps:[]uint8{0x1f, 0x9f, 0xe, 0x6, 0x6, 0x4, 0x7, 0x8, 0x3, 0x10, 0x3, 0x20, 0x3, 0xc0, 0x1, 0x80, 0x1, 0xc0, 0x3, 0xc0, 0x6, 0xe0, 0xc, 0x60, 0x18, 0x60, 0x30, 0x70, 0x70, 0x78, 0xf8, 0xfc}},
		/* Y */ tinyfont.Glyph{Width:0xd, Height:0x10, XAdvance:0xe, XOffset:3, YOffset:-15, Bitmaps:[]uint8{0xfc, 0xfb, 0x81, 0x8c, 0x8, 0x60, 0x83, 0x8c, 0xc, 0xc0, 0x64, 0x3, 0xc0, 0xc, 0x0, 0xe0, 0x7, 0x0, 0x30, 0x1, 0x80, 0x1c, 0x0, 0xc0, 0x1f, 0xc0}},
		/* Z */ tinyfont.Glyph{Width:0xf, Height:0x10, XAdvance:0xe, XOffset:0, YOffset:-15, Bitmaps:[]uint8{0x1f, 0xfe, 0x30, 0x38, 0xc0, 0xf1, 0x1, 0xc0, 0x7, 0x0, 0x1c, 0x0, 0x70, 0x1, 0xe0, 0x3, 0x80, 0xe, 0x0, 0x38, 0x0, 0xe0, 0x1, 0xc0, 0x47, 0x1, 0x1c, 0x6, 0x7f, 0xf8}},
		/* [ */ tinyfont.Glyph{Width:0x8, Height:0x14, XAdvance:0x9, XOffset:1, YOffset:-15, Bitmaps:[]uint8{0x7, 0x4, 0x8, 0x8, 0x8, 0x18, 0x10, 0x10, 0x10, 0x20, 0x20, 0x20, 0x20, 0x40, 0x40, 0x40, 0x80, 0x80, 0x80, 0xe0}},
		/* \ */ tinyfont.Glyph{Width:0x8, Height:0x10, XAdvance:0xc, XOffset:3, YOffset:-15, Bitmaps:[]uint8{0xc0, 0xc0, 0x40, 0x60, 0x20, 0x30, 0x30, 0x18, 0x18, 0x8, 0xc, 0x4, 0x6, 0x6, 0x3, 0x3}},
		/* ] */ tinyfont.Glyph{Width:0x7, Height:0x14, XAdvance:0x9, XOffset:1, YOffset:-15, Bitmaps:[]uint8{0xe, 0x4, 0x8, 0x10, 0x60, 0x81, 0x2, 0x4, 0x18, 0x20, 0x40, 0x81, 0x2, 0x8, 0x10, 0x20, 0x47, 0x80}},
		/* ^ */ tinyfont.Glyph{Width:0xa, Height:0x9, XAdvance:0xa, XOffset:0, YOffset:-15, Bitmaps:[]uint8{0xc, 0x3, 0x81, 0xe0, 0x4c, 0x33, 0x8, 0x66, 0x19, 0x3, 0xc0, 0xc0}},
		/* _ */ tinyfont.Glyph{Width:0xc, Height:0x1, XAdvance:0xc, XOffset:0, YOffset:3, Bitmaps:[]uint8{0xff, 0xf0}},
		/* ` */ tinyfont.Glyph{Width:0x4, Height:0x4, XAdvance:0x6, XOffset:3, YOffset:-15, Bitmaps:[]uint8{0xce, 0x63}},
		/* a */ tinyfont.Glyph{Width:0xc, Height:0xb, XAdvance:0xc, XOffset:0, YOffset:-10, Bitmaps:[]uint8{0x7, 0xa0, 0xce, 0x18, 0x63, 0x4, 0x60, 0xc6, 0xc, 0xc0, 0xcc, 0x18, 0xc3, 0x8c, 0x5a, 0x79, 0xc0}},
		/* b */ tinyfont.Glyph{Width:0xa, Height:0x10, XAdvance:0xb, XOffset:1, YOffset:-15, Bitmaps:[]uint8{0x38, 0x6, 0x1, 0x80, 0x40, 0x30, 0xc, 0xe3, 0xcc, 0xc3, 0x70, 0xd8, 0x36, 0x19, 0x6, 0xc3, 0x30, 0x8c, 0xc3, 0xe0}},
		/* c */ tinyfont.Glyph{Width:0x9, Height:0xb, XAdvance:0xa, XOffset:1, YOffset:-10, Bitmaps:[]uint8{0xf, 0xc, 0xcc, 0x6c, 0x6, 0x6, 0x3, 0x1, 0x80, 0xc0, 0x73, 0x1e, 0x0}},
		/* d */ tinyfont.Glyph{Width:0xd, Height:0x10, XAdvance:0xc, XOffset:0, YOffset:-15, Bitmaps:[]uint8{0x0, 0x70, 0x1, 0x80, 0xc, 0x0, 0x60, 0x2, 0x3, 0xf0, 0x31, 0x83, 0x8, 0x30, 0xc3, 0x6, 0x18, 0x31, 0x81, 0x8c, 0x18, 0x61, 0xcb, 0x16, 0x8f, 0x38}},
		/* e */ tinyfont.Glyph{Width:0x8, Height:0xb, XAdvance:0xa, XOffset:1, YOffset:-10, Bitmaps:[]uint8{0x7, 0x19, 0x31, 0x63, 0x62, 0xec, 0xd0, 0xc0, 0xc0, 0xe6, 0x78}},
		/* f */ tinyfont.Glyph{Width:0xe, Height:0x16, XAdvance:0xa, XOffset:-2, YOffset:-16, Bitmaps:[]uint8{0x0, 0x38, 0x1, 0x30, 0xc, 0x0, 0x20, 0x1, 0x80, 0x6, 0x0, 0xfe, 0x0, 0x40, 0x3, 0x0, 0xc, 0x0, 0x30, 0x0, 0x80, 0x6, 0x0, 0x18, 0x0, 0x60, 0x1, 0x80, 0x4, 0x0, 0x30, 0x0, 0xc0, 0x2, 0x0, 0x90, 0x3, 0x80, 0x0}},
		/* g */ tinyfont.Glyph{Width:0xc, Height:0x10, XAdvance:0xb, XOffset:-1, YOffset:-10, Bitmaps:[]uint8{0x7, 0xc0, 0xc7, 0x18, 0x61, 0x86, 0x18, 0xe1, 0x8c, 0x7, 0x80, 0x80, 0x1c, 0x0, 0xf0, 0x33, 0x84, 0x18, 0x80, 0x88, 0x8, 0x61, 0x3, 0xe0}},
		/* h */ tinyfont.Glyph{Width:0xc, Height:0x10, XAdvance:0xc, XOffset:0, YOffset:-15, Bitmaps:[]uint8{0x1c, 0x0, 0xc0, 0xc, 0x0, 0xc0, 0x18, 0x1, 0x8e, 0x1b, 0x61, 0xc6, 0x38, 0x63, 0x8c, 0x30, 0xc3, 0xc, 0x60, 0xc6, 0x1a, 0x61, 0xa4, 0x1c}},
		/* i */ tinyfont.Glyph{Width:0x5, Height:0x10, XAdvance:0x6, XOffset:1, YOffset:-15, Bitmaps:[]uint8{0x18, 0xc6, 0x0, 0xb, 0xc6, 0x23, 0x18, 0x8c, 0x63, 0x5c}},
		/* j */ tinyfont.Glyph{Width:0x9, Height:0x15, XAdvance:0x7, XOffset:-2, YOffset:-15, Bitmaps:[]uint8{0x1, 0x80, 0xc0, 0x60, 0x0, 0x0, 0xc, 0x1e, 0x2, 0x3, 0x1, 0x80, 0xc0, 0x40, 0x60, 0x30, 0x18, 0x8, 0xc, 0x6, 0x2, 0x1b, 0xf, 0x0}},
		/* k */ tinyfont.Glyph{Width:0xb, Height:0x10, XAdvance:0xb, XOffset:0, YOffset:-15, Bitmaps:[]uint8{0x1c, 0x1, 0x80, 0x30, 0x6, 0x1, 0x80, 0x33, 0xc6, 0x30, 0x88, 0x32, 0x6, 0x80, 0xf0, 0x1b, 0x6, 0x60, 0xc4, 0x18, 0xd2, 0xc}},
		/* l */ tinyfont.Glyph{Width:0x6, Height:0x10, XAdvance:0x6, XOffset:1, YOffset:-15, Bitmaps:[]uint8{0x3c, 0x61, 0x86, 0x18, 0xc3, 0xc, 0x21, 0x86, 0x18, 0x43, 0x2d, 0x38}},
		/* m */ tinyfont.Glyph{Width:0x11, Height:0xb, XAdvance:0x11, XOffset:0, YOffset:-10, Bitmaps:[]uint8{0x78, 0xe7, 0xd, 0xb5, 0x8d, 0x1c, 0xc7, 0xc, 0x63, 0x8e, 0x31, 0x86, 0x30, 0xc3, 0x18, 0xc1, 0xc, 0x61, 0x84, 0xb0, 0xc6, 0xb0, 0x63, 0x80}},
		/* n */ tinyfont.Glyph{Width:0xc, Height:0xb, XAdvance:0xc, XOffset:0, YOffset:-10, Bitmaps:[]uint8{0x78, 0xe1, 0xb6, 0x14, 0x63, 0x84, 0x38, 0xc3, 0xc, 0x70, 0x86, 0x18, 0x61, 0x96, 0x1a, 0xc1, 0xc0}},
		/* o */ tinyfont.Glyph{Width:0xa, Height:0xb, XAdvance:0xb, XOffset:1, YOffset:-10, Bitmaps:[]uint8{0xf, 0x6, 0x63, 0xd, 0x83, 0x60, 0xf0, 0x3c, 0x1b, 0x6, 0xc3, 0x39, 0x87, 0x80}},
		/* p */ tinyfont.Glyph{Width:0xd, Height:0x10, XAdvance:0xb, XOffset:-2, YOffset:-10, Bitmaps:[]uint8{0x1e, 0xf0, 0x39, 0xc1, 0x86, 0xc, 0x30, 0xc1, 0x86, 0xc, 0x30, 0xc3, 0x6, 0x18, 0x60, 0xc6, 0x7, 0xc0, 0x60, 0x3, 0x0, 0x18, 0x0, 0xc0, 0x1f, 0x0}},
		/* q */ tinyfont.Glyph{Width:0xb, Height:0x10, XAdvance:0xc, XOffset:0, YOffset:-10, Bitmaps:[]uint8{0x7, 0x81, 0x9c, 0x63, 0x98, 0x76, 0xc, 0xc1, 0xb0, 0x76, 0xe, 0xc3, 0x98, 0xb1, 0xe6, 0x0, 0x80, 0x30, 0x6, 0x0, 0xc0, 0xfc}},
		/* r */ tinyfont.Glyph{Width:0x9, Height:0xb, XAdvance:0x9, XOffset:0, YOffset:-10, Bitmaps:[]uint8{0x79, 0x8f, 0xc5, 0x7, 0x3, 0x1, 0x80, 0xc0, 0xc0, 0x60, 0x30, 0x10, 0x0}},
		/* s */ tinyfont.Glyph{Width:0x9, Height:0xb, XAdvance:0x8, XOffset:0, YOffset:-10, Bitmaps:[]uint8{0x1e, 0x98, 0xcc, 0x27, 0x11, 0x80, 0xe0, 0x39, 0xc, 0x86, 0x62, 0x2e, 0x0}},
		/* t */ tinyfont.Glyph{Width:0x6, Height:0xd, XAdvance:0x6, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0x8, 0x67, 0xcc, 0x30, 0xc6, 0x18, 0x61, 0x8c, 0x34, 0xe0}},
		/* u */ tinyfont.Glyph{Width:0xb, Height:0xb, XAdvance:0xc, XOffset:1, YOffset:-10, Bitmaps:[]uint8{0xf0, 0xcc, 0x19, 0x83, 0x30, 0xc6, 0x18, 0x87, 0x31, 0x66, 0x3c, 0xcb, 0x1a, 0x6b, 0x8e, 0x0}},
		/* v */ tinyfont.Glyph{Width:0xa, Height:0xb, XAdvance:0xb, XOffset:1, YOffset:-10, Bitmaps:[]uint8{0x70, 0xcc, 0x33, 0x4, 0xc2, 0x18, 0x86, 0x41, 0x90, 0x68, 0x1c, 0x6, 0x1, 0x0}},
		/* w */ tinyfont.Glyph{Width:0xe, Height:0xb, XAdvance:0x10, XOffset:2, YOffset:-10, Bitmaps:[]uint8{0x61, 0xf, 0x84, 0x36, 0x30, 0xdc, 0xc1, 0x35, 0x8, 0xd4, 0x23, 0x91, 0xe, 0x48, 0x30, 0xe0, 0xc3, 0x2, 0x8, 0x0}},
		/* x */ tinyfont.Glyph{Width:0xc, Height:0xb, XAdvance:0xa, XOffset:-1, YOffset:-10, Bitmaps:[]uint8{0xc, 0x63, 0x4a, 0x7, 0x0, 0x70, 0x6, 0x0, 0x20, 0x7, 0x0, 0xb0, 0xb, 0x21, 0x14, 0xe1, 0x80}},
		/* y */ tinyfont.Glyph{Width:0xb, Height:0x10, XAdvance:0xb, XOffset:0, YOffset:-10, Bitmaps:[]uint8{0x38, 0x63, 0xc, 0x30, 0x86, 0x10, 0xc4, 0xc, 0x81, 0xa0, 0x34, 0x7, 0x0, 0x60, 0x8, 0x2, 0x0, 0x40, 0x10, 0x4, 0x7, 0x0}},
		/* z */ tinyfont.Glyph{Width:0x9, Height:0xd, XAdvance:0x9, XOffset:0, YOffset:-10, Bitmaps:[]uint8{0x1f, 0x90, 0x80, 0x80, 0xc0, 0xc0, 0x40, 0x60, 0x60, 0x60, 0x38, 0x3e, 0x3, 0xa0, 0x60}},
		/* { */ tinyfont.Glyph{Width:0x9, Height:0x15, XAdvance:0xa, XOffset:1, YOffset:-16, Bitmaps:[]uint8{0x0, 0x83, 0x81, 0x1, 0x80, 0xc0, 0x40, 0x60, 0x30, 0x10, 0x10, 0x1c, 0x6, 0x3, 0x3, 0x1, 0x80, 0xc0, 0x40, 0x60, 0x30, 0x18, 0x7, 0x0}},
		/* | */ tinyfont.Glyph{Width:0x1, Height:0x10, XAdvance:0x7, XOffset:3, YOffset:-15, Bitmaps:[]uint8{0xff, 0xff}},
		/* } */ tinyfont.Glyph{Width:0x9, Height:0x15, XAdvance:0xa, XOffset:0, YOffset:-16, Bitmaps:[]uint8{0x7, 0x0, 0xc0, 0x60, 0x30, 0x10, 0x18, 0xc, 0x6, 0x6, 0x3, 0x1, 0x80, 0x60, 0x40, 0x60, 0x30, 0x10, 0x18, 0xc, 0x6, 0x6, 0x6, 0x0}},
		/* ~ */ tinyfont.Glyph{Width:0xb, Height:0x3, XAdvance:0xd, XOffset:1, YOffset:-6, Bitmaps:[]uint8{0x78, 0x18, 0x8c, 0xf, 0x0}},
	},

	RuneToIndex:[]tinyfont.RuneToIndex{
		/*   */ tinyfont.RuneToIndex{Rune:32, Index:0x0},
		/* ! */ tinyfont.RuneToIndex{Rune:33, Index:0x1},
		/* " */ tinyfont.RuneToIndex{Rune:34, Index:0x2},
		/* # */ tinyfont.RuneToIndex{Rune:35, Index:0x3},
		/* $ */ tinyfont.RuneToIndex{Rune:36, Index:0x4},
		/* % */ tinyfont.RuneToIndex{Rune:37, Index:0x5},
		/* & */ tinyfont.RuneToIndex{Rune:38, Index:0x6},
		/* ' */ tinyfont.RuneToIndex{Rune:39, Index:0x7},
		/* ( */ tinyfont.RuneToIndex{Rune:40, Index:0x8},
		/* ) */ tinyfont.RuneToIndex{Rune:41, Index:0x9},
		/* * */ tinyfont.RuneToIndex{Rune:42, Index:0xa},
		/* + */ tinyfont.RuneToIndex{Rune:43, Index:0xb},
		/* , */ tinyfont.RuneToIndex{Rune:44, Index:0xc},
		/* - */ tinyfont.RuneToIndex{Rune:45, Index:0xd},
		/* . */ tinyfont.RuneToIndex{Rune:46, Index:0xe},
		/* / */ tinyfont.RuneToIndex{Rune:47, Index:0xf},
		/* 0 */ tinyfont.RuneToIndex{Rune:48, Index:0x10},
		/* 1 */ tinyfont.RuneToIndex{Rune:49, Index:0x11},
		/* 2 */ tinyfont.RuneToIndex{Rune:50, Index:0x12},
		/* 3 */ tinyfont.RuneToIndex{Rune:51, Index:0x13},
		/* 4 */ tinyfont.RuneToIndex{Rune:52, Index:0x14},
		/* 5 */ tinyfont.RuneToIndex{Rune:53, Index:0x15},
		/* 6 */ tinyfont.RuneToIndex{Rune:54, Index:0x16},
		/* 7 */ tinyfont.RuneToIndex{Rune:55, Index:0x17},
		/* 8 */ tinyfont.RuneToIndex{Rune:56, Index:0x18},
		/* 9 */ tinyfont.RuneToIndex{Rune:57, Index:0x19},
		/* : */ tinyfont.RuneToIndex{Rune:58, Index:0x1a},
		/* ; */ tinyfont.RuneToIndex{Rune:59, Index:0x1b},
		/* < */ tinyfont.RuneToIndex{Rune:60, Index:0x1c},
		/* = */ tinyfont.RuneToIndex{Rune:61, Index:0x1d},
		/* > */ tinyfont.RuneToIndex{Rune:62, Index:0x1e},
		/* ? */ tinyfont.RuneToIndex{Rune:63, Index:0x1f},
		/* @ */ tinyfont.RuneToIndex{Rune:64, Index:0x20},
		/* A */ tinyfont.RuneToIndex{Rune:65, Index:0x21},
		/* B */ tinyfont.RuneToIndex{Rune:66, Index:0x22},
		/* C */ tinyfont.RuneToIndex{Rune:67, Index:0x23},
		/* D */ tinyfont.RuneToIndex{Rune:68, Index:0x24},
		/* E */ tinyfont.RuneToIndex{Rune:69, Index:0x25},
		/* F */ tinyfont.RuneToIndex{Rune:70, Index:0x26},
		/* G */ tinyfont.RuneToIndex{Rune:71, Index:0x27},
		/* H */ tinyfont.RuneToIndex{Rune:72, Index:0x28},
		/* I */ tinyfont.RuneToIndex{Rune:73, Index:0x29},
		/* J */ tinyfont.RuneToIndex{Rune:74, Index:0x2a},
		/* K */ tinyfont.RuneToIndex{Rune:75, Index:0x2b},
		/* L */ tinyfont.RuneToIndex{Rune:76, Index:0x2c},
		/* M */ tinyfont.RuneToIndex{Rune:77, Index:0x2d},
		/* N */ tinyfont.RuneToIndex{Rune:78, Index:0x2e},
		/* O */ tinyfont.RuneToIndex{Rune:79, Index:0x2f},
		/* P */ tinyfont.RuneToIndex{Rune:80, Index:0x30},
		/* Q */ tinyfont.RuneToIndex{Rune:81, Index:0x31},
		/* R */ tinyfont.RuneToIndex{Rune:82, Index:0x32},
		/* S */ tinyfont.RuneToIndex{Rune:83, Index:0x33},
		/* T */ tinyfont.RuneToIndex{Rune:84, Index:0x34},
		/* U */ tinyfont.RuneToIndex{Rune:85, Index:0x35},
		/* V */ tinyfont.RuneToIndex{Rune:86, Index:0x36},
		/* W */ tinyfont.RuneToIndex{Rune:87, Index:0x37},
		/* X */ tinyfont.RuneToIndex{Rune:88, Index:0x38},
		/* Y */ tinyfont.RuneToIndex{Rune:89, Index:0x39},
		/* Z */ tinyfont.RuneToIndex{Rune:90, Index:0x3a},
		/* [ */ tinyfont.RuneToIndex{Rune:91, Index:0x3b},
		/* \ */ tinyfont.RuneToIndex{Rune:92, Index:0x3c},
		/* ] */ tinyfont.RuneToIndex{Rune:93, Index:0x3d},
		/* ^ */ tinyfont.RuneToIndex{Rune:94, Index:0x3e},
		/* _ */ tinyfont.RuneToIndex{Rune:95, Index:0x3f},
		/* ` */ tinyfont.RuneToIndex{Rune:96, Index:0x40},
		/* a */ tinyfont.RuneToIndex{Rune:97, Index:0x41},
		/* b */ tinyfont.RuneToIndex{Rune:98, Index:0x42},
		/* c */ tinyfont.RuneToIndex{Rune:99, Index:0x43},
		/* d */ tinyfont.RuneToIndex{Rune:100, Index:0x44},
		/* e */ tinyfont.RuneToIndex{Rune:101, Index:0x45},
		/* f */ tinyfont.RuneToIndex{Rune:102, Index:0x46},
		/* g */ tinyfont.RuneToIndex{Rune:103, Index:0x47},
		/* h */ tinyfont.RuneToIndex{Rune:104, Index:0x48},
		/* i */ tinyfont.RuneToIndex{Rune:105, Index:0x49},
		/* j */ tinyfont.RuneToIndex{Rune:106, Index:0x4a},
		/* k */ tinyfont.RuneToIndex{Rune:107, Index:0x4b},
		/* l */ tinyfont.RuneToIndex{Rune:108, Index:0x4c},
		/* m */ tinyfont.RuneToIndex{Rune:109, Index:0x4d},
		/* n */ tinyfont.RuneToIndex{Rune:110, Index:0x4e},
		/* o */ tinyfont.RuneToIndex{Rune:111, Index:0x4f},
		/* p */ tinyfont.RuneToIndex{Rune:112, Index:0x50},
		/* q */ tinyfont.RuneToIndex{Rune:113, Index:0x51},
		/* r */ tinyfont.RuneToIndex{Rune:114, Index:0x52},
		/* s */ tinyfont.RuneToIndex{Rune:115, Index:0x53},
		/* t */ tinyfont.RuneToIndex{Rune:116, Index:0x54},
		/* u */ tinyfont.RuneToIndex{Rune:117, Index:0x55},
		/* v */ tinyfont.RuneToIndex{Rune:118, Index:0x56},
		/* w */ tinyfont.RuneToIndex{Rune:119, Index:0x57},
		/* x */ tinyfont.RuneToIndex{Rune:120, Index:0x58},
		/* y */ tinyfont.RuneToIndex{Rune:121, Index:0x59},
		/* z */ tinyfont.RuneToIndex{Rune:122, Index:0x5a},
		/* { */ tinyfont.RuneToIndex{Rune:123, Index:0x5b},
		/* | */ tinyfont.RuneToIndex{Rune:124, Index:0x5c},
		/* } */ tinyfont.RuneToIndex{Rune:125, Index:0x5d},
		/* ~ */ tinyfont.RuneToIndex{Rune:126, Index:0x5e},
	},

	YAdvance:0x1d,
}
