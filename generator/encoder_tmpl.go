package generator

import (
	"bytes"
	"compress/gzip"
	"io"
)

// encoder_tmpl returns raw, uncompressed file data.
func encoder_tmpl() []byte {
	gz, err := gzip.NewReader(bytes.NewBuffer([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xac, 0x96,
		0xcf, 0x6f, 0x9b, 0x30, 0x14, 0xc7, 0xcf, 0xf8, 0xaf, 0x78, 0x43, 0xdd,
		0x0a, 0x55, 0x45, 0xa6, 0xad, 0xea, 0xa1, 0x53, 0x77, 0xa8, 0xb4, 0x48,
		0xdb, 0x21, 0x93, 0x56, 0x55, 0x3b, 0x54, 0x3d, 0x10, 0x62, 0x12, 0xb7,
		0xc4, 0xce, 0xc0, 0xc0, 0x22, 0x8f, 0xff, 0x7d, 0xfe, 0x01, 0x49, 0xf8,
		0xdd, 0x4d, 0x5c, 0x82, 0xb1, 0x9f, 0xdf, 0xe7, 0xfb, 0xbe, 0xc0, 0x73,
		0x76, 0x7e, 0xf0, 0xe2, 0xaf, 0x31, 0x08, 0xe1, 0x2d, 0xfc, 0x2d, 0xd6,
		0x3f, 0x45, 0x81, 0x10, 0xd9, 0xee, 0x58, 0xcc, 0xc1, 0x41, 0x96, 0x4d,
		0x98, 0x2d, 0x7f, 0xd7, 0x84, 0x6f, 0xd2, 0xa5, 0x17, 0xb0, 0xed, 0x6c,
		0x89, 0xe9, 0xf2, 0x99, 0x6d, 0x68, 0xc2, 0xe8, 0x6c, 0x8b, 0xd7, 0xfe,
		0xb3, 0x1a, 0xe4, 0x31, 0xe1, 0x38, 0xb6, 0x91, 0x8b, 0x90, 0x10, 0xb1,
		0x4f, 0x65, 0x4e, 0xbe, 0xdf, 0xe1, 0x04, 0x3c, 0x99, 0x4e, 0x8d, 0x1a,
		0x88, 0x6f, 0xf7, 0xdf, 0x17, 0x5f, 0x68, 0xc0, 0x56, 0x38, 0x86, 0x84,
		0xc7, 0x69, 0xc0, 0x41, 0x20, 0x2b, 0x87, 0x0b, 0x93, 0xc8, 0xfb, 0xa9,
		0x2f, 0x48, 0x4a, 0x09, 0x53, 0x1a, 0xc0, 0x02, 0xe7, 0xbd, 0xfb, 0x9d,
		0x1c, 0x08, 0x2b, 0x37, 0xb8, 0x70, 0xd1, 0xcf, 0x91, 0x80, 0x18, 0xf3,
		0x34, 0xa6, 0xf0, 0xae, 0x37, 0x48, 0xe4, 0x37, 0x50, 0x4a, 0x90, 0x4c,
		0x93, 0xd4, 0xc9, 0xdd, 0x62, 0x58, 0xc9, 0x0f, 0x3f, 0x3f, 0x8a, 0xa9,
		0x97, 0x30, 0x8d, 0xa2, 0x23, 0xdf, 0xc1, 0x03, 0x09, 0x5d, 0x30, 0x03,
		0x27, 0x6b, 0x06, 0xb9, 0x80, 0xe3, 0x98, 0x69, 0x22, 0x09, 0xd5, 0x18,
		0x6e, 0x6e, 0x01, 0x7b, 0x07, 0xe1, 0x4e, 0xe6, 0x7e, 0xd2, 0xd3, 0x6f,
		0x6e, 0x81, 0x92, 0x48, 0xc5, 0x55, 0xd2, 0xe4, 0x2c, 0xb2, 0x8a, 0xfa,
		0xbe, 0xdc, 0x9b, 0x47, 0x69, 0xb2, 0x71, 0x46, 0x37, 0x95, 0xb7, 0x72,
		0xf5, 0xb5, 0x15, 0x9c, 0x48, 0x1a, 0x2e, 0x22, 0x83, 0xdb, 0x36, 0x56,
		0x2a, 0xd3, 0xbe, 0x2f, 0xd2, 0x28, 0x72, 0x5c, 0x25, 0xa1, 0x29, 0x5c,
		0x2f, 0xdf, 0xed, 0x39, 0x76, 0xce, 0xc5, 0xf9, 0x98, 0x7e, 0x64, 0x55,
		0x6f, 0xf3, 0x19, 0xa1, 0x2b, 0xfc, 0xfb, 0x12, 0xce, 0x42, 0x82, 0xa3,
		0x95, 0x4a, 0xa6, 0x07, 0xe6, 0x05, 0xb7, 0x64, 0x98, 0xa4, 0x98, 0x18,
		0x7d, 0xdf, 0x0f, 0xbd, 0xec, 0x80, 0xd6, 0xb1, 0x96, 0x65, 0x32, 0x62,
		0xba, 0x52, 0xdf, 0xa2, 0x65, 0xcd, 0x66, 0xa0, 0x13, 0xc0, 0x0b, 0xde,
		0x83, 0x4f, 0x57, 0x10, 0xb0, 0x88, 0x51, 0x0f, 0x75, 0x52, 0xee, 0x79,
		0x4c, 0xe8, 0xda, 0x11, 0x42, 0x06, 0x53, 0x69, 0x1b, 0x78, 0xf0, 0x07,
		0x76, 0x72, 0x8e, 0x87, 0x60, 0xbf, 0xfd, 0x65, 0x4b, 0x1b, 0xdb, 0xfc,
		0x1a, 0xbe, 0xe8, 0x4e, 0x6c, 0xe4, 0xdf, 0x74, 0xc9, 0x6f, 0x6c, 0x3f,
		0x95, 0x9c, 0xf9, 0x51, 0x8a, 0x95, 0x54, 0x1d, 0x98, 0xa9, 0x8c, 0x99,
		0x27, 0x84, 0x36, 0xcf, 0xc8, 0x33, 0x35, 0x1a, 0x07, 0x49, 0x22, 0x95,
		0x6e, 0x09, 0x27, 0x19, 0xd6, 0x4d, 0xc3, 0x98, 0x7b, 0x58, 0x34, 0x73,
		0x60, 0x27, 0xba, 0x46, 0xbb, 0x5c, 0x1c, 0x70, 0xa1, 0xe3, 0xad, 0x6e,
		0xbb, 0x6d, 0x2a, 0x3e, 0x3a, 0xde, 0x01, 0x94, 0xee, 0x0d, 0xd1, 0xbe,
		0x52, 0x3e, 0x25, 0xea, 0xfa, 0x6a, 0x04, 0x76, 0x7d, 0x35, 0x19, 0x2e,
		0x1d, 0x29, 0xed, 0x81, 0x4c, 0x58, 0x5b, 0x3a, 0x5a, 0xdc, 0x03, 0x99,
		0xb4, 0xba, 0x30, 0x62, 0x3e, 0xff, 0xf8, 0x61, 0x88, 0x38, 0x37, 0x21,
		0xd3, 0x22, 0x87, 0x8b, 0x9c, 0x9b, 0x90, 0xc9, 0x90, 0x4b, 0xc6, 0xa2,
		0x21, 0xde, 0x9d, 0x5c, 0xff, 0x6f, 0x58, 0x6d, 0x58, 0xc3, 0x5e, 0x54,
		0xcc, 0x23, 0x52, 0x9f, 0x8e, 0x49, 0xba, 0xac, 0x3e, 0xde, 0xc6, 0xe1,
		0x28, 0x15, 0xb9, 0xc3, 0x67, 0x4e, 0x5b, 0xd1, 0x88, 0x88, 0xc7, 0xa7,
		0xb6, 0x8a, 0x46, 0xcb, 0x7a, 0xec, 0xec, 0xb8, 0x6d, 0x8e, 0xbe, 0x84,
		0xf2, 0x7c, 0x29, 0x5b, 0xbd, 0x6e, 0x56, 0xa6, 0xfb, 0x67, 0xd5, 0x26,
		0xc5, 0x56, 0xab, 0xf0, 0x19, 0xde, 0x1f, 0x3c, 0xfc, 0xc7, 0x5e, 0xdf,
		0xe1, 0x7a, 0x69, 0xbb, 0xd5, 0x7a, 0x86, 0x53, 0x18, 0xda, 0xf3, 0x8c,
		0xcb, 0x82, 0x7b, 0xc5, 0x3f, 0xbd, 0xd2, 0xb6, 0xda, 0xe3, 0x29, 0xd0,
		0xc9, 0xc1, 0xd5, 0x9b, 0xba, 0x18, 0x3d, 0x78, 0xeb, 0x7f, 0x1c, 0xaa,
		0x94, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x5e, 0x00, 0x61, 0xaf,
		0x0a, 0x00, 0x00,
	}))

	if err != nil {
		panic("Decompression failed: " + err.Error())
	}

	var b bytes.Buffer
	io.Copy(&b, gz)
	gz.Close()

	return b.Bytes()
}
