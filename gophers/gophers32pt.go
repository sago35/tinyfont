package gophers

import (
	"tinygo.org/x/tinyfont"
)

var Regular32pt = sRegular32pt{}

type sRegular32pt struct {
}

const cRegular32pt = "" +
	/*   */ "\x00\x0A" + "\x00\x00\x00\x20" + "\x01\x01" + "\x0B\x00" + "\x00" + "\x00" +
	/* ! */ "\x00\x1E" + "\x00\x00\x00\x21" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* " */ "\x00\x1E" + "\x00\x00\x00\x22" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* # */ "\x00\x1E" + "\x00\x00\x00\x23" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* $ */ "\x00\x1E" + "\x00\x00\x00\x24" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* % */ "\x00\x1E" + "\x00\x00\x00\x25" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* & */ "\x00\x1E" + "\x00\x00\x00\x26" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* ' */ "\x00\x1E" + "\x00\x00\x00\x27" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* ( */ "\x00\x1E" + "\x00\x00\x00\x28" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* ) */ "\x00\x1E" + "\x00\x00\x00\x29" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* * */ "\x00\x1E" + "\x00\x00\x00\x2A" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* + */ "\x00\x1E" + "\x00\x00\x00\x2B" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* , */ "\x00\x1E" + "\x00\x00\x00\x2C" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* - */ "\x00\x1E" + "\x00\x00\x00\x2D" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* . */ "\x00\x1E" + "\x00\x00\x00\x2E" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* / */ "\x00\x1E" + "\x00\x00\x00\x2F" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* 0 */ "\x00\x1E" + "\x00\x00\x00\x30" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* 1 */ "\x00\x1E" + "\x00\x00\x00\x31" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* 2 */ "\x00\x1E" + "\x00\x00\x00\x32" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* 3 */ "\x00\x1E" + "\x00\x00\x00\x33" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* 4 */ "\x00\x1E" + "\x00\x00\x00\x34" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* 5 */ "\x00\x1E" + "\x00\x00\x00\x35" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* 6 */ "\x00\x1E" + "\x00\x00\x00\x36" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* 7 */ "\x00\x1E" + "\x00\x00\x00\x37" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* 8 */ "\x00\x1E" + "\x00\x00\x00\x38" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* 9 */ "\x00\x1E" + "\x00\x00\x00\x39" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* : */ "\x00\x1E" + "\x00\x00\x00\x3A" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* ; */ "\x00\x1E" + "\x00\x00\x00\x3B" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* < */ "\x00\x1E" + "\x00\x00\x00\x3C" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* = */ "\x00\x1E" + "\x00\x00\x00\x3D" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* > */ "\x00\x1E" + "\x00\x00\x00\x3E" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* ? */ "\x00\x1E" + "\x00\x00\x00\x3F" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* @ */ "\x00\x1E" + "\x00\x00\x00\x40" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* A */ "\x00\x3B" + "\x00\x00\x00\x41" + "\x12\x16" + "\x15\x00" + "\xEA" + "\x08\x00\x05\x73\xC1\x23\x20\x70\x38\x62\x31\x28\x90\xAB\x14\x6A\x89\x02\x82\x40\x91\x28\xE3\x95\xC8\x83\x02\x20\x80\xD8\x00\x36\x00\x08\x80\x04\x20\x01\x04\x00\x41\x00\x10\x20\x18\x07\xFC\x01\x00\x00" +
	/* B */ "\x00\x48" + "\x00\x00\x00\x42" + "\x14\x19" + "\x17\x00" + "\xE7" + "\x00\x00\x00\xE0\x38\x0D\xFE\x40\x80\x08\x3E\x07\xC4\x20\x84\x81\x10\x29\x11\x32\x99\x13\x28\x11\x02\x41\x48\x47\xEE\x7C\x40\xE0\x44\x00\x02\x40\x00\x54\x00\x06\xA0\x00\x46\x00\x04\x20\x00\x42\x00\x04\x10\x00\x40\x80\x08\x04\x03\x00\x3F\xD0\x03\x03\x00" +
	/* C */ "\x00\x4A" + "\x00\x00\x00\x43" + "\x1A\x14" + "\x1D\x00" + "\xE9" + "\x07\x00\x00\x02\x20\x01\xC0\xA9\xF9\x90\x11\x81\xF4\x00\x00\x0A\x03\x00\x0F\x83\x38\x0C\x11\x01\x04\x02\x90\x41\x00\xEE\x08\x9C\x1B\xC2\x2F\x06\x60\x89\x81\x80\x21\x00\xD0\x08\x60\x26\x04\x04\x30\x7E\x30\xF0\x00\x0E\x00\x00\x0C\x40\x00\x01\xE0\x00\x00\x70\x00" +
	/* D */ "\x00\x1B" + "\x00\x00\x00\x44" + "\x0B\x0D" + "\x0E\x00" + "\xF3" + "\x6F\x8A\x13\xDE\x6A\x69\x45\x5F\xB2\x1C\x02\x80\x48\x11\x02\x1D\x80\x40" +
	/* E */ "\x00\x43" + "\x00\x00\x00\x45" + "\x16\x15" + "\x1A\x01" + "\xEB" + "\x00\x00\x30\x00\x03\x23\x81\xC8\x49\xF9\x41\x30\x04\x89\x21\xF2\x2E\x49\xA5\x39\x26\x88\x84\x82\x42\x21\x19\x04\xB3\xA8\x1C\x40\xC0\xC0\x02\x03\x00\x08\x04\x00\x20\x10\x00\x80\x20\x04\x00\x80\x10\x01\x81\x80\x03\xFC\x00\x08\x00\x00" +
	/* F */ "\x00\x38" + "\x00\x00\x00\x46" + "\x11\x16" + "\x14\x00" + "\xEA" + "\x30\x00\x24\xCE\x0B\x9A\x04\x00\x84\x00\x44\x00\x22\x00\x09\x00\x04\x80\x02\x40\x01\xE0\x00\xD0\x00\x48\x00\x24\x00\x12\x20\x09\x1C\x04\x84\x04\x20\x02\x08\x02\x03\xFF\x01\x41\x00\x40\x00" +
	/* G */ "\x00\x3A" + "\x00\x00\x00\x47" + "\x11\x17" + "\x14\x00" + "\xE9" + "\x00\x00\x19\xF2\x0F\x07\x04\x01\x84\xE1\xC3\xA9\x53\x13\x65\x81\x82\xC1\x42\x50\x91\xA7\xB7\x70\x28\x38\x0C\x14\xE0\x0A\x60\x05\x00\x02\x80\x02\x20\x01\x10\x00\x86\x01\x80\xFF\x80\x30\xC0\x10\x00" +
	/* H */ "\x00\x27" + "\x00\x00\x00\x48" + "\x12\x0D" + "\x18\x03" + "\xEF" + "\x00\x01\x00\x01\xF0\x1E\x84\x04\x21\x03\x08\x40\x9F\x30\x2D\xF8\x28\x20\x31\x10\x1B\x38\x03\x60\x01\x60\x00\xB0\x00\x00" +
	/* I */ "\x00\x3F" + "\x00\x00\x00\x49" + "\x12\x18" + "\x15\x00" + "\xE8" + "\x00\x00\x08\x33\x85\xF3\xE0\xC0\x10\x20\x0E\x13\x0E\xC7\xF2\x11\x85\x02\x41\x42\x90\x50\xC8\x72\x11\x14\x7C\x72\xE1\x13\x38\x44\x06\x19\x00\x06\xA0\x01\x38\x00\x41\x00\x20\x20\x18\x07\xFE\x01\x40\x80\x50\x00\x08\x00" +
	/* J */ "\x00\x36" + "\x00\x00\x00\x4A" + "\x11\x15" + "\x14\x00" + "\xEA" + "\x08\x00\x0B\xFE\x04\x01\x03\x00\x86\x43\xA5\x12\x32\x89\x15\x04\x82\x44\x43\x1D\xFF\x08\x60\xEC\x00\x6A\x00\x23\xFF\xF0\x80\x08\x80\x0A\x54\x95\x4A\x4A\xB9\x25\x23\xFF\xE0\xC0\xC0" +
	/* K */ "\x00\x20" + "\x00\x00\x00\x4B" + "\x0E\x0D" + "\x13\x02" + "\xEA" + "\x04\x00\x7E\x01\xFC\x1F\xFC\x8F\x0A\x00\x28\x00\x90\x04\x20\x20\x7F\x01\x04\x04\x10\x0F\x80" +
	/* L */ "\x00\x20" + "\x00\x00\x00\x4C" + "\x0E\x0D" + "\x12\x01" + "\xEB" + "\x07\x80\x61\x86\x01\x10\x02\x8C\x0A\x58\x19\xE0\x61\x01\x80\x0A\x00\x24\x01\x08\x08\x1F\xC0" +
	/* M */ "\x00\x39" + "\x00\x00\x00\x4D" + "\x12\x15" + "\x16\x00" + "\xEA" + "\x30\x00\x16\xFE\x02\x41\x81\x00\x20\x9E\xF4\x28\x67\x0C\x51\xC2\x06\x10\xC2\x84\x31\x7F\x0B\x98\x42\x02\x1C\x80\x02\x28\x01\x79\x00\x4E\x00\x10\x40\x08\x18\x02\x01\x83\x00\x5F\xC0\x18\x00\x00" +
	/* N */ "\x00\x39" + "\x00\x00\x00\x4E" + "\x12\x15" + "\x15\x00" + "\xEB" + "\x00\x00\x07\x7B\x00\xA1\x40\xC0\x38\x4C\x31\x20\x90\x4B\x24\xD2\x85\x14\x81\x41\x19\xA8\xC5\x95\xD0\x82\x06\x60\x01\xB8\x00\x42\x00\x10\x80\x08\x10\x02\x04\x00\x80\xC0\xC0\x1F\xD0\x04\x08\x00" +
	/* O */ "\x00\x49" + "\x00\x00\x00\x4F" + "\x16\x17" + "\x1A\x02" + "\xE9" + "\x00\x00\x00\x07\xE0\x00\xE0\x60\x04\x00\x60\x20\x02\x41\x00\x1E\x84\x00\xF9\x20\x03\xE4\x80\x07\x0A\x00\x08\x28\x00\x00\xA0\x00\x02\x80\x00\x0A\x00\x00\x28\x00\x01\x20\x00\x04\x40\x00\x31\x00\x00\x82\x00\x06\x04\x00\x30\x08\x01\x80\x18\x18\x00\x1F\x80\x00" +
	/* P */ "\x00\x4A" + "\x00\x00\x00\x50" + "\x1A\x14" + "\x1C\xFF" + "\xEC" + "\x03\x00\x00\x01\x20\x03\x80\x5C\xFD\xA0\x08\xC0\x90\x02\x00\x04\x03\x40\x0E\x83\x0C\x0C\x71\x00\x84\x02\x48\x21\x00\xA7\x08\x83\x99\xC1\x20\xE6\x00\x84\x01\x40\x21\x00\x90\x10\x20\x23\x18\xE4\x18\xB8\x44\xFA\x20\x1A\x01\x04\x03\x00\x41\x00\x00\x10\x40\x00\x00" +
	/* Q */ "\x00\x39" + "\x00\x00\x00\x51" + "\x18\x10" + "\x1C\x01" + "\xF0" + "\x80\x00\x01\x80\x00\x01\x80\x00\x01\x80\x00\x01\x80\x00\x01\x80\x00\x01\x81\x80\x02\x81\x40\x02\x40\x80\x02\x40\x00\x04\x20\x00\x08\x18\x00\x10\x07\x00\xF0\x05\xFF\x28\x05\x00\x38\x02\x00\x00" +
	/* R */ "\x00\x3D" + "\x00\x00\x00\x52" + "\x12\x17" + "\x14\xFF" + "\xE9" + "\x00\xC0\x00\xC8\x00\x3A\x3C\x0C\x70\x01\x00\x01\xE0\x01\xC6\x00\xC0\x40\x40\x10\x10\x02\x08\xC0\x82\x70\x20\x8C\x08\x20\x04\x04\x01\x01\x00\x80\x20\x43\x86\x21\xE1\x70\x44\x40\x0E\x08\x02\x82\x00\x40\x80\x00" +
	/* S */ "\x00\x2D" + "\x00\x00\x00\x53" + "\x1A\x0B" + "\x1D\x00" + "\xF2" + "\x0F\x00\x3E\x0C\x30\x30\x44\x02\x10\x0A\x60\x88\x1A\xB8\x22\x0E\x66\x08\x83\x98\x02\x10\x05\x00\x84\x02\x40\x40\x80\x88\x20\x10\xC1\xF0\x03\xC0" +
	/* T */ "\x00\x32" + "\x00\x00\x00\x54" + "\x0E\x17" + "\x11\x00" + "\xE9" + "\x00\x40\x46\x8E\xEA\x00\x10\x00\x20\x0F\x80\xC1\x04\x12\x10\xC8\x43\x21\x00\x82\x02\x28\x11\xDF\xCD\x01\x18\x04\x00\x10\x00\x40\x00\x80\x02\x00\x08\x00\x20\x01\x00" +
	/* U */ "\x00\x2D" + "\x00\x00\x00\x55" + "\x1A\x0B" + "\x1D\x00" + "\xF2" + "\x0F\x00\x3E\x0C\x30\x30\x44\x02\x10\x0A\x06\x89\x82\x83\xA2\xE0\x60\x68\xB8\x18\x02\x10\x05\x00\x84\x02\x40\x40\x80\x88\x20\x10\xC1\xF0\x03\xC0" +
	/* V */ "\x00\x4B" + "\x00\x00\x00\x56" + "\x16\x18" + "\x1A\x01" + "\xE8" + "\x00\x00\x00\x00\x07\x80\x00\x22\x00\x00\x8C\x00\x02\xC8\xE0\x7A\x12\x7E\x50\x4C\x01\x22\x48\x7C\x8B\x92\x69\x4E\x49\xA2\x21\x20\x90\x88\x46\x41\x2C\xEA\x07\x10\x30\x30\x00\x80\xC0\x02\x01\x00\x08\x04\x00\x20\x08\x01\x00\x20\x04\x00\x60\x60\x00\xFF\x00\x02\x00\x00" +
	/* W */ "\x00\x48" + "\x00\x00\x00\x57" + "\x14\x19" + "\x17\x00" + "\xE7" + "\x00\x00\x00\xE0\x38\x0D\xFC\x80\x80\x08\x3E\x07\x84\x20\x84\x81\x10\x29\x11\x32\x99\x13\x28\x11\x02\x41\x08\x47\xEE\x7C\x40\xE0\xC4\x04\x36\x3F\xFC\x54\x00\x06\xC0\x00\x46\x00\x04\x20\x1F\xC1\xFE\x04\x10\x00\x40\x80\x08\x04\x01\x00\x3F\xF0\x03\x03\x00" +
	/* X */ "\x00\x4E" + "\x00\x00\x00\x58" + "\x16\x19" + "\x19\x00" + "\xE7" + "\x00\x00\x00\x38\x0E\x00\xDF\xE8\x02\x00\x20\x3E\x07\xC1\x08\x21\x08\x11\x02\x24\x44\xC8\x99\x13\x22\x04\x40\x84\x14\x84\x1F\xB9\xF0\x40\xE0\x4D\x00\x01\xE2\x03\x8F\x08\xFF\xE8\x67\xFC\x21\xBF\xF0\x82\xFF\x82\x09\xF8\x08\x13\xC0\x40\x20\x02\x00\x60\x10\x00\xFF\xC0\x03\x03\x00" +
	/* Y */ "\x00\x48" + "\x00\x00\x00\x59" + "\x14\x19" + "\x18\x01" + "\xE7" + "\x00\x00\x00\xE0\x38\x0D\xFE\x80\x80\x08\x1C\x00\x87\xE0\x7C\x7F\x1F\xCF\xF1\xFE\xFF\xFF\xEF\xF1\xFE\xFF\x0F\xC7\xE6\xFC\x40\xA0\x44\x06\x04\x40\x00\x64\x00\x06\xA0\x00\x46\x00\x04\x20\x00\x42\x00\x04\x10\x00\x40\x80\x08\x04\x03\x00\x3F\xF0\x03\x03\x00" +
	/* Z */ "\x00\x57" + "\x00\x00\x00\x5A" + "\x1A\x18" + "\x1B\xFE" + "\xE8" + "\x01\xC0\x30\x00\x4F\xF2\x00\x1C\x03\x00\x0E\x00\x20\x04\x41\xF8\x02\x08\xC3\x01\x12\x24\x40\x4C\x89\x90\x18\x22\x04\x03\x08\x42\x10\xBC\xE9\x88\x10\x31\xA4\x84\x04\x05\x19\x00\x02\x81\xC0\x00\xC0\x30\x00\x20\x04\x00\x04\x00\x80\x02\x00\x20\x00\x80\x04\x00\x20\x00\xC0\x70\x00\x1F\xF8\x00\x0E\x06\x00\x00\x00\x00" +
	/* [ */ "\x00\x1E" + "\x00\x00\x00\x5B" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* \ */ "\x00\x1E" + "\x00\x00\x00\x5C" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* ] */ "\x00\x1E" + "\x00\x00\x00\x5D" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* ^ */ "\x00\x1E" + "\x00\x00\x00\x5E" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* _ */ "\x00\x1E" + "\x00\x00\x00\x5F" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* ` */ "\x00\x1E" + "\x00\x00\x00\x60" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* a */ "\x00\x1E" + "\x00\x00\x00\x61" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* b */ "\x00\x1E" + "\x00\x00\x00\x62" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* c */ "\x00\x1E" + "\x00\x00\x00\x63" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* d */ "\x00\x1E" + "\x00\x00\x00\x64" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* e */ "\x00\x1E" + "\x00\x00\x00\x65" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* f */ "\x00\x1E" + "\x00\x00\x00\x66" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* g */ "\x00\x1E" + "\x00\x00\x00\x67" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* h */ "\x00\x1E" + "\x00\x00\x00\x68" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* i */ "\x00\x1E" + "\x00\x00\x00\x69" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* j */ "\x00\x1E" + "\x00\x00\x00\x6A" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* k */ "\x00\x1E" + "\x00\x00\x00\x6B" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* l */ "\x00\x1E" + "\x00\x00\x00\x6C" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* m */ "\x00\x1E" + "\x00\x00\x00\x6D" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* n */ "\x00\x1E" + "\x00\x00\x00\x6E" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* o */ "\x00\x1E" + "\x00\x00\x00\x6F" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* p */ "\x00\x1E" + "\x00\x00\x00\x70" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* q */ "\x00\x1E" + "\x00\x00\x00\x71" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* r */ "\x00\x1E" + "\x00\x00\x00\x72" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* s */ "\x00\x1E" + "\x00\x00\x00\x73" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* t */ "\x00\x1E" + "\x00\x00\x00\x74" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* u */ "\x00\x1E" + "\x00\x00\x00\x75" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* v */ "\x00\x1E" + "\x00\x00\x00\x76" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* w */ "\x00\x1E" + "\x00\x00\x00\x77" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* x */ "\x00\x1E" + "\x00\x00\x00\x78" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* y */ "\x00\x1E" + "\x00\x00\x00\x79" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* z */ "\x00\x1E" + "\x00\x00\x00\x7A" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* { */ "\x00\x1E" + "\x00\x00\x00\x7B" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* | */ "\x00\x1E" + "\x00\x00\x00\x7C" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	/* } */ "\x00\x1E" + "\x00\x00\x00\x7D" + "\x08\x15" + "\x0D\x01" + "\xEB" + "\xFF\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\x81\xFF" +
	""

func (f *sRegular32pt) GetGlyph(r rune) tinyfont.Glyph {
	idx := -1

	switch r {
	case 0x0020:
		idx = 0
	case 0x0021:
		idx = 12
	case 0x0022:
		idx = 44
	case 0x0023:
		idx = 76
	case 0x0024:
		idx = 108
	case 0x0025:
		idx = 140
	case 0x0026:
		idx = 172
	case 0x0027:
		idx = 204
	case 0x0028:
		idx = 236
	case 0x0029:
		idx = 268
	case 0x002A:
		idx = 300
	case 0x002B:
		idx = 332
	case 0x002C:
		idx = 364
	case 0x002D:
		idx = 396
	case 0x002E:
		idx = 428
	case 0x002F:
		idx = 460
	case 0x0030:
		idx = 492
	case 0x0031:
		idx = 524
	case 0x0032:
		idx = 556
	case 0x0033:
		idx = 588
	case 0x0034:
		idx = 620
	case 0x0035:
		idx = 652
	case 0x0036:
		idx = 684
	case 0x0037:
		idx = 716
	case 0x0038:
		idx = 748
	case 0x0039:
		idx = 780
	case 0x003A:
		idx = 812
	case 0x003B:
		idx = 844
	case 0x003C:
		idx = 876
	case 0x003D:
		idx = 908
	case 0x003E:
		idx = 940
	case 0x003F:
		idx = 972
	case 0x0040:
		idx = 1004
	case 0x0041:
		idx = 1036
	case 0x0042:
		idx = 1097
	case 0x0043:
		idx = 1171
	case 0x0044:
		idx = 1247
	case 0x0045:
		idx = 1276
	case 0x0046:
		idx = 1345
	case 0x0047:
		idx = 1403
	case 0x0048:
		idx = 1463
	case 0x0049:
		idx = 1504
	case 0x004A:
		idx = 1569
	case 0x004B:
		idx = 1625
	case 0x004C:
		idx = 1659
	case 0x004D:
		idx = 1693
	case 0x004E:
		idx = 1752
	case 0x004F:
		idx = 1811
	case 0x0050:
		idx = 1886
	case 0x0051:
		idx = 1962
	case 0x0052:
		idx = 2021
	case 0x0053:
		idx = 2084
	case 0x0054:
		idx = 2131
	case 0x0055:
		idx = 2183
	case 0x0056:
		idx = 2230
	case 0x0057:
		idx = 2307
	case 0x0058:
		idx = 2381
	case 0x0059:
		idx = 2461
	case 0x005A:
		idx = 2535
	case 0x005B:
		idx = 2624
	case 0x005C:
		idx = 2656
	case 0x005D:
		idx = 2688
	case 0x005E:
		idx = 2720
	case 0x005F:
		idx = 2752
	case 0x0060:
		idx = 2784
	case 0x0061:
		idx = 2816
	case 0x0062:
		idx = 2848
	case 0x0063:
		idx = 2880
	case 0x0064:
		idx = 2912
	case 0x0065:
		idx = 2944
	case 0x0066:
		idx = 2976
	case 0x0067:
		idx = 3008
	case 0x0068:
		idx = 3040
	case 0x0069:
		idx = 3072
	case 0x006A:
		idx = 3104
	case 0x006B:
		idx = 3136
	case 0x006C:
		idx = 3168
	case 0x006D:
		idx = 3200
	case 0x006E:
		idx = 3232
	case 0x006F:
		idx = 3264
	case 0x0070:
		idx = 3296
	case 0x0071:
		idx = 3328
	case 0x0072:
		idx = 3360
	case 0x0073:
		idx = 3392
	case 0x0074:
		idx = 3424
	case 0x0075:
		idx = 3456
	case 0x0076:
		idx = 3488
	case 0x0077:
		idx = 3520
	case 0x0078:
		idx = 3552
	case 0x0079:
		idx = 3584
	case 0x007A:
		idx = 3616
	case 0x007B:
		idx = 3648
	case 0x007C:
		idx = 3680
	case 0x007D:
		idx = 3712
	}

	if idx == -1 {
		return tinyfont.Glyph{
			Rune:     r,
			Width:    0,
			Height:   0,
			XAdvance: uint8(cRegular32pt[6]),
			XOffset:  int8(cRegular32pt[7]),
			YOffset:  int8(cRegular32pt[8]),
			Bitmaps:  []uint8{},
		}
	}

	length := int((uint16(cRegular32pt[idx+0]) << 8) + uint16(cRegular32pt[idx+1]))
	idx += 2
	ret := tinyfont.Glyph{
		Rune:     rune((uint32(cRegular32pt[idx+0]) << 24) + (uint32(cRegular32pt[idx+1]) << 16) + (uint32(cRegular32pt[idx+2]) << 8) + uint32(cRegular32pt[idx+3])),
		Width:    uint8(cRegular32pt[idx+4]),
		Height:   uint8(cRegular32pt[idx+5]),
		XAdvance: uint8(cRegular32pt[idx+6]),
		XOffset:  int8(cRegular32pt[idx+7]),
		YOffset:  int8(cRegular32pt[idx+8]),
		Bitmaps:  []uint8(cRegular32pt[idx+9 : idx+length]),
	}

	return ret
}

func (f *sRegular32pt) GetYAdvance() uint8 {
	return 0x1C
}
