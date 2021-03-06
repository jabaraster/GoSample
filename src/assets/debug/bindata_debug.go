package debug

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"path"
	"path/filepath"
)

// bindata_read reads the given file from disk. It returns an error on failure.
func bindata_read(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// assets_css_chat_css reads file data from disk. It returns an error on failure.
func assets_css_chat_css() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/css/chat.css"
	name := "assets/css/chat.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_common_bootstrap_min_css reads file data from disk. It returns an error on failure.
func assets_css_common_bootstrap_min_css() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/css/common/bootstrap.min.css"
	name := "assets/css/common/bootstrap.min.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_fonts_glyphicons_halflings_regular_eot reads file data from disk. It returns an error on failure.
func assets_css_fonts_glyphicons_halflings_regular_eot() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/css/fonts/glyphicons-halflings-regular.eot"
	name := "assets/css/fonts/glyphicons-halflings-regular.eot"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_fonts_glyphicons_halflings_regular_svg reads file data from disk. It returns an error on failure.
func assets_css_fonts_glyphicons_halflings_regular_svg() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/css/fonts/glyphicons-halflings-regular.svg"
	name := "assets/css/fonts/glyphicons-halflings-regular.svg"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_fonts_glyphicons_halflings_regular_ttf reads file data from disk. It returns an error on failure.
func assets_css_fonts_glyphicons_halflings_regular_ttf() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/css/fonts/glyphicons-halflings-regular.ttf"
	name := "assets/css/fonts/glyphicons-halflings-regular.ttf"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_fonts_glyphicons_halflings_regular_woff reads file data from disk. It returns an error on failure.
func assets_css_fonts_glyphicons_halflings_regular_woff() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/css/fonts/glyphicons-halflings-regular.woff"
	name := "assets/css/fonts/glyphicons-halflings-regular.woff"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_fonts_glyphicons_halflings_regular_woff2 reads file data from disk. It returns an error on failure.
func assets_css_fonts_glyphicons_halflings_regular_woff2() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/css/fonts/glyphicons-halflings-regular.woff2"
	name := "assets/css/fonts/glyphicons-halflings-regular.woff2"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_style_css reads file data from disk. It returns an error on failure.
func assets_css_style_css() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/css/style.css"
	name := "assets/css/style.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_wiki_css reads file data from disk. It returns an error on failure.
func assets_css_wiki_css() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/css/wiki.css"
	name := "assets/css/wiki.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_html_chat_html reads file data from disk. It returns an error on failure.
func assets_html_chat_html() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/html/chat.html"
	name := "assets/html/chat.html"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_html_echo_html reads file data from disk. It returns an error on failure.
func assets_html_echo_html() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/html/echo.html"
	name := "assets/html/echo.html"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_html_layout_html reads file data from disk. It returns an error on failure.
func assets_html_layout_html() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/html/layout.html"
	name := "assets/html/layout.html"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_html_wiki_edit_html reads file data from disk. It returns an error on failure.
func assets_html_wiki_edit_html() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/html/wiki/edit.html"
	name := "assets/html/wiki/edit.html"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_html_wiki_index_html reads file data from disk. It returns an error on failure.
func assets_html_wiki_index_html() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/html/wiki/index.html"
	name := "assets/html/wiki/index.html"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_html_wiki_new_html reads file data from disk. It returns an error on failure.
func assets_html_wiki_new_html() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/html/wiki/new.html"
	name := "assets/html/wiki/new.html"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_html_wiki_view_html reads file data from disk. It returns an error on failure.
func assets_html_wiki_view_html() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/html/wiki/view.html"
	name := "assets/html/wiki/view.html"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_chat_js reads file data from disk. It returns an error on failure.
func assets_js_chat_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/js/chat.js"
	name := "assets/js/chat.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_chat_min_js reads file data from disk. It returns an error on failure.
func assets_js_chat_min_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/js/chat.min.js"
	name := "assets/js/chat.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_common_jsxtransformer_js reads file data from disk. It returns an error on failure.
func assets_js_common_jsxtransformer_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/js/common/JSXTransformer.js"
	name := "assets/js/common/JSXTransformer.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_common_bootstrap_min_js reads file data from disk. It returns an error on failure.
func assets_js_common_bootstrap_min_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/js/common/bootstrap.min.js"
	name := "assets/js/common/bootstrap.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_common_jquery_min_js reads file data from disk. It returns an error on failure.
func assets_js_common_jquery_min_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/js/common/jquery.min.js"
	name := "assets/js/common/jquery.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_common_react_with_addons_js reads file data from disk. It returns an error on failure.
func assets_js_common_react_with_addons_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/js/common/react-with-addons.js"
	name := "assets/js/common/react-with-addons.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_common_react_with_addons_min_js reads file data from disk. It returns an error on failure.
func assets_js_common_react_with_addons_min_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/js/common/react-with-addons.min.js"
	name := "assets/js/common/react-with-addons.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_common_react_js reads file data from disk. It returns an error on failure.
func assets_js_common_react_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/js/common/react.js"
	name := "assets/js/common/react.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_common_react_min_js reads file data from disk. It returns an error on failure.
func assets_js_common_react_min_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/js/common/react.min.js"
	name := "assets/js/common/react.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_echo_js reads file data from disk. It returns an error on failure.
func assets_js_echo_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/js/echo.js"
	name := "assets/js/echo.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_echo_min_js reads file data from disk. It returns an error on failure.
func assets_js_echo_min_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/js/echo.min.js"
	name := "assets/js/echo.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_script_js reads file data from disk. It returns an error on failure.
func assets_js_script_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/js/script.js"
	name := "assets/js/script.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_script_min_js reads file data from disk. It returns an error on failure.
func assets_js_script_min_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/js/script.min.js"
	name := "assets/js/script.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_wiki_edit_js reads file data from disk. It returns an error on failure.
func assets_js_wiki_edit_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/js/wiki/edit.js"
	name := "assets/js/wiki/edit.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_wiki_edit_min_js reads file data from disk. It returns an error on failure.
func assets_js_wiki_edit_min_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/js/wiki/edit.min.js"
	name := "assets/js/wiki/edit.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_wiki_index_js reads file data from disk. It returns an error on failure.
func assets_js_wiki_index_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/js/wiki/index.js"
	name := "assets/js/wiki/index.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_wiki_index_min_js reads file data from disk. It returns an error on failure.
func assets_js_wiki_index_min_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/js/wiki/index.min.js"
	name := "assets/js/wiki/index.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_wiki_new_js reads file data from disk. It returns an error on failure.
func assets_js_wiki_new_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/js/wiki/new.js"
	name := "assets/js/wiki/new.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_wiki_new_min_js reads file data from disk. It returns an error on failure.
func assets_js_wiki_new_min_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/js/wiki/new.min.js"
	name := "assets/js/wiki/new.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_jsx_chat_js reads file data from disk. It returns an error on failure.
func assets_jsx_chat_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/jsx/chat.js"
	name := "assets/jsx/chat.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_jsx_echo_js reads file data from disk. It returns an error on failure.
func assets_jsx_echo_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/jsx/echo.js"
	name := "assets/jsx/echo.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_jsx_script_js reads file data from disk. It returns an error on failure.
func assets_jsx_script_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/jsx/script.js"
	name := "assets/jsx/script.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_jsx_wiki_edit_js reads file data from disk. It returns an error on failure.
func assets_jsx_wiki_edit_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/jsx/wiki/edit.js"
	name := "assets/jsx/wiki/edit.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_jsx_wiki_index_js reads file data from disk. It returns an error on failure.
func assets_jsx_wiki_index_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/jsx/wiki/index.js"
	name := "assets/jsx/wiki/index.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_jsx_wiki_new_js reads file data from disk. It returns an error on failure.
func assets_jsx_wiki_new_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/GoSample/assets/jsx/wiki/new.js"
	name := "assets/jsx/wiki/new.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
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
	"assets/css/chat.css": assets_css_chat_css,
	"assets/css/common/bootstrap.min.css": assets_css_common_bootstrap_min_css,
	"assets/css/fonts/glyphicons-halflings-regular.eot": assets_css_fonts_glyphicons_halflings_regular_eot,
	"assets/css/fonts/glyphicons-halflings-regular.svg": assets_css_fonts_glyphicons_halflings_regular_svg,
	"assets/css/fonts/glyphicons-halflings-regular.ttf": assets_css_fonts_glyphicons_halflings_regular_ttf,
	"assets/css/fonts/glyphicons-halflings-regular.woff": assets_css_fonts_glyphicons_halflings_regular_woff,
	"assets/css/fonts/glyphicons-halflings-regular.woff2": assets_css_fonts_glyphicons_halflings_regular_woff2,
	"assets/css/style.css": assets_css_style_css,
	"assets/css/wiki.css": assets_css_wiki_css,
	"assets/html/chat.html": assets_html_chat_html,
	"assets/html/echo.html": assets_html_echo_html,
	"assets/html/layout.html": assets_html_layout_html,
	"assets/html/wiki/edit.html": assets_html_wiki_edit_html,
	"assets/html/wiki/index.html": assets_html_wiki_index_html,
	"assets/html/wiki/new.html": assets_html_wiki_new_html,
	"assets/html/wiki/view.html": assets_html_wiki_view_html,
	"assets/js/chat.js": assets_js_chat_js,
	"assets/js/chat.min.js": assets_js_chat_min_js,
	"assets/js/common/JSXTransformer.js": assets_js_common_jsxtransformer_js,
	"assets/js/common/bootstrap.min.js": assets_js_common_bootstrap_min_js,
	"assets/js/common/jquery.min.js": assets_js_common_jquery_min_js,
	"assets/js/common/react-with-addons.js": assets_js_common_react_with_addons_js,
	"assets/js/common/react-with-addons.min.js": assets_js_common_react_with_addons_min_js,
	"assets/js/common/react.js": assets_js_common_react_js,
	"assets/js/common/react.min.js": assets_js_common_react_min_js,
	"assets/js/echo.js": assets_js_echo_js,
	"assets/js/echo.min.js": assets_js_echo_min_js,
	"assets/js/script.js": assets_js_script_js,
	"assets/js/script.min.js": assets_js_script_min_js,
	"assets/js/wiki/edit.js": assets_js_wiki_edit_js,
	"assets/js/wiki/edit.min.js": assets_js_wiki_edit_min_js,
	"assets/js/wiki/index.js": assets_js_wiki_index_js,
	"assets/js/wiki/index.min.js": assets_js_wiki_index_min_js,
	"assets/js/wiki/new.js": assets_js_wiki_new_js,
	"assets/js/wiki/new.min.js": assets_js_wiki_new_min_js,
	"assets/jsx/chat.js": assets_jsx_chat_js,
	"assets/jsx/echo.js": assets_jsx_echo_js,
	"assets/jsx/script.js": assets_jsx_script_js,
	"assets/jsx/wiki/edit.js": assets_jsx_wiki_edit_js,
	"assets/jsx/wiki/index.js": assets_jsx_wiki_index_js,
	"assets/jsx/wiki/new.js": assets_jsx_wiki_new_js,
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
			"chat.css": &_bintree_t{assets_css_chat_css, map[string]*_bintree_t{
			}},
			"common": &_bintree_t{nil, map[string]*_bintree_t{
				"bootstrap.min.css": &_bintree_t{assets_css_common_bootstrap_min_css, map[string]*_bintree_t{
				}},
			}},
			"fonts": &_bintree_t{nil, map[string]*_bintree_t{
				"glyphicons-halflings-regular.eot": &_bintree_t{assets_css_fonts_glyphicons_halflings_regular_eot, map[string]*_bintree_t{
				}},
				"glyphicons-halflings-regular.svg": &_bintree_t{assets_css_fonts_glyphicons_halflings_regular_svg, map[string]*_bintree_t{
				}},
				"glyphicons-halflings-regular.ttf": &_bintree_t{assets_css_fonts_glyphicons_halflings_regular_ttf, map[string]*_bintree_t{
				}},
				"glyphicons-halflings-regular.woff": &_bintree_t{assets_css_fonts_glyphicons_halflings_regular_woff, map[string]*_bintree_t{
				}},
				"glyphicons-halflings-regular.woff2": &_bintree_t{assets_css_fonts_glyphicons_halflings_regular_woff2, map[string]*_bintree_t{
				}},
			}},
			"style.css": &_bintree_t{assets_css_style_css, map[string]*_bintree_t{
			}},
			"wiki.css": &_bintree_t{assets_css_wiki_css, map[string]*_bintree_t{
			}},
		}},
		"html": &_bintree_t{nil, map[string]*_bintree_t{
			"chat.html": &_bintree_t{assets_html_chat_html, map[string]*_bintree_t{
			}},
			"echo.html": &_bintree_t{assets_html_echo_html, map[string]*_bintree_t{
			}},
			"layout.html": &_bintree_t{assets_html_layout_html, map[string]*_bintree_t{
			}},
			"wiki": &_bintree_t{nil, map[string]*_bintree_t{
				"edit.html": &_bintree_t{assets_html_wiki_edit_html, map[string]*_bintree_t{
				}},
				"index.html": &_bintree_t{assets_html_wiki_index_html, map[string]*_bintree_t{
				}},
				"new.html": &_bintree_t{assets_html_wiki_new_html, map[string]*_bintree_t{
				}},
				"view.html": &_bintree_t{assets_html_wiki_view_html, map[string]*_bintree_t{
				}},
			}},
		}},
		"js": &_bintree_t{nil, map[string]*_bintree_t{
			"chat.js": &_bintree_t{assets_js_chat_js, map[string]*_bintree_t{
			}},
			"chat.min.js": &_bintree_t{assets_js_chat_min_js, map[string]*_bintree_t{
			}},
			"common": &_bintree_t{nil, map[string]*_bintree_t{
				"JSXTransformer.js": &_bintree_t{assets_js_common_jsxtransformer_js, map[string]*_bintree_t{
				}},
				"bootstrap.min.js": &_bintree_t{assets_js_common_bootstrap_min_js, map[string]*_bintree_t{
				}},
				"jquery.min.js": &_bintree_t{assets_js_common_jquery_min_js, map[string]*_bintree_t{
				}},
				"react-with-addons.js": &_bintree_t{assets_js_common_react_with_addons_js, map[string]*_bintree_t{
				}},
				"react-with-addons.min.js": &_bintree_t{assets_js_common_react_with_addons_min_js, map[string]*_bintree_t{
				}},
				"react.js": &_bintree_t{assets_js_common_react_js, map[string]*_bintree_t{
				}},
				"react.min.js": &_bintree_t{assets_js_common_react_min_js, map[string]*_bintree_t{
				}},
			}},
			"echo.js": &_bintree_t{assets_js_echo_js, map[string]*_bintree_t{
			}},
			"echo.min.js": &_bintree_t{assets_js_echo_min_js, map[string]*_bintree_t{
			}},
			"script.js": &_bintree_t{assets_js_script_js, map[string]*_bintree_t{
			}},
			"script.min.js": &_bintree_t{assets_js_script_min_js, map[string]*_bintree_t{
			}},
			"wiki": &_bintree_t{nil, map[string]*_bintree_t{
				"edit.js": &_bintree_t{assets_js_wiki_edit_js, map[string]*_bintree_t{
				}},
				"edit.min.js": &_bintree_t{assets_js_wiki_edit_min_js, map[string]*_bintree_t{
				}},
				"index.js": &_bintree_t{assets_js_wiki_index_js, map[string]*_bintree_t{
				}},
				"index.min.js": &_bintree_t{assets_js_wiki_index_min_js, map[string]*_bintree_t{
				}},
				"new.js": &_bintree_t{assets_js_wiki_new_js, map[string]*_bintree_t{
				}},
				"new.min.js": &_bintree_t{assets_js_wiki_new_min_js, map[string]*_bintree_t{
				}},
			}},
		}},
		"jsx": &_bintree_t{nil, map[string]*_bintree_t{
			"chat.js": &_bintree_t{assets_jsx_chat_js, map[string]*_bintree_t{
			}},
			"echo.js": &_bintree_t{assets_jsx_echo_js, map[string]*_bintree_t{
			}},
			"script.js": &_bintree_t{assets_jsx_script_js, map[string]*_bintree_t{
			}},
			"wiki": &_bintree_t{nil, map[string]*_bintree_t{
				"edit.js": &_bintree_t{assets_jsx_wiki_edit_js, map[string]*_bintree_t{
				}},
				"index.js": &_bintree_t{assets_jsx_wiki_index_js, map[string]*_bintree_t{
				}},
				"new.js": &_bintree_t{assets_jsx_wiki_new_js, map[string]*_bintree_t{
				}},
			}},
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

