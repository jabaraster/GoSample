package static

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

var _html_layout_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x8f\x41\xce\x83\x20\x10\x85\xf7\xff\x29\xe6\x9f\xbd\xb2\xed\x02\xdc\xf4\x0a\xbd\x80\xd5\x31\x98\xa2\x36\xf2\x4c\x6a\x0c\x77\x2f\x42\x13\x64\xc3\x63\xe0\xfb\xf2\xd0\xff\xfd\xd2\x61\x7f\x0b\x59\x4c\xae\xf9\xd3\x79\xa3\xb8\xb4\x95\xb6\xcf\x31\x1d\x27\x41\x4b\x9d\x6d\x57\x2f\x30\xbc\x61\xa8\x6e\xac\x2e\xf7\x6e\x9c\x5f\xb4\x8a\x33\xec\xb1\x3b\xf1\x56\x04\x4c\xa7\xdb\x30\xe4\x03\xd5\x79\xcf\x64\x57\x19\x0c\xc7\xa8\xd2\xab\x3a\x0d\x2f\x16\x8c\x70\xd2\x1c\x47\xfd\x38\x43\x08\x5a\xe5\x49\xae\xa4\x4a\x27\xfd\x5c\xfa\xbd\x80\xf7\x65\x9b\x41\x91\x4b\x21\x04\xaa\x1a\x2a\x96\x1f\x9d\x91\x68\x39\x3f\xf9\x0d\x00\x00\xff\xff\x7c\xeb\x4d\x92\xfb\x00\x00\x00")

func html_layout_html_bytes() ([]byte, error) {
	return bindata_read(
		_html_layout_html,
		"html/layout.html",
	)
}

func html_layout_html() (*asset, error) {
	bytes, err := html_layout_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "html/layout.html", size: 251, mode: os.FileMode(420), modTime: time.Unix(1425938629, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _css_style_css = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\xca\x4f\xa9\xac\x4e\x4a\x4c\xce\x4e\x2f\xca\x2f\xcd\x4b\xb1\x52\x4e\x4b\x33\x33\x31\x33\xa9\x05\x04\x00\x00\xff\xff\x2a\xb4\x39\x75\x18\x00\x00\x00")

func css_style_css_bytes() ([]byte, error) {
	return bindata_read(
		_css_style_css,
		"css/style.css",
	)
}

func css_style_css() (*asset, error) {
	bytes, err := css_style_css_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "css/style.css", size: 24, mode: os.FileMode(420), modTime: time.Unix(1425939990, 0)}
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
	"html/layout.html": html_layout_html,
	"css/style.css": css_style_css,
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
	"css": &_bintree_t{nil, map[string]*_bintree_t{
		"style.css": &_bintree_t{css_style_css, map[string]*_bintree_t{
		}},
	}},
	"html": &_bintree_t{nil, map[string]*_bintree_t{
		"layout.html": &_bintree_t{html_layout_html, map[string]*_bintree_t{
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

