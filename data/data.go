package data

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func raw_data_yml() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xcc, 0x56,
		0xcd, 0x6e, 0xe3, 0x36, 0x10, 0xbe, 0xfb, 0x29, 0x06, 0x39, 0x27, 0xda,
		0xbb, 0x6f, 0x69, 0xe1, 0x6c, 0x0c, 0xb8, 0x89, 0x91, 0xa4, 0x5d, 0x14,
		0xe8, 0x85, 0x91, 0xc6, 0x12, 0x1b, 0x8a, 0x54, 0x49, 0xca, 0x8a, 0x7b,
		0x2a, 0xb0, 0x4f, 0xd1, 0x47, 0xe8, 0xb9, 0x8f, 0xb0, 0xd7, 0x7d, 0x8a,
		0x3e, 0x49, 0x87, 0xa4, 0x28, 0xcb, 0x92, 0xb5, 0x49, 0x80, 0xa4, 0x0d,
		0x60, 0x84, 0xe1, 0x70, 0xfe, 0xe7, 0x9b, 0x19, 0x55, 0x4a, 0x08, 0x2e,
		0xf3, 0x95, 0x4a, 0x99, 0xe5, 0x4a, 0xce, 0x67, 0x00, 0x16, 0x1f, 0xad,
		0x3b, 0x01, 0x50, 0x86, 0x13, 0xa0, 0xd2, 0xb8, 0xe1, 0x8f, 0x73, 0x38,
		0xf9, 0x59, 0xd5, 0x1a, 0x5a, 0x21, 0xa8, 0x04, 0x4b, 0x11, 0xb8, 0x99,
		0x9f, 0xb4, 0x6c, 0x05, 0xbd, 0x1a, 0xe2, 0xba, 0xf4, 0x67, 0xa0, 0xa2,
		0x19, 0x29, 0x31, 0x15, 0x93, 0xdc, 0x14, 0x67, 0xcf, 0x53, 0x16, 0xb9,
		0xf7, 0x4a, 0xad, 0xe6, 0x79, 0x8e, 0xda, 0x0c, 0xbd, 0x3c, 0xf3, 0xca,
		0x06, 0x56, 0xcf, 0xc0, 0x78, 0x6a, 0x5a, 0x30, 0x99, 0xe3, 0x8a, 0xfe,
		0xd4, 0x2c, 0xc7, 0xe3, 0x81, 0x06, 0xa9, 0x69, 0xf5, 0x28, 0x73, 0x41,
		0xae, 0x8c, 0x2c, 0xa0, 0xf3, 0xf1, 0xeb, 0x67, 0x25, 0x0e, 0x29, 0x74,
		0xe7, 0xd2, 0x6a, 0x35, 0x95, 0x55, 0x26, 0xc4, 0x1c, 0x3e, 0xa1, 0x48,
		0x55, 0x89, 0x60, 0x15, 0xd8, 0x02, 0xe1, 0x27, 0x65, 0x5d, 0x36, 0x96,
		0x72, 0xa3, 0x74, 0xe9, 0x6b, 0x02, 0x6b, 0xad, 0x7e, 0xc5, 0xd4, 0x82,
		0x50, 0xea, 0xa1, 0xae, 0x88, 0x51, 0x89, 0x04, 0x6e, 0xb0, 0x12, 0x3b,
		0x68, 0xb8, 0x2d, 0x60, 0xe7, 0xd2, 0xc8, 0xb2, 0x4c, 0xa3, 0x31, 0x4e,
		0xcd, 0x86, 0xcb, 0x2c, 0x10, 0x63, 0x6e, 0x45, 0x5b, 0x5e, 0x60, 0xf4,
		0xa2, 0xc8, 0x8c, 0xf6, 0x24, 0x01, 0x9c, 0xcc, 0x24, 0x83, 0x70, 0xbc,
		0x57, 0x31, 0xe9, 0xef, 0xc2, 0xbb, 0xe9, 0x82, 0x6c, 0x95, 0xc5, 0xee,
		0xd2, 0x90, 0x24, 0x8e, 0x6a, 0x43, 0x2c, 0xaa, 0xbb, 0x64, 0x4a, 0x66,
		0x7b, 0x81, 0xec, 0xcb, 0xdf, 0xee, 0xaa, 0x31, 0xe7, 0xc6, 0xea, 0xe7,
		0x35, 0xc0, 0x85, 0xd2, 0xd0, 0x17, 0xf0, 0x4e, 0xb6, 0xb9, 0x78, 0x1a,
		0xf1, 0x4f, 0x4a, 0x4f, 0x87, 0x1a, 0xc4, 0x50, 0xf7, 0x09, 0x83, 0xc7,
		0xa0, 0xf3, 0x28, 0xd1, 0x1b, 0x1a, 0x89, 0xf6, 0xf3, 0x14, 0xb9, 0x75,
		0x89, 0x33, 0x14, 0x93, 0x98, 0x3d, 0x9c, 0x04, 0xa1, 0x4e, 0x28, 0x08,
		0x00, 0xce, 0x88, 0xda, 0x6c, 0x78, 0xca, 0x5d, 0xe5, 0xf6, 0x6d, 0x5c,
		0x15, 0x4a, 0x22, 0xf1, 0xaf, 0xfd, 0x19, 0xa9, 0x58, 0x32, 0x4e, 0x30,
		0x3b, 0x59, 0xf8, 0xf3, 0x79, 0x93, 0xe2, 0x05, 0xc6, 0xa2, 0xdc, 0x71,
		0xa3, 0xf1, 0x75, 0x6f, 0x7c, 0x3a, 0xed, 0xd1, 0xca, 0xbe, 0xb3, 0xc7,
		0xf6, 0xfb, 0xa9, 0xed, 0x9e, 0x8f, 0x53, 0x27, 0xd5, 0xf5, 0x08, 0x6a,
		0x54, 0x9b, 0x4d, 0x2d, 0x53, 0xe2, 0x61, 0x9a, 0xab, 0x19, 0xbb, 0x57,
		0xb5, 0xfd, 0xe6, 0x44, 0xa1, 0x64, 0xfd, 0xf3, 0xc7, 0x9f, 0x1a, 0xa1,
		0x36, 0xae, 0xbd, 0x5e, 0xd2, 0xb7, 0x77, 0x0a, 0x04, 0x32, 0x2d, 0xa1,
		0x54, 0x1a, 0x4f, 0x61, 0xcb, 0x0d, 0xb7, 0x50, 0x58, 0x5b, 0xcd, 0x3f,
		0x7c, 0x68, 0x9a, 0x26, 0xd9, 0x7a, 0x3d, 0x2e, 0xb4, 0x2a, 0x88, 0x27,
		0x4a, 0x0f, 0xa1, 0x74, 0x30, 0x3f, 0xfe, 0x07, 0x5f, 0xa6, 0x6b, 0xd9,
		0x14, 0xcc, 0x76, 0x17, 0x9f, 0xc7, 0x51, 0xa6, 0x69, 0x07, 0xe9, 0x94,
		0x75, 0xd7, 0xdf, 0x6a, 0xec, 0xfd, 0xff, 0xe5, 0xaf, 0x59, 0x81, 0xa2,
		0x9a, 0x4a, 0x7e, 0x89, 0xb2, 0x26, 0x74, 0xdd, 0xd1, 0x13, 0xac, 0xaf,
		0x57, 0xab, 0xe3, 0xd3, 0xce, 0xef, 0xb9, 0xe4, 0x17, 0xe9, 0xd9, 0xae,
		0x2f, 0x2e, 0x96, 0xdf, 0x2f, 0xcf, 0x57, 0x40, 0xb9, 0x98, 0x40, 0x56,
		0x64, 0xbd, 0x59, 0x7c, 0x5c, 0xde, 0xde, 0x2d, 0x6e, 0x3c, 0xeb, 0xa8,
		0xb5, 0x93, 0x88, 0x71, 0xd1, 0x2e, 0x38, 0xda, 0x9c, 0x3f, 0xd0, 0x3c,
		0x65, 0xb0, 0xb8, 0x5d, 0x9f, 0x5f, 0x7d, 0xfd, 0x7c, 0xbd, 0x82, 0x8a,
		0x69, 0x06, 0x29, 0x2b, 0xef, 0x39, 0x73, 0xc6, 0x80, 0x67, 0x5c, 0x95,
		0x6c, 0x38, 0xfe, 0xdb, 0x28, 0x62, 0x01, 0xdf, 0x4d, 0x34, 0x5e, 0x6c,
		0x71, 0xf5, 0x71, 0xb5, 0xbc, 0xbd, 0x74, 0xbe, 0x84, 0x7d, 0xee, 0x11,
		0x15, 0xb9, 0xbe, 0xbd, 0x2b, 0x5c, 0xe9, 0xba, 0x8b, 0x8b, 0x72, 0x5c,
		0xfc, 0x5d, 0x9d, 0xb1, 0x19, 0x6a, 0xad, 0x82, 0xfc, 0xd1, 0x0e, 0x0b,
		0x9b, 0x6c, 0xcd, 0xb4, 0xc1, 0xc5, 0x23, 0xb9, 0x4d, 0x59, 0xf8, 0xd1,
		0xa0, 0x26, 0x07, 0x09, 0x5d, 0x34, 0x8b, 0x24, 0xe1, 0xdd, 0x02, 0xa3,
		0xa0, 0x52, 0x95, 0x4b, 0xfe, 0x3b, 0x66, 0x40, 0x1b, 0xb4, 0xa4, 0x5a,
		0x24, 0x21, 0x86, 0xcb, 0xc5, 0x6a, 0xed, 0x02, 0x30, 0x88, 0xae, 0x57,
		0x40, 0x55, 0x2e, 0x6e, 0x93, 0x1c, 0x31, 0x70, 0x85, 0x4d, 0xd0, 0xfd,
		0x09, 0x41, 0x22, 0x66, 0xa3, 0x65, 0x4a, 0xc8, 0xdf, 0xf2, 0x0c, 0x1d,
		0x3d, 0x6c, 0xdb, 0x30, 0x23, 0x43, 0x6b, 0xf4, 0x97, 0x4b, 0xe2, 0x54,
		0x64, 0x3c, 0x0b, 0xde, 0x75, 0xbe, 0x51, 0xfa, 0xc8, 0xe9, 0x56, 0xe1,
		0x29, 0x55, 0x13, 0x99, 0x21, 0xa2, 0xde, 0x01, 0xcb, 0x19, 0x97, 0xd1,
		0x27, 0xa9, 0x16, 0x6d, 0x35, 0x5d, 0xd7, 0xce, 0xe1, 0x4a, 0x75, 0xe5,
		0x35, 0xc0, 0xa8, 0xbb, 0x4d, 0x5a, 0x60, 0x56, 0x0b, 0xf2, 0xd0, 0xd5,
		0x33, 0x78, 0xa9, 0x91, 0xf9, 0x26, 0xa6, 0xa7, 0xf4, 0x81, 0xf8, 0x0d,
		0xfa, 0x05, 0x7d, 0x4a, 0xd6, 0xbb, 0x8f, 0x03, 0x46, 0x61, 0x35, 0xd1,
		0x7e, 0xb4, 0x96, 0xa3, 0x44, 0xcd, 0xc4, 0x77, 0x2c, 0x7d, 0x40, 0x99,
		0xcd, 0xe1, 0x56, 0x69, 0xbd, 0x3b, 0x85, 0x06, 0xe9, 0xe7, 0x26, 0x89,
		0x64, 0xf7, 0x02, 0x0f, 0xe1, 0xd8, 0x81, 0x2d, 0x63, 0xbb, 0xd1, 0x77,
		0xc5, 0x38, 0x8a, 0xeb, 0x16, 0x92, 0x3e, 0xb1, 0xf4, 0x4d, 0xe0, 0x93,
		0x52, 0xb0, 0x2d, 0x4e, 0xa0, 0x36, 0xce, 0x29, 0xfa, 0x6d, 0x38, 0xd9,
		0x3e, 0x0c, 0xb2, 0x53, 0x7f, 0xd3, 0xc3, 0x71, 0x48, 0xd4, 0x50, 0xfd,
		0xd4, 0xee, 0x7f, 0x42, 0x33, 0x95, 0xfe, 0x3c, 0xe4, 0xe8, 0x35, 0xa0,
		0x30, 0x9c, 0xd7, 0x93, 0x88, 0xee, 0x66, 0xc0, 0x1b, 0x22, 0x7b, 0xff,
		0xa1, 0xf9, 0xae, 0x10, 0x1e, 0xdd, 0xfa, 0x6f, 0x91, 0x1e, 0xad, 0xbe,
		0x25, 0xe2, 0x7b, 0x09, 0x7f, 0x4b, 0xe4, 0x4f, 0x99, 0x79, 0x85, 0x0e,
		0x78, 0x3d, 0xc8, 0x4c, 0xee, 0x0b, 0xd7, 0x1d, 0xff, 0x06, 0x00, 0x00,
		0xff, 0xff, 0xe6, 0xf9, 0x50, 0x2e, 0x1b, 0x0f, 0x00, 0x00,
	},
		"raw/data.yml",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"raw/data.yml": raw_data_yml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"raw": &_bintree_t{nil, map[string]*_bintree_t{
		"data.yml": &_bintree_t{raw_data_yml, map[string]*_bintree_t{}},
	}},
}}
