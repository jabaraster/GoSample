package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _assets_css_style_css = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\xca\x4f\xa9\xac\x4e\x4a\x4c\xce\x4e\x2f\xca\x2f\xcd\x4b\xb1\x52\x4e\x4b\x4b\xb6\x48\x35\xab\xcd\x30\x44\x16\xcd\x2f\x4a\xcc\x4b\x4f\xb5\x4e\xce\xcf\xc9\x2f\xb2\x2a\xcf\xc8\x2c\x49\xad\x4d\xc9\x2c\xd3\x4b\xca\x2f\x4a\x49\x2d\x4a\xa9\x86\xd0\x56\xc5\xf9\x39\x99\x29\x0a\x46\x05\x15\x0a\x45\xa9\x29\xb5\x80\x00\x00\x00\xff\xff\xaf\xae\x66\xf1\x5a\x00\x00\x00")

func assets_css_style_css_bytes() ([]byte, error) {
	return bindata_read(
		_assets_css_style_css,
		"assets/css/style.css",
	)
}

func assets_css_style_css() (*asset, error) {
	bytes, err := assets_css_style_css_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "assets/css/style.css", size: 90, mode: os.FileMode(420), modTime: time.Unix(1425963072, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _assets_js_script_js = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\x48\x2b\xcd\x4b\x2e\xc9\xcc\xcf\xd3\xd0\x54\xa8\xe6\x52\x00\x82\xc4\x9c\xd4\xa2\x12\x0d\x43\x43\x43\x4d\x6b\xae\x5a\x4d\x0d\x4d\x6b\x40\x00\x00\x00\xff\xff\x80\x0e\xab\xd3\x23\x00\x00\x00")

func assets_js_script_js_bytes() ([]byte, error) {
	return bindata_read(
		_assets_js_script_js,
		"assets/js/script.js",
	)
}

func assets_js_script_js() (*asset, error) {
	bytes, err := assets_js_script_js_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "assets/js/script.js", size: 35, mode: os.FileMode(420), modTime: time.Unix(1425962580, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _assets_layout_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x90\x31\x6e\x03\x21\x10\x45\xfb\x9c\x62\x3c\xbd\x8d\xdc\xa5\x00\x9a\x34\x39\x40\x2e\x80\x61\x1c\x56\xc1\x26\x82\x71\x94\x95\x45\x99\x3b\xe4\x7c\x39\x49\x60\x59\xef\x2e\x42\xe2\xf3\x47\xbc\xf9\x8c\xdc\xb9\x68\x79\xfc\x24\xf0\x7c\x09\xfa\x49\xf6\x03\xea\x92\x9e\x8c\xeb\x72\xba\x5e\x88\x0d\x58\x6f\x52\x26\x56\x78\xe3\xf3\xfe\x19\xc5\xa6\x1e\x86\xeb\x07\x24\x0a\x0a\x33\x8f\x81\xb2\x27\x62\x04\x9f\xe8\xac\x50\xd8\x9c\xc5\x64\x1f\xaa\x42\x68\x1d\x15\x32\x7d\x73\xab\xe0\x86\xc2\x03\x07\xd2\xf7\xfb\xe1\xad\x89\x52\xa4\xe8\x4e\x8f\x24\xd6\x4c\xf2\x14\xdd\xd8\xe5\xdf\xef\xcf\xb2\x17\xd2\x4b\xbc\x5d\x19\x2a\x68\x12\xa5\xc0\x5e\xc3\x8a\x85\xdd\xfc\xc7\xa3\x7e\x8d\xef\x54\xc1\xc7\x19\xeb\x86\x2f\xb0\xc1\xe4\xac\xf0\x14\x93\xa3\xe4\x50\x4b\x51\xdd\x47\x84\xde\xb7\xbe\x68\x93\xfa\x0f\x00\x00\xff\xff\xe0\x00\x48\x13\x40\x01\x00\x00")

func assets_layout_html_bytes() ([]byte, error) {
	return bindata_read(
		_assets_layout_html,
		"assets/layout.html",
	)
}

func assets_layout_html() (*asset, error) {
	bytes, err := assets_layout_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "assets/layout.html", size: 320, mode: os.FileMode(420), modTime: time.Unix(1425963647, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"assets/css/style.css": assets_css_style_css,
	"assets/js/script.js": assets_js_script_js,
	"assets/layout.html": assets_layout_html,
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
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"assets": &_bintree_t{nil, map[string]*_bintree_t{
		"css": &_bintree_t{nil, map[string]*_bintree_t{
			"style.css": &_bintree_t{assets_css_style_css, map[string]*_bintree_t{
			}},
		}},
		"js": &_bintree_t{nil, map[string]*_bintree_t{
			"script.js": &_bintree_t{assets_js_script_js, map[string]*_bintree_t{
			}},
		}},
		"layout.html": &_bintree_t{assets_layout_html, map[string]*_bintree_t{
		}},
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

