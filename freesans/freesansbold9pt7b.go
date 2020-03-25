package freesans

import (
	"tinygo.org/x/tinyfont"
)

var Bold9pt7b = tinyfont.Font{
	Glyphs:[]tinyfont.Glyph{
		/*   */ tinyfont.Glyph{Rune:32, Width:0x0, Height:0x0, XAdvance:0x5, XOffset:0, YOffset:1, Bitmaps:[]uint8{}},
		/* ! */ tinyfont.Glyph{Rune:33, Width:0x3, Height:0xd, XAdvance:0x6, XOffset:2, YOffset:-12, Bitmaps:[]uint8{0xff, 0xff, 0xfe, 0x48, 0x7e}},
		/* " */ tinyfont.Glyph{Rune:34, Width:0x7, Height:0x5, XAdvance:0x9, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xef, 0xdf, 0xbf, 0x74, 0x40}},
		/* # */ tinyfont.Glyph{Rune:35, Width:0xa, Height:0xc, XAdvance:0xa, XOffset:0, YOffset:-11, Bitmaps:[]uint8{0x19, 0x86, 0x67, 0xfd, 0xff, 0x33, 0xc, 0xc3, 0x33, 0xfe, 0xff, 0x99, 0x86, 0x61, 0x90}},
		/* $ */ tinyfont.Glyph{Rune:36, Width:0x9, Height:0xf, XAdvance:0xa, XOffset:1, YOffset:-13, Bitmaps:[]uint8{0x10, 0x1f, 0x1f, 0xde, 0xff, 0x3f, 0x83, 0xc0, 0xfc, 0x1f, 0x9, 0xfc, 0xfe, 0xf7, 0xf1, 0xe0, 0x40}},
		/* % */ tinyfont.Glyph{Rune:37, Width:0x10, Height:0xd, XAdvance:0x10, XOffset:0, YOffset:-12, Bitmaps:[]uint8{0x38, 0x10, 0x7c, 0x30, 0xc6, 0x20, 0xc6, 0x40, 0xc6, 0x40, 0x7c, 0x80, 0x39, 0x9c, 0x1, 0x3e, 0x3, 0x63, 0x2, 0x63, 0x4, 0x63, 0xc, 0x3e, 0x8, 0x1c}},
		/* & */ tinyfont.Glyph{Rune:38, Width:0xc, Height:0xd, XAdvance:0xd, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xe, 0x1, 0xf8, 0x3b, 0x83, 0xb8, 0x3f, 0x1, 0xe0, 0x3e, 0x67, 0x76, 0xe3, 0xee, 0x1c, 0xf3, 0xc7, 0xfe, 0x3f, 0x70}},
		/* ' */ tinyfont.Glyph{Rune:39, Width:0x3, Height:0x5, XAdvance:0x5, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xff, 0xf4}},
		/* ( */ tinyfont.Glyph{Rune:40, Width:0x6, Height:0x11, XAdvance:0x6, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0x18, 0x63, 0x1c, 0x73, 0x8e, 0x38, 0xe3, 0x8e, 0x18, 0x70, 0xc3, 0x6, 0x8}},
		/* ) */ tinyfont.Glyph{Rune:41, Width:0x6, Height:0x11, XAdvance:0x6, XOffset:0, YOffset:-12, Bitmaps:[]uint8{0x61, 0x83, 0xe, 0x38, 0x71, 0xc7, 0x1c, 0x71, 0xc6, 0x38, 0xe3, 0x18, 0x40}},
		/* * */ tinyfont.Glyph{Rune:42, Width:0x5, Height:0x6, XAdvance:0x7, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0x21, 0x3e, 0x45, 0x28}},
		/* + */ tinyfont.Glyph{Rune:43, Width:0x7, Height:0x8, XAdvance:0xb, XOffset:2, YOffset:-7, Bitmaps:[]uint8{0x38, 0x70, 0xe7, 0xff, 0xe7, 0xe, 0x1c}},
		/* , */ tinyfont.Glyph{Rune:44, Width:0x3, Height:0x5, XAdvance:0x4, XOffset:1, YOffset:-1, Bitmaps:[]uint8{0xfc, 0x9c}},
		/* - */ tinyfont.Glyph{Rune:45, Width:0x5, Height:0x2, XAdvance:0x6, XOffset:0, YOffset:-5, Bitmaps:[]uint8{0xff, 0xc0}},
		/* . */ tinyfont.Glyph{Rune:46, Width:0x3, Height:0x2, XAdvance:0x4, XOffset:1, YOffset:-1, Bitmaps:[]uint8{0xfc}},
		/* / */ tinyfont.Glyph{Rune:47, Width:0x5, Height:0xd, XAdvance:0x5, XOffset:0, YOffset:-12, Bitmaps:[]uint8{0x8, 0xc4, 0x23, 0x10, 0x84, 0x62, 0x11, 0x88, 0x0}},
		/* 0 */ tinyfont.Glyph{Rune:48, Width:0x9, Height:0xd, XAdvance:0xa, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0x3e, 0x3f, 0x9d, 0xdc, 0x7e, 0x3f, 0x1f, 0x8f, 0xc7, 0xe3, 0xf1, 0xdd, 0xcf, 0xe3, 0xe0}},
		/* 1 */ tinyfont.Glyph{Rune:49, Width:0x5, Height:0xd, XAdvance:0xa, XOffset:2, YOffset:-12, Bitmaps:[]uint8{0x8, 0xff, 0xf3, 0x9c, 0xe7, 0x39, 0xce, 0x73, 0x80}},
		/* 2 */ tinyfont.Glyph{Rune:50, Width:0x9, Height:0xd, XAdvance:0xa, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0x3e, 0x3f, 0xb8, 0xfc, 0x70, 0x38, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0xf, 0xf7, 0xf8}},
		/* 3 */ tinyfont.Glyph{Rune:51, Width:0x8, Height:0xd, XAdvance:0xa, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0x3c, 0x7f, 0xe7, 0xe7, 0x7, 0xc, 0xe, 0x7, 0x7, 0xe7, 0xe7, 0x7e, 0x3c}},
		/* 4 */ tinyfont.Glyph{Rune:52, Width:0x8, Height:0xd, XAdvance:0xa, XOffset:2, YOffset:-12, Bitmaps:[]uint8{0xe, 0x1e, 0x1e, 0x2e, 0x2e, 0x4e, 0x4e, 0x8e, 0xff, 0xff, 0xe, 0xe, 0xe}},
		/* 5 */ tinyfont.Glyph{Rune:53, Width:0x9, Height:0xd, XAdvance:0xa, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0x7f, 0x3f, 0x90, 0x18, 0xd, 0xe7, 0xfb, 0x9e, 0x7, 0x3, 0x81, 0xf1, 0xff, 0xe7, 0xc0}},
		/* 6 */ tinyfont.Glyph{Rune:54, Width:0x9, Height:0xd, XAdvance:0xa, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0x3e, 0x3f, 0x9c, 0xfc, 0xe, 0xe7, 0xfb, 0xdf, 0xc7, 0xe3, 0xf1, 0xdd, 0xef, 0xe3, 0xe0}},
		/* 7 */ tinyfont.Glyph{Rune:55, Width:0x9, Height:0xd, XAdvance:0xa, XOffset:0, YOffset:-12, Bitmaps:[]uint8{0xff, 0xff, 0xc0, 0xe0, 0xe0, 0x60, 0x70, 0x30, 0x38, 0x1c, 0xc, 0xe, 0x7, 0x3, 0x80}},
		/* 8 */ tinyfont.Glyph{Rune:56, Width:0xa, Height:0xd, XAdvance:0xa, XOffset:0, YOffset:-12, Bitmaps:[]uint8{0x3f, 0x1f, 0xee, 0x3f, 0x87, 0xe3, 0xcf, 0xc7, 0xfb, 0xcf, 0xe1, 0xf8, 0x7f, 0x3d, 0xfe, 0x3f, 0x0}},
		/* 9 */ tinyfont.Glyph{Rune:57, Width:0x9, Height:0xd, XAdvance:0xa, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0x3e, 0x3f, 0xbd, 0xdc, 0x7e, 0x3f, 0x1f, 0xde, 0xff, 0x3b, 0x81, 0xf9, 0xcf, 0xe3, 0xc0}},
		/* : */ tinyfont.Glyph{Rune:58, Width:0x3, Height:0x9, XAdvance:0x4, XOffset:1, YOffset:-8, Bitmaps:[]uint8{0xfc, 0x0, 0x7, 0xe0}},
		/* ; */ tinyfont.Glyph{Rune:59, Width:0x3, Height:0xc, XAdvance:0x4, XOffset:1, YOffset:-8, Bitmaps:[]uint8{0xfc, 0x0, 0x7, 0xe5, 0xe0}},
		/* < */ tinyfont.Glyph{Rune:60, Width:0x9, Height:0x9, XAdvance:0xb, XOffset:1, YOffset:-8, Bitmaps:[]uint8{0x0, 0x83, 0xc7, 0xdf, 0xc, 0x7, 0x80, 0xf8, 0x1f, 0x1, 0x80}},
		/* = */ tinyfont.Glyph{Rune:61, Width:0x9, Height:0x6, XAdvance:0xb, XOffset:1, YOffset:-6, Bitmaps:[]uint8{0xff, 0xff, 0xc0, 0x0, 0xf, 0xff, 0xfc}},
		/* > */ tinyfont.Glyph{Rune:62, Width:0x9, Height:0x9, XAdvance:0xb, XOffset:1, YOffset:-8, Bitmaps:[]uint8{0x0, 0x70, 0x3f, 0x3, 0xe0, 0x38, 0x7d, 0xf1, 0xe0, 0x80, 0x0}},
		/* ? */ tinyfont.Glyph{Rune:63, Width:0x9, Height:0xd, XAdvance:0xb, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0x3e, 0x3f, 0xb8, 0xfc, 0x70, 0x38, 0x1c, 0x1c, 0x1c, 0x1c, 0xe, 0x0, 0x3, 0x81, 0xc0}},
		/* @ */ tinyfont.Glyph{Rune:64, Width:0x10, Height:0xf, XAdvance:0x12, XOffset:0, YOffset:-12, Bitmaps:[]uint8{0x3, 0xf0, 0xf, 0xfc, 0x1e, 0xe, 0x38, 0x2, 0x70, 0xe9, 0x63, 0x19, 0xc2, 0x19, 0xc6, 0x11, 0xc6, 0x33, 0xc6, 0x32, 0x63, 0xfe, 0x73, 0xdc, 0x3c, 0x0, 0x1f, 0xf8, 0x7, 0xf0}},
		/* A */ tinyfont.Glyph{Rune:65, Width:0xc, Height:0xd, XAdvance:0xd, XOffset:0, YOffset:-12, Bitmaps:[]uint8{0x7, 0x0, 0xf0, 0xf, 0x80, 0xf8, 0x1d, 0x81, 0x9c, 0x19, 0xc3, 0x8c, 0x3f, 0xe7, 0xfe, 0x70, 0x66, 0x7, 0xe0, 0x70}},
		/* B */ tinyfont.Glyph{Rune:66, Width:0xb, Height:0xd, XAdvance:0xd, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xff, 0x9f, 0xfb, 0x83, 0xf0, 0x7e, 0xf, 0xff, 0x3f, 0xf7, 0x6, 0xe0, 0xfc, 0x1f, 0x83, 0xff, 0xef, 0xf8}},
		/* C */ tinyfont.Glyph{Rune:67, Width:0xc, Height:0xd, XAdvance:0xd, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0x1f, 0x83, 0xfe, 0x78, 0xe7, 0x7, 0xe0, 0xe, 0x0, 0xe0, 0xe, 0x0, 0xe0, 0x7, 0x7, 0x78, 0xf3, 0xfe, 0x1f, 0x80}},
		/* D */ tinyfont.Glyph{Rune:68, Width:0xc, Height:0xd, XAdvance:0xd, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xff, 0x8f, 0xfc, 0xe0, 0xee, 0xe, 0xe0, 0x7e, 0x7, 0xe0, 0x7e, 0x7, 0xe0, 0x7e, 0xe, 0xe0, 0xef, 0xfc, 0xff, 0x80}},
		/* E */ tinyfont.Glyph{Rune:69, Width:0x9, Height:0xd, XAdvance:0xc, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xff, 0xff, 0xf8, 0x1c, 0xe, 0x7, 0xfb, 0xfd, 0xc0, 0xe0, 0x70, 0x38, 0x1f, 0xff, 0xf8}},
		/* F */ tinyfont.Glyph{Rune:70, Width:0x9, Height:0xd, XAdvance:0xb, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xff, 0xff, 0xf8, 0x1c, 0xe, 0x7, 0xfb, 0xfd, 0xc0, 0xe0, 0x70, 0x38, 0x1c, 0xe, 0x0}},
		/* G */ tinyfont.Glyph{Rune:71, Width:0xb, Height:0xd, XAdvance:0xe, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xf, 0x87, 0xf9, 0xe3, 0xb8, 0x3e, 0x1, 0xc0, 0x38, 0xff, 0x1f, 0xe0, 0x6e, 0xd, 0xe3, 0x9f, 0xd0, 0xf2}},
		/* H */ tinyfont.Glyph{Rune:72, Width:0xb, Height:0xd, XAdvance:0xd, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xe0, 0xfc, 0x1f, 0x83, 0xf0, 0x7e, 0xf, 0xff, 0xff, 0xff, 0x7, 0xe0, 0xfc, 0x1f, 0x83, 0xf0, 0x7e, 0xe}},
		/* I */ tinyfont.Glyph{Rune:73, Width:0x3, Height:0xd, XAdvance:0x6, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xff, 0xff, 0xff, 0xff, 0xfe}},
		/* J */ tinyfont.Glyph{Rune:74, Width:0x8, Height:0xd, XAdvance:0xa, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0x7, 0x7, 0x7, 0x7, 0x7, 0x7, 0x7, 0x7, 0xe7, 0xe7, 0xe7, 0x7e, 0x3c}},
		/* K */ tinyfont.Glyph{Rune:75, Width:0xc, Height:0xd, XAdvance:0xd, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xe0, 0xee, 0x1c, 0xe3, 0x8e, 0x70, 0xee, 0xf, 0xc0, 0xfe, 0xf, 0x70, 0xe7, 0xe, 0x38, 0xe1, 0xce, 0xe, 0xe0, 0xe0}},
		/* L */ tinyfont.Glyph{Rune:76, Width:0x8, Height:0xd, XAdvance:0xb, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xe0, 0xe0, 0xe0, 0xe0, 0xe0, 0xe0, 0xe0, 0xe0, 0xe0, 0xe0, 0xe0, 0xff, 0xff}},
		/* M */ tinyfont.Glyph{Rune:77, Width:0xe, Height:0xd, XAdvance:0x10, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xf8, 0x7f, 0xe1, 0xff, 0x87, 0xfe, 0x1f, 0xec, 0x7f, 0xb3, 0x7e, 0xcd, 0xfb, 0x37, 0xec, 0xdf, 0x9e, 0x7e, 0x79, 0xf9, 0xe7, 0xe7, 0x9c}},
		/* N */ tinyfont.Glyph{Rune:78, Width:0xb, Height:0xd, XAdvance:0xe, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xe0, 0xfe, 0x1f, 0xc3, 0xfc, 0x7f, 0xcf, 0xd9, 0xfb, 0xbf, 0x37, 0xe7, 0xfc, 0x7f, 0x87, 0xf0, 0xfe, 0xe}},
		/* O */ tinyfont.Glyph{Rune:79, Width:0xd, Height:0xd, XAdvance:0xe, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xf, 0x81, 0xff, 0x1e, 0x3c, 0xe0, 0xee, 0x3, 0xf0, 0x1f, 0x80, 0xfc, 0x7, 0xe0, 0x3b, 0x83, 0x9e, 0x3c, 0x7f, 0xc0, 0xf8, 0x0}},
		/* P */ tinyfont.Glyph{Rune:80, Width:0xb, Height:0xd, XAdvance:0xc, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xff, 0x9f, 0xfb, 0x87, 0xf0, 0x7e, 0xf, 0xc3, 0xff, 0xf7, 0xfc, 0xe0, 0x1c, 0x3, 0x80, 0x70, 0xe, 0x0}},
		/* Q */ tinyfont.Glyph{Rune:81, Width:0xd, Height:0xe, XAdvance:0xe, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xf, 0x81, 0xff, 0x1e, 0x3c, 0xe0, 0xee, 0x3, 0xf0, 0x1f, 0x80, 0xfc, 0x7, 0xe1, 0xbb, 0x8f, 0x9e, 0x3c, 0x7f, 0xe0, 0xfb, 0x80, 0x8}},
		/* R */ tinyfont.Glyph{Rune:82, Width:0xc, Height:0xd, XAdvance:0xd, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xff, 0x8f, 0xfc, 0xe0, 0xee, 0xe, 0xe0, 0xee, 0xe, 0xff, 0xcf, 0xfc, 0xe0, 0xee, 0xe, 0xe0, 0xee, 0xe, 0xe0, 0xf0}},
		/* S */ tinyfont.Glyph{Rune:83, Width:0xb, Height:0xd, XAdvance:0xc, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0x3f, 0xf, 0xfb, 0xc7, 0xf0, 0x7e, 0x1, 0xfc, 0x1f, 0xf0, 0x3f, 0x0, 0xfc, 0x1d, 0xc7, 0xbf, 0xe1, 0xf8}},
		/* T */ tinyfont.Glyph{Rune:84, Width:0x9, Height:0xd, XAdvance:0xc, XOffset:2, YOffset:-12, Bitmaps:[]uint8{0xff, 0xff, 0xc7, 0x3, 0x81, 0xc0, 0xe0, 0x70, 0x38, 0x1c, 0xe, 0x7, 0x3, 0x81, 0xc0}},
		/* U */ tinyfont.Glyph{Rune:85, Width:0xb, Height:0xd, XAdvance:0xd, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xe0, 0xfc, 0x1f, 0x83, 0xf0, 0x7e, 0xf, 0xc1, 0xf8, 0x3f, 0x7, 0xe0, 0xfc, 0x1f, 0xc7, 0xbf, 0xe1, 0xf0}},
		/* V */ tinyfont.Glyph{Rune:86, Width:0xc, Height:0xd, XAdvance:0xc, XOffset:0, YOffset:-12, Bitmaps:[]uint8{0x60, 0x67, 0xe, 0x70, 0xe3, 0xc, 0x30, 0xc3, 0x9c, 0x19, 0x81, 0x98, 0x1f, 0x80, 0xf0, 0xf, 0x0, 0xf0, 0x6, 0x0}},
		/* W */ tinyfont.Glyph{Rune:87, Width:0x11, Height:0xd, XAdvance:0x11, XOffset:0, YOffset:-12, Bitmaps:[]uint8{0x61, 0xc3, 0xb8, 0xe1, 0x9c, 0x70, 0xce, 0x3c, 0xe3, 0x36, 0x71, 0x9b, 0x30, 0xed, 0x98, 0x36, 0x7c, 0x1b, 0x3c, 0xf, 0x1e, 0x7, 0x8f, 0x1, 0xc3, 0x80, 0xe1, 0x80}},
		/* X */ tinyfont.Glyph{Rune:88, Width:0xc, Height:0xd, XAdvance:0xc, XOffset:0, YOffset:-12, Bitmaps:[]uint8{0x70, 0xe7, 0x8e, 0x39, 0xc1, 0xf8, 0x1f, 0x80, 0xf0, 0x7, 0x0, 0xf0, 0x1f, 0x81, 0x9c, 0x39, 0xc7, 0xe, 0x70, 0xe0}},
		/* Y */ tinyfont.Glyph{Rune:89, Width:0xb, Height:0xd, XAdvance:0xc, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xe0, 0xfc, 0x39, 0xc7, 0x18, 0xc3, 0xb8, 0x36, 0x7, 0xc0, 0x70, 0xe, 0x1, 0xc0, 0x38, 0x7, 0x0, 0xe0}},
		/* Z */ tinyfont.Glyph{Rune:90, Width:0x9, Height:0xd, XAdvance:0xb, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xff, 0xff, 0xc0, 0xe0, 0xe0, 0xf0, 0x70, 0x70, 0x70, 0x78, 0x38, 0x38, 0x1f, 0xff, 0xf8}},
		/* [ */ tinyfont.Glyph{Rune:91, Width:0x4, Height:0x11, XAdvance:0x6, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xff, 0xee, 0xee, 0xee, 0xee, 0xee, 0xee, 0xef, 0xf0}},
		/* \ */ tinyfont.Glyph{Rune:92, Width:0x5, Height:0xd, XAdvance:0x5, XOffset:0, YOffset:-12, Bitmaps:[]uint8{0x86, 0x10, 0x86, 0x10, 0x84, 0x30, 0x84, 0x30, 0x80}},
		/* ] */ tinyfont.Glyph{Rune:93, Width:0x4, Height:0x11, XAdvance:0x6, XOffset:0, YOffset:-12, Bitmaps:[]uint8{0xff, 0x77, 0x77, 0x77, 0x77, 0x77, 0x77, 0x7f, 0xf0}},
		/* ^ */ tinyfont.Glyph{Rune:94, Width:0x8, Height:0x8, XAdvance:0xb, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0x18, 0x1c, 0x3c, 0x3e, 0x36, 0x66, 0x63, 0xc3}},
		/* _ */ tinyfont.Glyph{Rune:95, Width:0xa, Height:0x1, XAdvance:0xa, XOffset:0, YOffset:4, Bitmaps:[]uint8{0xff, 0xc0}},
		/* ` */ tinyfont.Glyph{Rune:96, Width:0x3, Height:0x2, XAdvance:0x5, XOffset:0, YOffset:-12, Bitmaps:[]uint8{0xcc}},
		/* a */ tinyfont.Glyph{Rune:97, Width:0xa, Height:0xa, XAdvance:0xa, XOffset:1, YOffset:-9, Bitmaps:[]uint8{0x3f, 0x1f, 0xee, 0x38, 0xe, 0x3f, 0x9e, 0xee, 0x3b, 0x9e, 0xff, 0x9e, 0xe0}},
		/* b */ tinyfont.Glyph{Rune:98, Width:0xa, Height:0xd, XAdvance:0xb, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xe0, 0x38, 0xe, 0x3, 0xbc, 0xff, 0xbc, 0xee, 0x1f, 0x87, 0xe1, 0xf8, 0x7f, 0x3b, 0xfe, 0xef, 0x0}},
		/* c */ tinyfont.Glyph{Rune:99, Width:0x9, Height:0xa, XAdvance:0xa, XOffset:1, YOffset:-9, Bitmaps:[]uint8{0x1f, 0x3f, 0xdc, 0x7c, 0xe, 0x7, 0x3, 0x80, 0xe3, 0x7f, 0x8f, 0x0}},
		/* d */ tinyfont.Glyph{Rune:100, Width:0x9, Height:0xd, XAdvance:0xb, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0x3, 0x81, 0xc0, 0xe7, 0x77, 0xfb, 0xbf, 0x8f, 0xc7, 0xe3, 0xf1, 0xfd, 0xef, 0xf3, 0xb8}},
		/* e */ tinyfont.Glyph{Rune:101, Width:0x9, Height:0xa, XAdvance:0xa, XOffset:1, YOffset:-9, Bitmaps:[]uint8{0x3e, 0x3f, 0x9c, 0xdc, 0x3f, 0xff, 0xff, 0x81, 0xc3, 0x7f, 0x8f, 0x0}},
		/* f */ tinyfont.Glyph{Rune:102, Width:0x5, Height:0xd, XAdvance:0x6, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0x3b, 0xdd, 0xff, 0xb9, 0xce, 0x73, 0x9c, 0xe7, 0x0}},
		/* g */ tinyfont.Glyph{Rune:103, Width:0x9, Height:0xe, XAdvance:0xb, XOffset:1, YOffset:-9, Bitmaps:[]uint8{0x3b, 0xbf, 0xdd, 0xfc, 0x7e, 0x3f, 0x1f, 0x8f, 0xef, 0x7f, 0x9d, 0xc0, 0xfc, 0x77, 0xf1, 0xf0}},
		/* h */ tinyfont.Glyph{Rune:104, Width:0x9, Height:0xd, XAdvance:0xb, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xe0, 0x70, 0x38, 0x1d, 0xef, 0xff, 0x9f, 0x8f, 0xc7, 0xe3, 0xf1, 0xf8, 0xfc, 0x7e, 0x38}},
		/* i */ tinyfont.Glyph{Rune:105, Width:0x3, Height:0xd, XAdvance:0x5, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xfc, 0x7f, 0xff, 0xff, 0xfe}},
		/* j */ tinyfont.Glyph{Rune:106, Width:0x4, Height:0x11, XAdvance:0x5, XOffset:0, YOffset:-12, Bitmaps:[]uint8{0x77, 0x7, 0x77, 0x77, 0x77, 0x77, 0x77, 0x7f, 0xe0}},
		/* k */ tinyfont.Glyph{Rune:107, Width:0x9, Height:0xd, XAdvance:0xa, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xe0, 0x70, 0x38, 0x1c, 0x7e, 0x77, 0x73, 0xf1, 0xf8, 0xfe, 0x77, 0x39, 0xdc, 0x6e, 0x38}},
		/* l */ tinyfont.Glyph{Rune:108, Width:0x3, Height:0xd, XAdvance:0x5, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0xff, 0xff, 0xff, 0xff, 0xfe}},
		/* m */ tinyfont.Glyph{Rune:109, Width:0xe, Height:0xa, XAdvance:0x10, XOffset:1, YOffset:-9, Bitmaps:[]uint8{0xef, 0x7b, 0xff, 0xfe, 0x39, 0xf8, 0xe7, 0xe3, 0x9f, 0x8e, 0x7e, 0x39, 0xf8, 0xe7, 0xe3, 0x9f, 0x8e, 0x70}},
		/* n */ tinyfont.Glyph{Rune:110, Width:0x9, Height:0xa, XAdvance:0xb, XOffset:1, YOffset:-9, Bitmaps:[]uint8{0xef, 0x7f, 0xf8, 0xfc, 0x7e, 0x3f, 0x1f, 0x8f, 0xc7, 0xe3, 0xf1, 0xc0}},
		/* o */ tinyfont.Glyph{Rune:111, Width:0xa, Height:0xa, XAdvance:0xb, XOffset:1, YOffset:-9, Bitmaps:[]uint8{0x1e, 0x1f, 0xe7, 0x3b, 0x87, 0xe1, 0xf8, 0x7e, 0x1d, 0xce, 0x7f, 0x87, 0x80}},
		/* p */ tinyfont.Glyph{Rune:112, Width:0xa, Height:0xe, XAdvance:0xb, XOffset:1, YOffset:-9, Bitmaps:[]uint8{0xef, 0x3f, 0xef, 0x3b, 0x87, 0xe1, 0xf8, 0x7e, 0x1f, 0xce, 0xff, 0xbb, 0xce, 0x3, 0x80, 0xe0, 0x38, 0x0}},
		/* q */ tinyfont.Glyph{Rune:113, Width:0x9, Height:0xe, XAdvance:0xb, XOffset:1, YOffset:-9, Bitmaps:[]uint8{0x3b, 0xbf, 0xfd, 0xfc, 0x7e, 0x3f, 0x1f, 0x8f, 0xef, 0x7f, 0x9d, 0xc0, 0xe0, 0x70, 0x38, 0x1c}},
		/* r */ tinyfont.Glyph{Rune:114, Width:0x6, Height:0xa, XAdvance:0x7, XOffset:1, YOffset:-9, Bitmaps:[]uint8{0xef, 0xff, 0x38, 0xe3, 0x8e, 0x38, 0xe3, 0x80}},
		/* s */ tinyfont.Glyph{Rune:115, Width:0x9, Height:0xa, XAdvance:0xa, XOffset:1, YOffset:-9, Bitmaps:[]uint8{0x3e, 0x3f, 0xb8, 0xfc, 0xf, 0xc3, 0xfc, 0x3f, 0xc7, 0xff, 0x1f, 0x0}},
		/* t */ tinyfont.Glyph{Rune:116, Width:0x5, Height:0xc, XAdvance:0x6, XOffset:1, YOffset:-11, Bitmaps:[]uint8{0x73, 0xbf, 0xf7, 0x39, 0xce, 0x73, 0x9e, 0x70}},
		/* u */ tinyfont.Glyph{Rune:117, Width:0x9, Height:0xa, XAdvance:0xb, XOffset:1, YOffset:-9, Bitmaps:[]uint8{0xe3, 0xf1, 0xf8, 0xfc, 0x7e, 0x3f, 0x1f, 0x8f, 0xc7, 0xff, 0xbd, 0xc0}},
		/* v */ tinyfont.Glyph{Rune:118, Width:0xa, Height:0xa, XAdvance:0xa, XOffset:0, YOffset:-9, Bitmaps:[]uint8{0xe1, 0x98, 0x67, 0x39, 0xcc, 0x33, 0xd, 0xc3, 0xe0, 0x78, 0x1e, 0x7, 0x0}},
		/* w */ tinyfont.Glyph{Rune:119, Width:0xe, Height:0xa, XAdvance:0xe, XOffset:0, YOffset:-9, Bitmaps:[]uint8{0xe3, 0x1d, 0x9e, 0x66, 0x79, 0x99, 0xe6, 0x77, 0xb8, 0xd2, 0xc3, 0xcf, 0xf, 0x3c, 0x3c, 0xf0, 0x73, 0x80}},
		/* x */ tinyfont.Glyph{Rune:120, Width:0xa, Height:0xa, XAdvance:0xa, XOffset:0, YOffset:-9, Bitmaps:[]uint8{0x73, 0x9c, 0xe3, 0xf0, 0x78, 0x1e, 0x7, 0x81, 0xe0, 0xfc, 0x73, 0x9c, 0xe0}},
		/* y */ tinyfont.Glyph{Rune:121, Width:0xa, Height:0xe, XAdvance:0xa, XOffset:0, YOffset:-9, Bitmaps:[]uint8{0xe1, 0xd8, 0x67, 0x39, 0xce, 0x33, 0xe, 0xc3, 0xe0, 0x78, 0x1e, 0x3, 0x0, 0xc0, 0x70, 0x38, 0xe, 0x0}},
		/* z */ tinyfont.Glyph{Rune:122, Width:0x8, Height:0xa, XAdvance:0x9, XOffset:1, YOffset:-9, Bitmaps:[]uint8{0xfe, 0xfe, 0xe, 0x1c, 0x38, 0x38, 0x70, 0xe0, 0xff, 0xff}},
		/* { */ tinyfont.Glyph{Rune:123, Width:0x4, Height:0x11, XAdvance:0x7, XOffset:1, YOffset:-12, Bitmaps:[]uint8{0x37, 0x66, 0x66, 0x6e, 0xe6, 0x66, 0x66, 0x67, 0x30}},
		/* | */ tinyfont.Glyph{Rune:124, Width:0x1, Height:0x11, XAdvance:0x5, XOffset:2, YOffset:-12, Bitmaps:[]uint8{0xff, 0xff, 0x80}},
		/* } */ tinyfont.Glyph{Rune:125, Width:0x4, Height:0x11, XAdvance:0x7, XOffset:2, YOffset:-12, Bitmaps:[]uint8{0xce, 0x66, 0x66, 0x67, 0x76, 0x66, 0x66, 0x6e, 0xc0}},
		/* ~ */ tinyfont.Glyph{Rune:126, Width:0x8, Height:0x2, XAdvance:0x9, XOffset:0, YOffset:-4, Bitmaps:[]uint8{0x71, 0x8e}},
	},

	YAdvance:0x16,
}
