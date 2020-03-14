package freeserif

import (
	"tinygo.org/x/tinyfont"
)

var Bold9pt7b = tinyfont.Font{
	Glyphs:[]tinyfont.Glyph{
		/*   */ tinyfont.Glyph{Width:0x0, Height:0x0, XAdvance:0x5, XOffset:0, YOffset:1, Bitmaps:[]uint8{}},
		/* ! */ tinyfont.Glyph{Width:0x3, Height:0xc, XAdvance:0x6, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0xff, 0xf4, 0x92, 0x1f, 0xf0}},
		/* " */ tinyfont.Glyph{Width:0x6, Height:0x5, XAdvance:0xa, XOffset:2, YOffset:-11, Bitmaps:[]uint8{0xcf, 0x3c, 0xe3, 0x88}},
		/* # */ tinyfont.Glyph{Width:0x9, Height:0xd, XAdvance:0x9, XOffset:0, YOffset:-12, Bitmaps:[]uint8{0x13, 0x9, 0x84, 0xc2, 0x47, 0xf9, 0x90, 0xc8, 0x4c, 0xff, 0x13, 0x9, 0xc, 0x86, 0x40}},
		/* $ */ tinyfont.Glyph{Width:0x8, Height:0xe, XAdvance:0x9, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0x10, 0x38, 0xd6, 0x92, 0xd2, 0xf0, 0x7c, 0x3e, 0x17, 0x93, 0x93, 0xd6, 0x7c, 0x10}},
		/* % */ tinyfont.Glyph{Width:0xe, Height:0xc, XAdvance:0x12, XOffset:2, YOffset:-11, Bitmaps:[]uint8{0x3c, 0x21, 0xcf, 0xe, 0x24, 0x30, 0xa0, 0xc5, 0x3, 0x34, 0xe7, 0x26, 0x40, 0xb9, 0x4, 0xc4, 0x23, 0x30, 0x8c, 0x84, 0x1c}},
		/* & */ tinyfont.Glyph{Width:0xd, Height:0xc, XAdvance:0xf, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0xf, 0x0, 0xcc, 0x6, 0x60, 0x3e, 0x0, 0xe7, 0x8f, 0x18, 0x9c, 0x8c, 0xe4, 0xe3, 0xc7, 0x9e, 0x3c, 0x72, 0xfd, 0xe0}},
		/* ' */ tinyfont.Glyph{Width:0x2, Height:0x5, XAdvance:0x5, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0xff, 0x80}},
		/* ( */ tinyfont.Glyph{Width:0x4, Height:0xf, XAdvance:0x6, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0x32, 0x44, 0xcc, 0xcc, 0xcc, 0xc4, 0x62, 0x10}},
		/* ) */ tinyfont.Glyph{Width:0x4, Height:0xf, XAdvance:0x6, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0x84, 0x22, 0x33, 0x33, 0x33, 0x32, 0x64, 0x80}},
		/* * */ tinyfont.Glyph{Width:0x7, Height:0x7, XAdvance:0x9, XOffset:2, YOffset:-11, Bitmaps:[]uint8{0x31, 0x6b, 0xb1, 0x8e, 0xd6, 0x8c, 0x0}},
		/* + */ tinyfont.Glyph{Width:0x9, Height:0x9, XAdvance:0xc, XOffset:1, YOffset:-8, Bitmaps:[]uint8{0x8, 0x4, 0x2, 0x1, 0xf, 0xf8, 0x40, 0x20, 0x10, 0x8, 0x0}},
		/* , */ tinyfont.Glyph{Width:0x3, Height:0x6, XAdvance:0x4, XOffset:1, YOffset:-2, Bitmaps:[]uint8{0xdf, 0x95, 0x0}},
		/* - */ tinyfont.Glyph{Width:0x4, Height:0x2, XAdvance:0x6, XOffset:1, YOffset:-4, Bitmaps:[]uint8{0xff}},
		/* . */ tinyfont.Glyph{Width:0x3, Height:0x3, XAdvance:0x4, XOffset:1, YOffset:-2, Bitmaps:[]uint8{0xff, 0x80}},
		/* / */ tinyfont.Glyph{Width:0x6, Height:0xd, XAdvance:0x5, XOffset:0, YOffset:-11, Bitmaps:[]uint8{0xc, 0x21, 0x86, 0x10, 0xc3, 0x8, 0x61, 0x84, 0x30, 0xc0}},
		/* 0 */ tinyfont.Glyph{Width:0x9, Height:0xc, XAdvance:0x9, XOffset:0, YOffset:-11, Bitmaps:[]uint8{0x1c, 0x33, 0x98, 0xdc, 0x7e, 0x3f, 0x1f, 0x8f, 0xc7, 0xe3, 0xb1, 0x98, 0xc3, 0x80}},
		/* 1 */ tinyfont.Glyph{Width:0x6, Height:0xc, XAdvance:0x9, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0x8, 0xe3, 0x8e, 0x38, 0xe3, 0x8e, 0x38, 0xe3, 0xbf}},
		/* 2 */ tinyfont.Glyph{Width:0x9, Height:0xc, XAdvance:0x9, XOffset:0, YOffset:-11, Bitmaps:[]uint8{0x3c, 0x3f, 0x23, 0xc0, 0xe0, 0x70, 0x30, 0x38, 0x18, 0x18, 0x18, 0x5f, 0xdf, 0xe0}},
		/* 3 */ tinyfont.Glyph{Width:0x8, Height:0xc, XAdvance:0x9, XOffset:0, YOffset:-11, Bitmaps:[]uint8{0x7c, 0x8e, 0xe, 0xe, 0xc, 0x1e, 0x7, 0x3, 0x3, 0x2, 0xe6, 0xf8}},
		/* 4 */ tinyfont.Glyph{Width:0x8, Height:0xc, XAdvance:0x9, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0x6, 0xe, 0xe, 0x3e, 0x2e, 0x4e, 0x8e, 0x8e, 0xff, 0xff, 0xe, 0xe}},
		/* 5 */ tinyfont.Glyph{Width:0x8, Height:0xc, XAdvance:0x9, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0x3f, 0x7e, 0x40, 0x40, 0xf8, 0xfc, 0x1e, 0x6, 0x2, 0x2, 0xe4, 0xf8}},
		/* 6 */ tinyfont.Glyph{Width:0x8, Height:0xc, XAdvance:0x9, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0x7, 0x1c, 0x30, 0x70, 0xfc, 0xe6, 0xe7, 0xe7, 0xe7, 0x67, 0x66, 0x3c}},
		/* 7 */ tinyfont.Glyph{Width:0x9, Height:0xc, XAdvance:0x9, XOffset:0, YOffset:-11, Bitmaps:[]uint8{0x7f, 0x3f, 0xa0, 0xd0, 0x40, 0x60, 0x30, 0x10, 0x18, 0xc, 0x4, 0x6, 0x3, 0x0}},
		/* 8 */ tinyfont.Glyph{Width:0x8, Height:0xc, XAdvance:0x9, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0x3c, 0xc6, 0xc6, 0xc6, 0xfc, 0x7c, 0x3e, 0xcf, 0xc7, 0xc7, 0xc6, 0x7c}},
		/* 9 */ tinyfont.Glyph{Width:0x9, Height:0xc, XAdvance:0x9, XOffset:0, YOffset:-11, Bitmaps:[]uint8{0x3e, 0x33, 0xb8, 0xdc, 0x7e, 0x3f, 0x1d, 0xce, 0x7f, 0x7, 0x7, 0xf, 0x1c, 0x0}},
		/* : */ tinyfont.Glyph{Width:0x3, Height:0x9, XAdvance:0x6, XOffset:1, YOffset:-8, Bitmaps:[]uint8{0xff, 0x80, 0x3f, 0xe0}},
		/* ; */ tinyfont.Glyph{Width:0x3, Height:0xc, XAdvance:0x6, XOffset:2, YOffset:-8, Bitmaps:[]uint8{0xff, 0x80, 0x37, 0xe5, 0x40}},
		/* < */ tinyfont.Glyph{Width:0xa, Height:0xa, XAdvance:0xc, XOffset:1, YOffset:-9, Bitmaps:[]uint8{0x0, 0x0, 0x70, 0x78, 0x78, 0x78, 0x38, 0x3, 0x80, 0x3c, 0x3, 0xc0, 0x30}},
		/* = */ tinyfont.Glyph{Width:0xa, Height:0x5, XAdvance:0xc, XOffset:1, YOffset:-6, Bitmaps:[]uint8{0xff, 0xc0, 0x0, 0x0, 0x0, 0xff, 0xc0}},
		/* > */ tinyfont.Glyph{Width:0xa, Height:0xa, XAdvance:0xc, XOffset:1, YOffset:-8, Bitmaps:[]uint8{0xc0, 0x3c, 0x3, 0xc0, 0x1c, 0x1, 0xc1, 0xe1, 0xe1, 0xe0, 0xe0, 0x0, 0x0}},
		/* ? */ tinyfont.Glyph{Width:0x7, Height:0xc, XAdvance:0x9, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0x3d, 0x9f, 0x3e, 0x70, 0xe1, 0x4, 0x8, 0x0, 0x70, 0xe1, 0xc0}},
		/* @ */ tinyfont.Glyph{Width:0xd, Height:0xc, XAdvance:0x11, XOffset:2, YOffset:-11, Bitmaps:[]uint8{0xf, 0x81, 0x83, 0x18, 0xc4, 0x89, 0x9c, 0x4c, 0xe4, 0x67, 0x22, 0x39, 0x22, 0x4f, 0xe3, 0x0, 0xc, 0x10, 0x1f, 0x0}},
		/* A */ tinyfont.Glyph{Width:0xd, Height:0xc, XAdvance:0xd, XOffset:0, YOffset:-11, Bitmaps:[]uint8{0x2, 0x0, 0x30, 0x1, 0xc0, 0xe, 0x0, 0xb8, 0x5, 0xc0, 0x4f, 0x2, 0x38, 0x3f, 0xe1, 0x7, 0x18, 0x3d, 0xe3, 0xf0}},
		/* B */ tinyfont.Glyph{Width:0xb, Height:0xc, XAdvance:0xc, XOffset:0, YOffset:-11, Bitmaps:[]uint8{0xff, 0x87, 0x1c, 0xe3, 0x9c, 0x73, 0x9c, 0x7f, 0xe, 0x71, 0xc7, 0x38, 0xe7, 0x1c, 0xe7, 0x7f, 0xc0}},
		/* C */ tinyfont.Glyph{Width:0xb, Height:0xc, XAdvance:0xd, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0x1f, 0x26, 0x1d, 0xc1, 0xb0, 0x1e, 0x1, 0xc0, 0x38, 0x7, 0x0, 0xe0, 0xe, 0x4, 0xe1, 0xf, 0xc0}},
		/* D */ tinyfont.Glyph{Width:0xb, Height:0xc, XAdvance:0xd, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0xff, 0xe, 0x71, 0xc7, 0x38, 0x77, 0xe, 0xe1, 0xdc, 0x3b, 0x87, 0x70, 0xce, 0x39, 0xc6, 0x7f, 0x80}},
		/* E */ tinyfont.Glyph{Width:0xb, Height:0xc, XAdvance:0xc, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0xff, 0xce, 0x19, 0xc1, 0x38, 0x87, 0x30, 0xfe, 0x1c, 0xc3, 0x88, 0x70, 0x2e, 0xd, 0xc3, 0x7f, 0xe0}},
		/* F */ tinyfont.Glyph{Width:0xa, Height:0xc, XAdvance:0xb, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0xff, 0xdc, 0x37, 0x5, 0xc4, 0x73, 0x1f, 0xc7, 0x31, 0xc4, 0x70, 0x1c, 0x7, 0x3, 0xe0}},
		/* G */ tinyfont.Glyph{Width:0xc, Height:0xc, XAdvance:0xe, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0x1f, 0x23, 0xe, 0x70, 0x6e, 0x2, 0xe0, 0xe, 0x0, 0xe1, 0xfe, 0xe, 0x60, 0xe7, 0xe, 0x38, 0xe0, 0xf8}},
		/* H */ tinyfont.Glyph{Width:0xc, Height:0xc, XAdvance:0xe, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0xf9, 0xf7, 0xe, 0x70, 0xe7, 0xe, 0x70, 0xe7, 0xfe, 0x70, 0xe7, 0xe, 0x70, 0xe7, 0xe, 0x70, 0xef, 0x9f}},
		/* I */ tinyfont.Glyph{Width:0x5, Height:0xc, XAdvance:0x7, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0xfb, 0x9c, 0xe7, 0x39, 0xce, 0x73, 0x9d, 0xf0}},
		/* J */ tinyfont.Glyph{Width:0x8, Height:0xe, XAdvance:0x9, XOffset:0, YOffset:-11, Bitmaps:[]uint8{0x1f, 0xe, 0xe, 0xe, 0xe, 0xe, 0xe, 0xe, 0xe, 0xe, 0xe, 0xce, 0xcc, 0x78}},
		/* K */ tinyfont.Glyph{Width:0xd, Height:0xc, XAdvance:0xe, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0xf9, 0xf3, 0x82, 0x1c, 0x20, 0xe2, 0x7, 0x20, 0x3f, 0x1, 0xdc, 0xe, 0x70, 0x73, 0xc3, 0x8f, 0x1c, 0x3d, 0xf3, 0xf0}},
		/* L */ tinyfont.Glyph{Width:0xb, Height:0xc, XAdvance:0xc, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0xf8, 0xe, 0x1, 0xc0, 0x38, 0x7, 0x0, 0xe0, 0x1c, 0x3, 0x80, 0x70, 0x2e, 0x9, 0xc3, 0x7f, 0xe0}},
		/* M */ tinyfont.Glyph{Width:0x10, Height:0xc, XAdvance:0x11, XOffset:0, YOffset:-11, Bitmaps:[]uint8{0xf8, 0xf, 0x3c, 0x1e, 0x3c, 0x1e, 0x2e, 0x2e, 0x2e, 0x2e, 0x26, 0x4e, 0x27, 0x4e, 0x27, 0x4e, 0x23, 0x8e, 0x23, 0x8e, 0x21, 0xe, 0x71, 0x1f}},
		/* N */ tinyfont.Glyph{Width:0xb, Height:0xc, XAdvance:0xd, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0xf0, 0xee, 0x9, 0xe1, 0x3e, 0x25, 0xe4, 0x9e, 0x91, 0xd2, 0x1e, 0x43, 0xc8, 0x39, 0x3, 0x70, 0x20}},
		/* O */ tinyfont.Glyph{Width:0xc, Height:0xc, XAdvance:0xe, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0x1f, 0x83, 0xc, 0x70, 0xee, 0x7, 0xe0, 0x7e, 0x7, 0xe0, 0x7e, 0x7, 0xe0, 0x77, 0xe, 0x30, 0xc1, 0xf8}},
		/* P */ tinyfont.Glyph{Width:0xa, Height:0xc, XAdvance:0xb, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0xff, 0x1c, 0xe7, 0x1d, 0xc7, 0x71, 0xdc, 0xe7, 0xf1, 0xc0, 0x70, 0x1c, 0x7, 0x3, 0xe0}},
		/* Q */ tinyfont.Glyph{Width:0xc, Height:0xf, XAdvance:0xe, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0xf, 0x83, 0x9c, 0x70, 0xe6, 0x6, 0xe0, 0x7e, 0x7, 0xe0, 0x7e, 0x7, 0xe0, 0x76, 0x6, 0x30, 0xc1, 0x98, 0xf, 0x0, 0x78, 0x3, 0xe0}},
		/* R */ tinyfont.Glyph{Width:0xc, Height:0xc, XAdvance:0xd, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0xff, 0x7, 0x38, 0x71, 0xc7, 0x1c, 0x71, 0xc7, 0x38, 0x7e, 0x7, 0x70, 0x77, 0x87, 0x3c, 0x71, 0xef, 0x8f}},
		/* S */ tinyfont.Glyph{Width:0x8, Height:0xc, XAdvance:0xa, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0x39, 0x47, 0xc1, 0xc0, 0xf0, 0x7c, 0x3e, 0xf, 0x83, 0xc3, 0xc6, 0xbc}},
		/* T */ tinyfont.Glyph{Width:0xc, Height:0xc, XAdvance:0xc, XOffset:0, YOffset:-11, Bitmaps:[]uint8{0xff, 0xfc, 0xe3, 0x8e, 0x10, 0xe0, 0xe, 0x0, 0xe0, 0xe, 0x0, 0xe0, 0xe, 0x0, 0xe0, 0xe, 0x1, 0xf0}},
		/* U */ tinyfont.Glyph{Width:0xb, Height:0xc, XAdvance:0xd, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0xf8, 0xee, 0x9, 0xc1, 0x38, 0x27, 0x4, 0xe0, 0x9c, 0x13, 0x82, 0x70, 0x4e, 0x8, 0xe2, 0xf, 0x80}},
		/* V */ tinyfont.Glyph{Width:0xd, Height:0xd, XAdvance:0xd, XOffset:0, YOffset:-11, Bitmaps:[]uint8{0xfc, 0x7b, 0xc1, 0xe, 0x8, 0x70, 0x81, 0xc4, 0xe, 0x20, 0x7a, 0x1, 0xd0, 0xe, 0x80, 0x38, 0x1, 0xc0, 0x4, 0x0, 0x20, 0x0}},
		/* W */ tinyfont.Glyph{Width:0x12, Height:0xc, XAdvance:0x12, XOffset:0, YOffset:-11, Bitmaps:[]uint8{0xfd, 0xfb, 0xdc, 0x38, 0x43, 0x87, 0x10, 0xe1, 0xc4, 0x38, 0xf2, 0x7, 0x2e, 0x81, 0xd3, 0xa0, 0x34, 0x70, 0xe, 0x1c, 0x3, 0x87, 0x0, 0x60, 0x80, 0x10, 0x20}},
		/* X */ tinyfont.Glyph{Width:0xd, Height:0xc, XAdvance:0xd, XOffset:0, YOffset:-11, Bitmaps:[]uint8{0xfe, 0xf3, 0xc3, 0xf, 0x10, 0x39, 0x0, 0xf0, 0x3, 0x80, 0x1e, 0x1, 0x70, 0x9, 0xc0, 0x8f, 0x8, 0x3d, 0xf3, 0xf0}},
		/* Y */ tinyfont.Glyph{Width:0xd, Height:0xc, XAdvance:0xd, XOffset:0, YOffset:-11, Bitmaps:[]uint8{0xfc, 0x7b, 0xc1, 0x8e, 0x8, 0x38, 0x81, 0xe8, 0x7, 0x40, 0x1c, 0x0, 0xe0, 0x7, 0x0, 0x38, 0x1, 0xc0, 0x1f, 0x0}},
		/* Z */ tinyfont.Glyph{Width:0xb, Height:0xc, XAdvance:0xc, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0xff, 0xd8, 0x72, 0x1e, 0x43, 0x80, 0xe0, 0x1c, 0x7, 0x1, 0xc0, 0x38, 0x2e, 0xf, 0x83, 0x7f, 0xe0}},
		/* [ */ tinyfont.Glyph{Width:0x4, Height:0xf, XAdvance:0x6, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0xfc, 0xcc, 0xcc, 0xcc, 0xcc, 0xcc, 0xcc, 0xf0}},
		/* \ */ tinyfont.Glyph{Width:0x6, Height:0xd, XAdvance:0x5, XOffset:0, YOffset:-11, Bitmaps:[]uint8{0xc1, 0x6, 0x18, 0x20, 0xc3, 0x4, 0x18, 0x60, 0x83, 0xc}},
		/* ] */ tinyfont.Glyph{Width:0x4, Height:0xf, XAdvance:0x6, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0xf3, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0xf0}},
		/* ^ */ tinyfont.Glyph{Width:0x8, Height:0x7, XAdvance:0xa, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0x18, 0x1c, 0x34, 0x26, 0x62, 0x43, 0xc1}},
		/* _ */ tinyfont.Glyph{Width:0x9, Height:0x1, XAdvance:0x9, XOffset:0, YOffset:3, Bitmaps:[]uint8{0xff, 0x80}},
		/* ` */ tinyfont.Glyph{Width:0x4, Height:0x3, XAdvance:0x6, XOffset:0, YOffset:-12, Bitmaps:[]uint8{0xc6, 0x30}},
		/* a */ tinyfont.Glyph{Width:0x9, Height:0x9, XAdvance:0x9, XOffset:0, YOffset:-8, Bitmaps:[]uint8{0x7c, 0x63, 0xb1, 0xc0, 0xe1, 0xf3, 0x3b, 0x9d, 0xce, 0xff, 0x80}},
		/* b */ tinyfont.Glyph{Width:0xa, Height:0xc, XAdvance:0xa, XOffset:0, YOffset:-11, Bitmaps:[]uint8{0xf0, 0x1c, 0x7, 0x1, 0xdc, 0x7b, 0x9c, 0x77, 0x1d, 0xc7, 0x71, 0xdc, 0x77, 0x39, 0x3c}},
		/* c */ tinyfont.Glyph{Width:0x7, Height:0x9, XAdvance:0x8, XOffset:0, YOffset:-8, Bitmaps:[]uint8{0x3c, 0xed, 0x9f, 0xe, 0x1c, 0x38, 0x39, 0x3c}},
		/* d */ tinyfont.Glyph{Width:0xa, Height:0xc, XAdvance:0xa, XOffset:0, YOffset:-11, Bitmaps:[]uint8{0x7, 0x80, 0xe0, 0x38, 0xee, 0x77, 0xb8, 0xee, 0x3b, 0x8e, 0xe3, 0xb8, 0xe7, 0x78, 0xef}},
		/* e */ tinyfont.Glyph{Width:0x8, Height:0x9, XAdvance:0x8, XOffset:0, YOffset:-8, Bitmaps:[]uint8{0x3c, 0x66, 0xe6, 0xfe, 0xe0, 0xe0, 0xe0, 0x72, 0x3c}},
		/* f */ tinyfont.Glyph{Width:0x7, Height:0xc, XAdvance:0x7, XOffset:0, YOffset:-11, Bitmaps:[]uint8{0x3e, 0xed, 0xc7, 0xc7, 0xe, 0x1c, 0x38, 0x70, 0xe1, 0xc7, 0xc0}},
		/* g */ tinyfont.Glyph{Width:0x7, Height:0xd, XAdvance:0x9, XOffset:1, YOffset:-8, Bitmaps:[]uint8{0x31, 0xdf, 0xbf, 0x7e, 0xe7, 0x90, 0x60, 0xfc, 0xfe, 0xc, 0x17, 0xc0}},
		/* h */ tinyfont.Glyph{Width:0xa, Height:0xc, XAdvance:0xa, XOffset:0, YOffset:-11, Bitmaps:[]uint8{0xf0, 0x1c, 0x7, 0x1, 0xdc, 0x7b, 0x9c, 0xe7, 0x39, 0xce, 0x73, 0x9c, 0xe7, 0x3b, 0xff}},
		/* i */ tinyfont.Glyph{Width:0x5, Height:0xc, XAdvance:0x5, XOffset:0, YOffset:-11, Bitmaps:[]uint8{0x73, 0x9d, 0xe7, 0x39, 0xce, 0x73, 0x9d, 0xf0}},
		/* j */ tinyfont.Glyph{Width:0x6, Height:0x10, XAdvance:0x7, XOffset:0, YOffset:-11, Bitmaps:[]uint8{0x1c, 0x71, 0xcf, 0x1c, 0x71, 0xc7, 0x1c, 0x71, 0xc7, 0x1c, 0x7d, 0xbe}},
		/* k */ tinyfont.Glyph{Width:0xa, Height:0xc, XAdvance:0xa, XOffset:0, YOffset:-11, Bitmaps:[]uint8{0xf0, 0x1c, 0x7, 0x1, 0xce, 0x71, 0x1c, 0x87, 0x41, 0xf8, 0x77, 0x1c, 0xe7, 0x1b, 0xef}},
		/* l */ tinyfont.Glyph{Width:0x5, Height:0xc, XAdvance:0x5, XOffset:0, YOffset:-11, Bitmaps:[]uint8{0xf3, 0x9c, 0xe7, 0x39, 0xce, 0x73, 0x9d, 0xf0}},
		/* m */ tinyfont.Glyph{Width:0xf, Height:0x9, XAdvance:0xf, XOffset:0, YOffset:-8, Bitmaps:[]uint8{0xf7, 0x38, 0xf7, 0xb9, 0xce, 0x73, 0x9c, 0xe7, 0x39, 0xce, 0x73, 0x9c, 0xe7, 0x39, 0xce, 0xff, 0xfe}},
		/* n */ tinyfont.Glyph{Width:0xa, Height:0x9, XAdvance:0xa, XOffset:0, YOffset:-8, Bitmaps:[]uint8{0xf7, 0x1e, 0xe7, 0x39, 0xce, 0x73, 0x9c, 0xe7, 0x39, 0xce, 0xff, 0xc0}},
		/* o */ tinyfont.Glyph{Width:0x9, Height:0x9, XAdvance:0x9, XOffset:0, YOffset:-8, Bitmaps:[]uint8{0x3e, 0x31, 0xb8, 0xfc, 0x7e, 0x3f, 0x1f, 0x8e, 0xc6, 0x3e, 0x0}},
		/* p */ tinyfont.Glyph{Width:0xa, Height:0xd, XAdvance:0xa, XOffset:0, YOffset:-8, Bitmaps:[]uint8{0xf7, 0x1e, 0xe7, 0x1d, 0xc7, 0x71, 0xdc, 0x77, 0x1d, 0xce, 0x7f, 0x1c, 0x7, 0x1, 0xc0, 0xf8, 0x0}},
		/* q */ tinyfont.Glyph{Width:0xa, Height:0xd, XAdvance:0xa, XOffset:0, YOffset:-8, Bitmaps:[]uint8{0x3c, 0x9c, 0xee, 0x3b, 0x8e, 0xe3, 0xb8, 0xee, 0x39, 0xce, 0x3f, 0x80, 0xe0, 0x38, 0xe, 0x7, 0xc0}},
		/* r */ tinyfont.Glyph{Width:0x8, Height:0x9, XAdvance:0x8, XOffset:0, YOffset:-8, Bitmaps:[]uint8{0xf7, 0x7b, 0x70, 0x70, 0x70, 0x70, 0x70, 0x70, 0xf8}},
		/* s */ tinyfont.Glyph{Width:0x5, Height:0x9, XAdvance:0x7, XOffset:1, YOffset:-8, Bitmaps:[]uint8{0x7e, 0x73, 0xc7, 0x8e, 0x39, 0xb0}},
		/* t */ tinyfont.Glyph{Width:0x6, Height:0xb, XAdvance:0x6, XOffset:0, YOffset:-10, Bitmaps:[]uint8{0x10, 0xcf, 0x9c, 0x71, 0xc7, 0x1c, 0x71, 0xd3, 0x80}},
		/* u */ tinyfont.Glyph{Width:0xa, Height:0x9, XAdvance:0xa, XOffset:0, YOffset:-8, Bitmaps:[]uint8{0xf7, 0x9c, 0xe7, 0x39, 0xce, 0x73, 0x9c, 0xe7, 0x39, 0xce, 0x3f, 0xc0}},
		/* v */ tinyfont.Glyph{Width:0x9, Height:0x9, XAdvance:0x9, XOffset:0, YOffset:-8, Bitmaps:[]uint8{0xfb, 0xb8, 0x8c, 0x87, 0x43, 0xc0, 0xe0, 0x70, 0x10, 0x8, 0x0}},
		/* w */ tinyfont.Glyph{Width:0xc, Height:0x9, XAdvance:0xd, XOffset:0, YOffset:-8, Bitmaps:[]uint8{0xf7, 0xb6, 0x31, 0x73, 0xa3, 0x3a, 0x3d, 0xa3, 0xdc, 0x18, 0xc1, 0x88, 0x10, 0x80}},
		/* x */ tinyfont.Glyph{Width:0x9, Height:0x9, XAdvance:0x9, XOffset:0, YOffset:-8, Bitmaps:[]uint8{0xfb, 0xb8, 0x8e, 0x83, 0x81, 0xc0, 0xf0, 0x98, 0xce, 0xef, 0x80}},
		/* y */ tinyfont.Glyph{Width:0x8, Height:0xd, XAdvance:0x9, XOffset:0, YOffset:-8, Bitmaps:[]uint8{0xf7, 0x62, 0x72, 0x34, 0x34, 0x3c, 0x18, 0x18, 0x10, 0x10, 0x10, 0xe0, 0xe0}},
		/* z */ tinyfont.Glyph{Width:0x7, Height:0x9, XAdvance:0x8, XOffset:1, YOffset:-8, Bitmaps:[]uint8{0xff, 0x1c, 0x70, 0xe3, 0x87, 0x1c, 0x71, 0xfe}},
		/* { */ tinyfont.Glyph{Width:0x5, Height:0x10, XAdvance:0x7, XOffset:0, YOffset:-12, Bitmaps:[]uint8{0x19, 0x8c, 0x63, 0x18, 0xcc, 0x61, 0x8c, 0x63, 0x18, 0xc3}},
		/* | */ tinyfont.Glyph{Width:0x1, Height:0xd, XAdvance:0x4, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0xff, 0xf8}},
		/* } */ tinyfont.Glyph{Width:0x5, Height:0x10, XAdvance:0x7, XOffset:2, YOffset:-12, Bitmaps:[]uint8{0xc3, 0x18, 0xc6, 0x31, 0x86, 0x33, 0x18, 0xc6, 0x31, 0x98}},
		/* ~ */ tinyfont.Glyph{Width:0x8, Height:0x2, XAdvance:0x9, XOffset:1, YOffset:-4, Bitmaps:[]uint8{0xf0, 0x8e}},
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

	YAdvance:0x16,
}
