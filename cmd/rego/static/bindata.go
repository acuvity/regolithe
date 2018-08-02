// Code generated by go-bindata.
// sources:
// templates/spec-md.gotpl
// templates/toc-md.gotpl
// specset/.gitignore
// specset/Gopkg.toml
// specset/specs/.regolithe-gen-cmd
// specset/specs/@identifiable.abs
// specset/specs/_api.info
// specset/specs/_type.mapping
// specset/specs/monolithe.ini
// specset/specs/object.spec
// specset/specs/root.spec
// specset/specs/type_mapping.ini
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

var _templatesSpecMdGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x52\x4b\x4e\xc3\x30\x10\xdd\xe7\x14\xa3\x86\x05\x2c\xea\x03\x54\x02\x09\xa9\xdd\x01\x0b\xa8\x58\xc7\x75\x06\x6a\x94\xc6\xc6\x33\x95\x5a\x45\xbe\x3b\xf2\x27\x4e\x52\xb1\xb3\x9f\xde\xc7\xf3\x3c\x75\x0d\xc3\x00\xe2\xc3\xa2\x12\xaf\xa6\xc5\x4e\xec\x7a\xd6\x7c\x7d\x93\x27\x04\xef\xab\x61\x80\xbb\x4e\x32\x12\x7f\xa2\x23\x6d\x7a\xd8\x3c\x66\xfa\x4b\x84\x9f\x99\x9d\x3e\x9c\x19\x69\x24\x24\x15\x3b\x7d\x22\x2b\x15\x2e\xcc\xb7\x48\xca\x69\xcb\x99\x17\xed\xf1\x22\x4f\xb6\xc3\x60\x3c\x1e\xa3\xe4\x36\x38\xf9\xea\xaf\x49\xe1\x7d\x55\xd7\x35\xec\xd2\xb5\xaa\x9a\xa6\xf9\x21\xd3\x2f\x5c\xbd\x0f\x70\x80\xb0\x6f\x4b\xa6\xb1\x14\xf2\x8c\x45\x27\xc3\x63\x28\x47\x8a\x77\xec\x12\x70\xd4\x36\x80\xc8\xb3\x5c\x63\x29\x47\x16\xda\x64\x97\x68\xf3\x10\x59\xaa\x99\x4a\xdb\x5d\xac\x21\x6c\xa7\xd6\xfe\x9d\x72\x1d\xe3\x66\xfa\x60\x19\x72\x27\x5d\x8c\x70\xb2\xff\xc6\x05\x71\x9d\x1e\x58\x43\x13\xbe\x35\xff\x22\xdc\xa7\x01\xf0\x17\xc4\xfe\x6a\x11\x56\x78\x61\x74\xbd\xec\x56\xe0\x7d\x20\x46\xd4\xfb\x4d\xdc\x85\xf3\x21\x5f\xc3\x3c\x1d\xe1\x92\x53\x86\x7c\x68\x72\x2f\x62\x8b\xd6\xa1\x92\x8c\x71\xf6\x27\xd8\x1f\x35\x41\x79\x14\x68\x82\xb6\x30\x66\x25\x2d\xb7\xe4\x66\x35\x86\x01\xd4\x51\x3a\xa9\x18\x9d\x26\xd6\x8a\x40\x8c\xe5\x14\x7d\x39\xfe\x05\x00\x00\xff\xff\xd4\xf1\xef\xd5\xc7\x02\x00\x00")

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

	info := bindataFileInfo{name: "templates/spec-md.gotpl", size: 711, mode: os.FileMode(420), modTime: time.Unix(1533099905, 0)}
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

	info := bindataFileInfo{name: "templates/toc-md.gotpl", size: 186, mode: os.FileMode(420), modTime: time.Unix(1528494787, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _specsetGitignore = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xca\x41\x0a\x84\x30\x0c\x05\xd0\xfd\x3f\x4a\x61\x72\xa1\x61\x16\x43\xf2\x0d\xc5\x9a\x48\x2d\x45\x6f\x2f\xe2\xe6\xad\x5e\x91\xfd\xfa\x6a\xda\x0f\x45\xe8\xfe\xfa\xa9\xb1\x24\x34\x8d\xce\xc0\x53\x14\xde\xc9\x51\xc3\xb7\x34\xb6\x03\x93\x61\xd9\x51\xa4\xa5\xae\xd0\x9c\xec\x7f\xa7\x8c\x73\xe0\x0e\x00\x00\xff\xff\x66\xf6\x5a\x96\x53\x00\x00\x00")

func specsetGitignoreBytes() ([]byte, error) {
	return bindataRead(
		_specsetGitignore,
		"specset/.gitignore",
	)
}

func specsetGitignore() (*asset, error) {
	bytes, err := specsetGitignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "specset/.gitignore", size: 83, mode: os.FileMode(420), modTime: time.Unix(1528494787, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _specsetGopkgToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\x2e\x28\x2a\xcd\x4b\x8d\xe5\x52\x50\x48\xcf\xd7\x2d\x49\x2d\x2e\x29\x56\xb0\x55\x28\x29\x2a\x4d\xe5\x52\x50\x28\xcd\x2b\x2d\x4e\x4d\xd1\x2d\x48\x4c\xce\x4e\x4c\x4f\x85\x4b\x00\x02\x00\x00\xff\xff\x34\x36\xb3\x89\x33\x00\x00\x00")

func specsetGopkgTomlBytes() ([]byte, error) {
	return bindataRead(
		_specsetGopkgToml,
		"specset/Gopkg.toml",
	)
}

func specsetGopkgToml() (*asset, error) {
	bytes, err := specsetGopkgTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "specset/Gopkg.toml", size: 51, mode: os.FileMode(420), modTime: time.Unix(1528494787, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _specsetSpecsRegolitheGenCmd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\x41\x4e\xc5\x30\x0c\x44\xf7\x3e\xc5\xf0\xf9\x0b\x40\x4a\x22\x0e\x00\x77\x49\x63\x37\x8d\x48\xe2\x2a\x0d\x88\x45\x0f\x8f\x5a\x5a\xca\xc2\x0b\x6b\xde\xb3\xe7\xf1\xc1\x0d\xa9\xba\xc1\x2f\x13\x95\x0f\x4e\x0d\x66\x86\xb5\x8e\x35\x10\xfb\xee\xdf\xee\x4f\x4d\xa2\x82\x35\xc0\x30\x2c\xd6\x15\xf2\x9d\x3a\x5e\x9f\x49\xc2\xa4\xb8\xdd\x37\xec\x86\xf7\xc3\xda\xe6\xb3\x48\xed\xbe\x27\xad\xb6\x30\x05\x86\xfd\xe7\x91\x64\x89\x52\x31\x6a\x66\x69\xdb\xd1\x65\x96\xb0\xc0\x28\x82\xf2\x1e\x5d\xec\xd5\x29\x6a\xf6\x35\x52\x2b\x30\x6d\x84\x75\xbf\xbb\x7b\xb1\x51\xa9\x7c\x9d\xaa\x93\x2c\xfb\xf3\xbc\x27\x7f\xdc\xe9\x1d\x18\xfd\x04\x00\x00\xff\xff\x25\x5b\x5e\xc5\xf7\x00\x00\x00")

func specsetSpecsRegolitheGenCmdBytes() ([]byte, error) {
	return bindataRead(
		_specsetSpecsRegolitheGenCmd,
		"specset/specs/.regolithe-gen-cmd",
	)
}

func specsetSpecsRegolitheGenCmd() (*asset, error) {
	bytes, err := specsetSpecsRegolitheGenCmdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "specset/specs/.regolithe-gen-cmd", size: 247, mode: os.FileMode(420), modTime: time.Unix(1528494787, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _specsetSpecsIdentifiableAbs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xce\xc1\x6a\xc3\x30\x10\x04\xd0\xbb\xbf\x62\xa0\xe7\x16\x7a\xd5\xad\xd0\x4b\xbf\xa2\xc8\xd6\xc8\xd9\x60\x6b\xcd\x6a\x1d\xe2\xbf\x0f\x56\x20\xb1\x2f\xcb\xf2\x98\x81\xf9\xc0\x8f\xbb\x49\xbf\x3a\x6b\x17\x5f\x6f\xe8\x80\xdb\xf7\x7e\x3f\x51\xe2\xcc\x80\xbf\xdf\x0e\x00\x12\xeb\x60\xb2\xb8\x68\xd9\x0d\x52\xe1\x17\x42\x12\x8b\x4b\x16\x1a\x34\x37\xd1\xfe\xca\xc1\xbf\x5a\xc9\xb7\x85\x01\xd5\x4d\xca\xd8\x80\xf7\x45\x2b\x53\x80\xdb\xca\x26\xd5\xd5\x4e\x60\x8c\xe9\x5f\xcb\xb4\x1d\x2c\xae\xae\x23\x0b\x2d\xfa\x29\x9b\x65\x72\x5a\xec\x27\x1e\x51\x6d\x8e\x1e\x90\x8d\x4f\x78\x6f\x3c\xa4\xd4\xd2\xa9\xf9\x08\x00\x00\xff\xff\x5f\x95\x51\x9b\x10\x01\x00\x00")

func specsetSpecsIdentifiableAbsBytes() ([]byte, error) {
	return bindataRead(
		_specsetSpecsIdentifiableAbs,
		"specset/specs/@identifiable.abs",
	)
}

func specsetSpecsIdentifiableAbs() (*asset, error) {
	bytes, err := specsetSpecsIdentifiableAbsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "specset/specs/@identifiable.abs", size: 272, mode: os.FileMode(420), modTime: time.Unix(1533098329, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _specsetSpecs_apiInfo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x28\x4a\x4d\xcb\xac\xb0\x52\x48\x2c\xc8\xe4\x2a\xca\xcf\x2f\xb1\x52\x00\x91\x5c\x65\xa9\x45\xc5\x99\xf9\x79\x56\x0a\x86\x5c\x80\x00\x00\x00\xff\xff\x0c\x97\x42\xd8\x22\x00\x00\x00")

func specsetSpecs_apiInfoBytes() ([]byte, error) {
	return bindataRead(
		_specsetSpecs_apiInfo,
		"specset/specs/_api.info",
	)
}

func specsetSpecs_apiInfo() (*asset, error) {
	bytes, err := specsetSpecs_apiInfoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "specset/specs/_api.info", size: 34, mode: os.FileMode(420), modTime: time.Unix(1528494787, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _specsetSpecs_typeMapping = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\xcd\x49\xcd\x4d\xcd\x2b\x49\xcc\xb1\x52\xa8\xae\xe5\x02\x04\x00\x00\xff\xff\x14\xdf\xfc\xd6\x0e\x00\x00\x00")

func specsetSpecs_typeMappingBytes() ([]byte, error) {
	return bindataRead(
		_specsetSpecs_typeMapping,
		"specset/specs/_type.mapping",
	)
}

func specsetSpecs_typeMapping() (*asset, error) {
	bytes, err := specsetSpecs_typeMappingBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "specset/specs/_type.mapping", size: 14, mode: os.FileMode(420), modTime: time.Unix(1528494787, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _specsetSpecsMonolitheIni = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xcc\x41\xaa\xc3\x30\x0c\x04\xd0\xbd\x4e\xa1\x13\xf8\xff\x1e\xc0\xd0\x7b\x84\x50\x84\xab\x36\x86\x48\x32\x8a\xdc\xe0\xdb\x77\xe1\x2e\xda\xdd\xc0\xcc\xbc\x45\x4c\x6d\xaf\xb1\xf1\x0a\xcd\xed\xde\x4b\xdc\x94\x84\x31\xa3\x0c\x6a\x15\x8a\xb5\xe1\xf5\xb9\x05\x66\x1c\xd6\xbd\x98\x34\xd2\x01\xb0\x84\x93\x1e\x0f\x73\x61\x5f\xe1\xe7\xd3\x7d\xc7\x8c\xe7\x79\x26\x19\x9f\x7d\x52\x8e\xbf\xd9\x52\x8f\xcd\x7c\x72\xc0\x42\x75\x9f\xf9\xfa\xc5\xa7\x62\x02\x2f\xf6\xa3\x9a\x62\xc6\x4b\xfa\x87\x77\x00\x00\x00\xff\xff\x89\xa9\xf2\x82\xaa\x00\x00\x00")

func specsetSpecsMonolitheIniBytes() ([]byte, error) {
	return bindataRead(
		_specsetSpecsMonolitheIni,
		"specset/specs/monolithe.ini",
	)
}

func specsetSpecsMonolitheIni() (*asset, error) {
	bytes, err := specsetSpecsMonolitheIniBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "specset/specs/monolithe.ini", size: 170, mode: os.FileMode(420), modTime: time.Unix(1528853434, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _specsetSpecsObjectSpec = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x31\x6b\xc3\x30\x10\x85\x77\xfd\x8a\x07\x19\x32\xb5\xd0\x55\x53\x0b\x85\x4e\xa5\x4b\x3a\x17\xc5\x7a\xb6\xd5\xda\x92\xd1\x9d\x4a\xf2\xef\x8b\xec\x10\x62\x37\xcb\x21\xbe\xfb\xf4\xee\xa4\x1d\xde\x93\xe7\x60\xc6\x5a\xad\x01\x32\x45\xbf\xa2\x1b\x69\x91\x8e\xdf\x6c\x74\x61\xa9\xe4\x86\x2b\x2e\x06\x60\xd4\xa0\xe7\x0b\xfe\xb8\xe2\xc9\x35\x3f\xae\xa3\x85\x67\xeb\xca\x50\x23\x3c\xa5\xc9\x61\xd2\x90\xa2\xc5\xa1\x0f\x82\x20\xc8\x2e\xfa\x34\x5e\xf2\x1e\x0d\xd0\x51\xeb\x0e\x1b\xfd\x8d\x2a\xd0\x9e\x37\x62\x99\xbc\x53\xde\x71\x3f\xe7\xc6\x56\xf7\x1c\x78\x57\x7f\x9d\x1b\x5b\x9d\x27\x65\xf4\x52\xfd\x07\xec\x9f\x83\xaf\xef\x6c\x83\x3b\x0e\xdc\x1b\xb3\xc3\x8b\x6a\x0e\xc7\xa2\x14\xe3\xae\xc7\x6a\xff\x3e\x2d\x77\x96\x0f\xa9\xf5\xff\xc8\x43\xcf\xb9\x83\xd4\x6e\xc6\x02\x7a\x9e\x68\x21\x9a\x43\xec\x66\xc0\xd3\x94\x84\xde\x42\x73\x59\xb2\x44\x53\x5e\x81\x36\x0c\xca\x5c\x77\xbb\x81\x29\xfb\x15\xfb\x0b\x00\x00\xff\xff\x4a\xf2\xda\x7e\xe7\x01\x00\x00")

func specsetSpecsObjectSpecBytes() ([]byte, error) {
	return bindataRead(
		_specsetSpecsObjectSpec,
		"specset/specs/object.spec",
	)
}

func specsetSpecsObjectSpec() (*asset, error) {
	bytes, err := specsetSpecsObjectSpecBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "specset/specs/object.spec", size: 487, mode: os.FileMode(420), modTime: time.Unix(1533220356, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _specsetSpecsRootSpec = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\x31\x8e\xc2\x30\x10\x45\x7b\x9f\x62\xa4\xd4\xbb\x07\x70\xbb\xf5\x36\xb9\x00\x32\xce\x27\x18\x12\x4f\x34\x33\x01\x71\x7b\x64\x1c\x20\x20\x68\x2c\xeb\xcd\xff\x7a\xbf\xa1\x7f\xee\x30\xb8\xb1\xbc\xde\x11\x09\xd4\x36\x39\x8c\xf0\x24\xcc\x56\x09\xcf\x12\xf1\x4a\x91\x2d\xd9\x65\x61\x6d\x65\x53\x88\xc7\xd0\x3f\x33\x1d\x34\x4a\x9a\x2c\x71\xae\x8c\x78\x7b\x40\xb4\x5f\x47\xd4\xc3\x8a\xee\x2d\xd4\xc3\x94\x6c\x8f\x55\xb0\xf4\x3c\x99\xcc\x70\xae\xa1\x16\x43\x28\x51\x75\x72\xff\x79\xf7\xb3\x5e\x5d\x9b\x5f\x0d\x2d\x4c\x12\x4e\xa8\x9a\x21\xa9\x11\xef\x96\x92\x16\x5f\x14\x04\xc3\x87\xe6\xdf\xed\xa0\x14\x28\xe3\xfc\x18\x78\x0d\x00\x00\xff\xff\x22\x7b\xf8\x38\x41\x01\x00\x00")

func specsetSpecsRootSpecBytes() ([]byte, error) {
	return bindataRead(
		_specsetSpecsRootSpec,
		"specset/specs/root.spec",
	)
}

func specsetSpecsRootSpec() (*asset, error) {
	bytes, err := specsetSpecsRootSpecBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "specset/specs/root.spec", size: 321, mode: os.FileMode(420), modTime: time.Unix(1533220356, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _specsetSpecsType_mappingIni = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe2\x02\x04\x00\x00\xff\xff\x93\x06\xd7\x32\x01\x00\x00\x00")

func specsetSpecsType_mappingIniBytes() ([]byte, error) {
	return bindataRead(
		_specsetSpecsType_mappingIni,
		"specset/specs/type_mapping.ini",
	)
}

func specsetSpecsType_mappingIni() (*asset, error) {
	bytes, err := specsetSpecsType_mappingIniBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "specset/specs/type_mapping.ini", size: 1, mode: os.FileMode(420), modTime: time.Unix(1533220356, 0)}
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
	"specset/.gitignore": specsetGitignore,
	"specset/Gopkg.toml": specsetGopkgToml,
	"specset/specs/.regolithe-gen-cmd": specsetSpecsRegolitheGenCmd,
	"specset/specs/@identifiable.abs": specsetSpecsIdentifiableAbs,
	"specset/specs/_api.info": specsetSpecs_apiInfo,
	"specset/specs/_type.mapping": specsetSpecs_typeMapping,
	"specset/specs/monolithe.ini": specsetSpecsMonolitheIni,
	"specset/specs/object.spec": specsetSpecsObjectSpec,
	"specset/specs/root.spec": specsetSpecsRootSpec,
	"specset/specs/type_mapping.ini": specsetSpecsType_mappingIni,
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
	"specset": &bintree{nil, map[string]*bintree{
		".gitignore": &bintree{specsetGitignore, map[string]*bintree{}},
		"Gopkg.toml": &bintree{specsetGopkgToml, map[string]*bintree{}},
		"specs": &bintree{nil, map[string]*bintree{
			".regolithe-gen-cmd": &bintree{specsetSpecsRegolitheGenCmd, map[string]*bintree{}},
			"@identifiable.abs": &bintree{specsetSpecsIdentifiableAbs, map[string]*bintree{}},
			"_api.info": &bintree{specsetSpecs_apiInfo, map[string]*bintree{}},
			"_type.mapping": &bintree{specsetSpecs_typeMapping, map[string]*bintree{}},
			"monolithe.ini": &bintree{specsetSpecsMonolitheIni, map[string]*bintree{}},
			"object.spec": &bintree{specsetSpecsObjectSpec, map[string]*bintree{}},
			"root.spec": &bintree{specsetSpecsRootSpec, map[string]*bintree{}},
			"type_mapping.ini": &bintree{specsetSpecsType_mappingIni, map[string]*bintree{}},
		}},
	}},
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

