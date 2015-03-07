package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
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

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _file_error_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb2\x51\x74\xf1\x77\x0e\x89\x0c\x70\x55\xc8\x28\xc9\xcd\xb1\xe3\xb2\x81\x50\x0a\x0a\x36\x19\xa9\x89\x29\x20\x06\x90\x99\x9b\x5a\x92\xa8\x90\x9c\x91\x58\x54\x9c\x5a\x62\xab\x14\x1a\xe2\xa6\x6b\xa1\x04\x56\xa3\x0f\x53\x64\x93\x94\x9f\x52\x09\x55\x5d\x60\x17\x9c\x5f\x54\x54\x69\xa3\x5f\x00\x51\x03\x91\x02\xaa\x05\x9b\x0c\x08\x00\x00\xff\xff\xe7\x62\x65\x61\x71\x00\x00\x00")

func file_error_html_bytes() ([]byte, error) {
	return bindata_read(
		_file_error_html,
		"file/error.html",
	)
}

func file_error_html() (*asset, error) {
	bytes, err := file_error_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "file/error.html", size: 113, mode: os.FileMode(420), modTime: time.Unix(1416479991, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _file_index_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb2\x51\x74\xf1\x77\x0e\x89\x0c\x70\x55\xc8\x28\xc9\xcd\xb1\xe3\xb2\x81\x50\x0a\x0a\x36\x19\xa9\x89\x29\x20\x06\x90\x99\x9b\x5a\x92\xa8\x90\x9c\x91\x58\x54\x9c\x5a\x62\xab\x14\x1a\xe2\xa6\x6b\xa1\x84\x2c\x95\x51\x52\x52\xa0\x9b\x5a\x58\x9a\x59\x66\xab\x54\x94\x9a\x56\x94\x5a\x9c\xa1\xa4\x90\x9c\x9f\x57\x92\x9a\x07\x54\x6f\x60\x5d\x5a\x94\x63\x5b\x5d\xad\xa0\x17\x1a\xe4\xa3\x50\x5b\x0b\xd6\x6a\xa3\x0f\x33\xdf\x26\x29\x3f\xa5\x12\x6a\x5a\x01\x84\x06\xb2\x80\x86\x02\x4d\xb2\x55\x42\xd6\xf6\xa2\x7d\xd5\xd3\xee\xa9\x8f\x1b\x97\x3f\x6e\x5e\xf5\xb8\xb9\xe1\x71\xd3\x92\xc7\xcd\x6b\x1e\x37\xad\x7f\xdc\xdc\xf1\xb8\x71\xfa\xe3\xc6\x55\x8f\x1b\x5b\x9e\x2e\xd8\xf2\x74\x02\x90\xbb\xfe\x71\xe3\xe4\xc7\x8d\x0b\x1f\x37\x75\xda\xe8\x27\x42\x4d\xd7\x2f\x80\xd8\x0c\xb1\x10\xe8\x02\xb0\x57\x01\x01\x00\x00\xff\xff\xcc\x3a\x77\xce\x02\x01\x00\x00")

func file_index_html_bytes() ([]byte, error) {
	return bindata_read(
		_file_index_html,
		"file/index.html",
	)
}

func file_index_html() (*asset, error) {
	bytes, err := file_index_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "file/index.html", size: 258, mode: os.FileMode(420), modTime: time.Unix(1416670353, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _file_seelog_xml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb2\x29\x4e\x4d\xcd\xc9\x4f\x57\x28\xa9\x2c\x48\xb5\x55\x2a\xae\xcc\x4b\x56\xb2\xe3\x52\x00\x02\x9b\xfc\xd2\x92\x82\xd2\x92\x62\x85\xb4\xfc\xa2\xdc\xc4\x92\xcc\x14\x5b\xa5\xdc\xc4\xcc\x3c\xa8\x2c\x58\x45\x72\x7e\x5e\x71\x7e\x4e\xaa\x3e\x54\x83\x3e\x54\x07\x94\x0b\xd1\x57\x8c\xa4\x1e\x22\xa2\x00\x37\x0a\x6a\xb4\xad\x92\xaa\x4b\x62\x49\xaa\x82\x6a\x48\x66\x6e\xaa\x42\xb4\xaa\x8f\x6b\x58\xac\x82\xaa\x6f\x71\xba\x6a\x9e\x12\xdc\x6c\xb8\x69\x36\xfa\x10\x17\xdb\x25\x70\x01\x02\x00\x00\xff\xff\x86\xc7\x20\x95\xbc\x00\x00\x00")

func file_seelog_xml_bytes() ([]byte, error) {
	return bindata_read(
		_file_seelog_xml,
		"file/seelog.xml",
	)
}

func file_seelog_xml() (*asset, error) {
	bytes, err := file_seelog_xml_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "file/seelog.xml", size: 188, mode: os.FileMode(420), modTime: time.Unix(1425711610, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
var _bindata = map[string]func() (*asset, error){
	"file/error.html": file_error_html,
	"file/index.html": file_index_html,
	"file/seelog.xml": file_seelog_xml,
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
	Func     func() (*asset, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"file": &_bintree_t{nil, map[string]*_bintree_t{
		"error.html": &_bintree_t{file_error_html, map[string]*_bintree_t{}},
		"index.html": &_bintree_t{file_index_html, map[string]*_bintree_t{}},
		"seelog.xml": &_bintree_t{file_seelog_xml, map[string]*_bintree_t{}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	if err != nil { // File
		return RestoreAsset(dir, name)
	} else { // Dir
		for _, child := range children {
			err = RestoreAssets(dir, path.Join(name, child))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
