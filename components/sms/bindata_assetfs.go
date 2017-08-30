// Code generated by go-bindata.
// sources:
// templates/views/index.html
// DO NOT EDIT!

package sms

import (
	"github.com/elazarl/go-bindata-assetfs"
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

var _templatesViewsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\x5d\x6f\xeb\x36\x0c\x7d\xcf\xaf\xe0\x55\x8b\xda\x19\x6e\xec\x35\x29\xee\xee\x5a\x3b\xc3\x30\xec\xb1\x5b\x81\x00\x7b\x19\x86\x42\xb1\xe8\x44\x85\x2c\xf9\x4a\x72\xda\x34\xc8\x7f\x1f\xe4\xaf\x3a\x89\x9b\x04\x58\xb7\xe9\x25\xb6\x78\x24\x92\x87\xc7\x64\x36\x1b\x60\x98\x72\x89\x40\x12\x25\x2d\x4a\x4b\x60\xbb\x1d\x44\x8c\xaf\x20\x11\xd4\x98\x98\x68\xf5\x4c\xa6\x03\x00\x80\xee\x6e\xa2\xc4\x28\x63\xa3\xeb\x31\xb8\x27\x93\x35\x4f\x2f\x66\x74\x3d\xae\xf1\xfb\x67\x5e\x1e\x73\x2a\x51\x74\xac\x87\x08\xcb\xad\xc0\x3d\x44\x89\x5a\x8e\xa7\x33\x94\x0c\x66\xf7\xb3\x28\x5c\x8e\x7b\x10\xdd\xe8\x04\x52\x9d\xf2\x17\x32\x8d\x42\xc6\x57\x7b\x0e\x7b\xb6\x76\x62\x68\x98\xe8\xf1\x91\x4f\x7f\x29\xb4\x46\x69\x61\x4e\x05\x95\x09\x02\x37\x10\x25\x8a\xe1\x74\xb3\x81\xa0\xde\xfc\x83\x8a\x02\x61\xbb\x8d\xc2\xd2\x02\x57\x17\x5f\x27\x5f\xaf\xef\xa2\x30\x9f\x0e\x0e\xef\x4c\x95\xce\x1a\xe7\xee\x79\xb4\x54\x9a\xbf\x2a\x69\xa9\x80\xf2\x5d\xd0\x39\x8a\x91\xc0\xd4\x12\xd0\x4a\x60\x05\x23\x90\xa1\x5d\x2a\x16\x93\x5c\x19\x4b\x80\xb3\x98\x18\x94\x6c\x64\x32\x43\x80\x26\x96\x2b\x19\x93\x0b\x02\x52\xad\xa8\xe0\x8c\x5a\x3c\x4c\x68\x3f\x79\x6e\x31\xab\x9c\x2e\xb4\x2a\xf2\x1e\x0a\xda\x53\x65\x54\x0e\x1b\x93\x7c\xa9\x24\x92\x37\x69\x48\xab\x95\xa8\xc2\x86\x5a\x28\x93\x46\x27\x93\x5e\x99\xf4\xad\x07\x77\x2b\x44\x26\xa7\xb2\xd5\x22\x7e\x2b\xb8\x46\x46\xa6\xdf\x45\xa1\x33\x1c\x89\x2f\x2c\xfd\x1f\x01\x1c\xaa\xf9\x4b\x13\xe4\x97\xb3\x83\x8c\xb8\xcc\x0b\x0b\x76\x9d\x63\x4c\x2c\xbe\x58\xb2\x53\xca\x9a\x8b\xaa\x3a\x35\x4d\x92\x66\xd8\xbe\xe4\x82\x26\xb8\x54\x82\xa1\x8e\xc9\x0f\x3f\x7e\x7f\x3d\x1e\x8f\x27\x93\x9b\x1b\x02\x2b\x27\xa2\x98\x10\x68\x92\xee\xa6\x7f\x24\xed\x03\x71\x9f\x63\xfa\xc7\x12\xc8\xd0\x18\xba\xf8\x78\x11\xdc\x57\xf7\xfe\xef\x65\x76\xa5\xa5\x1a\x69\x5f\x75\x41\xa3\xe1\xaf\x74\x2e\xf0\xb1\x81\xb9\x0f\xf5\xd9\xc4\xe4\xa6\x29\x77\xcb\x8f\x13\x42\xfb\xd2\x57\xda\x28\x6c\x2e\xf9\x37\xab\x2c\xe4\xa3\x51\x82\xb3\xde\x0e\xd9\x77\xe0\x3c\x45\xbc\x47\x76\xc6\x46\x2a\x4d\x0d\xda\xd1\xe4\x14\xd3\xf3\xc2\x5a\x25\xdb\x76\x46\xea\x6f\xcb\x14\xf3\x8c\xbf\x7d\x5d\x73\x2b\x61\x6e\xe5\xc8\x14\x49\x82\xc6\x90\x72\x38\x44\x61\x75\xf8\x23\x89\x8b\x42\x97\xf9\xd1\x09\xd2\x79\xad\x1f\xeb\x9f\xcd\x06\xdc\xc8\xda\x6e\x07\x83\xce\xa0\x7d\x32\xd5\x8c\x35\x89\xe6\xb9\x05\xa3\x93\x98\x84\x8c\x9a\xe5\x5c\x51\xcd\x42\x6a\x0c\x5a\x13\xae\x50\x32\xa5\x4d\x58\xf7\x6e\xa5\xc3\xa7\xce\x4b\x90\x71\x19\x3c\x99\x9f\x56\xb1\x1b\x3b\x3f\xe7\xb9\xe0\x09\x75\x1d\x3f\x98\x17\x5c\x38\x97\xae\xb0\x95\x87\xe9\xa0\xf5\x25\xa8\x5c\x14\x74\x81\x31\x79\xa2\x2b\x5a\x6d\xd6\xf5\xb8\xf4\x99\x4a\x8a\x0c\xa5\x1d\x06\x1a\x29\x5b\xfb\x69\x21\xcb\x21\x02\xfe\x10\x36\x6d\xb2\x9b\x0d\xf0\xb4\x9d\x74\xbf\x6a\xad\xb4\xcb\xa6\x31\x4b\x7c\x86\x87\xdf\x94\xe5\xe9\xda\xdf\xec\x90\x56\x4e\xf6\x5b\xf0\xca\x23\xde\xe7\x5d\x1b\xbe\xd8\x5b\xf0\x3a\x23\xb4\xb9\x78\x1f\xb8\xce\xdd\x1d\xd8\x73\xc7\x92\x33\xbc\x85\x94\x0a\x83\xbb\x06\x63\xd7\x82\xcb\xc5\x2d\x78\x73\xa5\xac\xb1\x9a\xe6\x13\xaf\x45\x6c\x87\x77\xdd\xdc\x9a\x7a\x35\x5b\x2b\xaa\xcb\x66\x08\x31\x5c\xfa\xde\x45\x33\x60\xbd\xe1\xdd\x1b\xc6\xd9\x83\x4a\x9f\x1d\xd2\xb0\xcb\x9a\x5b\x3c\x05\xff\x53\x89\x65\xd4\x52\xdf\x2b\x8b\xe9\x0d\xf7\x61\x6e\x69\xb4\x85\x96\x55\x32\x77\x3b\xd6\xed\xee\x1f\x88\xcb\xc0\x0d\x7f\xbf\xbc\x95\x5a\xab\x7d\xaf\x1a\xfc\xde\xf0\x73\x1d\x16\x6a\x4e\x05\x7f\x45\x7f\x38\x0c\x98\x92\xd8\x46\xe8\xbb\x28\xfa\x9c\xbb\x38\x9d\x2d\x58\x52\xf3\xfb\xb3\x7c\xd0\x2a\x47\x6d\xd7\x7e\xcd\xfa\x10\xae\xae\xc0\xd9\xff\xac\x37\xfe\x82\x4f\x71\x0c\xb2\x10\xe2\xc0\x12\x08\x94\x0b\xbb\xec\xf3\x02\x47\xc5\xd2\x5d\xc7\x84\xb3\x83\x2b\x45\xb4\xe3\xff\x08\xf8\x7d\x21\x75\xd7\xbb\xa2\xea\xae\xe3\x02\xeb\xae\xae\xd8\xda\x3d\x40\x61\xf0\x43\x18\x9a\x55\xcd\xf0\x24\x47\xde\xec\x7e\x06\xf5\x10\x02\x27\x69\x30\xa7\x4f\x56\x84\x9d\x06\xfe\x17\x94\xed\x7e\x12\xc3\x20\xa5\x5c\xbc\x09\xbb\x4f\x6e\xa7\x89\x3c\x47\x66\x3d\xf4\x51\xc6\xc0\x2a\xf8\x56\x60\x81\xe0\xe2\x40\xf6\xde\xe1\xd3\x92\x3b\xc9\xdd\x79\xbc\xed\x73\xb6\xed\x76\x2b\x78\xb7\xbf\x34\xc7\xdc\xef\xdb\xf0\x68\x9b\xe2\xdf\x01\x00\x00\xff\xff\x12\xe7\x07\x47\x2c\x0e\x00\x00")

func templatesViewsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsIndexHtml,
		"templates/views/index.html",
	)
}

func templatesViewsIndexHtml() (*asset, error) {
	bytes, err := templatesViewsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/index.html", size: 3628, mode: os.FileMode(420), modTime: time.Unix(1504098955, 0)}
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
	"templates/views/index.html": templatesViewsIndexHtml,
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
		"views": &bintree{nil, map[string]*bintree{
			"index.html": &bintree{templatesViewsIndexHtml, map[string]*bintree{}},
		}},
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


func assetFS() *assetfs.AssetFS {
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: k}
	}
	panic("unreachable")
}
