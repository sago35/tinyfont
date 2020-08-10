package freeserif

import (
	"tinygo.org/x/tinyfont"
)

var Bold18pt7b = sBold18pt7b{}

type sBold18pt7b struct {
}

const cBold18pt7b = "" +
	/*   */ "\x00\x09" + "\x00\x00\x00\x20" + "\x00\x00" + "\x09\x00" + "\x01" + "" +
	/* ! */ "\x00\x1B" + "\x00\x00\x00\x21" + "\x06\x18" + "\x0C\x03" + "\xE9" + "\x7B\xEF\xFF\xFF\xF7\x9E\x71\xC7\x0C\x20\x82\x00\x00\x07\x3E\xFF\xFF\xDC" +
	/* " */ "\x00\x1A" + "\x00\x00\x00\x22" + "\x0D\x0A" + "\x13\x03" + "\xE9" + "\x60\x37\x83\xFC\x1F\xE0\xFF\x07\xB8\x3D\xC0\xCC\x06\x20\x31\x01\x80" +
	/* # */ "\x00\x3F" + "\x00\x00\x00\x23" + "\x12\x18" + "\x11\x00" + "\xE9" + "\x03\x8E\x00\xC3\x80\x30\xE0\x1C\x38\x07\x0E\x01\xC3\x87\xFF\xFD\xFF\xFF\x7F\xFF\xC1\x87\x00\xE1\xC0\x38\x70\x0E\x1C\x03\x86\x0F\xFF\xF3\xFF\xFC\xFF\xFF\x07\x0E\x01\xC3\x80\x70\xE0\x1C\x30\x07\x0C\x01\x87\x00\x61\xC0" +
	/* $ */ "\x00\x3E" + "\x00\x00\x00\x24" + "\x0F\x1C" + "\x11\x01" + "\xE7" + "\x02\x00\x04\x00\x08\x00\xFF\x03\x27\x8C\x47\x38\x86\x71\x0C\xF2\x09\xF4\x03\xF8\x03\xF8\x07\xFC\x03\xFC\x03\xFE\x01\xFE\x03\xFC\x04\xFC\x08\xFA\x10\xF4\x21\xEC\x43\xD8\x8F\x3D\x3C\x3F\xF0\x1F\x00\x08\x00\x10\x00" +
	/* % */ "\x00\x5A" + "\x00\x00\x00\x25" + "\x1B\x18" + "\x23\x04" + "\xE9" + "\x03\xC0\x18\x01\xFE\x0F\x00\x7C\xFF\xC0\x1F\x0F\x90\x07\xC1\x06\x00\xF0\x21\x80\x3E\x04\x30\x07\x81\x8C\x00\xF0\x21\x80\x1E\x0C\x60\x03\xC1\x18\x1E\x3C\xE3\x0F\xE7\xF8\xC3\xE6\x3C\x18\xF8\x40\x06\x3E\x08\x01\x87\x81\x00\x31\xF0\x20\x0C\x3E\x04\x01\x87\x81\x00\x60\xF0\x60\x18\x1E\x08\x03\x03\xC7\x00\xC0\x3F\xC0\x18\x03\xE0" +
	/* & */ "\x00\x5B" + "\x00\x00\x00\x26" + "\x1A\x19" + "\x1D\x02" + "\xE9" + "\x00\x7E\x00\x00\x7F\xE0\x00\x38\xF8\x00\x1E\x1F\x00\x07\x83\xC0\x01\xF0\xF0\x00\x7C\x38\x00\x1F\x9C\x00\x03\xFC\x00\x00\xFE\x0F\xF0\x3F\x80\xF0\x1F\xF0\x18\x1C\xFE\x0C\x0E\x1F\xC3\x07\x87\xF1\x81\xE0\xFE\x40\xF8\x1F\xF0\x3F\x07\xF8\x0F\xC0\xFE\x03\xF8\x1F\xC0\xFE\x07\xF8\x9F\xE3\xFF\xE7\xFF\x9F\xF0\xFF\xC3\xF8\x0F\x80\x3C\x00" +
	/* ' */ "\x00\x0E" + "\x00\x00\x00\x27" + "\x04\x0A" + "\x0A\x03" + "\xE9" + "\x6F\xFF\xFF\x66\x66" +
	/* ( */ "\x00\x2B" + "\x00\x00\x00\x28" + "\x09\x1E" + "\x0C\x02" + "\xE9" + "\x00\x81\x81\x81\x81\x80\xC0\xE0\x70\x70\x38\x3C\x1E\x0F\x07\x83\xC1\xE0\xF0\x78\x3C\x0E\x07\x03\x80\xE0\x70\x18\x06\x01\x00\x40\x10\x04" +
	/* ) */ "\x00\x2B" + "\x00\x00\x00\x29" + "\x09\x1E" + "\x0C\x01" + "\xE9" + "\x80\x30\x0C\x03\x00\xC0\x60\x38\x1C\x07\x03\x81\xC0\xF0\x78\x3C\x1E\x0F\x07\x83\xC1\xE0\xE0\x70\x38\x38\x1C\x0C\x0C\x06\x04\x04\x04\x00" +
	/* * */ "\x00\x24" + "\x00\x00\x00\x2A" + "\x0E\x0F" + "\x12\x02" + "\xE9" + "\x03\x00\x1E\x00\x78\x1D\xE6\xFB\x3D\xED\xF3\xFF\x01\xC0\x7F\xF3\xED\xFF\x33\xD9\xE6\x07\x80\x1E\x00\x30\x00" +
	/* + */ "\x00\x37" + "\x00\x00\x00\x2B" + "\x13\x13" + "\x18\x02" + "\xEF" + "\x00\xE0\x00\x1C\x00\x03\x80\x00\x70\x00\x0E\x00\x01\xC0\x00\x38\x00\x07\x00\xFF\xFF\xFF\xFF\xFF\xFF\xFF\x80\x70\x00\x0E\x00\x01\xC0\x00\x38\x00\x07\x00\x00\xE0\x00\x1C\x00\x03\x80\x00" +
	/* , */ "\x00\x12" + "\x00\x00\x00\x2C" + "\x06\x0C" + "\x09\x01" + "\xFB" + "\x73\xEF\xFF\xFD\xF0\xC2\x18\xC6\x30" +
	/* - */ "\x00\x0D" + "\x00\x00\x00\x2D" + "\x08\x04" + "\x0C\x02" + "\xF7" + "\xFF\xFF\xFF\xFF" +
	/* . */ "\x00\x0E" + "\x00\x00\x00\x2E" + "\x06\x06" + "\x09\x01" + "\xFB" + "\x7B\xFF\xFF\xFD\xE0" +
	/* / */ "\x00\x2C" + "\x00\x00\x00\x2F" + "\x0B\x19" + "\x0A\xFF" + "\xE9" + "\x00\xE0\x3C\x07\x00\xE0\x1C\x07\x00\xE0\x1C\x07\x00\xE0\x1C\x07\x00\xE0\x1C\x07\x00\xE0\x1C\x07\x00\xE0\x1C\x07\x00\xE0\x1C\x07\x00\xE0\x00" +
	/* 0 */ "\x00\x39" + "\x00\x00\x00\x30" + "\x10\x18" + "\x12\x01" + "\xE9" + "\x03\xC0\x0E\x70\x1E\x78\x3C\x3C\x3C\x3C\x7C\x3E\x7C\x3E\x7C\x3E\xFC\x3F\xFC\x3F\xFC\x3F\xFC\x3F\xFC\x3F\xFC\x3F\xFC\x3F\xFC\x3F\xFC\x3E\x7C\x3E\x7C\x3E\x3C\x3C\x3C\x3C\x1E\x78\x0E\x70\x03\xC0" +
	/* 1 */ "\x00\x2D" + "\x00\x00\x00\x31" + "\x0C\x18" + "\x12\x03" + "\xE9" + "\x00\xC0\x3C\x0F\xC3\xFC\x4F\xC0\xFC\x0F\xC0\xFC\x0F\xC0\xFC\x0F\xC0\xFC\x0F\xC0\xFC\x0F\xC0\xFC\x0F\xC0\xFC\x0F\xC0\xFC\x0F\xC0\xFC\x1F\xEF\xFF" +
	/* 2 */ "\x00\x39" + "\x00\x00\x00\x32" + "\x10\x18" + "\x11\x01" + "\xE9" + "\x03\xE0\x0F\xF8\x1F\xFC\x3F\xFC\x30\xFE\x60\x7E\x40\x3E\x00\x3E\x00\x3E\x00\x3C\x00\x3C\x00\x78\x00\x70\x00\xE0\x00\xC0\x01\x80\x03\x00\x06\x01\x0C\x03\x1F\xFF\x1F\xFF\x3F\xFE\x7F\xFE\xFF\xFE" +
	/* 3 */ "\x00\x39" + "\x00\x00\x00\x33" + "\x10\x18" + "\x12\x00" + "\xE9" + "\x03\xF0\x0F\xF8\x3F\xFC\x21\xFE\x40\xFE\x00\x7E\x00\x7E\x00\x7C\x00\x78\x00\xF0\x01\xFC\x03\xFE\x00\x7E\x00\x3F\x00\x1F\x00\x0F\x00\x0F\x00\x0F\x00\x0E\x70\x0E\xFC\x1C\xFE\x38\x7F\xE0\x3F\x80" +
	/* 4 */ "\x00\x36" + "\x00\x00\x00\x34" + "\x0F\x18" + "\x12\x01" + "\xE9" + "\x00\x38\x00\xF0\x03\xE0\x07\xC0\x1F\x80\x5F\x00\xBE\x02\x7C\x08\xF8\x31\xF0\x43\xE1\x07\xC4\x0F\x88\x1F\x20\x3E\x7F\xFF\xFF\xFF\xFF\xFF\xFF\xF8\x07\xC0\x0F\x80\x1F\x00\x3E\x00\x7C" +
	/* 5 */ "\x00\x36" + "\x00\x00\x00\x35" + "\x0F\x18" + "\x12\x01" + "\xE9" + "\x0F\xFE\x1F\xF8\x7F\xF0\xFF\xE1\x80\x03\x00\x0C\x00\x18\x00\x3F\x80\xFF\xC1\xFF\xC3\xFF\xC3\xFF\x80\x3F\x80\x0F\x00\x0E\x00\x1C\x00\x18\x00\x37\x80\x4F\x81\x9F\xC6\x3F\xF8\x1F\x80" +
	/* 6 */ "\x00\x39" + "\x00\x00\x00\x36" + "\x10\x18" + "\x12\x01" + "\xE9" + "\x00\x07\x00\x7C\x01\xF0\x03\xC0\x0F\x80\x1F\x00\x1F\x00\x3E\x00\x7E\x00\x7F\xF0\x7F\xFC\xFC\x7E\xFC\x7E\xFC\x3F\xFC\x3F\xFC\x3F\xFC\x3F\xFC\x3F\x7C\x3F\x7C\x3E\x3C\x3E\x3E\x3C\x1E\x78\x07\xE0" +
	/* 7 */ "\x00\x39" + "\x00\x00\x00\x37" + "\x10\x18" + "\x11\x01" + "\xE9" + "\x7F\xFF\x7F\xFE\x7F\xFE\xFF\xFE\xFF\xFC\xC0\x1C\x80\x18\x80\x38\x00\x38\x00\x70\x00\x70\x00\x70\x00\xE0\x00\xE0\x00\xE0\x01\xC0\x01\xC0\x01\xC0\x03\x80\x03\x80\x07\x80\x07\x00\x07\x00\x0F\x00" +
	/* 8 */ "\x00\x39" + "\x00\x00\x00\x38" + "\x10\x18" + "\x11\x01" + "\xE9" + "\x0F\xE0\x38\x78\x70\x3C\xF0\x1E\xF0\x1E\xF8\x1E\xF8\x1E\xFE\x3C\x7F\xB0\x7F\xE0\x3F\xF0\x0F\xF8\x1F\xFC\x39\xFE\x70\xFF\xF0\x3F\xF0\x3F\xF0\x1F\xF0\x1F\xF0\x1E\x78\x3E\x7C\x7C\x3F\xF8\x0F\xE0" +
	/* 9 */ "\x00\x39" + "\x00\x00\x00\x39" + "\x10\x18" + "\x12\x01" + "\xE9" + "\x07\xE0\x1E\x78\x3C\x7C\x7C\x3C\x7C\x3E\xFC\x3E\xFC\x3F\xFC\x3F\xFC\x3F\xFC\x3F\xFC\x3F\x7E\x3F\x7E\x3F\x3F\xFE\x0F\xFE\x00\x7E\x00\x7C\x00\xF8\x00\xF8\x01\xF0\x03\xC0\x0F\x80\x3E\x00\xE0\x00" +
	/* : */ "\x00\x15" + "\x00\x00\x00\x3A" + "\x06\x10" + "\x0C\x03" + "\xF1" + "\x7B\xFF\xFF\xFD\xE0\x00\x00\x07\xBF\xFF\xFF\xDE" +
	/* ; */ "\x00\x1D" + "\x00\x00\x00\x3B" + "\x07\x16" + "\x0C\x02" + "\xF1" + "\x39\xFB\xF7\xEF\xC7\x00\x00\x00\x01\xE7\xEF\xFF\xFF\xBF\x06\x08\x30\xC2\x18\x00" +
	/* < */ "\x00\x39" + "\x00\x00\x00\x3C" + "\x13\x14" + "\x18\x02" + "\xEE" + "\x00\x00\x00\x00\x0C\x00\x0F\x80\x07\xF0\x03\xFC\x01\xFE\x00\xFE\x00\x7F\x00\x3F\x80\x1F\xC0\x03\xF8\x00\x1F\xC0\x00\xFE\x00\x07\xF0\x00\x3F\x80\x01\xFE\x00\x0F\xE0\x00\x7C\x00\x01\x80\x00\x00" +
	/* = */ "\x00\x26" + "\x00\x00\x00\x3D" + "\x13\x0C" + "\x18\x02" + "\xF2" + "\xFF\xFF\xFF\xFF\xFF\xFF\xFF\x80\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x1F\xFF\xFF\xFF\xFF\xFF\xFF\xF0" +
	/* > */ "\x00\x39" + "\x00\x00\x00\x3E" + "\x13\x14" + "\x18\x03" + "\xEE" + "\x00\x00\x18\x00\x03\xE0\x00\x7F\x00\x07\xF8\x00\x1F\xC0\x00\xFE\x00\x07\xF0\x00\x3F\x80\x01\xFC\x00\x3F\x80\x1F\xC0\x0F\xE0\x07\xF0\x07\xF8\x03\xFC\x00\xFE\x00\x1F\x00\x03\x00\x00\x00\x00\x00" +
	/* ? */ "\x00\x33" + "\x00\x00\x00\x3F" + "\x0E\x18" + "\x12\x02" + "\xE9" + "\x0F\xC0\xFF\xC7\x1F\xB8\x3E\xF0\xFF\xC3\xFF\x0F\xD8\x3F\x00\xF8\x07\xC0\x1E\x00\x60\x03\x00\x08\x00\x20\x00\x80\x00\x00\x00\x00\x70\x03\xE0\x1F\x80\x7E\x01\xF8\x01\xC0" +
	/* @ */ "\x00\x54" + "\x00\x00\x00\x40" + "\x18\x19" + "\x21\x04" + "\xE9" + "\x00\x7F\x00\x01\xFF\xE0\x07\xC0\xF0\x0F\x00\x38\x1E\x00\x0C\x3C\x07\x06\x38\x1F\x72\x78\x3C\xF3\x78\x78\xE1\xF0\x70\xE1\xF0\xF0\xE1\xF0\xE0\xC1\xF1\xE1\xC1\xF1\xC1\xC1\xF1\xC3\x82\xF1\xC3\x86\x71\xC7\x8C\x79\xFB\xF8\x78\xF1\xF0\x3C\x00\x00\x1E\x00\x00\x0F\x00\x00\x07\xC0\x78\x03\xFF\xE0\x00\x7F\x80" +
	/* A */ "\x00\x51" + "\x00\x00\x00\x41" + "\x18\x18" + "\x19\x01" + "\xE9" + "\x00\x10\x00\x00\x38\x00\x00\x38\x00\x00\x78\x00\x00\x7C\x00\x00\x7C\x00\x00\xFE\x00\x00\xFE\x00\x01\xBF\x00\x01\xBF\x00\x01\x1F\x00\x03\x1F\x80\x02\x1F\x80\x06\x0F\xC0\x06\x0F\xC0\x04\x07\xE0\x0F\xFF\xE0\x0F\xFF\xE0\x18\x03\xF0\x18\x03\xF0\x30\x01\xF8\x30\x01\xF8\x70\x01\xFC\xFE\x0F\xFF" +
	/* B */ "\x00\x48" + "\x00\x00\x00\x42" + "\x15\x18" + "\x17\x01" + "\xE9" + "\xFF\xFE\x07\xFF\xFE\x0F\xE1\xF8\x3F\x07\xC1\xF8\x3F\x0F\xC1\xF8\x7E\x0F\xC3\xF0\x7E\x1F\x87\xE0\xFC\x7C\x07\xFF\x00\x3F\xFF\x01\xF8\xFE\x0F\xC1\xF8\x7E\x0F\xC3\xF0\x3F\x1F\x81\xF8\xFC\x0F\xC7\xE0\x7E\x3F\x03\xF1\xF8\x3F\x0F\xC3\xF0\xFF\xFF\x1F\xFF\xC0" +
	/* C */ "\x00\x51" + "\x00\x00\x00\x43" + "\x17\x19" + "\x19\x01" + "\xE9" + "\x00\x7E\x04\x07\xFF\x18\x1F\x07\xF0\x7C\x03\xE1\xF0\x03\xC7\xC0\x03\x9F\x80\x03\x3F\x00\x06\x7C\x00\x05\xF8\x00\x03\xF0\x00\x07\xE0\x00\x0F\xC0\x00\x1F\x80\x00\x3F\x00\x00\x7E\x00\x00\xFC\x00\x00\xFC\x00\x01\xF8\x00\x01\xF0\x00\x23\xF0\x00\xC3\xF0\x07\x03\xF0\x3C\x01\xFF\xE0\x00\xFF\x00" +
	/* D */ "\x00\x4E" + "\x00\x00\x00\x44" + "\x17\x18" + "\x1A\x01" + "\xE9" + "\xFF\xFE\x00\x7F\xFF\x00\x7E\x1F\x80\xFC\x1F\x81\xF8\x1F\x83\xF0\x1F\x07\xE0\x3F\x0F\xC0\x7E\x1F\x80\x7E\x3F\x00\xFC\x7E\x01\xF8\xFC\x03\xF1\xF8\x07\xE3\xF0\x0F\xC7\xE0\x1F\x8F\xC0\x3F\x1F\x80\x7C\x3F\x01\xF8\x7E\x03\xE0\xFC\x0F\x81\xF8\x1F\x03\xF0\xFC\x0F\xFF\xE0\x7F\xFF\x00" +
	/* E */ "\x00\x48" + "\x00\x00\x00\x45" + "\x15\x18" + "\x17\x02" + "\xE9" + "\xFF\xFF\xE3\xFF\xFF\x0F\xC0\x78\x7E\x01\xC3\xF0\x06\x1F\x80\x10\xFC\x10\x87\xE0\x80\x3F\x0C\x01\xF8\xE0\x0F\xFF\x00\x7F\xF8\x03\xF1\xC0\x1F\x86\x00\xFC\x10\x07\xE0\x80\x3F\x00\x09\xF8\x00\xCF\xC0\x0C\x7E\x00\x63\xF0\x0F\x1F\x81\xFB\xFF\xFF\xDF\xFF\xFC" +
	/* F */ "\x00\x42" + "\x00\x00\x00\x46" + "\x13\x18" + "\x16\x02" + "\xE9" + "\xFF\xFF\xEF\xFF\xFC\xFC\x0F\x9F\x80\x73\xF0\x06\x7E\x00\x4F\xC1\x09\xF8\x20\x3F\x0C\x07\xE3\x80\xFF\xF0\x1F\xFE\x03\xF1\xC0\x7E\x18\x0F\xC1\x01\xF8\x20\x3F\x00\x07\xE0\x00\xFC\x00\x1F\x80\x03\xF0\x00\x7E\x00\x1F\xE0\x07\xFF\x00" +
	/* G */ "\x00\x58" + "\x00\x00\x00\x47" + "\x19\x19" + "\x1B\x01" + "\xE9" + "\x00\x7E\x02\x01\xFF\xE3\x01\xF0\x3F\x81\xF0\x07\xC1\xF0\x01\xE1\xF0\x00\x71\xF8\x00\x18\xFC\x00\x0C\x7C\x00\x02\x7E\x00\x00\x3F\x00\x00\x1F\x80\x00\x0F\xC0\x00\x07\xE0\x00\x03\xF0\x0F\xFF\xF8\x01\xFE\x7C\x00\x7E\x3F\x00\x3F\x1F\x80\x1F\x87\xC0\x0F\xC1\xF0\x07\xE0\xFC\x03\xF0\x1F\x83\xF0\x07\xFF\xE0\x00\x7F\x80\x00" +
	/* H */ "\x00\x51" + "\x00\x00\x00\x48" + "\x18\x18" + "\x1B\x02" + "\xE9" + "\xFF\xC3\xFF\x7F\x81\xFE\x3F\x00\xFC\x3F\x00\xFC\x3F\x00\xFC\x3F\x00\xFC\x3F\x00\xFC\x3F\x00\xFC\x3F\x00\xFC\x3F\x00\xFC\x3F\xFF\xFC\x3F\xFF\xFC\x3F\x00\xFC\x3F\x00\xFC\x3F\x00\xFC\x3F\x00\xFC\x3F\x00\xFC\x3F\x00\xFC\x3F\x00\xFC\x3F\x00\xFC\x3F\x00\xFC\x3F\x00\xFC\x7F\x81\xFE\xFF\xC3\xFF" +
	/* I */ "\x00\x2A" + "\x00\x00\x00\x49" + "\x0B\x18" + "\x0E\x02" + "\xE9" + "\xFF\xEF\xF0\xFC\x1F\x83\xF0\x7E\x0F\xC1\xF8\x3F\x07\xE0\xFC\x1F\x83\xF0\x7E\x0F\xC1\xF8\x3F\x07\xE0\xFC\x1F\x83\xF0\x7E\x1F\xE7\xFF" +
	/* J */ "\x00\x3F" + "\x00\x00\x00\x4A" + "\x10\x1B" + "\x12\x00" + "\xE9" + "\x07\xFF\x01\xFE\x00\xFC\x00\xFC\x00\xFC\x00\xFC\x00\xFC\x00\xFC\x00\xFC\x00\xFC\x00\xFC\x00\xFC\x00\xFC\x00\xFC\x00\xFC\x00\xFC\x00\xFC\x00\xFC\x00\xFC\x00\xFC\x70\xFC\xF8\xFC\xF8\xF8\xF0\xF8\x71\xF0\x7F\xE0\x1F\x80" +
	/* K */ "\x00\x54" + "\x00\x00\x00\x4B" + "\x19\x18" + "\x1B\x02" + "\xE9" + "\xFF\xC3\xFF\x3F\xC0\x3E\x0F\xC0\x1C\x07\xE0\x18\x03\xF0\x18\x01\xF8\x18\x00\xFC\x18\x00\x7E\x18\x00\x3F\x18\x00\x1F\x9C\x00\x0F\xDF\x00\x07\xFF\xC0\x03\xFF\xF0\x01\xF9\xF8\x00\xFC\xFE\x00\x7E\x3F\x80\x3F\x0F\xE0\x1F\x83\xF8\x0F\xC0\xFC\x07\xE0\x7F\x03\xF0\x1F\xC1\xF8\x07\xF1\xFE\x03\xFD\xFF\x8F\xFF" +
	/* L */ "\x00\x48" + "\x00\x00\x00\x4C" + "\x15\x18" + "\x17\x02" + "\xE9" + "\xFF\xE0\x03\xFC\x00\x0F\xC0\x00\x7E\x00\x03\xF0\x00\x1F\x80\x00\xFC\x00\x07\xE0\x00\x3F\x00\x01\xF8\x00\x0F\xC0\x00\x7E\x00\x03\xF0\x00\x1F\x80\x00\xFC\x00\x07\xE0\x01\x3F\x00\x19\xF8\x00\xCF\xC0\x0C\x7E\x00\x63\xF0\x0F\x1F\x81\xFB\xFF\xFF\xDF\xFF\xFE" +
	/* M */ "\x00\x66" + "\x00\x00\x00\x4D" + "\x1F\x18" + "\x21\x01" + "\xE9" + "\xFF\x80\x03\xFE\x7F\x00\x07\xF8\x7E\x00\x0F\xE0\xFE\x00\x3F\xC1\x7C\x00\x5F\x82\xFC\x01\xBF\x05\xF8\x02\x7E\x09\xF8\x0C\xFC\x13\xF0\x11\xF8\x23\xE0\x23\xF0\x47\xE0\xC7\xE0\x87\xC1\x0F\xC1\x0F\xC6\x1F\x82\x0F\x88\x3F\x04\x1F\xB0\x7E\x08\x3F\x60\xFC\x10\x3E\x81\xF8\x20\x7F\x03\xF0\x40\x7C\x07\xE0\x80\xF8\x0F\xC1\x00\xE0\x1F\x82\x01\xC0\x3F\x0E\x03\x80\xFF\x7F\x82\x03\xFF" +
	/* N */ "\x00\x4E" + "\x00\x00\x00\x4E" + "\x17\x18" + "\x19\x01" + "\xE9" + "\xFE\x00\xFE\xFE\x00\x70\xFE\x00\x40\xFE\x00\x81\xFC\x01\x03\xFC\x02\x05\xFC\x04\x09\xFC\x08\x11\xFC\x10\x23\xF8\x20\x43\xF8\x40\x83\xF8\x81\x03\xF9\x02\x03\xFA\x04\x03\xF4\x08\x07\xF8\x10\x07\xF0\x20\x07\xE0\x40\x07\xC0\x80\x07\x81\x00\x0F\x02\x00\x0E\x0E\x00\x0C\x7F\x00\x08" +
	/* O */ "\x00\x58" + "\x00\x00\x00\x4F" + "\x19\x19" + "\x1B\x01" + "\xE9" + "\x00\x7F\x00\x01\xFF\xF0\x01\xF0\x7C\x01\xF0\x1F\x01\xF0\x07\xC1\xF0\x01\xF1\xF8\x00\xFC\xFC\x00\x7E\x7C\x00\x1F\x7E\x00\x0F\xFF\x00\x07\xFF\x80\x03\xFF\xC0\x01\xFF\xE0\x00\xFF\xF0\x00\x7F\xF8\x00\x3F\x7C\x00\x1F\x3E\x00\x1F\x9F\x80\x0F\xC7\xC0\x07\xC1\xF0\x07\xC0\xFC\x07\xE0\x3F\x07\xC0\x07\xFF\xC0\x00\x7F\x00\x00" +
	/* P */ "\x00\x42" + "\x00\x00\x00\x50" + "\x13\x18" + "\x16\x02" + "\xE9" + "\xFF\xFC\x0F\xFF\xE0\xFC\x7E\x1F\x87\xE3\xF0\x7E\x7E\x0F\xCF\xC1\xF9\xF8\x3F\x3F\x07\xE7\xE0\xFC\xFC\x3F\x1F\x8F\xC3\xFF\xF0\x7F\xF8\x0F\xC0\x01\xF8\x00\x3F\x00\x07\xE0\x00\xFC\x00\x1F\x80\x03\xF0\x00\x7E\x00\x1F\xE0\x07\xFE\x00" +
	/* Q */ "\x00\x67" + "\x00\x00\x00\x51" + "\x19\x1E" + "\x1B\x01" + "\xE9" + "\x00\x7F\x00\x01\xFF\xF0\x01\xF0\x7C\x01\xF0\x1F\x01\xF0\x07\xC1\xF0\x01\xF1\xF8\x00\xFC\xFC\x00\x7E\x7C\x00\x1F\x7E\x00\x0F\xFF\x00\x07\xFF\x80\x03\xFF\xC0\x01\xFF\xE0\x00\xFF\xF0\x00\x7F\xF8\x00\x3F\x7C\x00\x1F\x3E\x00\x0F\x9F\x80\x0F\xC7\xC0\x07\xC1\xF0\x07\xC0\x78\x03\xC0\x1E\x07\xC0\x03\xFF\x80\x00\x7F\x00\x00\x3F\xC0\x00\x0F\xF0\x00\x03\xFE\x00\x00\xFF\xF8\x00\x0F\xE0" +
	/* R */ "\x00\x4E" + "\x00\x00\x00\x52" + "\x17\x18" + "\x19\x02" + "\xE9" + "\xFF\xFE\x00\xFF\xFF\x00\xFC\x3F\x01\xF8\x3F\x03\xF0\x3F\x07\xE0\x7E\x0F\xC0\xFC\x1F\x81\xF8\x3F\x03\xF0\x7E\x07\xC0\xFC\x1F\x81\xF8\x7E\x03\xFF\xF0\x07\xFF\xC0\x0F\xDF\xC0\x1F\x9F\x80\x3F\x1F\x80\x7E\x3F\x80\xFC\x3F\x81\xF8\x3F\x03\xF0\x7F\x07\xE0\x7F\x1F\xE0\x7F\x7F\xE0\xFF" +
	/* S */ "\x00\x3B" + "\x00\x00\x00\x53" + "\x10\x19" + "\x14\x02" + "\xE9" + "\x07\xC2\x1F\xF2\x3C\x3E\x70\x0E\xF0\x06\xF0\x06\xF0\x02\xF8\x00\xFE\x00\xFF\x80\x7F\xE0\x3F\xF8\x1F\xFC\x0F\xFE\x03\xFE\x00\xFF\x00\x3F\x80\x1F\xC0\x0F\xC0\x0F\xE0\x0E\xF0\x1E\xF8\x3C\x9F\xF8\x87\xE0" +
	/* T */ "\x00\x48" + "\x00\x00\x00\x54" + "\x15\x18" + "\x17\x01" + "\xE9" + "\xFF\xFF\xFF\xFF\xFF\xFC\x7E\x3F\x83\xF0\x7C\x1F\x81\xC0\xFC\x06\x07\xE0\x20\x3F\x00\x01\xF8\x00\x0F\xC0\x00\x7E\x00\x03\xF0\x00\x1F\x80\x00\xFC\x00\x07\xE0\x00\x3F\x00\x01\xF8\x00\x0F\xC0\x00\x7E\x00\x03\xF0\x00\x1F\x80\x00\xFC\x00\x0F\xF0\x01\xFF\xE0" +
	/* U */ "\x00\x4E" + "\x00\x00\x00\x55" + "\x16\x19" + "\x19\x02" + "\xE9" + "\xFF\xC1\xFD\xFE\x01\xC3\xF0\x02\x0F\xC0\x08\x3F\x00\x20\xFC\x00\x83\xF0\x02\x0F\xC0\x08\x3F\x00\x20\xFC\x00\x83\xF0\x02\x0F\xC0\x08\x3F\x00\x20\xFC\x00\x83\xF0\x02\x0F\xC0\x08\x3F\x00\x20\xFC\x00\x83\xF0\x02\x0F\xC0\x18\x1F\x80\x40\x7E\x03\x00\xFC\x18\x01\xFF\xC0\x00\xFC\x00" +
	/* V */ "\x00\x51" + "\x00\x00\x00\x56" + "\x18\x18" + "\x19\x00" + "\xE9" + "\xFF\xF0\x7F\x3F\xC0\x1E\x1F\x80\x0C\x1F\x80\x08\x0F\xC0\x18\x0F\xC0\x18\x07\xE0\x10\x07\xE0\x30\x07\xE0\x20\x03\xF0\x60\x03\xF0\x60\x01\xF8\x40\x01\xF8\xC0\x00\xF8\x80\x00\xFC\x80\x00\xFD\x80\x00\x7F\x00\x00\x7F\x00\x00\x3E\x00\x00\x3E\x00\x00\x1E\x00\x00\x1C\x00\x00\x1C\x00\x00\x0C\x00" +
	/* W */ "\x00\x74" + "\x00\x00\x00\x57" + "\x22\x19" + "\x22\x00" + "\xE9" + "\xFF\xE7\xFF\x0F\xCF\xE0\x7F\x00\xE1\xF8\x0F\xC0\x30\x7E\x03\xF0\x0C\x1F\x80\x7C\x02\x03\xE0\x1F\x81\x80\xFC\x07\xE0\x60\x3F\x03\xF8\x10\x07\xC0\xBF\x0C\x01\xF8\x2F\xC3\x00\x7E\x19\xF0\x80\x0F\x84\x7C\x60\x03\xF3\x0F\x98\x00\xFC\xC3\xE4\x00\x1F\x20\xFB\x00\x07\xF8\x1F\xC0\x00\xFC\x07\xE0\x00\x3F\x01\xF8\x00\x0F\xC0\x3E\x00\x01\xE0\x0F\x00\x00\x78\x03\xC0\x00\x1C\x00\x70\x00\x03\x00\x18\x00\x00\xC0\x06\x00\x00\x20\x00\x80\x00" +
	/* X */ "\x00\x51" + "\x00\x00\x00\x58" + "\x18\x18" + "\x19\x01" + "\xE9" + "\xFF\xF3\xFE\x7F\x80\x78\x3F\x80\x70\x1F\xC0\x60\x0F\xC0\xC0\x0F\xE1\x80\x07\xF1\x00\x03\xF3\x00\x03\xFE\x00\x01\xFC\x00\x00\xFC\x00\x00\xFE\x00\x00\x7E\x00\x00\x7F\x00\x00\xFF\x80\x00\x9F\x80\x01\x8F\xC0\x03\x0F\xE0\x06\x07\xE0\x06\x07\xF0\x0C\x03\xF8\x1C\x03\xF8\x3C\x03\xFC\xFF\x0F\xFF" +
	/* Y */ "\x00\x51" + "\x00\x00\x00\x59" + "\x18\x18" + "\x19\x01" + "\xE9" + "\xFF\xF0\xFF\x7F\x80\x1E\x3F\x80\x1C\x1F\x80\x18\x1F\xC0\x10\x0F\xC0\x30\x07\xE0\x20\x07\xE0\x60\x03\xF0\xC0\x03\xF0\x80\x01\xF9\x80\x01\xFF\x00\x00\xFF\x00\x00\xFE\x00\x00\x7E\x00\x00\x7E\x00\x00\x7E\x00\x00\x7E\x00\x00\x7E\x00\x00\x7E\x00\x00\x7E\x00\x00\x7E\x00\x00\xFF\x00\x01\xFF\x80" +
	/* Z */ "\x00\x48" + "\x00\x00\x00\x5A" + "\x15\x18" + "\x17\x01" + "\xE9" + "\x7F\xFF\xF3\xFF\xFF\x9F\x01\xF8\xE0\x1F\x86\x01\xFC\x20\x0F\xC1\x00\xFC\x00\x07\xE0\x00\x7E\x00\x07\xE0\x00\x3F\x00\x03\xF0\x00\x3F\x80\x01\xF8\x00\x1F\x80\x01\xFC\x01\x0F\xC0\x18\xFC\x00\xC7\xE0\x06\x7E\x00\x77\xF0\x07\x3F\x00\xFB\xFF\xFF\xDF\xFF\xFE" +
	/* [ */ "\x00\x26" + "\x00\x00\x00\x5B" + "\x08\x1D" + "\x0C\x02" + "\xE9" + "\xFF\xFF\xF0\xF0\xF0\xF0\xF0\xF0\xF0\xF0\xF0\xF0\xF0\xF0\xF0\xF0\xF0\xF0\xF0\xF0\xF0\xF0\xF0\xF0\xF0\xF0\xF0\xFF\xFF" +
	/* \ */ "\x00\x2C" + "\x00\x00\x00\x5C" + "\x0B\x19" + "\x0A\xFF" + "\xE9" + "\xE0\x1E\x01\xC0\x38\x07\x80\x70\x0E\x01\xC0\x1C\x03\x80\x70\x07\x00\xE0\x1C\x01\xC0\x38\x07\x00\x70\x0E\x01\xC0\x1C\x03\x80\x70\x0F\x00\xE0" +
	/* ] */ "\x00\x26" + "\x00\x00\x00\x5D" + "\x08\x1D" + "\x0C\x02" + "\xE9" + "\xFF\xFF\x0F\x0F\x0F\x0F\x0F\x0F\x0F\x0F\x0F\x0F\x0F\x0F\x0F\x0F\x0F\x0F\x0F\x0F\x0F\x0F\x0F\x0F\x0F\x0F\x0F\xFF\xFF" +
	/* ^ */ "\x00\x22" + "\x00\x00\x00\x5E" + "\x0F\x0D" + "\x14\x03" + "\xE9" + "\x03\x80\x0F\x00\x1F\x00\x7E\x00\xEE\x03\x9C\x07\x1C\x1C\x38\x38\x38\xE0\x71\xC0\x77\x00\xEE\x00\xE0" +
	/* _ */ "\x00\x10" + "\x00\x00\x00\x5F" + "\x12\x03" + "\x11\x00" + "\x03" + "\xFF\xFF\xFF\xFF\xFF\xFF\xFC" +
	/* ` */ "\x00\x0F" + "\x00\x00\x00\x60" + "\x08\x06" + "\x0C\x00" + "\xE9" + "\xE0\xF0\x78\x3C\x0E\x07" +
	/* a */ "\x00\x29" + "\x00\x00\x00\x61" + "\x10\x10" + "\x12\x01" + "\xF1" + "\x0F\xE0\x3F\xF0\x78\xF8\x78\x7C\x78\x7C\x38\x7C\x00\x7C\x03\xFC\x1E\x7C\x7C\x7C\xFC\x7C\xFC\x7C\xFC\xFC\xFF\xFD\x7F\x7F\x3C\x3C" +
	/* b */ "\x00\x3F" + "\x00\x00\x00\x62" + "\x12\x18" + "\x13\x01" + "\xE9" + "\xFC\x00\x1F\x00\x07\xC0\x01\xF0\x00\x7C\x00\x1F\x00\x07\xC0\x01\xF0\x00\x7C\xF8\x1F\x7F\x87\xE3\xF1\xF0\x7E\x7C\x0F\x9F\x03\xF7\xC0\xFD\xF0\x3F\x7C\x0F\xDF\x03\xF7\xC0\xFD\xF0\x3E\x7C\x1F\x1F\x8F\xC6\x7F\xC1\x07\xC0" +
	/* c */ "\x00\x25" + "\x00\x00\x00\x63" + "\x0E\x10" + "\x0F\x01" + "\xF1" + "\x07\xC0\x7F\xC3\xC7\x9F\x1E\x78\x7B\xE1\xCF\x80\x3E\x00\xF8\x03\xE0\x0F\x80\x3F\x00\x7C\x00\xFC\x61\xFF\x03\xF0" +
	/* d */ "\x00\x3F" + "\x00\x00\x00\x64" + "\x12\x18" + "\x13\x01" + "\xE9" + "\x00\x7F\x00\x07\xC0\x01\xF0\x00\x7C\x00\x1F\x00\x07\xC0\x01\xF0\x00\x7C\x07\x9F\x07\xF7\xC3\xE3\xF1\xF8\x7C\x7C\x1F\x3F\x07\xCF\xC1\xF3\xF0\x7C\xFC\x1F\x3F\x07\xCF\xC1\xF1\xF0\x7C\x7E\x1F\x0F\x8F\xC1\xFD\xFC\x3E\x70" +
	/* e */ "\x00\x25" + "\x00\x00\x00\x65" + "\x0E\x10" + "\x10\x01" + "\xF1" + "\x0F\xC0\x7F\xC3\xC7\x1E\x1E\xF8\x7B\xFF\xFF\xFF\xFE\x00\xF8\x03\xE0\x0F\xC0\x1F\x03\x7E\x18\xFF\xC1\xFE\x03\xF0" +
	/* f */ "\x00\x2A" + "\x00\x00\x00\x66" + "\x0B\x18" + "\x0E\x02" + "\xE9" + "\x0F\x83\xF8\xF3\xBE\xF7\xDC\xF8\x1F\x03\xE0\xFF\x1F\xE1\xF0\x3E\x07\xC0\xF8\x1F\x03\xE0\x7C\x0F\x81\xF0\x3E\x07\xC0\xF8\x1F\x07\xF8" +
	/* g */ "\x00\x3A" + "\x00\x00\x00\x67" + "\x11\x17" + "\x11\x01" + "\xF1" + "\x0F\xC0\x1F\xFF\xDF\x1F\xFF\x07\x8F\x83\xE7\xC1\xF3\xE0\xF9\xF0\x7C\x78\x3C\x1E\x3E\x03\xFC\x03\x00\x07\x00\x07\x80\x03\xFF\xF1\xFF\xFE\x7F\xFF\x8F\xFF\xF8\x01\xFC\x00\x7F\x00\x73\xFF\xF0\x7F\xC0" +
	/* h */ "\x00\x3C" + "\x00\x00\x00\x68" + "\x11\x18" + "\x13\x01" + "\xE9" + "\xFC\x00\x3E\x00\x1F\x00\x0F\x80\x07\xC0\x03\xE0\x01\xF0\x00\xF8\x00\x7C\x7C\x3E\xFF\x1F\xCF\xCF\x83\xE7\xC1\xF3\xE0\xF9\xF0\x7C\xF8\x3E\x7C\x1F\x3E\x0F\x9F\x07\xCF\x83\xE7\xC1\xF3\xE0\xF9\xF0\x7D\xFC\x7F" +
	/* i */ "\x00\x1E" + "\x00\x00\x00\x69" + "\x07\x18" + "\x0A\x02" + "\xE9" + "\x39\xFB\xF7\xE7\x80\x00\x00\xFC\xF9\xF3\xE7\xCF\x9F\x3E\x7C\xF9\xF3\xE7\xCF\x9F\x7F" +
	/* j */ "\x00\x34" + "\x00\x00\x00\x6A" + "\x0B\x1F" + "\x0E\x00" + "\xE9" + "\x03\xC0\xFC\x1F\x83\xF0\x3C\x00\x00\x00\x00\x0F\xE0\x7C\x0F\x81\xF0\x3E\x07\xC0\xF8\x1F\x03\xE0\x7C\x0F\x81\xF0\x3E\x07\xC0\xF8\x1F\x03\xE0\x7D\xCF\xF9\xEE\x7C\xFF\x0F\x80" +
	/* k */ "\x00\x3F" + "\x00\x00\x00\x6B" + "\x12\x18" + "\x13\x01" + "\xE9" + "\xFC\x00\x1F\x00\x07\xC0\x01\xF0\x00\x7C\x00\x1F\x00\x07\xC0\x01\xF0\x00\x7C\x7F\x9F\x07\x87\xC1\x81\xF0\xC0\x7C\x60\x1F\x30\x07\xDE\x01\xFF\xC0\x7F\xF0\x1F\x3E\x07\xCF\xC1\xF1\xF8\x7C\x3E\x1F\x07\xC7\xC1\xFB\xF9\xFF" +
	/* l */ "\x00\x1E" + "\x00\x00\x00\x6C" + "\x07\x18" + "\x0A\x01" + "\xE9" + "\xFC\xF9\xF3\xE7\xCF\x9F\x3E\x7C\xF9\xF3\xE7\xCF\x9F\x3E\x7C\xF9\xF3\xE7\xCF\x9F\x7F" +
	/* m */ "\x00\x3F" + "\x00\x00\x00\x6D" + "\x1B\x10" + "\x1D\x01" + "\xF1" + "\xFC\x7C\x1F\x0F\xBF\xCF\xF1\xF8\xFF\x3F\x3E\x0F\x83\xE7\xC1\xF0\x7C\xF8\x3E\x0F\x9F\x07\xC1\xF3\xE0\xF8\x3E\x7C\x1F\x07\xCF\x83\xE0\xF9\xF0\x7C\x1F\x3E\x0F\x83\xE7\xC1\xF0\x7C\xF8\x3E\x0F\x9F\x07\xC1\xF7\xF1\xFC\x7F" +
	/* n */ "\x00\x2B" + "\x00\x00\x00\x6E" + "\x11\x10" + "\x13\x01" + "\xF1" + "\xFC\x7C\x3E\xFF\x1F\xCF\xCF\x83\xE7\xC1\xF3\xE0\xF9\xF0\x7C\xF8\x3E\x7C\x1F\x3E\x0F\x9F\x07\xCF\x83\xE7\xC1\xF3\xE0\xF9\xF0\x7D\xFC\x7F" +
	/* o */ "\x00\x2B" + "\x00\x00\x00\x6F" + "\x11\x10" + "\x12\x01" + "\xF1" + "\x07\xF0\x0F\xFE\x0F\x8F\x8F\x87\xE7\xC1\xF7\xE0\xFF\xF0\x7F\xF8\x3F\xFC\x1F\xFE\x0F\xFF\x07\xEF\x83\xE7\xC1\xF1\xF1\xF0\x7F\xF0\x0F\xE0" +
	/* p */ "\x00\x40" + "\x00\x00\x00\x70" + "\x13\x17" + "\x13\x00" + "\xF1" + "\xFE\x7C\x07\xDF\xE0\xFE\x3E\x1F\x07\xE3\xE0\x7C\x7C\x0F\xCF\x81\xF9\xF0\x3F\x3E\x07\xE7\xC0\xFC\xF8\x1F\x9F\x03\xE3\xE0\xFC\x7E\x3F\x0F\xBF\xC1\xF3\xE0\x3E\x00\x07\xC0\x00\xF8\x00\x1F\x00\x03\xE0\x00\x7E\x00\x1F\xE0\x00" +
	/* q */ "\x00\x3A" + "\x00\x00\x00\x71" + "\x11\x17" + "\x13\x01" + "\xF1" + "\x07\xC1\x0F\xF9\x8F\xCD\xCF\xC3\xE7\xC1\xF7\xE0\xFB\xF0\x7D\xF8\x3E\xFC\x1F\x7E\x0F\xBF\x07\xDF\x83\xE7\xE1\xF1\xF1\xF8\x7F\x7C\x1F\x3E\x00\x1F\x00\x0F\x80\x07\xC0\x03\xE0\x01\xF0\x01\xF8\x01\xFE" +
	/* r */ "\x00\x23" + "\x00\x00\x00\x72" + "\x0D\x10" + "\x0F\x01" + "\xF1" + "\xFC\x73\xEF\xDF\xFE\xFC\xF7\xC3\xBE\x01\xF0\x0F\x80\x7C\x03\xE0\x1F\x00\xF8\x07\xC0\x3E\x01\xF0\x1F\xE0" +
	/* s */ "\x00\x21" + "\x00\x00\x00\x73" + "\x0C\x10" + "\x0E\x01" + "\xF1" + "\x1E\x23\xFE\x70\xEE\x06\xE0\x2F\x80\xFF\x07\xFC\x3F\xE0\xFF\x81\xF8\x07\xC0\x7E\x0E\xBF\xC8\xF8" +
	/* t */ "\x00\x24" + "\x00\x00\x00\x74" + "\x0A\x15" + "\x0C\x01" + "\xEC" + "\x04\x03\x01\xC0\xF0\x7C\x3F\xEF\xF9\xF0\x7C\x1F\x07\xC1\xF0\x7C\x1F\x07\xC1\xF0\x7C\x5F\x37\xF8\xFE\x1E\x00" +
	/* u */ "\x00\x2D" + "\x00\x00\x00\x75" + "\x12\x10" + "\x14\x01" + "\xF1" + "\xFC\x7F\x1F\x07\xC7\xC1\xF1\xF0\x7C\x7C\x1F\x1F\x07\xC7\xC1\xF1\xF0\x7C\x7C\x1F\x1F\x07\xC7\xC1\xF1\xF0\x7C\x7C\x1F\x1F\x8F\xC3\xFD\xFC\x7C\x60" +
	/* v */ "\x00\x2B" + "\x00\x00\x00\x76" + "\x11\x10" + "\x11\x00" + "\xF1" + "\xFF\x9F\xBF\x83\x0F\x81\x87\xE0\x81\xF0\x40\xF8\x40\x3E\x20\x1F\x30\x07\xD0\x03\xF8\x00\xF8\x00\x7C\x00\x3C\x00\x0E\x00\x07\x00\x01\x00" +
	/* w */ "\x00\x39" + "\x00\x00\x00\x77" + "\x18\x10" + "\x19\x00" + "\xF1" + "\xFF\x3F\xCF\x7E\x1F\x06\x3E\x0F\x06\x3E\x0F\x84\x1F\x0F\x8C\x1F\x1F\x88\x0F\x17\xC8\x0F\x97\xD8\x0F\xB3\xD0\x07\xE3\xF0\x07\xE3\xE0\x03\xC1\xE0\x03\xC1\xE0\x03\x81\xC0\x01\x80\xC0\x01\x80\x80" +
	/* x */ "\x00\x29" + "\x00\x00\x00\x78" + "\x10\x10" + "\x12\x01" + "\xF1" + "\xFF\x3F\x7E\x0C\x3E\x08\x3F\x18\x1F\x30\x0F\xE0\x0F\xC0\x07\xE0\x03\xE0\x03\xF0\x05\xF8\x0C\xF8\x18\xFC\x30\x7E\x70\x7E\xFC\xFF" +
	/* y */ "\x00\x37" + "\x00\x00\x00\x79" + "\x10\x17" + "\x11\x00" + "\xF1" + "\xFF\x3F\x7E\x0C\x7C\x0C\x3E\x08\x3E\x08\x1E\x18\x1F\x10\x0F\x30\x0F\xA0\x0F\xA0\x07\xE0\x07\xC0\x03\xC0\x03\x80\x01\x80\x01\x80\x01\x00\x01\x00\x61\x00\xF2\x00\xF6\x00\xFC\x00\x78\x00" +
	/* z */ "\x00\x25" + "\x00\x00\x00\x7A" + "\x0E\x10" + "\x10\x00" + "\xF1" + "\x7F\xFD\xFF\xF7\x0F\xD0\x3E\x01\xF0\x0F\xC0\x3E\x01\xF0\x0F\xC0\x3E\x01\xF8\x0F\xC1\x3E\x05\xF8\x7F\xFF\xFF\xFF" +
	/* { */ "\x00\x34" + "\x00\x00\x00\x7B" + "\x0B\x1F" + "\x0E\x01" + "\xE8" + "\x01\xE0\xF8\x3E\x07\x80\xF0\x1E\x03\xC0\x78\x0F\x01\xE0\x3C\x07\x80\xF0\x1E\x07\x87\x80\x1E\x01\xE0\x3C\x07\x80\xF0\x1E\x03\xC0\x78\x0F\x01\xE0\x3C\x07\x80\xF8\x0F\x80\x78" +
	/* | */ "\x00\x13" + "\x00\x00\x00\x7C" + "\x03\x19" + "\x08\x02" + "\xE9" + "\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xE0" +
	/* } */ "\x00\x34" + "\x00\x00\x00\x7D" + "\x0B\x1F" + "\x0E\x03" + "\xE8" + "\xF0\x0F\x80\xF0\x0F\x01\xE0\x3C\x07\x80\xF0\x1E\x03\xC0\x78\x0F\x01\xE0\x3C\x03\xC0\x0F\x0F\x03\xC0\x78\x0F\x01\xE0\x3C\x07\x80\xF0\x1E\x03\xC0\x78\x0F\x03\xE0\xF8\x3C\x00" +
	/* ~ */ "\x00\x13" + "\x00\x00\x00\x7E" + "\x10\x05" + "\x12\x01" + "\xF5" + "\x3E\x00\x7F\xC6\xFF\xFF\x61\xFE\x00\x7C" +
	""

func (f *sBold18pt7b) GetGlyph(r rune) tinyfont.Glyph {
	idx := -1

	switch r {
	case 0x0020:
		idx = 0
	case 0x0021:
		idx = 11
	case 0x0022:
		idx = 40
	case 0x0023:
		idx = 68
	case 0x0024:
		idx = 133
	case 0x0025:
		idx = 197
	case 0x0026:
		idx = 289
	case 0x0027:
		idx = 382
	case 0x0028:
		idx = 398
	case 0x0029:
		idx = 443
	case 0x002A:
		idx = 488
	case 0x002B:
		idx = 526
	case 0x002C:
		idx = 583
	case 0x002D:
		idx = 603
	case 0x002E:
		idx = 618
	case 0x002F:
		idx = 634
	case 0x0030:
		idx = 680
	case 0x0031:
		idx = 739
	case 0x0032:
		idx = 786
	case 0x0033:
		idx = 845
	case 0x0034:
		idx = 904
	case 0x0035:
		idx = 960
	case 0x0036:
		idx = 1016
	case 0x0037:
		idx = 1075
	case 0x0038:
		idx = 1134
	case 0x0039:
		idx = 1193
	case 0x003A:
		idx = 1252
	case 0x003B:
		idx = 1275
	case 0x003C:
		idx = 1306
	case 0x003D:
		idx = 1365
	case 0x003E:
		idx = 1405
	case 0x003F:
		idx = 1464
	case 0x0040:
		idx = 1517
	case 0x0041:
		idx = 1603
	case 0x0042:
		idx = 1686
	case 0x0043:
		idx = 1760
	case 0x0044:
		idx = 1843
	case 0x0045:
		idx = 1923
	case 0x0046:
		idx = 1997
	case 0x0047:
		idx = 2065
	case 0x0048:
		idx = 2155
	case 0x0049:
		idx = 2238
	case 0x004A:
		idx = 2282
	case 0x004B:
		idx = 2347
	case 0x004C:
		idx = 2433
	case 0x004D:
		idx = 2507
	case 0x004E:
		idx = 2611
	case 0x004F:
		idx = 2691
	case 0x0050:
		idx = 2781
	case 0x0051:
		idx = 2849
	case 0x0052:
		idx = 2954
	case 0x0053:
		idx = 3034
	case 0x0054:
		idx = 3095
	case 0x0055:
		idx = 3169
	case 0x0056:
		idx = 3249
	case 0x0057:
		idx = 3332
	case 0x0058:
		idx = 3450
	case 0x0059:
		idx = 3533
	case 0x005A:
		idx = 3616
	case 0x005B:
		idx = 3690
	case 0x005C:
		idx = 3730
	case 0x005D:
		idx = 3776
	case 0x005E:
		idx = 3816
	case 0x005F:
		idx = 3852
	case 0x0060:
		idx = 3870
	case 0x0061:
		idx = 3887
	case 0x0062:
		idx = 3930
	case 0x0063:
		idx = 3995
	case 0x0064:
		idx = 4034
	case 0x0065:
		idx = 4099
	case 0x0066:
		idx = 4138
	case 0x0067:
		idx = 4182
	case 0x0068:
		idx = 4242
	case 0x0069:
		idx = 4304
	case 0x006A:
		idx = 4336
	case 0x006B:
		idx = 4390
	case 0x006C:
		idx = 4455
	case 0x006D:
		idx = 4487
	case 0x006E:
		idx = 4552
	case 0x006F:
		idx = 4597
	case 0x0070:
		idx = 4642
	case 0x0071:
		idx = 4708
	case 0x0072:
		idx = 4768
	case 0x0073:
		idx = 4805
	case 0x0074:
		idx = 4840
	case 0x0075:
		idx = 4878
	case 0x0076:
		idx = 4925
	case 0x0077:
		idx = 4970
	case 0x0078:
		idx = 5029
	case 0x0079:
		idx = 5072
	case 0x007A:
		idx = 5129
	case 0x007B:
		idx = 5168
	case 0x007C:
		idx = 5222
	case 0x007D:
		idx = 5243
	case 0x007E:
		idx = 5297
	}

	if idx == -1 {
		return tinyfont.Glyph{
			Rune:     r,
			Width:    0,
			Height:   0,
			XAdvance: uint8(cBold18pt7b[6]),
			XOffset:  int8(cBold18pt7b[7]),
			YOffset:  int8(cBold18pt7b[8]),
			Bitmaps:  []uint8{},
		}
	}

	length := int((uint16(cBold18pt7b[idx+0]) << 8) + uint16(cBold18pt7b[idx+1]))
	idx += 2
	ret := tinyfont.Glyph{
		Rune:     rune((uint32(cBold18pt7b[idx+0]) << 24) + (uint32(cBold18pt7b[idx+1]) << 16) + (uint32(cBold18pt7b[idx+2]) << 8) + uint32(cBold18pt7b[idx+3])),
		Width:    uint8(cBold18pt7b[idx+4]),
		Height:   uint8(cBold18pt7b[idx+5]),
		XAdvance: uint8(cBold18pt7b[idx+6]),
		XOffset:  int8(cBold18pt7b[idx+7]),
		YOffset:  int8(cBold18pt7b[idx+8]),
		Bitmaps:  []uint8(cBold18pt7b[idx+9 : idx+length]),
	}

	return ret
}

func (f *sBold18pt7b) GetYAdvance() uint8 {
	return 0x2A
}
