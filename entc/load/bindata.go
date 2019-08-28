// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// template/main.tmpl
// schema.go
package load

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

var _templateMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x91\x41\x6b\xdc\x30\x10\x85\xcf\x9a\x5f\x31\x15\x5b\x22\x81\xa3\xd0\x6b\x61\x4f\xcd\x1e\x52\x68\x52\xd8\x40\x0f\xdb\x25\xc8\xf6\x78\x23\x6a\xcb\xae\xa4\x94\x06\xa1\xff\x5e\x24\xd9\x0b\x3d\xd9\xf2\x7b\xfa\xe6\x3d\x4f\x8c\xd8\xd3\x60\x2c\x21\x9f\xb4\xb1\x1c\x53\x82\xbb\x3b\xfc\x32\xf7\x84\x17\xb2\xe4\x74\xa0\x1e\xdb\x77\xbc\x21\x1b\xba\xeb\xa7\x1b\x85\xf7\x4f\xf8\xf8\xf4\x8c\x87\xfb\x87\x67\x05\x8b\xee\x7e\xe9\x0b\x61\x66\x00\x98\x69\x99\x5d\x40\x01\x8c\xcf\x9e\x03\xe3\xed\x7b\xa0\xfc\x12\x23\x06\x9a\x96\x51\x07\x42\x5e\x5d\xbe\x8c\x2c\xd2\xe2\x8c\x0d\x03\xf2\x8f\xbf\x39\xaa\xef\x2b\x31\x25\x90\x00\x31\xe2\xae\xd5\x9e\xf0\xf3\x1e\xcb\x73\xd3\xf3\xdd\x3f\xda\xa1\xef\x5e\x69\xd2\x1e\xf7\x78\x3a\x93\x0d\xea\xc1\x06\x72\x83\xee\x28\x16\xb4\xd3\xf6\x42\xb8\x7b\x69\x70\x67\xf5\x54\x30\xea\x51\x4f\xe4\x33\x9f\xb1\x18\x6f\x57\x7e\x4a\x2a\x1f\xae\x51\x7c\x4c\x7c\xbd\x93\x52\x53\x58\x64\x7b\xbc\x4d\x09\x12\xc0\xf0\x66\xbb\xd2\x59\x48\x8c\xc0\x72\x90\xd1\x58\xf2\x78\x3a\x9f\xce\xb9\x34\xb0\x61\x76\xf8\xd2\xac\xf9\xf2\xdc\x1a\x65\xcb\x1b\x81\xb1\xb6\x41\x72\x2e\x6b\xdf\xb4\xf3\xaf\x7a\x3c\x16\x51\x54\x8f\x04\xc6\xcc\x50\x1c\x1f\xf6\x68\xcd\x58\xee\xb0\x41\x9b\x51\x90\x73\x59\xce\x15\xea\xdc\x3d\xea\x65\x21\xdb\x8b\x72\x6c\xb0\x95\x90\xd5\xd9\xab\x63\xe8\xe7\xb7\xa0\x7e\x38\x13\x48\x94\x7d\xa8\xaf\xb3\xb1\x9b\xb1\xc6\x15\xfc\xa7\xe5\x52\xca\x6b\xb7\x6d\x4a\x1e\x3f\xbb\x52\xb2\xb2\xc8\xb9\xca\x3a\x06\x67\xec\x25\x7b\xd4\x21\x7b\x84\x94\xc5\x73\xf8\x6b\x82\xf8\x54\x48\xff\x6d\xbd\x96\xaa\x4b\x5f\x7f\x66\x4a\xf0\x2f\x00\x00\xff\xff\xb5\xb1\x2f\xf6\x87\x02\x00\x00")

func templateMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMainTmpl,
		"template/main.tmpl",
	)
}

func templateMainTmpl() (*asset, error) {
	bytes, err := templateMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/main.tmpl", size: 647, mode: os.FileMode(420), modTime: time.Unix(1566131522, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x4b\x6f\xdb\x38\x10\x3e\x4b\xbf\x62\x6a\xa0\x81\x14\x78\xe5\x6e\x6f\xab\xc2\x87\xa2\x4d\x81\xec\x23\x5d\x34\xdd\xbd\x14\x45\x4b\x4b\x43\x9b\xad\x44\xaa\x24\xed\x34\x0d\xf2\xdf\x17\xc3\x87\x1e\xb6\x93\xdd\x6d\x91\x00\x81\xa5\x99\xf9\x86\x9c\x6f\x1e\xa4\x16\x0b\x78\xa1\xba\x6b\x2d\xd6\x1b\x0b\x4f\x9f\xfc\xfc\xcb\x4f\x9d\x46\x83\xd2\xc2\x2b\x56\xe1\x4a\xa9\xcf\x70\x2e\xab\x02\x9e\x37\x0d\x38\x23\x03\xa4\xd7\x3b\xac\x8b\x74\xb1\x80\xb7\x1b\x61\xc0\xa8\xad\xae\x10\x2a\x55\x23\x08\x03\x8d\xa8\x50\x1a\xac\x61\x2b\x6b\xd4\x60\x37\x08\xcf\x3b\x56\x6d\x10\x9e\x16\x4f\xa2\x16\xb8\xda\xca\x9a\x5c\x08\xe9\x4c\x7e\x3f\x7f\x71\x76\x71\x79\x06\x5c\x34\x18\x65\x5a\x29\x0b\xb5\xd0\x58\x59\xa5\xaf\x41\x71\xb0\xa3\xf5\xac\x46\x2c\xd2\xb4\x63\xd5\x67\xb6\x46\x68\x14\xab\xd3\x54\xb4\x9d\xd2\x16\xb2\x34\x99\xa1\xac\x54\x2d\xe4\x7a\xf1\xc9\x28\x39\x4b\x93\x19\x6f\x2d\xfd\x68\xe4\x0d\x56\x76\x96\xa6\xc9\x6c\x2d\xec\x66\xbb\x2a\x2a\xd5\x2e\x78\x08\x58\xc8\x6a\xbb\x62\x56\xe9\x05\x4a\xbb\x30\xd5\x06\x5b\xb6\xc0\x7a\x8d\xff\x09\x30\xfb\x1f\x4e\xb9\xc0\xa6\x9e\xa5\x79\x4a\x34\x5c\x3a\x19\x68\x0c\x09\x30\xc0\x24\xa0\xb4\x45\x50\xd8\x0d\xb3\x70\xc5\x8c\x8b\x13\x6b\xe0\x5a\xb5\xc0\xa0\x52\x6d\xd7\x08\x22\xdb\xa0\x86\xc0\x45\x91\xda\xeb\x0e\xa3\x4b\x63\xf5\xb6\xb2\x70\x93\x26\x17\xac\x45\x00\x20\x89\x90\x6b\x00\xf8\x48\xd4\x94\x33\xc9\x5a\x9c\xab\x56\x58\x6c\x3b\x7b\x3d\xfb\x98\x26\x67\xf5\x1a\x0d\x00\xbc\x7b\x7f\x4a\x8f\xbd\x25\xf1\x60\xa6\xa6\xaf\x28\x0a\xe3\x4c\xdd\x63\x34\x75\xd1\xed\xd9\x9e\xcb\x1a\xbf\xa2\x21\x5b\xf7\x18\x6d\x85\x97\x4f\x8c\x6f\x1d\x2d\xde\xe5\x21\x2b\x5e\xfe\x1d\xa4\x78\xe0\x21\x27\xfe\xaf\x67\xe6\x1e\x6e\xde\x92\x9b\xfe\xcf\x85\x59\x38\x59\x40\xd0\x32\x7b\x08\xb6\x86\xfb\xd6\xb0\x6c\x3d\x05\x5c\x8a\x6f\xa3\x25\x4e\x85\xb4\xe1\x31\x00\x8c\xf8\xb6\xb7\xc4\x8b\x0d\xd3\x06\xa3\xd9\xe9\xb0\x46\x40\x54\x5e\xbf\x07\x52\x4d\xc3\xac\x50\xf2\x2e\x50\xd4\x4f\x61\x7f\x49\xf1\x65\xdb\xef\x6f\xa5\x54\x33\xdd\xdd\xd6\xe9\xa7\x98\x0b\xd1\x34\x6c\xd5\xe0\x5d\x18\x19\xf4\x53\xd4\xeb\x8e\x56\x67\xcd\x5d\x28\x15\xf4\x53\xd4\x4b\xe4\x6c\xdb\xd8\x3b\xf7\x57\x7b\xfd\x5e\x50\x5d\xcd\x2c\x46\xe8\x91\xa0\x9c\xfe\xc3\x51\xec\x79\xdb\x6e\x6d\x1f\xdd\x21\x56\x44\xfd\x14\xf6\x37\x6b\x44\x4d\x43\x81\x3a\x0d\x86\x24\x47\xd8\xae\xd7\x1f\x69\x0c\xd7\x96\x87\x7d\xe1\xc4\xdf\xd1\x16\x0e\x77\xa4\x2b\x42\x4d\xfc\x7b\x33\x4c\x0d\xef\xe9\x81\x3d\xc3\xfd\xd2\x7f\x83\xdc\x2f\x3e\xb5\xd3\xc8\x3f\x1c\xae\xfe\x06\x79\x28\xf8\xc9\x94\xd2\xc8\xef\xa8\xd8\x90\x9b\x7b\x0a\xf5\x5c\xee\x50\x1b\xdc\x37\x15\x5e\xbc\xbf\xfc\x97\xad\xd0\x58\xef\xd9\xea\x20\x3e\x92\x35\x3f\xf5\x0e\xd3\xe6\xe5\xdf\x91\x37\x0f\x1c\x12\x17\x22\xed\x6b\xf0\x9e\x48\xc3\x8c\x7f\xf7\x7e\xca\xf4\xdd\x23\x7e\xdf\xf2\xc8\x84\xf7\x51\x5e\xe0\x95\xcb\x47\xa5\x91\x59\x74\x41\x86\x88\xc8\xb9\x0f\xcb\x3d\xd5\x68\x2a\x2d\x3a\xab\x74\x91\xf2\xad\xac\x22\x32\xc3\x1a\x4e\xc9\xa2\x78\xd9\x5b\xe4\x21\xc9\x37\x69\x22\x11\xca\x25\x9c\xd0\xeb\x4d\x9a\x50\x69\x95\xbe\x0c\xb0\x2e\xde\xb2\xf5\x9c\x64\xd7\x1d\x96\xbd\x8c\xaa\x31\x4d\x5c\x55\xf7\x42\x7a\x21\xa1\x67\xac\xf4\x42\xff\x42\xe2\x50\x07\xa5\x13\x87\x17\x92\xc7\x9c\x97\x24\x8f\x2f\x5e\xc1\x83\x7f\xa7\xe0\xc1\xff\x6d\x9a\x08\x0e\x1a\x39\x6d\xd9\x6b\x9e\xb9\xd7\x47\x4b\x90\xa2\xa1\x70\x12\x89\x24\x86\x65\x1f\xbe\x46\x9e\x3b\xa8\x46\xbb\xd5\x12\x24\x06\x66\xff\x60\xda\x6c\x58\x13\x4e\x76\x77\xc3\x41\x77\x55\x1a\xdd\x14\x84\xb4\xa8\xe9\xe2\x41\x4f\x0a\x18\xfc\x7a\xf9\xfa\x82\xc0\xae\xbc\x2a\x26\x61\x45\xcc\x13\xb4\xf6\x26\xe4\x20\x80\xd5\xea\x13\x56\x36\xfc\x84\xa4\x4c\x16\xcd\x4c\x5c\x9b\xaa\x36\xac\x94\x43\xb6\x82\x77\xef\x57\xd7\x16\xe7\x80\x5a\xd3\x3f\x65\xec\x26\x4d\x8c\x4b\x95\xc7\xde\x78\x82\x84\xf4\x77\xba\x2c\xdc\xc4\x5c\x7e\x5e\xf3\xe0\x39\xcf\x5d\x6a\xb2\xfc\x36\x4d\x42\x85\x39\x97\xe5\x12\x0c\xe3\xe8\x6b\x31\xda\x3a\x72\x49\x3b\x62\x33\x72\x26\x9a\x39\xf0\xd6\x16\x67\xb4\x17\x9e\xcd\xc2\xc6\x1f\x7f\x29\xe1\xf1\x6e\x36\x07\xe3\x4b\x80\xe0\x9e\x6c\xae\x34\x7c\x98\x83\xcb\x94\x66\x92\x2a\xd5\x17\x3e\x79\xe5\x35\x89\xf9\xa8\x20\xb3\x3c\x4d\x12\xe3\xac\x4f\xdc\xae\xc8\x6c\x54\x63\xfe\x7a\x30\x14\xda\xa8\x26\xa3\x2a\x16\xe6\xa8\x84\x7b\x95\xaf\xe3\x51\x79\x46\xcd\x50\xa3\xfd\xa9\x5a\x0e\x8b\xc5\x73\x94\xd4\xf1\xf8\x1c\xd4\x51\xe2\xd4\xfd\xb1\x55\x46\x75\x2f\x71\xfa\xe1\x7c\x72\x06\x0d\xca\x8c\xd7\xc5\x20\xcd\x9d\x55\x38\x32\xcb\x61\x83\xf1\x10\xf5\x39\xf1\x51\x8c\x4f\xd7\xd2\x45\x31\x39\x6f\x07\xd3\xdb\x34\xa1\x9c\xf2\xba\x70\x97\xa0\x47\x4b\x78\xe2\xf8\x4f\x0c\xf7\x92\x25\x9c\x04\xe5\xc4\x3a\x5e\x80\x1e\x2d\x61\x36\xeb\x11\x51\xea\x41\xe1\x6d\x8a\xeb\xef\x40\x7b\xc8\x5e\x1e\xb0\xf1\x3d\xa0\x4d\x11\xa6\xe2\x12\x58\xd7\xa1\xac\xb3\x28\x99\x83\x09\xdd\xeb\x47\xe9\xb8\x7a\xdd\xcc\x7d\xc8\xe2\xc5\xa1\x78\xdd\xea\xce\xa9\x29\xfc\xac\x1f\x6d\xf5\xcc\x6f\xad\x9f\xb7\x93\xba\xce\xbd\xcb\x78\x2b\x1f\x07\x10\x6e\xf0\x0f\x19\x82\xa8\xbf\x0e\x41\x84\x3d\x38\xc7\x41\x21\xea\xaf\x07\x5d\x58\xc4\x2f\x8b\x51\x88\xe7\x71\xfb\x27\xee\xc9\x25\xd6\x85\x5d\x82\xf3\xe1\x29\x20\xa9\xcf\x5b\xe9\xa4\x21\x87\xe3\xd6\x23\xf1\xd0\x74\xb7\x93\xc9\x4c\x27\x61\x11\x06\x64\x66\xf2\x30\xa6\x87\x41\x05\x57\x9a\x75\xc6\x4d\x58\x1f\x7f\x2c\x9b\x16\xed\x46\xd5\x70\x25\xec\x06\x34\x56\x6a\x47\x9f\xca\x0a\x50\x9a\xad\x46\x90\x0a\x3a\x26\x45\x65\xe8\x53\xb8\xf5\xee\x85\x5c\x87\x81\x7c\x30\x07\x0f\xa6\x31\x8f\x27\x76\xff\xad\xb4\x3f\x97\x6b\xe4\xa8\x81\xdc\x65\xb9\x67\x97\xc3\xce\xf1\xee\x37\x93\xe5\xcf\x60\x37\x4e\x6b\x42\xf8\xe5\x91\x8c\xc6\x88\xfc\x86\x43\x72\x77\x94\x96\x30\xbf\xc1\x39\xf1\x7d\x73\x4b\xf9\x0a\xdc\x4d\xe0\x59\x3e\x77\x56\x03\x81\xbe\x66\x0f\xf8\xf3\xe2\x1f\xa5\x6f\xdc\x88\x07\xec\xf9\xce\xf1\xe4\x91\xe1\x03\x72\xe7\xa3\x39\x42\x1d\x86\x8e\xbd\x8f\x39\x1f\xc4\x01\x71\xb1\x17\x0e\xa8\x8b\x8a\x1f\x25\x6f\x3a\x04\x0e\xe8\x13\xfd\x57\x7e\x7f\xb5\x7d\x40\x06\x63\x50\x47\x38\x14\xfd\x50\xb8\x8f\xc5\x18\xcd\xc0\xa3\x0b\xb4\xbf\xa5\x58\x18\xdf\x53\xf2\xc9\x1b\xed\x8d\xc6\x96\x2d\x7e\x13\xb2\xce\x72\x58\x2e\x7b\xfd\x9f\x56\xbb\xad\xd3\xe9\x63\x8b\xb3\x06\xdb\x6c\x32\x3a\x6c\x7a\x9b\xfe\x13\x00\x00\xff\xff\x2b\xd1\x21\x50\x80\x13\x00\x00")

func schemaGoBytes() ([]byte, error) {
	return bindataRead(
		_schemaGo,
		"schema.go",
	)
}

func schemaGo() (*asset, error) {
	bytes, err := schemaGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.go", size: 4992, mode: os.FileMode(420), modTime: time.Unix(1566995215, 0)}
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
	"template/main.tmpl": templateMainTmpl,
	"schema.go":          schemaGo,
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
	"schema.go": &bintree{schemaGo, map[string]*bintree{}},
	"template": &bintree{nil, map[string]*bintree{
		"main.tmpl": &bintree{templateMainTmpl, map[string]*bintree{}},
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
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
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
