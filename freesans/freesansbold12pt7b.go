package freesans

import (
	"tinygo.org/x/tinyfont"
)

var Bold12pt7b = tinyfont.Font{
	Glyphs:[]tinyfont.Glyph{
		/*   */ tinyfont.Glyph{Width:0x0, Height:0x0, XAdvance:0x7, XOffset:0, YOffset:1, Bitmaps:[]uint8{}},
		/* ! */ tinyfont.Glyph{Width:0x4, Height:0x11, XAdvance:0x8, XOffset:3, YOffset:-16, Bitmaps:[]uint8{0xff, 0xff, 0xff, 0xff, 0x76, 0x66, 0x60, 0xff, 0xf0}},
		/* " */ tinyfont.Glyph{Width:0xa, Height:0x6, XAdvance:0xb, XOffset:1, YOffset:-17, Bitmaps:[]uint8{0xf3, 0xfc, 0xff, 0x3f, 0xcf, 0x61, 0x98, 0x60}},
		/* # */ tinyfont.Glyph{Width:0xd, Height:0x10, XAdvance:0xd, XOffset:0, YOffset:-15, Bitmaps:[]uint8{0xe, 0x70, 0x73, 0x83, 0x18, 0xff, 0xf7, 0xff, 0xbf, 0xfc, 0x73, 0x83, 0x18, 0x18, 0xc7, 0xff, 0xbf, 0xfd, 0xff, 0xe3, 0x18, 0x39, 0xc1, 0xce, 0xe, 0x70}},
		/* $ */ tinyfont.Glyph{Width:0xd, Height:0x14, XAdvance:0xd, XOffset:0, YOffset:-17, Bitmaps:[]uint8{0x2, 0x0, 0x7e, 0xf, 0xf8, 0x7f, 0xe7, 0xaf, 0xb9, 0x3d, 0xc8, 0xf, 0x40, 0x3f, 0x0, 0xff, 0x0, 0xfc, 0x5, 0xff, 0x27, 0xf9, 0x3f, 0xeb, 0xef, 0xfe, 0x3f, 0xe0, 0x7c, 0x0, 0x80, 0x4, 0x0}},
		/* % */ tinyfont.Glyph{Width:0x13, Height:0x11, XAdvance:0x15, XOffset:1, YOffset:-16, Bitmaps:[]uint8{0x3c, 0x6, 0xf, 0xc1, 0x81, 0xfc, 0x30, 0x73, 0x8c, 0xc, 0x31, 0x81, 0xce, 0x60, 0x1f, 0xcc, 0x3, 0xf3, 0x0, 0x3c, 0x67, 0x80, 0x19, 0xf8, 0x2, 0x7f, 0x80, 0xce, 0x70, 0x11, 0x86, 0x6, 0x39, 0xc1, 0x87, 0xf8, 0x30, 0x7e, 0xc, 0x7, 0x80}},
		/* & */ tinyfont.Glyph{Width:0x10, Height:0x11, XAdvance:0x11, XOffset:1, YOffset:-16, Bitmaps:[]uint8{0x7, 0x80, 0x1f, 0xc0, 0x3f, 0xe0, 0x3c, 0xe0, 0x3c, 0xe0, 0x3e, 0xe0, 0xf, 0xc0, 0x7, 0x0, 0x3f, 0x8c, 0x7f, 0xcc, 0xf1, 0xfc, 0xf0, 0xf8, 0xf0, 0x78, 0xf8, 0xf8, 0x7f, 0xfc, 0x3f, 0xde, 0x1f, 0x8e}},
		/* ' */ tinyfont.Glyph{Width:0x4, Height:0x6, XAdvance:0x6, XOffset:1, YOffset:-17, Bitmaps:[]uint8{0xff, 0xff, 0x66}},
		/* ( */ tinyfont.Glyph{Width:0x6, Height:0x16, XAdvance:0x8, XOffset:1, YOffset:-17, Bitmaps:[]uint8{0xc, 0x73, 0x8e, 0x71, 0xc7, 0x38, 0xe3, 0x8e, 0x38, 0xe3, 0x8e, 0x1c, 0x71, 0xc3, 0x8e, 0x18, 0x70}},
		/* ) */ tinyfont.Glyph{Width:0x6, Height:0x16, XAdvance:0x8, XOffset:1, YOffset:-17, Bitmaps:[]uint8{0xc3, 0x87, 0x1c, 0x38, 0xe3, 0x87, 0x1c, 0x71, 0xc7, 0x1c, 0x71, 0xce, 0x38, 0xe7, 0x1c, 0x63, 0x80}},
		/* * */ tinyfont.Glyph{Width:0x7, Height:0x8, XAdvance:0x9, XOffset:1, YOffset:-17, Bitmaps:[]uint8{0x10, 0x23, 0x5f, 0xf3, 0x87, 0x1b, 0x14}},
		/* + */ tinyfont.Glyph{Width:0xb, Height:0xb, XAdvance:0xe, XOffset:2, YOffset:-10, Bitmaps:[]uint8{0xe, 0x1, 0xc0, 0x38, 0x7, 0xf, 0xff, 0xff, 0xff, 0xf8, 0x70, 0xe, 0x1, 0xc0, 0x38, 0x0}},
		/* , */ tinyfont.Glyph{Width:0x4, Height:0x7, XAdvance:0x6, XOffset:1, YOffset:-2, Bitmaps:[]uint8{0xff, 0xf3, 0x36, 0xc0}},
		/* - */ tinyfont.Glyph{Width:0x6, Height:0x3, XAdvance:0x8, XOffset:1, YOffset:-7, Bitmaps:[]uint8{0xff, 0xff, 0xc0}},
		/* . */ tinyfont.Glyph{Width:0x4, Height:0x3, XAdvance:0x6, XOffset:1, YOffset:-2, Bitmaps:[]uint8{0xff, 0xf0}},
		/* / */ tinyfont.Glyph{Width:0x6, Height:0x11, XAdvance:0x7, XOffset:0, YOffset:-16, Bitmaps:[]uint8{0xc, 0x30, 0x86, 0x18, 0x61, 0xc, 0x30, 0xc2, 0x18, 0x61, 0x84, 0x30, 0xc0}},
		/* 0 */ tinyfont.Glyph{Width:0xc, Height:0x11, XAdvance:0xd, XOffset:1, YOffset:-16, Bitmaps:[]uint8{0x1f, 0x83, 0xfc, 0x7f, 0xe7, 0x9e, 0xf0, 0xff, 0xf, 0xf0, 0xff, 0xf, 0xf0, 0xff, 0xf, 0xf0, 0xff, 0xf, 0xf0, 0xf7, 0x9e, 0x7f, 0xe3, 0xfc, 0xf, 0x0}},
		/* 1 */ tinyfont.Glyph{Width:0x7, Height:0x11, XAdvance:0xe, XOffset:3, YOffset:-16, Bitmaps:[]uint8{0x6, 0x1c, 0x7f, 0xff, 0xe3, 0xc7, 0x8f, 0x1e, 0x3c, 0x78, 0xf1, 0xe3, 0xc7, 0x8f, 0x1e}},
		/* 2 */ tinyfont.Glyph{Width:0xc, Height:0x11, XAdvance:0xd, XOffset:1, YOffset:-16, Bitmaps:[]uint8{0x1f, 0x83, 0xfc, 0x7f, 0xef, 0x9f, 0xf0, 0xff, 0xf, 0x0, 0xf0, 0xf, 0x1, 0xe0, 0x3c, 0xf, 0x81, 0xe0, 0x3c, 0x3, 0x80, 0x7f, 0xf7, 0xff, 0x7f, 0xf0}},
		/* 3 */ tinyfont.Glyph{Width:0xc, Height:0x11, XAdvance:0xd, XOffset:1, YOffset:-16, Bitmaps:[]uint8{0x1f, 0x7, 0xfc, 0xff, 0xef, 0x1e, 0xf1, 0xe0, 0x1e, 0x3, 0xc0, 0x78, 0x7, 0xc0, 0x1e, 0x0, 0xf0, 0xf, 0xf0, 0xff, 0x1f, 0x7f, 0xe7, 0xfc, 0x1f, 0x80}},
		/* 4 */ tinyfont.Glyph{Width:0xb, Height:0x11, XAdvance:0xd, XOffset:1, YOffset:-16, Bitmaps:[]uint8{0x3, 0xc0, 0xf8, 0x1f, 0x7, 0xe1, 0xbc, 0x27, 0x8c, 0xf3, 0x1e, 0x63, 0xd8, 0x7b, 0xff, 0xff, 0xff, 0xfe, 0x7, 0x80, 0xf0, 0x1e, 0x3, 0xc0}},
		/* 5 */ tinyfont.Glyph{Width:0xc, Height:0x11, XAdvance:0xd, XOffset:1, YOffset:-16, Bitmaps:[]uint8{0x3f, 0xe7, 0xfe, 0x7f, 0xe7, 0x0, 0x60, 0x6, 0xf8, 0x7f, 0xcf, 0xfe, 0xf1, 0xf0, 0xf, 0x0, 0xf0, 0xf, 0x0, 0xfe, 0x1e, 0xff, 0xe7, 0xfc, 0x3f, 0x0}},
		/* 6 */ tinyfont.Glyph{Width:0xc, Height:0x11, XAdvance:0xd, XOffset:1, YOffset:-16, Bitmaps:[]uint8{0xf, 0x83, 0xfc, 0x7f, 0xe7, 0x9f, 0xf0, 0xf, 0x78, 0xff, 0xcf, 0xfe, 0xf9, 0xff, 0xf, 0xf0, 0xff, 0xf, 0xf0, 0xf7, 0x9f, 0x7f, 0xe3, 0xfc, 0xf, 0x80}},
		/* 7 */ tinyfont.Glyph{Width:0xb, Height:0x11, XAdvance:0xd, XOffset:1, YOffset:-16, Bitmaps:[]uint8{0xff, 0xff, 0xff, 0xff, 0x80, 0xe0, 0x1c, 0x7, 0x1, 0xe0, 0x38, 0xf, 0x1, 0xc0, 0x78, 0xf, 0x1, 0xe0, 0x38, 0xf, 0x1, 0xe0, 0x3c, 0x0}},
		/* 8 */ tinyfont.Glyph{Width:0xc, Height:0x11, XAdvance:0xd, XOffset:1, YOffset:-16, Bitmaps:[]uint8{0xf, 0x3, 0xfc, 0x7f, 0xc7, 0x9e, 0x70, 0xe7, 0xe, 0x39, 0xc1, 0xf8, 0x3f, 0xc7, 0x9e, 0xf0, 0xff, 0xf, 0xf0, 0xff, 0x9f, 0x7f, 0xe3, 0xfc, 0x1f, 0x80}},
		/* 9 */ tinyfont.Glyph{Width:0xc, Height:0x11, XAdvance:0xd, XOffset:1, YOffset:-16, Bitmaps:[]uint8{0x1f, 0x3, 0xfc, 0x7f, 0xef, 0x9e, 0xf0, 0xef, 0xf, 0xf0, 0xff, 0xf, 0xf9, 0xf7, 0xff, 0x3f, 0xf1, 0xef, 0x0, 0xef, 0x1e, 0x7f, 0xe7, 0xfc, 0x1f, 0x0}},
		/* : */ tinyfont.Glyph{Width:0x4, Height:0xc, XAdvance:0x6, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0xff, 0xf0, 0x0, 0x0, 0xf, 0xff}},
		/* ; */ tinyfont.Glyph{Width:0x4, Height:0x10, XAdvance:0x6, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0xff, 0xf0, 0x0, 0x0, 0xf, 0xff, 0x11, 0x6c}},
		/* < */ tinyfont.Glyph{Width:0xc, Height:0xc, XAdvance:0xe, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0x0, 0x10, 0x7, 0x3, 0xf1, 0xfc, 0x7e, 0xf, 0x80, 0xe0, 0xf, 0xc0, 0x3f, 0x80, 0x7f, 0x0, 0xf0, 0x3}},
		/* = */ tinyfont.Glyph{Width:0xc, Height:0x9, XAdvance:0xe, XOffset:1, YOffset:-9, Bitmaps:[]uint8{0xff, 0xff, 0xff, 0xff, 0xf0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xff, 0xff, 0xff, 0xf0}},
		/* > */ tinyfont.Glyph{Width:0xc, Height:0xc, XAdvance:0xe, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0x0, 0xe, 0x0, 0xfc, 0x7, 0xf0, 0xf, 0xe0, 0x1f, 0x0, 0xf0, 0x7f, 0x1f, 0x8f, 0xe0, 0xf0, 0x8, 0x0}},
		/* ? */ tinyfont.Glyph{Width:0xc, Height:0x12, XAdvance:0xf, XOffset:2, YOffset:-17, Bitmaps:[]uint8{0x1f, 0x7, 0xfc, 0x7f, 0xef, 0x9f, 0xf0, 0xff, 0xf, 0x0, 0xf0, 0xf, 0x1, 0xe0, 0x3c, 0x7, 0x80, 0xf0, 0xe, 0x0, 0xe0, 0x0, 0x0, 0xf0, 0xf, 0x0, 0xf0}},
		/* @ */ tinyfont.Glyph{Width:0x15, Height:0x15, XAdvance:0x17, XOffset:1, YOffset:-17, Bitmaps:[]uint8{0x0, 0xfe, 0x0, 0x1f, 0xfc, 0x3, 0xc0, 0xf0, 0x38, 0x1, 0xc3, 0x80, 0x7, 0x18, 0x3d, 0x99, 0x87, 0xec, 0x6c, 0x71, 0xc3, 0xc3, 0x6, 0x1e, 0x18, 0x30, 0xf1, 0x81, 0x87, 0x8c, 0x18, 0x7c, 0x60, 0xc3, 0x63, 0x8e, 0x3b, 0x8f, 0xdf, 0x8c, 0x3c, 0xf0, 0x70, 0x0, 0x1, 0xc0, 0x0, 0x7, 0x80, 0x80, 0x1f, 0xfe, 0x0, 0x1f, 0xc0, 0x0}},
		/* A */ tinyfont.Glyph{Width:0x10, Height:0x12, XAdvance:0x11, XOffset:0, YOffset:-17, Bitmaps:[]uint8{0x3, 0xe0, 0x3, 0xe0, 0x3, 0xe0, 0x7, 0xf0, 0x7, 0xf0, 0x7, 0x70, 0xf, 0x78, 0xe, 0x78, 0xe, 0x38, 0x1e, 0x3c, 0x1c, 0x3c, 0x3f, 0xfc, 0x3f, 0xfe, 0x3f, 0xfe, 0x78, 0xe, 0x78, 0xf, 0x70, 0xf, 0xf0, 0x7}},
		/* B */ tinyfont.Glyph{Width:0xe, Height:0x12, XAdvance:0x11, XOffset:2, YOffset:-17, Bitmaps:[]uint8{0xff, 0xc3, 0xff, 0xcf, 0xff, 0x3c, 0x3e, 0xf0, 0x7b, 0xc1, 0xef, 0xf, 0xbf, 0xfc, 0xff, 0xe3, 0xff, 0xcf, 0x7, 0xbc, 0xf, 0xf0, 0x3f, 0xc0, 0xff, 0x7, 0xff, 0xfe, 0xff, 0xfb, 0xff, 0x80}},
		/* C */ tinyfont.Glyph{Width:0x10, Height:0x12, XAdvance:0x11, XOffset:1, YOffset:-17, Bitmaps:[]uint8{0x7, 0xe0, 0x1f, 0xf8, 0x3f, 0xfc, 0x7c, 0x3e, 0x78, 0x1f, 0xf8, 0xf, 0xf0, 0x0, 0xf0, 0x0, 0xf0, 0x0, 0xf0, 0x0, 0xf0, 0x0, 0xf0, 0x0, 0xf8, 0xf, 0x78, 0x1f, 0x7c, 0x3e, 0x3f, 0xfe, 0x1f, 0xfc, 0x7, 0xf0}},
		/* D */ tinyfont.Glyph{Width:0xf, Height:0x12, XAdvance:0x11, XOffset:2, YOffset:-17, Bitmaps:[]uint8{0xff, 0xe1, 0xff, 0xe3, 0xff, 0xe7, 0x83, 0xef, 0x3, 0xde, 0x7, 0xfc, 0x7, 0xf8, 0xf, 0xf0, 0x1f, 0xe0, 0x3f, 0xc0, 0x7f, 0x80, 0xff, 0x3, 0xfe, 0x7, 0xbc, 0x1f, 0x7f, 0xfc, 0xff, 0xf1, 0xff, 0x80}},
		/* E */ tinyfont.Glyph{Width:0xd, Height:0x12, XAdvance:0x10, XOffset:2, YOffset:-17, Bitmaps:[]uint8{0xff, 0xf7, 0xff, 0xbf, 0xfd, 0xe0, 0xf, 0x0, 0x78, 0x3, 0xc0, 0x1f, 0xfc, 0xff, 0xe7, 0xff, 0x3c, 0x1, 0xe0, 0xf, 0x0, 0x78, 0x3, 0xc0, 0x1f, 0xff, 0xff, 0xff, 0xff, 0xc0}},
		/* F */ tinyfont.Glyph{Width:0xc, Height:0x12, XAdvance:0xf, XOffset:2, YOffset:-17, Bitmaps:[]uint8{0xff, 0xff, 0xff, 0xff, 0xff, 0x0, 0xf0, 0xf, 0x0, 0xf0, 0xf, 0xfe, 0xff, 0xef, 0xfe, 0xf0, 0xf, 0x0, 0xf0, 0xf, 0x0, 0xf0, 0xf, 0x0, 0xf0, 0xf, 0x0}},
		/* G */ tinyfont.Glyph{Width:0x10, Height:0x12, XAdvance:0x12, XOffset:1, YOffset:-17, Bitmaps:[]uint8{0x3, 0xf0, 0xf, 0xfc, 0x3f, 0xfe, 0x3e, 0x1f, 0x78, 0x7, 0x78, 0x0, 0xf0, 0x0, 0xf0, 0x0, 0xf0, 0x7f, 0xf0, 0x7f, 0xf0, 0x7f, 0xf0, 0x7, 0x78, 0x7, 0x7c, 0xf, 0x3e, 0x1f, 0x3f, 0xfb, 0xf, 0xfb, 0x3, 0xe3}},
		/* H */ tinyfont.Glyph{Width:0xe, Height:0x12, XAdvance:0x12, XOffset:2, YOffset:-17, Bitmaps:[]uint8{0xf0, 0x3f, 0xc0, 0xff, 0x3, 0xfc, 0xf, 0xf0, 0x3f, 0xc0, 0xff, 0x3, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x3, 0xfc, 0xf, 0xf0, 0x3f, 0xc0, 0xff, 0x3, 0xfc, 0xf, 0xf0, 0x3f, 0xc0, 0xf0}},
		/* I */ tinyfont.Glyph{Width:0x4, Height:0x12, XAdvance:0x7, XOffset:2, YOffset:-17, Bitmaps:[]uint8{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		/* J */ tinyfont.Glyph{Width:0xb, Height:0x12, XAdvance:0xe, XOffset:1, YOffset:-17, Bitmaps:[]uint8{0x1, 0xe0, 0x3c, 0x7, 0x80, 0xf0, 0x1e, 0x3, 0xc0, 0x78, 0xf, 0x1, 0xe0, 0x3c, 0x7, 0xf8, 0xff, 0x1f, 0xe3, 0xfc, 0x7b, 0xfe, 0x7f, 0xc3, 0xe0}},
		/* K */ tinyfont.Glyph{Width:0x10, Height:0x12, XAdvance:0x11, XOffset:2, YOffset:-17, Bitmaps:[]uint8{0xf0, 0x3e, 0xf0, 0x3c, 0xf0, 0x78, 0xf0, 0xf0, 0xf1, 0xe0, 0xf3, 0xc0, 0xf7, 0x80, 0xff, 0x0, 0xff, 0x80, 0xff, 0x80, 0xfb, 0xc0, 0xf1, 0xe0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0x78, 0xf0, 0x3c, 0xf0, 0x3e, 0xf0, 0x1e}},
		/* L */ tinyfont.Glyph{Width:0xb, Height:0x12, XAdvance:0xf, XOffset:2, YOffset:-17, Bitmaps:[]uint8{0xf0, 0x1e, 0x3, 0xc0, 0x78, 0xf, 0x1, 0xe0, 0x3c, 0x7, 0x80, 0xf0, 0x1e, 0x3, 0xc0, 0x78, 0xf, 0x1, 0xe0, 0x3c, 0x7, 0xff, 0xff, 0xff, 0xfc}},
		/* M */ tinyfont.Glyph{Width:0x11, Height:0x12, XAdvance:0x15, XOffset:2, YOffset:-17, Bitmaps:[]uint8{0xf8, 0x1f, 0xfe, 0xf, 0xff, 0xf, 0xff, 0x87, 0xff, 0xc3, 0xff, 0xe1, 0xff, 0xf9, 0xff, 0xfc, 0xef, 0xfe, 0x77, 0xfb, 0x3b, 0xfd, 0xdd, 0xfe, 0xfc, 0xff, 0x7e, 0x7f, 0x9f, 0x3f, 0xcf, 0x9f, 0xe7, 0x8f, 0xf3, 0xc7, 0xf8, 0xe3, 0xc0}},
		/* N */ tinyfont.Glyph{Width:0xf, Height:0x12, XAdvance:0x12, XOffset:2, YOffset:-17, Bitmaps:[]uint8{0xf0, 0x1f, 0xf0, 0x3f, 0xf0, 0x7f, 0xe0, 0xff, 0xe1, 0xff, 0xc3, 0xfd, 0xc7, 0xfb, 0x8f, 0xf3, 0x9f, 0xe7, 0x3f, 0xc7, 0x7f, 0x8f, 0xff, 0xf, 0xfe, 0x1f, 0xfc, 0x1f, 0xf8, 0x1f, 0xf0, 0x3f, 0xe0, 0x3c}},
		/* O */ tinyfont.Glyph{Width:0x11, Height:0x12, XAdvance:0x13, XOffset:1, YOffset:-17, Bitmaps:[]uint8{0x3, 0xe0, 0xf, 0xfc, 0xf, 0xff, 0x87, 0xc7, 0xc7, 0x80, 0xf3, 0xc0, 0x7b, 0xc0, 0x1f, 0xe0, 0xf, 0xf0, 0x7, 0xf8, 0x3, 0xfc, 0x1, 0xfe, 0x0, 0xf7, 0x80, 0xf3, 0xc0, 0x78, 0xf0, 0xf8, 0x7f, 0xfc, 0x1f, 0xfc, 0x3, 0xf8, 0x0}},
		/* P */ tinyfont.Glyph{Width:0xe, Height:0x12, XAdvance:0x10, XOffset:2, YOffset:-17, Bitmaps:[]uint8{0xff, 0xe3, 0xff, 0xef, 0xff, 0xbc, 0x1f, 0xf0, 0x3f, 0xc0, 0xff, 0x3, 0xfc, 0x1f, 0xff, 0xfb, 0xff, 0xcf, 0xfe, 0x3c, 0x0, 0xf0, 0x3, 0xc0, 0xf, 0x0, 0x3c, 0x0, 0xf0, 0x3, 0xc0, 0x0}},
		/* Q */ tinyfont.Glyph{Width:0x11, Height:0x13, XAdvance:0x13, XOffset:1, YOffset:-17, Bitmaps:[]uint8{0x3, 0xe0, 0xf, 0xfc, 0xf, 0xff, 0x87, 0xc7, 0xc7, 0x80, 0xf3, 0xc0, 0x7b, 0xc0, 0x1f, 0xe0, 0xf, 0xf0, 0x7, 0xf8, 0x3, 0xfc, 0x1, 0xfe, 0x4, 0xf7, 0x87, 0xf3, 0xc3, 0xf8, 0xf0, 0xf8, 0x7f, 0xfc, 0x1f, 0xff, 0x83, 0xf1, 0x80, 0x0, 0x0}},
		/* R */ tinyfont.Glyph{Width:0x10, Height:0x12, XAdvance:0x11, XOffset:2, YOffset:-17, Bitmaps:[]uint8{0xff, 0xf8, 0xff, 0xfc, 0xff, 0xfc, 0xf0, 0x3e, 0xf0, 0x1e, 0xf0, 0x1e, 0xf0, 0x1e, 0xf0, 0x3c, 0xff, 0xf8, 0xff, 0xf0, 0xff, 0xf8, 0xf0, 0x3c, 0xf0, 0x3c, 0xf0, 0x3c, 0xf0, 0x3c, 0xf0, 0x3c, 0xf0, 0x3c, 0xf0, 0x1f}},
		/* S */ tinyfont.Glyph{Width:0xf, Height:0x12, XAdvance:0x10, XOffset:1, YOffset:-17, Bitmaps:[]uint8{0xf, 0xc0, 0x7f, 0xe1, 0xff, 0xe7, 0xc3, 0xef, 0x3, 0xde, 0x0, 0x3c, 0x0, 0x7f, 0x0, 0x7f, 0xf0, 0x3f, 0xf8, 0xf, 0xf8, 0x1, 0xf0, 0x1, 0xfe, 0x3, 0xde, 0xf, 0xbf, 0xfe, 0x3f, 0xf8, 0x1f, 0xc0}},
		/* T */ tinyfont.Glyph{Width:0xc, Height:0x12, XAdvance:0xf, XOffset:2, YOffset:-17, Bitmaps:[]uint8{0xff, 0xff, 0xff, 0xff, 0xf0, 0xf0, 0xf, 0x0, 0xf0, 0xf, 0x0, 0xf0, 0xf, 0x0, 0xf0, 0xf, 0x0, 0xf0, 0xf, 0x0, 0xf0, 0xf, 0x0, 0xf0, 0xf, 0x0, 0xf0}},
		/* U */ tinyfont.Glyph{Width:0xe, Height:0x12, XAdvance:0x12, XOffset:2, YOffset:-17, Bitmaps:[]uint8{0xf0, 0x3f, 0xc0, 0xff, 0x3, 0xfc, 0xf, 0xf0, 0x3f, 0xc0, 0xff, 0x3, 0xfc, 0xf, 0xf0, 0x3f, 0xc0, 0xff, 0x3, 0xfc, 0xf, 0xf0, 0x3f, 0xc0, 0xf7, 0x87, 0x9f, 0xfe, 0x3f, 0xf0, 0x3f, 0x0}},
		/* V */ tinyfont.Glyph{Width:0xf, Height:0x12, XAdvance:0x10, XOffset:0, YOffset:-17, Bitmaps:[]uint8{0x70, 0xe, 0xf0, 0x3d, 0xe0, 0x79, 0xc0, 0xe3, 0x81, 0xc7, 0x87, 0x87, 0xe, 0xe, 0x1c, 0x1e, 0x78, 0x1c, 0xe0, 0x39, 0xc0, 0x73, 0x80, 0x7e, 0x0, 0xfc, 0x1, 0xf8, 0x1, 0xe0, 0x3, 0xc0, 0x7, 0x80}},
		/* W */ tinyfont.Glyph{Width:0x17, Height:0x12, XAdvance:0x17, XOffset:0, YOffset:-17, Bitmaps:[]uint8{0x70, 0x38, 0x1c, 0xe0, 0xf0, 0x79, 0xe1, 0xf0, 0xf3, 0xc3, 0xe1, 0xe3, 0x87, 0xc3, 0x87, 0xf, 0x87, 0xe, 0x3b, 0x9e, 0x1e, 0x77, 0x38, 0x1c, 0xee, 0x70, 0x39, 0xcc, 0xe0, 0x73, 0x99, 0xc0, 0x6e, 0x3f, 0x0, 0xfc, 0x7e, 0x1, 0xf8, 0xfc, 0x3, 0xf0, 0xf8, 0x3, 0xe1, 0xe0, 0x7, 0x83, 0xc0, 0xf, 0x7, 0x80}},
		/* X */ tinyfont.Glyph{Width:0xf, Height:0x12, XAdvance:0x10, XOffset:1, YOffset:-17, Bitmaps:[]uint8{0xf0, 0x3c, 0xf0, 0xf9, 0xe1, 0xe1, 0xe7, 0x83, 0xcf, 0x3, 0xfc, 0x3, 0xf0, 0x7, 0xe0, 0x7, 0x80, 0xf, 0x0, 0x3f, 0x0, 0xff, 0x1, 0xfe, 0x7, 0x9e, 0xf, 0x1e, 0x3c, 0x3c, 0xf8, 0x3d, 0xe0, 0x78}},
		/* Y */ tinyfont.Glyph{Width:0x10, Height:0x12, XAdvance:0xf, XOffset:0, YOffset:-17, Bitmaps:[]uint8{0xf0, 0x1e, 0x78, 0x1e, 0x78, 0x3c, 0x3c, 0x3c, 0x3c, 0x78, 0x1e, 0x78, 0xe, 0x70, 0xf, 0xf0, 0x7, 0xe0, 0x7, 0xe0, 0x3, 0xc0, 0x3, 0xc0, 0x3, 0xc0, 0x3, 0xc0, 0x3, 0xc0, 0x3, 0xc0, 0x3, 0xc0, 0x3, 0xc0}},
		/* Z */ tinyfont.Glyph{Width:0xd, Height:0x12, XAdvance:0xf, XOffset:1, YOffset:-17, Bitmaps:[]uint8{0xff, 0xff, 0xff, 0xff, 0xfe, 0x1, 0xf0, 0xf, 0x0, 0xf0, 0xf, 0x0, 0xf8, 0x7, 0x80, 0x78, 0x7, 0x80, 0x7c, 0x3, 0xc0, 0x3c, 0x3, 0xc0, 0x1f, 0xff, 0xff, 0xff, 0xff, 0xc0}},
		/* [ */ tinyfont.Glyph{Width:0x6, Height:0x17, XAdvance:0x8, XOffset:2, YOffset:-17, Bitmaps:[]uint8{0xff, 0xff, 0xfc, 0xf3, 0xcf, 0x3c, 0xf3, 0xcf, 0x3c, 0xf3, 0xcf, 0x3c, 0xf3, 0xcf, 0x3c, 0xff, 0xff, 0xc0}},
		/* \ */ tinyfont.Glyph{Width:0x7, Height:0x11, XAdvance:0x7, XOffset:0, YOffset:-16, Bitmaps:[]uint8{0xc1, 0x81, 0x3, 0x6, 0x4, 0xc, 0x18, 0x10, 0x30, 0x60, 0x40, 0xc1, 0x81, 0x3, 0x6}},
		/* ] */ tinyfont.Glyph{Width:0x6, Height:0x17, XAdvance:0x8, XOffset:0, YOffset:-17, Bitmaps:[]uint8{0xff, 0xff, 0xcf, 0x3c, 0xf3, 0xcf, 0x3c, 0xf3, 0xcf, 0x3c, 0xf3, 0xcf, 0x3c, 0xf3, 0xcf, 0xff, 0xff, 0xc0}},
		/* ^ */ tinyfont.Glyph{Width:0xc, Height:0xb, XAdvance:0xe, XOffset:1, YOffset:-16, Bitmaps:[]uint8{0xf, 0x0, 0xf0, 0xf, 0x1, 0xf8, 0x1b, 0x83, 0x9c, 0x39, 0xc3, 0xc, 0x70, 0xe7, 0xe, 0xe0, 0x70}},
		/* _ */ tinyfont.Glyph{Width:0xf, Height:0x2, XAdvance:0xd, XOffset:-1, YOffset:4, Bitmaps:[]uint8{0xff, 0xff, 0xff, 0xfc}},
		/* ` */ tinyfont.Glyph{Width:0x4, Height:0x3, XAdvance:0x6, XOffset:0, YOffset:-17, Bitmaps:[]uint8{0xe6, 0x30}},
		/* a */ tinyfont.Glyph{Width:0xd, Height:0xd, XAdvance:0xe, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0x1f, 0x83, 0xff, 0x1f, 0xfd, 0xe1, 0xe0, 0xf, 0x3, 0xf9, 0xff, 0xdf, 0x1e, 0xf0, 0xf7, 0x8f, 0xbf, 0xfc, 0xff, 0xe3, 0xcf, 0x80}},
		/* b */ tinyfont.Glyph{Width:0xd, Height:0x12, XAdvance:0xf, XOffset:2, YOffset:-17, Bitmaps:[]uint8{0xf0, 0x7, 0x80, 0x3c, 0x1, 0xe0, 0xf, 0x0, 0x7b, 0xc3, 0xff, 0x9f, 0xfe, 0xf8, 0xf7, 0x83, 0xfc, 0x1f, 0xe0, 0xff, 0x7, 0xf8, 0x3f, 0xe3, 0xdf, 0xfe, 0xff, 0xe7, 0xbe, 0x0}},
		/* c */ tinyfont.Glyph{Width:0xc, Height:0xd, XAdvance:0xd, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xf, 0x83, 0xfe, 0x7f, 0xf7, 0x8f, 0xf0, 0x7f, 0x0, 0xf0, 0xf, 0x0, 0xf0, 0x77, 0x8f, 0x7f, 0xf3, 0xfe, 0xf, 0x80}},
		/* d */ tinyfont.Glyph{Width:0xd, Height:0x12, XAdvance:0xf, XOffset:1, YOffset:-17, Bitmaps:[]uint8{0x0, 0x78, 0x3, 0xc0, 0x1e, 0x0, 0xf0, 0x7, 0x8f, 0xbc, 0xff, 0xef, 0xff, 0x78, 0xff, 0x83, 0xfc, 0x1f, 0xe0, 0xff, 0x7, 0xf8, 0x3d, 0xe3, 0xef, 0xff, 0x3f, 0xf8, 0xfb, 0xc0}},
		/* e */ tinyfont.Glyph{Width:0xd, Height:0xd, XAdvance:0xe, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0x1f, 0x81, 0xfe, 0x1f, 0xf9, 0xf1, 0xcf, 0x7, 0x7f, 0xfb, 0xff, 0xde, 0x0, 0xf0, 0x3, 0xc3, 0x9f, 0xfc, 0x7f, 0xc0, 0xf8, 0x0}},
		/* f */ tinyfont.Glyph{Width:0x7, Height:0x12, XAdvance:0x8, XOffset:1, YOffset:-17, Bitmaps:[]uint8{0x3e, 0xfd, 0xfb, 0xc7, 0x9f, 0xbf, 0x3c, 0x78, 0xf1, 0xe3, 0xc7, 0x8f, 0x1e, 0x3c, 0x78, 0xf0}},
		/* g */ tinyfont.Glyph{Width:0xd, Height:0x12, XAdvance:0xf, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0x1e, 0x79, 0xfb, 0xdf, 0xfe, 0xf1, 0xff, 0x7, 0xf8, 0x3f, 0xc1, 0xfe, 0xf, 0xf0, 0x7f, 0xc7, 0xdf, 0xfe, 0x7f, 0xf1, 0xf7, 0x80, 0x3c, 0x1, 0xff, 0x1e, 0x7f, 0xf0, 0xfe, 0x0}},
		/* h */ tinyfont.Glyph{Width:0xc, Height:0x12, XAdvance:0xe, XOffset:2, YOffset:-17, Bitmaps:[]uint8{0xf0, 0xf, 0x0, 0xf0, 0xf, 0x0, 0xf0, 0xf, 0x7c, 0xff, 0xef, 0xff, 0xf9, 0xff, 0xf, 0xf0, 0xff, 0xf, 0xf0, 0xff, 0xf, 0xf0, 0xff, 0xf, 0xf0, 0xff, 0xf}},
		/* i */ tinyfont.Glyph{Width:0x4, Height:0x12, XAdvance:0x7, XOffset:2, YOffset:-17, Bitmaps:[]uint8{0xff, 0xf0, 0xf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		/* j */ tinyfont.Glyph{Width:0x6, Height:0x17, XAdvance:0x7, XOffset:0, YOffset:-17, Bitmaps:[]uint8{0x3c, 0xf3, 0xc0, 0x0, 0xf3, 0xcf, 0x3c, 0xf3, 0xcf, 0x3c, 0xf3, 0xcf, 0x3c, 0xf3, 0xcf, 0xff, 0xff, 0x80}},
		/* k */ tinyfont.Glyph{Width:0xc, Height:0x12, XAdvance:0xe, XOffset:2, YOffset:-17, Bitmaps:[]uint8{0xf0, 0xf, 0x0, 0xf0, 0xf, 0x0, 0xf0, 0xf, 0xf, 0xf1, 0xef, 0x3c, 0xf7, 0x8f, 0xf0, 0xff, 0xf, 0xf8, 0xff, 0x8f, 0x3c, 0xf1, 0xcf, 0x1e, 0xf0, 0xef, 0xf}},
		/* l */ tinyfont.Glyph{Width:0x4, Height:0x12, XAdvance:0x6, XOffset:2, YOffset:-17, Bitmaps:[]uint8{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		/* m */ tinyfont.Glyph{Width:0x13, Height:0xd, XAdvance:0x15, XOffset:2, YOffset:-12, Bitmaps:[]uint8{0xf7, 0x8f, 0x9f, 0xfb, 0xfb, 0xff, 0xff, 0xfc, 0xf8, 0xff, 0x1e, 0x1f, 0xe3, 0xc3, 0xfc, 0x78, 0x7f, 0x8f, 0xf, 0xf1, 0xe1, 0xfe, 0x3c, 0x3f, 0xc7, 0x87, 0xf8, 0xf0, 0xff, 0x1e, 0x1e}},
		/* n */ tinyfont.Glyph{Width:0xc, Height:0xd, XAdvance:0xf, XOffset:2, YOffset:-12, Bitmaps:[]uint8{0xf7, 0xcf, 0xfe, 0xff, 0xff, 0x9f, 0xf0, 0xff, 0xf, 0xf0, 0xff, 0xf, 0xf0, 0xff, 0xf, 0xf0, 0xff, 0xf, 0xf0, 0xf0}},
		/* o */ tinyfont.Glyph{Width:0xd, Height:0xd, XAdvance:0xf, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xf, 0x81, 0xff, 0x1f, 0xfc, 0xf1, 0xef, 0x7, 0xf8, 0x3f, 0xc1, 0xfe, 0xf, 0xf0, 0x7b, 0xc7, 0x9f, 0xfc, 0x7f, 0xc0, 0xf8, 0x0}},
		/* p */ tinyfont.Glyph{Width:0xd, Height:0x12, XAdvance:0xf, XOffset:2, YOffset:-12, Bitmaps:[]uint8{0xf7, 0xc7, 0xff, 0x3f, 0xfd, 0xf1, 0xef, 0x7, 0xf8, 0x3f, 0xc1, 0xfe, 0xf, 0xf0, 0x7f, 0xc7, 0xbf, 0xfd, 0xff, 0xcf, 0x78, 0x78, 0x3, 0xc0, 0x1e, 0x0, 0xf0, 0x7, 0x80, 0x0}},
		/* q */ tinyfont.Glyph{Width:0xd, Height:0x12, XAdvance:0xf, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xf, 0x79, 0xff, 0xdf, 0xfe, 0xf1, 0xff, 0x7, 0xf8, 0x3f, 0xc1, 0xfe, 0xf, 0xf0, 0x7b, 0xc7, 0xdf, 0xfe, 0x7f, 0xf1, 0xf7, 0x80, 0x3c, 0x1, 0xe0, 0xf, 0x0, 0x78, 0x3, 0xc0}},
		/* r */ tinyfont.Glyph{Width:0x8, Height:0xd, XAdvance:0x9, XOffset:2, YOffset:-12, Bitmaps:[]uint8{0xf3, 0xf7, 0xff, 0xf8, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0}},
		/* s */ tinyfont.Glyph{Width:0xc, Height:0xd, XAdvance:0xd, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0x1f, 0x87, 0xfc, 0xff, 0xef, 0xf, 0xf8, 0xf, 0xf0, 0x7f, 0xe0, 0xff, 0x1, 0xff, 0xf, 0xff, 0xe7, 0xfe, 0x1f, 0x80}},
		/* t */ tinyfont.Glyph{Width:0x6, Height:0xf, XAdvance:0x8, XOffset:1, YOffset:-14, Bitmaps:[]uint8{0x79, 0xe7, 0xbf, 0xfd, 0xe7, 0x9e, 0x79, 0xe7, 0x9e, 0x7d, 0xf3, 0xc0}},
		/* u */ tinyfont.Glyph{Width:0xc, Height:0xd, XAdvance:0xf, XOffset:2, YOffset:-12, Bitmaps:[]uint8{0xf0, 0xff, 0xf, 0xf0, 0xff, 0xf, 0xf0, 0xff, 0xf, 0xf0, 0xff, 0xf, 0xf0, 0xff, 0x1f, 0xff, 0xf7, 0xff, 0x3e, 0xf0}},
		/* v */ tinyfont.Glyph{Width:0xd, Height:0xd, XAdvance:0xd, XOffset:0, YOffset:-12, Bitmaps:[]uint8{0xf0, 0x7b, 0x83, 0x9e, 0x1c, 0xf1, 0xe3, 0x8e, 0x1c, 0x70, 0x77, 0x83, 0xb8, 0x1d, 0xc0, 0x7e, 0x3, 0xe0, 0x1f, 0x0, 0x70, 0x0}},
		/* w */ tinyfont.Glyph{Width:0x12, Height:0xd, XAdvance:0x13, XOffset:0, YOffset:-12, Bitmaps:[]uint8{0xf0, 0xe1, 0xdc, 0x78, 0x77, 0x1f, 0x3d, 0xe7, 0xcf, 0x79, 0xb3, 0x8e, 0x6c, 0xe3, 0xbb, 0x38, 0xee, 0xfc, 0x1f, 0x3f, 0x7, 0xc7, 0xc1, 0xf1, 0xf0, 0x7c, 0x78, 0xe, 0x1e, 0x0}},
		/* x */ tinyfont.Glyph{Width:0xd, Height:0xd, XAdvance:0xd, XOffset:0, YOffset:-12, Bitmaps:[]uint8{0x78, 0xf3, 0xc7, 0x8f, 0x78, 0x3b, 0x81, 0xfc, 0x7, 0xc0, 0x1e, 0x1, 0xf0, 0x1f, 0xc0, 0xef, 0xf, 0x78, 0xf1, 0xe7, 0x87, 0x0}},
		/* y */ tinyfont.Glyph{Width:0xd, Height:0x12, XAdvance:0xd, XOffset:0, YOffset:-12, Bitmaps:[]uint8{0xf0, 0x7b, 0x83, 0x9e, 0x1c, 0x71, 0xe3, 0x8e, 0x1e, 0x70, 0x73, 0x83, 0xb8, 0x1f, 0xc0, 0x7e, 0x3, 0xe0, 0xf, 0x0, 0x70, 0x3, 0x80, 0x3c, 0x7, 0xc0, 0x3e, 0x1, 0xe0, 0x0}},
		/* z */ tinyfont.Glyph{Width:0xa, Height:0xd, XAdvance:0xc, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xff, 0xff, 0xff, 0xfc, 0xf, 0x7, 0x83, 0xc1, 0xe0, 0xf0, 0x78, 0x3c, 0xf, 0xff, 0xff, 0xff, 0xc0}},
		/* { */ tinyfont.Glyph{Width:0x6, Height:0x17, XAdvance:0x9, XOffset:1, YOffset:-17, Bitmaps:[]uint8{0x1c, 0xf3, 0xce, 0x38, 0xe3, 0x8e, 0x38, 0xe3, 0xbc, 0xf0, 0xe3, 0x8e, 0x38, 0xe3, 0x8e, 0x3c, 0xf1, 0xc0}},
		/* | */ tinyfont.Glyph{Width:0x2, Height:0x16, XAdvance:0x7, XOffset:2, YOffset:-17, Bitmaps:[]uint8{0xff, 0xff, 0xff, 0xff, 0xff, 0xf0}},
		/* } */ tinyfont.Glyph{Width:0x6, Height:0x17, XAdvance:0x9, XOffset:3, YOffset:-17, Bitmaps:[]uint8{0xe3, 0x8f, 0x1c, 0x71, 0xc7, 0x1c, 0x71, 0xc7, 0xf, 0x3d, 0xc7, 0x1c, 0x71, 0xc7, 0x1c, 0xf3, 0xce, 0x0}},
		/* ~ */ tinyfont.Glyph{Width:0xc, Height:0x5, XAdvance:0xc, XOffset:0, YOffset:-7, Bitmaps:[]uint8{0x78, 0xf, 0xe0, 0xcf, 0x30, 0x7f, 0x1, 0xe0}},
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
