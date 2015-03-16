package assets
import (
    "html/template"
    "net/http"
    "time"
    "fmt"
    "../env"
    "../assets/debug"
)

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
        // panic(err) // テスト中はこれでもいいのだが・・・
        fmt.Println(err)
        http.NotFound(w, r)
        return
    }
    w.Header().Add("content-type", "text/css") // これ大事！
    _, err2 := w.Write(data)
    if err2 != nil {
        // panic(err) // テスト中はこれでもいいのだが・・・
        fmt.Println(err2)
        http.NotFound(w, r)
        return
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
    if env.IsProductionMode() {
        return Asset("assets/" + path)
    }
    return debug.Asset("assets/" + path)
}