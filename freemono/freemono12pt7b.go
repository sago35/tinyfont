package freemono

import (
	"tinygo.org/x/tinyfont"
)

var Regular12pt7b = tinyfont.Font{
	Glyphs: []tinyfont.Glyph{
		/*   */ tinyfont.Glyph{Rune: 32, Width: 0x0, Height: 0x0, XAdvance: 0xe, XOffset: 0, YOffset: 1, Bitmaps: []uint8{}},
		/* ! */ tinyfont.Glyph{Rune: 33, Width: 0x3, Height: 0xf, XAdvance: 0xe, XOffset: 6, YOffset: -14, Bitmaps: []uint8{0x49, 0x24, 0x92, 0x48, 0x1, 0xf8}},
		/* " */ tinyfont.Glyph{Rune: 34, Width: 0x8, Height: 0x7, XAdvance: 0xe, XOffset: 3, YOffset: -14, Bitmaps: []uint8{0xe7, 0xe7, 0x67, 0x42, 0x42, 0x42, 0x42}},
		/* # */ tinyfont.Glyph{Rune: 35, Width: 0xa, Height: 0x10, XAdvance: 0xe, XOffset: 2, YOffset: -14, Bitmaps: []uint8{0x9, 0x2, 0x41, 0x10, 0x44, 0x11, 0x1f, 0xf1, 0x10, 0x4c, 0x12, 0x3f, 0xe1, 0x20, 0x48, 0x12, 0x4, 0x81, 0x20, 0x48}},
		/* $ */ tinyfont.Glyph{Rune: 36, Width: 0xa, Height: 0x11, XAdvance: 0xe, XOffset: 2, YOffset: -14, Bitmaps: []uint8{0x4, 0x7, 0xa2, 0x19, 0x2, 0x40, 0x10, 0x3, 0x0, 0x3c, 0x0, 0x80, 0x10, 0x6, 0x1, 0xe0, 0xa7, 0xc0, 0x40, 0x10, 0x4, 0x0}},
		/* % */ tinyfont.Glyph{Rune: 37, Width: 0xa, Height: 0xf, XAdvance: 0xe, XOffset: 2, YOffset: -14, Bitmaps: []uint8{0x3c, 0x19, 0x84, 0x21, 0x8, 0x66, 0xf, 0x0, 0xc, 0x1c, 0x78, 0x1, 0xe0, 0xcc, 0x21, 0x8, 0x43, 0x30, 0x78}},
		/* & */ tinyfont.Glyph{Rune: 38, Width: 0x9, Height: 0xc, XAdvance: 0xe, XOffset: 3, YOffset: -11, Bitmaps: []uint8{0x3e, 0x30, 0x10, 0x8, 0x2, 0x3, 0x3, 0x47, 0x14, 0x8a, 0x43, 0x11, 0x8f, 0x60}},
		/* ' */ tinyfont.Glyph{Rune: 39, Width: 0x3, Height: 0x7, XAdvance: 0xe, XOffset: 5, YOffset: -14, Bitmaps: []uint8{0xfd, 0xa4, 0x90}},
		/* ( */ tinyfont.Glyph{Rune: 40, Width: 0x3, Height: 0x12, XAdvance: 0xe, XOffset: 7, YOffset: -14, Bitmaps: []uint8{0x5, 0x25, 0x24, 0x92, 0x48, 0x92, 0x24}},
		/* ) */ tinyfont.Glyph{Rune: 41, Width: 0x3, Height: 0x12, XAdvance: 0xe, XOffset: 4, YOffset: -14, Bitmaps: []uint8{0x11, 0x24, 0x89, 0x24, 0x92, 0x92, 0x90}},
		/* * */ tinyfont.Glyph{Rune: 42, Width: 0x9, Height: 0x9, XAdvance: 0xe, XOffset: 3, YOffset: -14, Bitmaps: []uint8{0x0, 0x4, 0x2, 0x11, 0x7, 0xf0, 0xc0, 0x50, 0x48, 0x42, 0x0}},
		/* + */ tinyfont.Glyph{Rune: 43, Width: 0x9, Height: 0xb, XAdvance: 0xe, XOffset: 3, YOffset: -11, Bitmaps: []uint8{0x8, 0x4, 0x2, 0x1, 0x0, 0x87, 0xfc, 0x20, 0x10, 0x8, 0x4, 0x2, 0x0}},
		/* , */ tinyfont.Glyph{Rune: 44, Width: 0x5, Height: 0x7, XAdvance: 0xe, XOffset: 3, YOffset: -3, Bitmaps: []uint8{0x3b, 0x9c, 0xce, 0x62, 0x0}},
		/* - */ tinyfont.Glyph{Rune: 45, Width: 0xb, Height: 0x1, XAdvance: 0xe, XOffset: 2, YOffset: -6, Bitmaps: []uint8{0xff, 0xe0}},
		/* . */ tinyfont.Glyph{Rune: 46, Width: 0x3, Height: 0x3, XAdvance: 0xe, XOffset: 5, YOffset: -2, Bitmaps: []uint8{0xff, 0x80}},
		/* / */ tinyfont.Glyph{Rune: 47, Width: 0x9, Height: 0x12, XAdvance: 0xe, XOffset: 3, YOffset: -15, Bitmaps: []uint8{0x0, 0x80, 0xc0, 0x40, 0x20, 0x20, 0x10, 0x10, 0x8, 0x8, 0x4, 0x4, 0x2, 0x2, 0x1, 0x1, 0x0, 0x80, 0x80, 0x40, 0x0}},
		/* 0 */ tinyfont.Glyph{Rune: 48, Width: 0x9, Height: 0xf, XAdvance: 0xe, XOffset: 3, YOffset: -14, Bitmaps: []uint8{0x1c, 0x31, 0x90, 0x58, 0x38, 0xc, 0x6, 0x3, 0x1, 0x80, 0xc0, 0x60, 0x30, 0x34, 0x13, 0x18, 0x70}},
		/* 1 */ tinyfont.Glyph{Rune: 49, Width: 0x7, Height: 0xe, XAdvance: 0xe, XOffset: 4, YOffset: -13, Bitmaps: []uint8{0x30, 0xe1, 0x44, 0x81, 0x2, 0x4, 0x8, 0x10, 0x20, 0x40, 0x81, 0x1f, 0xc0}},
		/* 2 */ tinyfont.Glyph{Rune: 50, Width: 0x9, Height: 0xf, XAdvance: 0xe, XOffset: 2, YOffset: -14, Bitmaps: []uint8{0x1e, 0x10, 0x90, 0x68, 0x10, 0x8, 0xc, 0x4, 0x4, 0x4, 0x6, 0x6, 0x6, 0x6, 0xe, 0x7, 0xfe}},
		/* 3 */ tinyfont.Glyph{Rune: 51, Width: 0xa, Height: 0xf, XAdvance: 0xe, XOffset: 2, YOffset: -14, Bitmaps: []uint8{0x3e, 0x10, 0x40, 0x8, 0x2, 0x0, 0x80, 0x40, 0xe0, 0x4, 0x0, 0x80, 0x10, 0x4, 0x1, 0x0, 0xd8, 0x63, 0xe0}},
		/* 4 */ tinyfont.Glyph{Rune: 52, Width: 0x8, Height: 0xf, XAdvance: 0xe, XOffset: 3, YOffset: -14, Bitmaps: []uint8{0x6, 0xa, 0xa, 0x12, 0x22, 0x22, 0x42, 0x42, 0x82, 0x82, 0xff, 0x2, 0x2, 0x2, 0xf}},
		/* 5 */ tinyfont.Glyph{Rune: 53, Width: 0x9, Height: 0xf, XAdvance: 0xe, XOffset: 3, YOffset: -14, Bitmaps: []uint8{0x7f, 0x20, 0x10, 0x8, 0x4, 0x2, 0xf1, 0x8c, 0x3, 0x0, 0x80, 0x40, 0x20, 0x18, 0x16, 0x18, 0xf0}},
		/* 6 */ tinyfont.Glyph{Rune: 54, Width: 0x9, Height: 0xf, XAdvance: 0xe, XOffset: 3, YOffset: -14, Bitmaps: []uint8{0xf, 0x8c, 0x8, 0x8, 0x4, 0x4, 0x2, 0x79, 0x46, 0xc1, 0xe0, 0x60, 0x28, 0x14, 0x19, 0x8, 0x78}},
		/* 7 */ tinyfont.Glyph{Rune: 55, Width: 0x8, Height: 0xf, XAdvance: 0xe, XOffset: 3, YOffset: -14, Bitmaps: []uint8{0xff, 0x81, 0x81, 0x2, 0x2, 0x2, 0x2, 0x4, 0x4, 0x4, 0x4, 0x8, 0x8, 0x8, 0x8}},
		/* 8 */ tinyfont.Glyph{Rune: 56, Width: 0x9, Height: 0xf, XAdvance: 0xe, XOffset: 3, YOffset: -14, Bitmaps: []uint8{0x3e, 0x31, 0xb0, 0x70, 0x18, 0xc, 0x5, 0x8c, 0x38, 0x63, 0x40, 0x60, 0x30, 0x18, 0x1b, 0x18, 0xf8}},
		/* 9 */ tinyfont.Glyph{Rune: 57, Width: 0x9, Height: 0xf, XAdvance: 0xe, XOffset: 3, YOffset: -14, Bitmaps: []uint8{0x3c, 0x31, 0x30, 0x50, 0x28, 0xc, 0xf, 0x6, 0x85, 0x3c, 0x80, 0x40, 0x40, 0x20, 0x20, 0x63, 0xe0}},
		/* : */ tinyfont.Glyph{Rune: 58, Width: 0x3, Height: 0xa, XAdvance: 0xe, XOffset: 5, YOffset: -9, Bitmaps: []uint8{0xff, 0x80, 0x7, 0xfc}},
		/* ; */ tinyfont.Glyph{Rune: 59, Width: 0x5, Height: 0xd, XAdvance: 0xe, XOffset: 3, YOffset: -9, Bitmaps: []uint8{0x39, 0xce, 0x0, 0x0, 0x6, 0x33, 0x98, 0xc4, 0x0}},
		/* < */ tinyfont.Glyph{Rune: 60, Width: 0xb, Height: 0xb, XAdvance: 0xe, XOffset: 2, YOffset: -11, Bitmaps: []uint8{0x0, 0xc0, 0x60, 0x18, 0xc, 0x6, 0x1, 0x80, 0xc, 0x0, 0x60, 0x3, 0x0, 0x30, 0x1, 0x0}},
		/* = */ tinyfont.Glyph{Rune: 61, Width: 0xc, Height: 0x4, XAdvance: 0xe, XOffset: 1, YOffset: -8, Bitmaps: []uint8{0xff, 0xf0, 0x0, 0x0, 0xf, 0xff}},
		/* > */ tinyfont.Glyph{Rune: 62, Width: 0xb, Height: 0xb, XAdvance: 0xe, XOffset: 2, YOffset: -11, Bitmaps: []uint8{0xc0, 0x6, 0x0, 0x30, 0x1, 0x80, 0x18, 0x1, 0x80, 0xc0, 0x30, 0x18, 0xc, 0x2, 0x0, 0x0}},
		/* ? */ tinyfont.Glyph{Rune: 63, Width: 0x9, Height: 0xe, XAdvance: 0xe, XOffset: 3, YOffset: -13, Bitmaps: []uint8{0x3e, 0x60, 0xa0, 0x20, 0x10, 0x8, 0x8, 0x18, 0x10, 0x8, 0x0, 0x0, 0x0, 0x1, 0xc0, 0xe0}},
		/* @ */ tinyfont.Glyph{Rune: 64, Width: 0x9, Height: 0x10, XAdvance: 0xe, XOffset: 3, YOffset: -14, Bitmaps: []uint8{0x1c, 0x31, 0x10, 0x50, 0x28, 0x14, 0x3a, 0x25, 0x22, 0x91, 0x4c, 0xa3, 0xf0, 0x8, 0x2, 0x1, 0x80, 0x7c}},
		/* A */ tinyfont.Glyph{Rune: 65, Width: 0xe, Height: 0xe, XAdvance: 0xe, XOffset: 0, YOffset: -13, Bitmaps: []uint8{0x3f, 0x0, 0xc, 0x0, 0x48, 0x1, 0x20, 0x4, 0x40, 0x21, 0x0, 0x84, 0x4, 0x8, 0x1f, 0xe0, 0x40, 0x82, 0x1, 0x8, 0x4, 0x20, 0x13, 0xe1, 0xf0}},
		/* B */ tinyfont.Glyph{Rune: 66, Width: 0xb, Height: 0xe, XAdvance: 0xe, XOffset: 2, YOffset: -13, Bitmaps: []uint8{0xff, 0x8, 0x11, 0x1, 0x20, 0x24, 0x4, 0x81, 0x1f, 0xc2, 0x6, 0x40, 0x68, 0x5, 0x0, 0xa0, 0x14, 0x5, 0xff, 0x0}},
		/* C */ tinyfont.Glyph{Rune: 67, Width: 0xa, Height: 0xe, XAdvance: 0xe, XOffset: 2, YOffset: -13, Bitmaps: []uint8{0x1e, 0x48, 0x74, 0x5, 0x1, 0x80, 0x20, 0x8, 0x2, 0x0, 0x80, 0x20, 0x4, 0x1, 0x1, 0x30, 0x87, 0xc0}},
		/* D */ tinyfont.Glyph{Rune: 68, Width: 0xa, Height: 0xe, XAdvance: 0xe, XOffset: 2, YOffset: -13, Bitmaps: []uint8{0xfe, 0x10, 0x44, 0x9, 0x2, 0x40, 0x50, 0x14, 0x5, 0x1, 0x40, 0x50, 0x14, 0xd, 0x2, 0x41, 0x3f, 0x80}},
		/* E */ tinyfont.Glyph{Rune: 69, Width: 0xb, Height: 0xe, XAdvance: 0xe, XOffset: 2, YOffset: -13, Bitmaps: []uint8{0xff, 0xc8, 0x9, 0x1, 0x20, 0x4, 0x0, 0x88, 0x1f, 0x2, 0x20, 0x40, 0x8, 0x1, 0x0, 0xa0, 0x14, 0x3, 0xff, 0xc0}},
		/* F */ tinyfont.Glyph{Rune: 70, Width: 0xb, Height: 0xe, XAdvance: 0xe, XOffset: 2, YOffset: -13, Bitmaps: []uint8{0xff, 0xe8, 0x5, 0x0, 0xa0, 0x4, 0x0, 0x88, 0x1f, 0x2, 0x20, 0x40, 0x8, 0x1, 0x0, 0x20, 0x4, 0x1, 0xf0, 0x0}},
		/* G */ tinyfont.Glyph{Rune: 71, Width: 0xb, Height: 0xe, XAdvance: 0xe, XOffset: 2, YOffset: -13, Bitmaps: []uint8{0x1f, 0x46, 0x19, 0x1, 0x60, 0x28, 0x1, 0x0, 0x20, 0x4, 0x0, 0x83, 0xf0, 0xb, 0x1, 0x20, 0x23, 0xc, 0x3e, 0x0}},
		/* H */ tinyfont.Glyph{Rune: 72, Width: 0xa, Height: 0xe, XAdvance: 0xe, XOffset: 2, YOffset: -13, Bitmaps: []uint8{0xe1, 0xd0, 0x24, 0x9, 0x2, 0x40, 0x90, 0x27, 0xf9, 0x2, 0x40, 0x90, 0x24, 0x9, 0x2, 0x40, 0xb8, 0x70}},
		/* I */ tinyfont.Glyph{Rune: 73, Width: 0x7, Height: 0xe, XAdvance: 0xe, XOffset: 4, YOffset: -13, Bitmaps: []uint8{0xfe, 0x20, 0x40, 0x81, 0x2, 0x4, 0x8, 0x10, 0x20, 0x40, 0x81, 0x1f, 0xc0}},
		/* J */ tinyfont.Glyph{Rune: 74, Width: 0xb, Height: 0xe, XAdvance: 0xe, XOffset: 2, YOffset: -13, Bitmaps: []uint8{0xf, 0xe0, 0x10, 0x2, 0x0, 0x40, 0x8, 0x1, 0x0, 0x20, 0x4, 0x80, 0x90, 0x12, 0x2, 0x40, 0xc6, 0x30, 0x7c, 0x0}},
		/* K */ tinyfont.Glyph{Rune: 75, Width: 0xc, Height: 0xe, XAdvance: 0xe, XOffset: 2, YOffset: -13, Bitmaps: []uint8{0xf1, 0xe4, 0xc, 0x41, 0x4, 0x20, 0x44, 0x4, 0x80, 0x5c, 0x6, 0x60, 0x43, 0x4, 0x10, 0x40, 0x84, 0x8, 0x40, 0xcf, 0x7}},
		/* L */ tinyfont.Glyph{Rune: 76, Width: 0xb, Height: 0xe, XAdvance: 0xe, XOffset: 2, YOffset: -13, Bitmaps: []uint8{0xf8, 0x4, 0x0, 0x80, 0x10, 0x2, 0x0, 0x40, 0x8, 0x1, 0x0, 0x20, 0x4, 0x4, 0x80, 0x90, 0x12, 0x3, 0xff, 0xc0}},
		/* M */ tinyfont.Glyph{Rune: 77, Width: 0xd, Height: 0xe, XAdvance: 0xe, XOffset: 1, YOffset: -13, Bitmaps: []uint8{0xe0, 0x3b, 0x1, 0x94, 0x14, 0xa0, 0xa4, 0x89, 0x24, 0x49, 0x14, 0x48, 0xa2, 0x45, 0x12, 0x10, 0x90, 0x4, 0x80, 0x24, 0x1, 0x78, 0x3c}},
		/* N */ tinyfont.Glyph{Rune: 78, Width: 0xc, Height: 0xe, XAdvance: 0xe, XOffset: 1, YOffset: -13, Bitmaps: []uint8{0xe0, 0xf6, 0x2, 0x50, 0x25, 0x2, 0x48, 0x24, 0xc2, 0x44, 0x24, 0x22, 0x43, 0x24, 0x12, 0x40, 0xa4, 0xa, 0x40, 0x6f, 0x6}},
		/* O */ tinyfont.Glyph{Rune: 79, Width: 0xc, Height: 0xe, XAdvance: 0xe, XOffset: 1, YOffset: -13, Bitmaps: []uint8{0xf, 0x3, 0xc, 0x60, 0x64, 0x2, 0x80, 0x18, 0x1, 0x80, 0x18, 0x1, 0x80, 0x18, 0x1, 0x40, 0x26, 0x6, 0x30, 0xc0, 0xf0}},
		/* P */ tinyfont.Glyph{Rune: 80, Width: 0xa, Height: 0xe, XAdvance: 0xe, XOffset: 2, YOffset: -13, Bitmaps: []uint8{0xff, 0x10, 0x64, 0x5, 0x1, 0x40, 0x50, 0x34, 0x19, 0xfc, 0x40, 0x10, 0x4, 0x1, 0x0, 0x40, 0x3e, 0x0}},
		/* Q */ tinyfont.Glyph{Rune: 81, Width: 0xc, Height: 0x11, XAdvance: 0xe, XOffset: 1, YOffset: -13, Bitmaps: []uint8{0xf, 0x3, 0xc, 0x60, 0x64, 0x2, 0x80, 0x18, 0x1, 0x80, 0x18, 0x1, 0x80, 0x18, 0x1, 0x40, 0x26, 0x6, 0x30, 0xc1, 0xf0, 0xc, 0x1, 0xf1, 0x30, 0xe0}},
		/* R */ tinyfont.Glyph{Rune: 82, Width: 0xc, Height: 0xe, XAdvance: 0xe, XOffset: 2, YOffset: -13, Bitmaps: []uint8{0xff, 0x4, 0x18, 0x40, 0xc4, 0x4, 0x40, 0x44, 0xc, 0x41, 0x87, 0xe0, 0x43, 0x4, 0x10, 0x40, 0x84, 0x4, 0x40, 0x4f, 0x3}},
		/* S */ tinyfont.Glyph{Rune: 83, Width: 0xa, Height: 0xe, XAdvance: 0xe, XOffset: 2, YOffset: -13, Bitmaps: []uint8{0x1f, 0x48, 0x34, 0x5, 0x1, 0x40, 0x8, 0x1, 0xc0, 0xe, 0x0, 0x40, 0x18, 0x6, 0x1, 0xe1, 0xa7, 0xc0}},
		/* T */ tinyfont.Glyph{Rune: 84, Width: 0xb, Height: 0xe, XAdvance: 0xe, XOffset: 2, YOffset: -13, Bitmaps: []uint8{0xff, 0xf0, 0x86, 0x10, 0x82, 0x0, 0x40, 0x8, 0x1, 0x0, 0x20, 0x4, 0x0, 0x80, 0x10, 0x2, 0x0, 0x40, 0x7f, 0x0}},
		/* U */ tinyfont.Glyph{Rune: 85, Width: 0xc, Height: 0xe, XAdvance: 0xe, XOffset: 1, YOffset: -13, Bitmaps: []uint8{0xf0, 0xf4, 0x2, 0x40, 0x24, 0x2, 0x40, 0x24, 0x2, 0x40, 0x24, 0x2, 0x40, 0x24, 0x2, 0x40, 0x22, 0x4, 0x30, 0xc0, 0xf0}},
		/* V */ tinyfont.Glyph{Rune: 86, Width: 0xe, Height: 0xe, XAdvance: 0xe, XOffset: 0, YOffset: -13, Bitmaps: []uint8{0xf8, 0x7c, 0x80, 0x22, 0x1, 0x4, 0x4, 0x10, 0x20, 0x40, 0x80, 0x82, 0x2, 0x10, 0x8, 0x40, 0x11, 0x0, 0x48, 0x1, 0xa0, 0x3, 0x0, 0xc, 0x0}},
		/* W */ tinyfont.Glyph{Rune: 87, Width: 0xe, Height: 0xe, XAdvance: 0xe, XOffset: 0, YOffset: -13, Bitmaps: []uint8{0xf8, 0x7c, 0x80, 0x22, 0x0, 0x88, 0xc2, 0x23, 0x10, 0x8e, 0x42, 0x29, 0x9, 0x24, 0x24, 0x90, 0x91, 0x41, 0x85, 0x6, 0x14, 0x18, 0x70, 0x60, 0x80}},
		/* X */ tinyfont.Glyph{Rune: 88, Width: 0xc, Height: 0xe, XAdvance: 0xe, XOffset: 1, YOffset: -13, Bitmaps: []uint8{0xf0, 0xf2, 0x6, 0x30, 0x41, 0x8, 0x9, 0x80, 0x50, 0x6, 0x0, 0x60, 0xd, 0x0, 0x88, 0x10, 0xc2, 0x4, 0x60, 0x2f, 0xf}},
		/* Y */ tinyfont.Glyph{Rune: 89, Width: 0xc, Height: 0xe, XAdvance: 0xe, XOffset: 1, YOffset: -13, Bitmaps: []uint8{0xf0, 0xf2, 0x2, 0x10, 0x41, 0x4, 0x8, 0x80, 0x50, 0x5, 0x0, 0x20, 0x2, 0x0, 0x20, 0x2, 0x0, 0x20, 0x2, 0x1, 0xfc}},
		/* Z */ tinyfont.Glyph{Rune: 90, Width: 0x9, Height: 0xe, XAdvance: 0xe, XOffset: 3, YOffset: -13, Bitmaps: []uint8{0xff, 0x40, 0xa0, 0x90, 0x40, 0x40, 0x40, 0x20, 0x20, 0x20, 0x10, 0x50, 0x30, 0x18, 0xf, 0xfc}},
		/* [ */ tinyfont.Glyph{Rune: 91, Width: 0x3, Height: 0x12, XAdvance: 0xe, XOffset: 7, YOffset: -14, Bitmaps: []uint8{0xf2, 0x49, 0x24, 0x92, 0x49, 0x24, 0x9c}},
		/* \ */ tinyfont.Glyph{Rune: 92, Width: 0x9, Height: 0x12, XAdvance: 0xe, XOffset: 3, YOffset: -15, Bitmaps: []uint8{0x80, 0x60, 0x10, 0x8, 0x2, 0x1, 0x0, 0x40, 0x20, 0x8, 0x4, 0x1, 0x0, 0x80, 0x20, 0x10, 0x4, 0x2, 0x0, 0x80, 0x40}},
		/* ] */ tinyfont.Glyph{Rune: 93, Width: 0x3, Height: 0x12, XAdvance: 0xe, XOffset: 5, YOffset: -14, Bitmaps: []uint8{0xe4, 0x92, 0x49, 0x24, 0x92, 0x49, 0x3c}},
		/* ^ */ tinyfont.Glyph{Rune: 94, Width: 0x9, Height: 0x6, XAdvance: 0xe, XOffset: 3, YOffset: -14, Bitmaps: []uint8{0x8, 0xc, 0x9, 0xc, 0x4c, 0x14, 0x4}},
		/* _ */ tinyfont.Glyph{Rune: 95, Width: 0xe, Height: 0x1, XAdvance: 0xe, XOffset: 0, YOffset: 3, Bitmaps: []uint8{0xff, 0xfc}},
		/* ` */ tinyfont.Glyph{Rune: 96, Width: 0x4, Height: 0x4, XAdvance: 0xe, XOffset: 4, YOffset: -15, Bitmaps: []uint8{0x84, 0x21}},
		/* a */ tinyfont.Glyph{Rune: 97, Width: 0xa, Height: 0xa, XAdvance: 0xe, XOffset: 2, YOffset: -9, Bitmaps: []uint8{0x3e, 0x0, 0x60, 0x8, 0x2, 0x3f, 0x98, 0x28, 0xa, 0x2, 0xc3, 0x9f, 0x30}},
		/* b */ tinyfont.Glyph{Rune: 98, Width: 0xd, Height: 0xf, XAdvance: 0xe, XOffset: 0, YOffset: -14, Bitmaps: []uint8{0xe0, 0x1, 0x0, 0x8, 0x0, 0x40, 0x2, 0x0, 0x13, 0xe0, 0xa0, 0x86, 0x2, 0x20, 0x9, 0x0, 0x48, 0x2, 0x40, 0x13, 0x1, 0x14, 0x1b, 0x9f, 0x0}},
		/* c */ tinyfont.Glyph{Rune: 99, Width: 0xb, Height: 0xa, XAdvance: 0xe, XOffset: 2, YOffset: -9, Bitmaps: []uint8{0x1f, 0x4c, 0x19, 0x1, 0x40, 0x28, 0x1, 0x0, 0x20, 0x2, 0x0, 0x60, 0x43, 0xf0}},
		/* d */ tinyfont.Glyph{Rune: 100, Width: 0xb, Height: 0xf, XAdvance: 0xe, XOffset: 2, YOffset: -14, Bitmaps: []uint8{0x0, 0xc0, 0x8, 0x1, 0x0, 0x20, 0x4, 0x3c, 0x98, 0x52, 0x6, 0x80, 0x50, 0xa, 0x1, 0x40, 0x24, 0xc, 0xc2, 0x87, 0x98}},
		/* e */ tinyfont.Glyph{Rune: 101, Width: 0xa, Height: 0xa, XAdvance: 0xe, XOffset: 2, YOffset: -9, Bitmaps: []uint8{0x3f, 0x18, 0x68, 0x6, 0x1, 0xff, 0xe0, 0x8, 0x3, 0x0, 0x60, 0xc7, 0xc0}},
		/* f */ tinyfont.Glyph{Rune: 102, Width: 0x9, Height: 0xf, XAdvance: 0xe, XOffset: 4, YOffset: -14, Bitmaps: []uint8{0xf, 0x98, 0x8, 0x4, 0x2, 0x7, 0xf8, 0x80, 0x40, 0x20, 0x10, 0x8, 0x4, 0x2, 0x1, 0x3, 0xf8}},
		/* g */ tinyfont.Glyph{Rune: 103, Width: 0xb, Height: 0xe, XAdvance: 0xe, XOffset: 2, YOffset: -9, Bitmaps: []uint8{0x1e, 0x6c, 0x39, 0x3, 0x40, 0x28, 0x5, 0x0, 0xa0, 0x12, 0x6, 0x61, 0x43, 0xc8, 0x1, 0x0, 0x20, 0x8, 0x3e, 0x0}},
		/* h */ tinyfont.Glyph{Rune: 104, Width: 0xa, Height: 0xf, XAdvance: 0xe, XOffset: 2, YOffset: -14, Bitmaps: []uint8{0xc0, 0x10, 0x4, 0x1, 0x0, 0x40, 0x13, 0x87, 0x11, 0x82, 0x40, 0x90, 0x24, 0x9, 0x2, 0x40, 0x90, 0x2e, 0x1c}},
		/* i */ tinyfont.Glyph{Rune: 105, Width: 0x9, Height: 0xf, XAdvance: 0xe, XOffset: 3, YOffset: -14, Bitmaps: []uint8{0x8, 0x4, 0x2, 0x0, 0x0, 0x3, 0xc0, 0x20, 0x10, 0x8, 0x4, 0x2, 0x1, 0x0, 0x80, 0x43, 0xfe}},
		/* j */ tinyfont.Glyph{Rune: 106, Width: 0x7, Height: 0x13, XAdvance: 0xe, XOffset: 3, YOffset: -14, Bitmaps: []uint8{0x4, 0x8, 0x10, 0x0, 0x1f, 0xc0, 0x81, 0x2, 0x4, 0x8, 0x10, 0x20, 0x40, 0x81, 0x2, 0xb, 0xe0}},
		/* k */ tinyfont.Glyph{Rune: 107, Width: 0xc, Height: 0xf, XAdvance: 0xe, XOffset: 1, YOffset: -14, Bitmaps: []uint8{0xe0, 0x2, 0x0, 0x20, 0x2, 0x0, 0x20, 0x2, 0x3c, 0x21, 0x2, 0x60, 0x2c, 0x3, 0x80, 0x24, 0x2, 0x20, 0x21, 0x2, 0x8, 0xe1, 0xf0}},
		/* l */ tinyfont.Glyph{Rune: 108, Width: 0x9, Height: 0xf, XAdvance: 0xe, XOffset: 3, YOffset: -14, Bitmaps: []uint8{0x78, 0x4, 0x2, 0x1, 0x0, 0x80, 0x40, 0x20, 0x10, 0x8, 0x4, 0x2, 0x1, 0x0, 0x80, 0x43, 0xfe}},
		/* m */ tinyfont.Glyph{Rune: 109, Width: 0xd, Height: 0xa, XAdvance: 0xe, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0xdc, 0xe3, 0x19, 0x90, 0x84, 0x84, 0x24, 0x21, 0x21, 0x9, 0x8, 0x48, 0x42, 0x42, 0x17, 0x18, 0xc0}},
		/* n */ tinyfont.Glyph{Rune: 110, Width: 0xc, Height: 0xa, XAdvance: 0xe, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0x67, 0x83, 0x84, 0x20, 0x22, 0x2, 0x20, 0x22, 0x2, 0x20, 0x22, 0x2, 0x20, 0x2f, 0x7}},
		/* o */ tinyfont.Glyph{Rune: 111, Width: 0xb, Height: 0xa, XAdvance: 0xe, XOffset: 2, YOffset: -9, Bitmaps: []uint8{0x1f, 0x4, 0x11, 0x1, 0x40, 0x18, 0x3, 0x0, 0x60, 0xa, 0x2, 0x20, 0x83, 0xe0}},
		/* p */ tinyfont.Glyph{Rune: 112, Width: 0xc, Height: 0xe, XAdvance: 0xe, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0xcf, 0x85, 0x6, 0x60, 0x24, 0x1, 0x40, 0x14, 0x1, 0x40, 0x16, 0x2, 0x50, 0x44, 0xf8, 0x40, 0x4, 0x0, 0x40, 0xf, 0x0}},
		/* q */ tinyfont.Glyph{Rune: 113, Width: 0xb, Height: 0xe, XAdvance: 0xe, XOffset: 2, YOffset: -9, Bitmaps: []uint8{0x1e, 0x6c, 0x3b, 0x3, 0x40, 0x28, 0x5, 0x0, 0xa0, 0x12, 0x6, 0x61, 0x43, 0xc8, 0x1, 0x0, 0x20, 0x4, 0x3, 0xc0}},
		/* r */ tinyfont.Glyph{Rune: 114, Width: 0xa, Height: 0xa, XAdvance: 0xe, XOffset: 3, YOffset: -9, Bitmaps: []uint8{0xe3, 0x8b, 0x13, 0x80, 0x80, 0x20, 0x8, 0x2, 0x0, 0x80, 0x20, 0x3f, 0x80}},
		/* s */ tinyfont.Glyph{Rune: 115, Width: 0xa, Height: 0xa, XAdvance: 0xe, XOffset: 2, YOffset: -9, Bitmaps: []uint8{0x1f, 0x58, 0x34, 0x5, 0x80, 0x1e, 0x0, 0x60, 0x6, 0x1, 0xc0, 0xaf, 0xc0}},
		/* t */ tinyfont.Glyph{Rune: 116, Width: 0xb, Height: 0xe, XAdvance: 0xe, XOffset: 1, YOffset: -13, Bitmaps: []uint8{0x20, 0x4, 0x0, 0x80, 0x10, 0xf, 0xf0, 0x40, 0x8, 0x1, 0x0, 0x20, 0x4, 0x0, 0x80, 0x10, 0x3, 0x4, 0x3f, 0x0}},
		/* u */ tinyfont.Glyph{Rune: 117, Width: 0xb, Height: 0xa, XAdvance: 0xe, XOffset: 2, YOffset: -9, Bitmaps: []uint8{0xc1, 0xc8, 0x9, 0x1, 0x20, 0x24, 0x4, 0x80, 0x90, 0x12, 0x2, 0x61, 0xc7, 0xcc}},
		/* v */ tinyfont.Glyph{Rune: 118, Width: 0xd, Height: 0xa, XAdvance: 0xe, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0xf8, 0xf9, 0x1, 0x8, 0x10, 0x60, 0x81, 0x8, 0x8, 0x40, 0x22, 0x1, 0x20, 0x5, 0x0, 0x30, 0x0}},
		/* w */ tinyfont.Glyph{Rune: 119, Width: 0xd, Height: 0xa, XAdvance: 0xe, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0xf0, 0x7a, 0x1, 0x10, 0x8, 0x8c, 0x42, 0x62, 0x12, 0x90, 0xa5, 0x5, 0x18, 0x28, 0xc0, 0x86, 0x0}},
		/* x */ tinyfont.Glyph{Rune: 120, Width: 0xc, Height: 0xa, XAdvance: 0xe, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0x78, 0xf3, 0x4, 0x18, 0x80, 0xd0, 0x6, 0x0, 0x70, 0x9, 0x81, 0xc, 0x20, 0x6f, 0x8f}},
		/* y */ tinyfont.Glyph{Rune: 121, Width: 0xc, Height: 0xe, XAdvance: 0xe, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0xf0, 0xf2, 0x2, 0x20, 0x41, 0x4, 0x10, 0x80, 0x88, 0x9, 0x0, 0x50, 0x6, 0x0, 0x20, 0x4, 0x0, 0x40, 0x8, 0xf, 0xe0}},
		/* z */ tinyfont.Glyph{Rune: 122, Width: 0x9, Height: 0xa, XAdvance: 0xe, XOffset: 3, YOffset: -9, Bitmaps: []uint8{0xff, 0x41, 0x0, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x40, 0xbf, 0xc0}},
		/* { */ tinyfont.Glyph{Rune: 123, Width: 0x5, Height: 0x12, XAdvance: 0xe, XOffset: 5, YOffset: -14, Bitmaps: []uint8{0x19, 0x8, 0x42, 0x10, 0x84, 0x64, 0x18, 0x42, 0x10, 0x84, 0x20, 0xc0}},
		/* | */ tinyfont.Glyph{Rune: 124, Width: 0x1, Height: 0x12, XAdvance: 0xe, XOffset: 7, YOffset: -14, Bitmaps: []uint8{0xff, 0xff, 0xc0}},
		/* } */ tinyfont.Glyph{Rune: 125, Width: 0x5, Height: 0x12, XAdvance: 0xe, XOffset: 5, YOffset: -14, Bitmaps: []uint8{0xc1, 0x8, 0x42, 0x10, 0x84, 0x10, 0x4c, 0x42, 0x10, 0x84, 0x26, 0x0}},
		/* ~ */ tinyfont.Glyph{Rune: 126, Width: 0xa, Height: 0x3, XAdvance: 0xe, XOffset: 2, YOffset: -7, Bitmaps: []uint8{0x38, 0x13, 0x38, 0x38}},
	},

	YAdvance: 0x18,
}
