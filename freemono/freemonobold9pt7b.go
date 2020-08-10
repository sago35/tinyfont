package freemono

import (
	"tinygo.org/x/tinyfont"
)

var Bold9pt7b = sBold9pt7b{}

type sBold9pt7b struct {
}

const cBold9pt7b = "" +
	/*   */ "\x00\x09" + "\x00\x00\x00\x20" + "\x00\x00" + "\x0B\x00" + "\x01" + "" +
	/* ! */ "\x00\x0E" + "\x00\x00\x00\x21" + "\x03\x0B" + "\x0B\x04" + "\xF6" + "\xFF\xFF\xD2\x1F\x80" +
	/* " */ "\x00\x0E" + "\x00\x00\x00\x22" + "\x07\x05" + "\x0B\x02" + "\xF6" + "\xEC\x89\x12\x24\x40" +
	/* # */ "\x00\x15" + "\x00\x00\x00\x23" + "\x08\x0C" + "\x0B\x01" + "\xF6" + "\x36\x36\x36\x7F\x7F\x36\xFF\xFF\x3C\x3C\x3C\x00" +
	/* $ */ "\x00\x16" + "\x00\x00\x00\x24" + "\x07\x0E" + "\x0B\x02" + "\xF5" + "\x18\xFF\xFE\x3C\x1F\x1F\x83\x46\x8D\xF0\xC1\x83\x00" +
	/* % */ "\x00\x13" + "\x00\x00\x00\x25" + "\x07\x0B" + "\x0B\x02" + "\xF6" + "\x61\x22\x44\x86\x67\x37\x11\x22\x4C\x70" +
	/* & */ "\x00\x13" + "\x00\x00\x00\x26" + "\x08\x0A" + "\x0B\x01" + "\xF7" + "\x3C\x7E\x60\x60\x30\x7B\xDF\xCE\xFF\x7F" +
	/* ' */ "\x00\x0B" + "\x00\x00\x00\x27" + "\x03\x05" + "\x0B\x04" + "\xF6" + "\xC9\x24" +
	/* ( */ "\x00\x10" + "\x00\x00\x00\x28" + "\x04\x0E" + "\x0B\x05" + "\xF6" + "\x37\x66\xCC\xCC\xCC\x66\x31" +
	/* ) */ "\x00\x10" + "\x00\x00\x00\x29" + "\x04\x0E" + "\x0B\x02" + "\xF6" + "\xCE\x66\x33\x33\x33\x66\xC8" +
	/* * */ "\x00\x10" + "\x00\x00\x00\x2A" + "\x08\x07" + "\x0B\x02" + "\xF6" + "\x18\x18\xFF\xFF\x3C\x3C\x66" +
	/* + */ "\x00\x12" + "\x00\x00\x00\x2B" + "\x08\x09" + "\x0B\x02" + "\xF8" + "\x18\x18\x18\xFF\xFF\x18\x18\x18\x18" +
	/* , */ "\x00\x0B" + "\x00\x00\x00\x2C" + "\x03\x05" + "\x0B\x03" + "\xFF" + "\x6B\x48" +
	/* - */ "\x00\x0C" + "\x00\x00\x00\x2D" + "\x09\x02" + "\x0B\x01" + "\xFB" + "\xFF\xFF\xC0" +
	/* . */ "\x00\x0A" + "\x00\x00\x00\x2E" + "\x02\x02" + "\x0B\x04" + "\xFF" + "\xF0" +
	/* / */ "\x00\x17" + "\x00\x00\x00\x2F" + "\x07\x0F" + "\x0B\x02" + "\xF4" + "\x02\x0C\x18\x60\xC3\x06\x0C\x30\x61\x83\x0C\x18\x20\x00" +
	/* 0 */ "\x00\x14" + "\x00\x00\x00\x30" + "\x07\x0C" + "\x0B\x02" + "\xF5" + "\x38\xFB\xBE\x3C\x78\xF1\xE3\xC7\xDD\xF1\xC0" +
	/* 1 */ "\x00\x13" + "\x00\x00\x00\x31" + "\x07\x0B" + "\x0B\x02" + "\xF6" + "\x38\xF3\x60\xC1\x83\x06\x0C\x18\xFD\xF8" +
	/* 2 */ "\x00\x15" + "\x00\x00\x00\x32" + "\x08\x0C" + "\x0B\x01" + "\xF5" + "\x3C\xFE\xC7\x03\x03\x06\x0C\x18\x70\xE3\xFF\xFF" +
	/* 3 */ "\x00\x15" + "\x00\x00\x00\x33" + "\x08\x0C" + "\x0B\x02" + "\xF5" + "\x7C\xFE\x03\x03\x03\x1E\x1E\x07\x03\x03\xFE\x7C" +
	/* 4 */ "\x00\x12" + "\x00\x00\x00\x34" + "\x07\x0A" + "\x0B\x02" + "\xF7" + "\x1C\x38\xB1\x64\xD9\xBF\xFF\x3E\x7C" +
	/* 5 */ "\x00\x16" + "\x00\x00\x00\x35" + "\x09\x0B" + "\x0B\x01" + "\xF6" + "\x7E\x3F\x18\x0F\xC7\xF3\x1C\x06\x03\xC3\xFF\x9F\x80" +
	/* 6 */ "\x00\x15" + "\x00\x00\x00\x36" + "\x08\x0C" + "\x0B\x02" + "\xF5" + "\x0F\x3F\x30\x60\x60\xDC\xFE\xE3\xC3\x63\x7E\x3C" +
	/* 7 */ "\x00\x14" + "\x00\x00\x00\x37" + "\x08\x0B" + "\x0B\x01" + "\xF6" + "\xFF\xFF\xC3\x03\x06\x06\x06\x0C\x0C\x0C\x18" +
	/* 8 */ "\x00\x14" + "\x00\x00\x00\x38" + "\x07\x0C" + "\x0B\x02" + "\xF5" + "\x38\xFB\x1E\x3C\x6F\x9F\x63\xC7\x8F\xF1\xC0" +
	/* 9 */ "\x00\x15" + "\x00\x00\x00\x39" + "\x08\x0C" + "\x0B\x02" + "\xF5" + "\x3C\x7E\xE6\xC3\xC3\xE7\x7F\x3B\x06\x0E\xFC\xF0" +
	/* : */ "\x00\x0B" + "\x00\x00\x00\x3A" + "\x02\x08" + "\x0B\x04" + "\xF9" + "\xF0\x0F" +
	/* ; */ "\x00\x0E" + "\x00\x00\x00\x3B" + "\x03\x0B" + "\x0B\x03" + "\xF9" + "\x6C\x00\x1A\xD2\x00" +
	/* < */ "\x00\x12" + "\x00\x00\x00\x3C" + "\x09\x08" + "\x0B\x01" + "\xF8" + "\x01\x83\x87\x0E\x0F\x80\xE0\x1C\x03" +
	/* = */ "\x00\x10" + "\x00\x00\x00\x3D" + "\x09\x06" + "\x0B\x01" + "\xF9" + "\xFF\xFF\xC0\x00\x0F\xFF\xFC" +
	/* > */ "\x00\x12" + "\x00\x00\x00\x3E" + "\x09\x08" + "\x0B\x01" + "\xF8" + "\xC0\x78\x0F\x01\xE0\xF9\xE3\xC1\x80" +
	/* ? */ "\x00\x14" + "\x00\x00\x00\x3F" + "\x08\x0B" + "\x0B\x02" + "\xF6" + "\x7C\xFE\xC7\x03\x0E\x1C\x00\x00\x00\x30\x30" +
	/* @ */ "\x00\x1A" + "\x00\x00\x00\x40" + "\x09\x0F" + "\x0B\x01" + "\xF5" + "\x1E\x1F\x1C\xDC\x6C\x76\x7B\x6D\xB6\xDB\x6F\xF3\xFC\x06\x33\xF8\x78" +
	/* A */ "\x00\x19" + "\x00\x00\x00\x41" + "\x0B\x0B" + "\x0B\x00" + "\xF6" + "\x3C\x07\xC0\x38\x05\x81\xB0\x36\x0F\xE1\xFC\x71\xDF\x7F\xEF\x80" +
	/* B */ "\x00\x17" + "\x00\x00\x00\x42" + "\x0A\x0B" + "\x0B\x01" + "\xF6" + "\xFF\x3F\xE6\x19\x86\x7F\x1F\xE6\x1D\x83\x60\xFF\xFF\xF0" +
	/* C */ "\x00\x16" + "\x00\x00\x00\x43" + "\x09\x0B" + "\x0B\x01" + "\xF6" + "\x1F\xBF\xD8\xF8\x3C\x06\x03\x01\x80\x61\xBF\xC7\xC0" +
	/* D */ "\x00\x17" + "\x00\x00\x00\x44" + "\x0A\x0B" + "\x0B\x00" + "\xF6" + "\xFE\x3F\xE6\x19\x83\x60\xD8\x36\x0D\x83\x61\xBF\xEF\xE0" +
	/* E */ "\x00\x16" + "\x00\x00\x00\x45" + "\x09\x0B" + "\x0B\x01" + "\xF6" + "\xFF\xFF\xD8\x6D\xB7\xC3\xE1\xB0\xC3\x61\xFF\xFF\xE0" +
	/* F */ "\x00\x16" + "\x00\x00\x00\x46" + "\x09\x0B" + "\x0B\x01" + "\xF6" + "\xFF\xFF\xD8\x6D\xB7\xC3\xE1\xB0\xC0\x60\x7C\x3E\x00" +
	/* G */ "\x00\x17" + "\x00\x00\x00\x47" + "\x0A\x0B" + "\x0B\x01" + "\xF6" + "\x1F\x9F\xE6\x1B\x06\xC0\x30\x0C\x7F\x1F\xE1\x9F\xE3\xF0" +
	/* H */ "\x00\x16" + "\x00\x00\x00\x48" + "\x09\x0B" + "\x0B\x01" + "\xF6" + "\xF7\xFB\xD8\xCC\x66\x33\xF9\xFC\xC6\x63\x7B\xFD\xE0" +
	/* I */ "\x00\x12" + "\x00\x00\x00\x49" + "\x06\x0B" + "\x0B\x03" + "\xF6" + "\xFF\xF3\x0C\x30\xC3\x0C\x33\xFF\xC0" +
	/* J */ "\x00\x17" + "\x00\x00\x00\x4A" + "\x0A\x0B" + "\x0B\x01" + "\xF6" + "\x1F\xC7\xF0\x30\x0C\x03\x00\xCC\x33\x0C\xC7\x3F\x87\xC0" +
	/* K */ "\x00\x17" + "\x00\x00\x00\x4B" + "\x0A\x0B" + "\x0B\x01" + "\xF6" + "\xF7\xBD\xE6\x61\xB0\x78\x1F\x06\xE1\x98\x63\x3C\xFF\x3C" +
	/* L */ "\x00\x16" + "\x00\x00\x00\x4C" + "\x09\x0B" + "\x0B\x01" + "\xF6" + "\xFC\x7E\x0C\x06\x03\x01\x80\xC6\x63\x31\xFF\xFF\xE0" +
	/* M */ "\x00\x19" + "\x00\x00\x00\x4D" + "\x0B\x0B" + "\x0B\x00" + "\xF6" + "\xE0\xFE\x3D\xC7\x3D\xE7\xBC\xD7\x9B\xB3\x76\x60\xDE\x3F\xC7\x80" +
	/* N */ "\x00\x19" + "\x00\x00\x00\x4E" + "\x0B\x0B" + "\x0B\x00" + "\xF6" + "\xE1\xFE\x3D\xE3\x3C\x66\xCC\xDD\x99\xB3\x1E\x63\xDE\x3B\xC3\x00" +
	/* O */ "\x00\x19" + "\x00\x00\x00\x4F" + "\x0B\x0B" + "\x0B\x00" + "\xF6" + "\x1F\x07\xF1\xC7\x70\x7C\x07\x80\xF0\x1F\x07\x71\xC7\xF0\x7C\x00" +
	/* P */ "\x00\x16" + "\x00\x00\x00\x50" + "\x09\x0B" + "\x0B\x01" + "\xF6" + "\xFE\x7F\x98\x6C\x36\x1B\xF9\xF8\xC0\x60\x7C\x3E\x00" +
	/* Q */ "\x00\x1D" + "\x00\x00\x00\x51" + "\x0B\x0E" + "\x0B\x00" + "\xF6" + "\x1F\x07\xF1\xC7\x70\x7C\x07\x80\xF0\x1F\x07\x71\xC7\xF0\x7C\x0C\x33\xFE\x7F\x80" +
	/* R */ "\x00\x16" + "\x00\x00\x00\x52" + "\x09\x0B" + "\x0B\x01" + "\xF6" + "\xFC\x7F\x18\xCC\x66\x73\xF1\xF0\xCC\x63\x7D\xFE\x60" +
	/* S */ "\x00\x16" + "\x00\x00\x00\x53" + "\x09\x0B" + "\x0B\x01" + "\xF6" + "\x3F\xBF\xF0\x78\x0F\x03\xF8\x3F\x83\xC3\xFF\xBF\x80" +
	/* T */ "\x00\x16" + "\x00\x00\x00\x54" + "\x09\x0B" + "\x0B\x01" + "\xF6" + "\xFF\xFF\xF6\x7B\x3D\x98\xC0\x60\x30\x18\x3F\x1F\x80" +
	/* U */ "\x00\x19" + "\x00\x00\x00\x55" + "\x0B\x0B" + "\x0B\x00" + "\xF6" + "\xF1\xFE\x3D\x83\x30\x66\x0C\xC1\x98\x33\x06\x60\xC7\xF0\x7C\x00" +
	/* V */ "\x00\x19" + "\x00\x00\x00\x56" + "\x0B\x0B" + "\x0B\x00" + "\xF6" + "\xFB\xFF\x7D\xC3\x18\xC3\x18\x36\x06\xC0\x50\x0E\x01\xC0\x10\x00" +
	/* W */ "\x00\x17" + "\x00\x00\x00\x57" + "\x0A\x0B" + "\x0B\x00" + "\xF6" + "\xFB\xFE\xF6\x0D\x93\x6E\xDB\xB7\xAD\xEE\x7B\x8E\xE3\x18" +
	/* X */ "\x00\x17" + "\x00\x00\x00\x58" + "\x0A\x0B" + "\x0B\x00" + "\xF6" + "\xF3\xFC\xF7\x38\xFC\x1E\x03\x01\xE0\xCC\x73\xBC\xFF\x3C" +
	/* Y */ "\x00\x17" + "\x00\x00\x00\x59" + "\x0A\x0B" + "\x0B\x00" + "\xF6" + "\xF3\xFC\xF7\x38\xCC\x1E\x07\x80\xC0\x30\x0C\x0F\xC3\xF0" +
	/* Z */ "\x00\x14" + "\x00\x00\x00\x5A" + "\x08\x0B" + "\x0B\x02" + "\xF6" + "\xFE\xFE\xC6\xCC\x18\x18\x30\x63\xC3\xFF\xFF" +
	/* [ */ "\x00\x10" + "\x00\x00\x00\x5B" + "\x04\x0E" + "\x0B\x05" + "\xF6" + "\xFF\xCC\xCC\xCC\xCC\xCC\xFF" +
	/* \ */ "\x00\x17" + "\x00\x00\x00\x5C" + "\x07\x0F" + "\x0B\x02" + "\xF4" + "\x01\x03\x06\x06\x0C\x0C\x18\x18\x30\x30\x60\x60\xC0\x80" +
	/* ] */ "\x00\x10" + "\x00\x00\x00\x5D" + "\x04\x0E" + "\x0B\x02" + "\xF6" + "\xFF\x33\x33\x33\x33\x33\xFF" +
	/* ^ */ "\x00\x0F" + "\x00\x00\x00\x5E" + "\x07\x06" + "\x0B\x02" + "\xF5" + "\x10\x71\xE3\x6C\x70\x40" +
	/* _ */ "\x00\x0C" + "\x00\x00\x00\x5F" + "\x0B\x02" + "\x0B\x00" + "\x03" + "\xFF\xFF\xFC" +
	/* ` */ "\x00\x0B" + "\x00\x00\x00\x60" + "\x03\x03" + "\x0B\x03" + "\xF5" + "\x88\x80" +
	/* a */ "\x00\x12" + "\x00\x00\x00\x61" + "\x09\x08" + "\x0B\x01" + "\xF9" + "\x7E\x3F\x8F\xCF\xEE\x36\x1B\xFE\xFF" +
	/* b */ "\x00\x17" + "\x00\x00\x00\x62" + "\x0A\x0B" + "\x0B\x00" + "\xF6" + "\xE0\x38\x06\x01\xBC\x7F\x9C\x76\x0D\x83\x71\xFF\xEE\xF0" +
	/* c */ "\x00\x12" + "\x00\x00\x00\x63" + "\x09\x08" + "\x0B\x01" + "\xF9" + "\x3F\xBF\xF8\x78\x3C\x07\x05\xFE\x7E" +
	/* d */ "\x00\x17" + "\x00\x00\x00\x64" + "\x0A\x0B" + "\x0B\x01" + "\xF6" + "\x03\x80\xE0\x18\xF6\x7F\xB8\xEC\x1B\x06\xE3\x9F\xF3\xFC" +
	/* e */ "\x00\x12" + "\x00\x00\x00\x65" + "\x09\x08" + "\x0B\x01" + "\xF9" + "\x3E\x3F\xB0\xFF\xFF\xFE\x01\xFE\x7E" +
	/* f */ "\x00\x14" + "\x00\x00\x00\x66" + "\x08\x0B" + "\x0B\x02" + "\xF6" + "\x1F\x3F\x30\x7E\x7E\x30\x30\x30\x30\xFE\xFE" +
	/* g */ "\x00\x17" + "\x00\x00\x00\x67" + "\x09\x0C" + "\x0B\x01" + "\xF9" + "\x3F\xBF\xF9\xD8\x6C\x37\x39\xFC\x76\x03\x01\x8F\xC7\xC0" +
	/* h */ "\x00\x16" + "\x00\x00\x00\x68" + "\x09\x0B" + "\x0B\x01" + "\xF6" + "\xE0\x70\x18\x0D\xC7\xF3\x99\x8C\xC6\x63\x7B\xFD\xE0" +
	/* i */ "\x00\x14" + "\x00\x00\x00\x69" + "\x08\x0B" + "\x0B\x02" + "\xF6" + "\x18\x18\x00\x78\x78\x18\x18\x18\x18\xFF\xFF" +
	/* j */ "\x00\x15" + "\x00\x00\x00\x6A" + "\x06\x0F" + "\x0B\x02" + "\xF6" + "\x18\x60\x3F\xFC\x30\xC3\x0C\x30\xC3\x0F\xFF\x80" +
	/* k */ "\x00\x16" + "\x00\x00\x00\x6B" + "\x09\x0B" + "\x0B\x01" + "\xF6" + "\xE0\x70\x18\x0D\xE6\xF3\xE1\xE0\xF8\x6E\x73\xF9\xE0" +
	/* l */ "\x00\x14" + "\x00\x00\x00\x6C" + "\x08\x0B" + "\x0B\x02" + "\xF6" + "\x78\x78\x18\x18\x18\x18\x18\x18\x18\xFF\xFF" +
	/* m */ "\x00\x14" + "\x00\x00\x00\x6D" + "\x0B\x08" + "\x0B\x00" + "\xF9" + "\xFD\x9F\xF9\x9B\x33\x66\x6C\xCD\xBD\xFF\xBF" +
	/* n */ "\x00\x12" + "\x00\x00\x00\x6E" + "\x09\x08" + "\x0B\x01" + "\xF9" + "\xEE\x7F\x98\xCC\x66\x33\x1B\xDF\xEF" +
	/* o */ "\x00\x12" + "\x00\x00\x00\x6F" + "\x09\x08" + "\x0B\x01" + "\xF9" + "\x3E\x3F\xB8\xF8\x3C\x1F\x1D\xFC\x7C" +
	/* p */ "\x00\x1A" + "\x00\x00\x00\x70" + "\x0B\x0C" + "\x0B\x00" + "\xF9" + "\xEF\x1F\xF9\xC3\xB0\x36\x06\xE1\xDF\xF3\x78\x60\x0C\x03\xE0\x7C\x00" +
	/* q */ "\x00\x1A" + "\x00\x00\x00\x71" + "\x0B\x0C" + "\x0B\x00" + "\xF9" + "\x1E\xEF\xFF\x87\x60\x6C\x0D\xC3\x9F\xF0\xF6\x00\xC0\x18\x0F\x81\xF0" +
	/* r */ "\x00\x12" + "\x00\x00\x00\x72" + "\x09\x08" + "\x0B\x01" + "\xF9" + "\x77\xBF\xCF\x06\x03\x01\x83\xF9\xFC" +
	/* s */ "\x00\x11" + "\x00\x00\x00\x73" + "\x08\x08" + "\x0B\x02" + "\xF9" + "\x3F\xFF\xC3\xFC\x3F\xC3\xFF\xFC" +
	/* t */ "\x00\x14" + "\x00\x00\x00\x74" + "\x08\x0B" + "\x0B\x01" + "\xF6" + "\x60\x60\x60\xFE\xFE\x60\x60\x60\x61\x7F\x3E" +
	/* u */ "\x00\x12" + "\x00\x00\x00\x75" + "\x09\x08" + "\x0B\x01" + "\xF9" + "\xE7\x73\x98\xCC\x66\x33\x19\xFE\x7F" +
	/* v */ "\x00\x14" + "\x00\x00\x00\x76" + "\x0B\x08" + "\x0B\x00" + "\xF9" + "\xFB\xFF\x7C\xC6\x18\xC1\xB0\x36\x03\x80\x70" +
	/* w */ "\x00\x14" + "\x00\x00\x00\x77" + "\x0B\x08" + "\x0B\x00" + "\xF9" + "\xF1\xFE\x3D\xBB\x37\x63\xF8\x77\x0E\xE1\x8C" +
	/* x */ "\x00\x12" + "\x00\x00\x00\x78" + "\x09\x08" + "\x0B\x01" + "\xF9" + "\xF7\xFB\xCD\x83\x83\xC3\xBB\xDF\xEF" +
	/* y */ "\x00\x18" + "\x00\x00\x00\x79" + "\x0A\x0C" + "\x0B\x00" + "\xF9" + "\xF3\xFC\xF6\x18\xCC\x33\x07\x81\xE0\x30\x0C\x06\x0F\xC3\xF0" +
	/* z */ "\x00\x10" + "\x00\x00\x00\x7A" + "\x07\x08" + "\x0B\x02" + "\xF9" + "\xFF\xFF\x30\xC3\x0C\x7F\xFF" +
	/* { */ "\x00\x10" + "\x00\x00\x00\x7B" + "\x04\x0E" + "\x0B\x03" + "\xF6" + "\x37\x66\x66\xCC\x66\x66\x73" +
	/* | */ "\x00\x0D" + "\x00\x00\x00\x7C" + "\x02\x0E" + "\x0B\x05" + "\xF6" + "\xFF\xFF\xFF\xF0" +
	/* } */ "\x00\x10" + "\x00\x00\x00\x7D" + "\x04\x0E" + "\x0B\x04" + "\xF6" + "\xCE\x66\x66\x33\x66\x66\xEC" +
	/* ~ */ "\x00\x0E" + "\x00\x00\x00\x7E" + "\x09\x04" + "\x0B\x01" + "\xFA" + "\x70\x7C\xF3\xC0\xC0" +
	""

func (f *sBold9pt7b) GetGlyph(r rune) tinyfont.Glyph {
	idx := -1

	switch r {
	case 0x0020:
		idx = 0
	case 0x0021:
		idx = 11
	case 0x0022:
		idx = 27
	case 0x0023:
		idx = 43
	case 0x0024:
		idx = 66
	case 0x0025:
		idx = 90
	case 0x0026:
		idx = 111
	case 0x0027:
		idx = 132
	case 0x0028:
		idx = 145
	case 0x0029:
		idx = 163
	case 0x002A:
		idx = 181
	case 0x002B:
		idx = 199
	case 0x002C:
		idx = 219
	case 0x002D:
		idx = 232
	case 0x002E:
		idx = 246
	case 0x002F:
		idx = 258
	case 0x0030:
		idx = 283
	case 0x0031:
		idx = 305
	case 0x0032:
		idx = 326
	case 0x0033:
		idx = 349
	case 0x0034:
		idx = 372
	case 0x0035:
		idx = 392
	case 0x0036:
		idx = 416
	case 0x0037:
		idx = 439
	case 0x0038:
		idx = 461
	case 0x0039:
		idx = 483
	case 0x003A:
		idx = 506
	case 0x003B:
		idx = 519
	case 0x003C:
		idx = 535
	case 0x003D:
		idx = 555
	case 0x003E:
		idx = 573
	case 0x003F:
		idx = 593
	case 0x0040:
		idx = 615
	case 0x0041:
		idx = 643
	case 0x0042:
		idx = 670
	case 0x0043:
		idx = 695
	case 0x0044:
		idx = 719
	case 0x0045:
		idx = 744
	case 0x0046:
		idx = 768
	case 0x0047:
		idx = 792
	case 0x0048:
		idx = 817
	case 0x0049:
		idx = 841
	case 0x004A:
		idx = 861
	case 0x004B:
		idx = 886
	case 0x004C:
		idx = 911
	case 0x004D:
		idx = 935
	case 0x004E:
		idx = 962
	case 0x004F:
		idx = 989
	case 0x0050:
		idx = 1016
	case 0x0051:
		idx = 1040
	case 0x0052:
		idx = 1071
	case 0x0053:
		idx = 1095
	case 0x0054:
		idx = 1119
	case 0x0055:
		idx = 1143
	case 0x0056:
		idx = 1170
	case 0x0057:
		idx = 1197
	case 0x0058:
		idx = 1222
	case 0x0059:
		idx = 1247
	case 0x005A:
		idx = 1272
	case 0x005B:
		idx = 1294
	case 0x005C:
		idx = 1312
	case 0x005D:
		idx = 1337
	case 0x005E:
		idx = 1355
	case 0x005F:
		idx = 1372
	case 0x0060:
		idx = 1386
	case 0x0061:
		idx = 1399
	case 0x0062:
		idx = 1419
	case 0x0063:
		idx = 1444
	case 0x0064:
		idx = 1464
	case 0x0065:
		idx = 1489
	case 0x0066:
		idx = 1509
	case 0x0067:
		idx = 1531
	case 0x0068:
		idx = 1556
	case 0x0069:
		idx = 1580
	case 0x006A:
		idx = 1602
	case 0x006B:
		idx = 1625
	case 0x006C:
		idx = 1649
	case 0x006D:
		idx = 1671
	case 0x006E:
		idx = 1693
	case 0x006F:
		idx = 1713
	case 0x0070:
		idx = 1733
	case 0x0071:
		idx = 1761
	case 0x0072:
		idx = 1789
	case 0x0073:
		idx = 1809
	case 0x0074:
		idx = 1828
	case 0x0075:
		idx = 1850
	case 0x0076:
		idx = 1870
	case 0x0077:
		idx = 1892
	case 0x0078:
		idx = 1914
	case 0x0079:
		idx = 1934
	case 0x007A:
		idx = 1960
	case 0x007B:
		idx = 1978
	case 0x007C:
		idx = 1996
	case 0x007D:
		idx = 2011
	case 0x007E:
		idx = 2029
	}

	if idx == -1 {
		return tinyfont.Glyph{
			Rune:     r,
			Width:    0,
			Height:   0,
			XAdvance: uint8(cBold9pt7b[6]),
			XOffset:  int8(cBold9pt7b[7]),
			YOffset:  int8(cBold9pt7b[8]),
			Bitmaps:  []uint8{},
		}
	}

	length := int((uint16(cBold9pt7b[idx+0]) << 8) + uint16(cBold9pt7b[idx+1]))
	idx += 2
	ret := tinyfont.Glyph{
		Rune:     rune((uint32(cBold9pt7b[idx+0]) << 24) + (uint32(cBold9pt7b[idx+1]) << 16) + (uint32(cBold9pt7b[idx+2]) << 8) + uint32(cBold9pt7b[idx+3])),
		Width:    uint8(cBold9pt7b[idx+4]),
		Height:   uint8(cBold9pt7b[idx+5]),
		XAdvance: uint8(cBold9pt7b[idx+6]),
		XOffset:  int8(cBold9pt7b[idx+7]),
		YOffset:  int8(cBold9pt7b[idx+8]),
		Bitmaps:  []uint8(cBold9pt7b[idx+9 : idx+length]),
	}

	return ret
}

func (f *sBold9pt7b) GetYAdvance() uint8 {
	return 0x12
}
