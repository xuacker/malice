// Code generated by go-bindata.
// sources:
// config/config.toml
// DO NOT EDIT!

package config

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

var _configConfigToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x53\x49\x6f\xdb\x3c\x10\xbd\xf3\x57\x0c\xe4\xcb\xf7\x01\x8d\x2c\x4b\x59\x5c\x01\x3e\x04\x41\x0e\x29\x9a\x16\x48\x8e\x46\x50\x8c\xa4\xb1\x44\x98\x1b\x48\xca\x59\x7e\x7d\x31\x94\xb3\x18\xf5\xa1\x05\x42\x1f\x68\xce\x9b\x99\xf7\xe6\x8d\x3d\x83\x2b\xeb\x9e\xbd\xec\x87\x08\xff\xb5\xff\x43\x59\x2c\x2a\x38\xe1\x6b\x09\x8d\xc2\x76\x1b\xad\x83\x6f\x36\x0c\x23\xc2\x2d\x4a\x43\x5f\xe0\x52\x29\xb8\xe3\x82\x00\x77\x14\xc8\xef\xa8\xcb\xc5\x0c\xee\x89\xe0\xfb\xcd\xd5\xf5\x8f\xfb\x6b\xd8\x58\x0f\x4a\xb6\x64\x02\x81\x34\x1b\xeb\x35\x46\x69\x4d\x2e\xc4\xec\x73\x8e\x98\xc1\xed\x25\xb3\xc1\x95\x35\x1b\xd9\x8f\x3e\x11\xc0\xbf\xf7\xf9\x24\x3d\x22\xca\xa8\x08\x56\x90\xdd\x22\x4f\x0e\x77\xa3\x89\x52\xd3\xa1\xbe\x4c\xec\xc8\x07\x16\xba\x82\x6c\x57\xe4\x55\xbe\x38\xcb\x84\x58\xe3\x18\x07\xeb\x1f\x04\x80\x41\x9d\xba\xbc\x7a\x9f\x09\x00\xeb\x7b\x34\xf2\x65\x9a\xf0\x8d\xe1\xe6\x27\x57\x3e\x52\xc3\x65\xa3\x57\x8c\x14\x79\xfa\xd4\xcb\x82\xeb\xb0\xd3\xd2\xfc\xda\x43\x8b\xf2\x22\x81\x8b\xba\xaa\xaa\x8a\x4b\x49\xa3\x54\x5c\x3c\xd8\x10\x39\x25\xe8\xe8\x72\x7a\x42\xed\x14\xe5\xad\xd5\xdc\xc3\x59\xcf\x58\x79\xc6\x24\x81\x3c\xe7\xf1\xcd\x3a\x13\x8e\x21\x70\x8c\xef\x47\xeb\x3b\x6e\xdc\x61\xc4\x06\x03\x7d\x9c\x47\x27\xcd\x27\xa4\x30\x44\xd9\x72\xa5\xd4\xd8\x7f\x80\xe6\x7b\x28\x10\xfa\x76\xa8\xcf\xf3\x8a\x93\xd2\xcf\x2b\x91\x2a\xdb\xa2\x62\xa5\xaf\xaa\x98\x76\xfd\xb5\x2c\x0a\xa6\x61\xab\xed\x98\x94\x16\x02\x80\x0c\x36\x8a\x3a\x58\x41\xf4\x23\x09\xb1\x1e\xe5\x11\x31\x5b\xd9\xa0\xc1\x63\x5a\x26\xe4\x6f\x45\x9c\x9e\x56\x0f\xc7\x48\xc9\xec\xa4\xb7\x46\x93\x89\x8c\xfb\x31\x6d\xaf\xa3\x1d\x29\xeb\x38\x9a\xcc\xb2\xed\x96\xd2\xea\x35\xb6\x83\x34\x74\x72\xa8\x32\x4b\x9d\x3b\x67\xa5\x49\x4b\x8a\xad\xab\xe7\xf3\x37\x21\x75\x59\x5d\x9c\x67\x07\x0e\x2c\x92\x05\x8d\x34\x5d\x78\x6f\x53\xcf\x35\xaa\x47\xf4\x54\x7b\xcb\xe9\x4a\x9a\x6d\xf8\x73\x31\xf5\xc1\x16\x38\xb1\x75\x23\xac\xe0\xac\xd8\x1f\xd6\x49\xda\xfa\x67\x0e\x96\xa7\xe5\x72\xc9\x41\xb1\x56\xb6\xef\xa7\x31\x36\x52\xd1\xe1\x08\xb9\xb2\x7d\x96\x06\x7c\x0a\xf2\x85\x81\x45\x31\x3d\x27\xd7\xab\xfd\xab\xc1\x76\x3b\x3a\x56\x75\xc1\x0a\x79\xc4\xf4\x17\x5a\xc1\x06\x55\x60\x47\x9d\xb7\x4f\xcf\xef\x5e\xbf\x21\x00\x43\x8c\x8e\x19\xb3\xfd\xf7\x30\x3d\x7e\x07\x00\x00\xff\xff\x32\x81\x18\x7e\xdf\x04\x00\x00")

func configConfigTomlBytes() ([]byte, error) {
	return bindataRead(
		_configConfigToml,
		"config/config.toml",
	)
}

func configConfigToml() (*asset, error) {
	bytes, err := configConfigTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/config.toml", size: 1247, mode: os.FileMode(420), modTime: time.Unix(1534725911, 0)}
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
	"config/config.toml": configConfigToml,
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
	"config": &bintree{nil, map[string]*bintree{
		"config.toml": &bintree{configConfigToml, map[string]*bintree{}},
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

