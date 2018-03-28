// Code generated by go-bindata.
// sources:
// templates/spec-md.gotpl
// templates/toc-md.gotpl
// DO NOT EDIT!

package static

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesSpecMdGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x52\xcb\x4e\xeb\x30\x10\xdd\xe7\x2b\x46\xcd\x5d\x5c\x16\xf5\x07\x54\x02\xa9\x52\xbb\x03\x16\x50\xb1\x8e\xeb\x0c\xd4\x28\x8d\x8d\x67\x2a\xb5\x8a\xfc\xef\xc8\x8f\x38\x49\xc5\xce\x3e\x3e\x8f\x99\xf1\xd4\x35\x0c\x03\x88\x77\x8b\x4a\xbc\x98\x16\x3b\xb1\xef\x59\xf3\xed\x55\x9e\x11\xbc\xaf\x86\x01\xfe\x75\x92\x91\xf8\x03\x1d\x69\xd3\xc3\xe6\x31\xd3\x9f\x23\xbc\x65\x76\xfa\x78\x61\x1c\xdf\x93\x88\x9d\x3e\x93\x95\x0a\x17\xde\x3b\x24\xe5\xb4\xe5\xcc\x8b\xee\x78\x95\x67\xdb\x61\xf0\x1d\x8f\x51\x72\x9f\x9b\x7c\xf5\xe7\xa4\xf0\xbe\xaa\xeb\x1a\xf6\xe9\x5a\x55\x4d\xd3\x7c\x93\xe9\x17\xae\xde\x07\x38\x40\xd8\xb7\x25\xd3\x58\x0a\x79\xc6\xa2\x93\xa1\x18\xca\x91\xe2\x0d\xbb\x04\x9c\xb4\x0d\x20\xf2\x2c\xd7\x58\xca\x91\x85\x36\xd9\x25\xda\x3c\x44\x8e\x93\xa1\x69\x66\xfb\xab\x35\x84\xed\x76\x7a\xfa\xab\xcb\x75\x8c\x9b\xe9\x83\x65\xc8\x9d\x74\x31\xc2\xc9\xfe\x0b\x17\xc4\x75\x2a\xb0\x86\x26\xfc\x6a\xfe\x44\xf8\x9f\x1a\xc0\x1f\x10\x87\x9b\x45\x58\xe1\x95\xd1\xf5\xb2\x5b\x81\xf7\x81\x18\x51\xef\x37\x71\x15\x2e\xc7\x7c\x0d\xfd\x74\x84\x4b\x4e\x69\xf2\xa1\xc9\x73\x11\x3b\xb4\x0e\x95\x64\x8c\xbd\x3f\xc1\xe1\xa4\x09\x4a\x51\xa0\x09\xda\xc2\x98\x0d\x69\xb9\x25\x77\xab\x31\x0c\xa0\x4e\xd2\x49\xc5\xe8\x34\xb1\x56\x04\x62\x1c\x4e\xd1\x97\xe3\x6f\x00\x00\x00\xff\xff\x1d\x03\xd2\x10\xc6\x02\x00\x00")

func templatesSpecMdGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSpecMdGotpl,
		"templates/spec-md.gotpl",
	)
}

func templatesSpecMdGotpl() (*asset, error) {
	bytes, err := templatesSpecMdGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/spec-md.gotpl", size: 710, mode: os.FileMode(420), modTime: time.Unix(1522194570, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTocMdGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xcd\x31\x6b\xc3\x30\x10\x05\xe0\x5d\xbf\xe2\xaa\x2e\xed\x20\x51\x4a\xbb\x94\x62\x30\xf5\x62\x4a\xc0\x10\xc8\x1a\x14\xf9\x9c\x1c\x91\xee\x8c\x2c\x13\x82\xd1\x7f\x0f\x36\x59\xde\xf0\xe0\x7b\xef\xf7\xc5\x18\x88\x2e\x5d\x7b\xb9\x71\x20\xce\xa6\xa7\xc9\x9d\x02\xc2\xae\xf9\xf8\xfc\xda\xf2\x1b\x8c\xa9\x94\x7a\x85\x65\x81\x4c\x39\x20\xbc\xd9\x3f\xe1\x81\xce\x73\x72\x99\x84\xed\x3f\xde\x41\x47\x61\x09\x94\x2f\xa8\x41\x8f\x49\xfa\xd9\xe7\x23\xbb\x88\xfa\x1d\x4a\x81\xba\x6b\xa1\x11\x3f\x47\xe4\xbc\x21\xa5\x2a\x38\x60\x9a\x48\xf8\x67\x5d\xb6\x75\xd7\xb6\x3c\x88\x7d\x96\x50\x8a\x52\xeb\xa3\x78\xb0\xfb\x11\x3d\x0d\xe4\x37\x39\x81\x29\x45\x3d\x02\x00\x00\xff\xff\x2f\xc9\x8f\x8f\xba\x00\x00\x00")

func templatesTocMdGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesTocMdGotpl,
		"templates/toc-md.gotpl",
	)
}

func templatesTocMdGotpl() (*asset, error) {
	bytes, err := templatesTocMdGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/toc-md.gotpl", size: 186, mode: os.FileMode(420), modTime: time.Unix(1521249819, 0)}
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

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
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
	"templates/spec-md.gotpl": templatesSpecMdGotpl,
	"templates/toc-md.gotpl": templatesTocMdGotpl,
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
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"spec-md.gotpl": &bintree{templatesSpecMdGotpl, map[string]*bintree{}},
		"toc-md.gotpl": &bintree{templatesTocMdGotpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

