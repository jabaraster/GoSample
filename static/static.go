package static
import (
    "html/template"
    "net/http"
    "fmt"
)

var ns = template.New("complex")

func ParseAsset(path string) (*template.Template, error) {
    src, err := Asset(path)
    if err != nil {
        return nil, err
    }
    return ns.Parse(string(src))
}


func CssHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf(r.URL.Path[1:])
    data, err := Asset(r.URL.Path[1:])
    if err != nil {
        panic(err)
    }
    _, err2 := w.Write(data)
    if err2 != nil {
        panic(err2)
    }
}