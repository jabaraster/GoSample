package app_web
import (
    "net/http"
    "fmt"
    "../assets"
    "github.com/zenazn/goji/web"
)

type PageData struct {
    Title string
    Count int
}

func Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World! ->", r.URL.Path[0:])
}

func ViewHandler2(c web.C, w http.ResponseWriter, r *http.Request) {
    ViewHandler(w, r)
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := assets.ParseAsset("html/layout.html")
    if err != nil {
        panic(err)
    }
    pd := PageData{"タイトルやで", 10}
    err = tmpl.Execute(w, pd)
    if err != nil {
        panic(err)
    }
}

func ChatHnaler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := assets.ParseAsset("html/chat.html")
    if err != nil {
        panic(err)
    }
    err = tmpl.Execute(w, nil)
    if err != nil {
        panic(err)
    }
}