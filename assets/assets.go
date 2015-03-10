package assets
import (
    "html/template"
    "net/http"
    "time"
)

//var ns = template.New("complex")

func ParseAsset(path string) (*template.Template, error) {
    src, err := getData(path)
    if err != nil {
        return nil, err
    }
    return template.New(time.Now().String()).Parse(string(src))
}

func CssHandler(w http.ResponseWriter, r *http.Request) {
    data, err := getData(r.URL.Path[1:])
    if err != nil {
        panic(err)
    }
    w.Header().Add("content-type", "text/css")
    _, err2 := w.Write(data)
    if err2 != nil {
        panic(err2)
    }
}

func AssetsHandler(w http.ResponseWriter, r *http.Request) {
    data, err := getData(r.URL.Path[1:])
    if err != nil {
        panic(err)
    }
    _, err2 := w.Write(data)
    if err2 != nil {
        panic(err2)
    }
}

func getData(path string) ([]byte, error) {
    return Asset("assets/" + path)
}