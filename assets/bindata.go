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

var _assets_js_script_js = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\xcd\x41\x0a\xc2\x30\x10\x85\xe1\x7d\x4e\x31\x94\x40\x67\x36\x5e\x20\x78\x08\x8f\x10\x43\xc4\xd1\x90\x42\xf2\x22\xa8\xe4\xee\xa6\x5d\x88\xd2\xb7\x19\x18\x3e\xf8\xf9\xd2\x72\x80\x2e\x99\xad\xd0\xdb\x4c\xad\x46\xaa\x28\x1a\x30\x39\x63\x1e\xbe\x90\x66\x85\xfa\xa4\xaf\x48\x47\xfa\xea\x15\xd3\x98\xe5\xf9\xdc\x80\x25\xcf\x72\x08\x49\xc3\x9d\x77\x64\x9d\x4f\xb1\x80\x71\xd5\x2a\x6e\x7b\xf6\x71\xfb\x28\xd8\x3f\xff\x13\x63\x71\x1b\x32\x5d\xf8\x76\x6a\xb1\x3c\xc5\x7d\x02\x00\x00\xff\xff\xab\x71\x5c\x65\xad\x00\x00\x00")

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

	info := bindata_file_info{name: "assets/js/script.js", size: 173, mode: os.FileMode(420), modTime: time.Unix(1425964130, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _assets_layout_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x52\x41\xae\x9b\x30\x10\xdd\xf7\x14\xfe\xde\x63\x8b\x64\xd3\x85\x61\xd3\x4d\x0f\xd0\x0b\x18\x7b\x12\x9c\x00\xa6\xf6\x50\x05\x45\x48\x55\x72\x93\xaa\xdb\x4a\xbd\x13\x17\xa9\x0d\x4d\x20\xb4\xdf\x12\x62\x3c\xf3\xe6\xcd\x7b\x23\x8b\x37\x6d\x15\xf6\x2d\x90\x12\xeb\x2a\xff\x20\xe6\x1f\x09\x47\x94\x20\xf5\x1c\x4e\xd7\x1a\x50\x12\x55\x4a\xe7\x01\x33\xda\xe1\x21\xf9\x48\xf9\xaa\x5e\x99\xe6\x4c\x4a\x07\x87\x8c\x72\x5e\xcb\x8b\xd2\x0d\x2b\xac\x45\x8f\x4e\xb6\xf1\xa2\x6c\xcd\x9f\x09\xbe\x67\x7b\xb6\xe3\xca\xfb\x25\xc7\x6a\x13\x50\xde\x53\xe2\xa0\xca\xa8\xc7\xbe\x02\x5f\x02\x20\xdd\x8e\xd9\xd6\x1f\x73\x23\xdd\x94\x9e\x69\xa2\xb1\x8c\x22\x5c\x30\x56\xd6\x2c\x68\xb0\x82\xfc\x7a\x65\x5f\x62\x30\x0c\x82\xcf\x99\x05\xe1\x95\x33\x2d\x12\xef\x54\xf4\xa3\xac\x06\x76\xfa\xda\x81\xeb\x27\x1f\x73\x98\xa4\x2c\x4d\xd9\x6e\xd2\x7d\x0a\xfc\x82\xcf\x5d\xef\xd1\x9c\xfc\x5f\xc0\x7f\xd0\x82\x2f\xfb\x16\x85\xd5\xfd\x8a\x44\x9b\x6f\x44\x55\xd2\xfb\x8c\x2a\xdb\xa0\x34\x0d\xb8\x95\x9b\x78\x3e\xd9\xae\x41\x12\x0c\x4d\xc1\x30\x90\x24\x27\x8b\x3d\xf2\xf6\x02\x16\x65\x9a\x7f\xb6\x47\x08\x33\xd3\x57\x9a\xf5\xa8\xc2\x3a\x0d\x4e\x47\xa1\x21\xbb\xc1\x15\x1d\xa2\x6d\x9e\x50\x6c\x48\xf8\x92\xd6\x99\x5a\xba\x7e\xa3\x6d\x6a\x30\x0f\xec\xb1\xea\xdb\xd2\x04\x1f\xe4\x19\x25\xf6\x1c\xa7\x98\x7f\xdb\xc6\xfb\xf7\xf1\xf6\x73\xbc\xfd\x18\xef\xbf\xc6\xdb\xef\x57\x0d\x7c\x16\xb1\x5a\xd4\x22\x34\x14\xa7\x1d\x06\x8b\xf1\x45\xff\x09\x00\x00\xff\xff\x65\x0d\xbd\xc1\xe8\x02\x00\x00")

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

	info := bindata_file_info{name: "assets/layout.html", size: 744, mode: os.FileMode(420), modTime: time.Unix(1425964085, 0)}
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

