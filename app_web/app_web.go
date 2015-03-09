package app_web
import (
    "net/http"
    "fmt"
    "../static"
)

type PageData struct {
    Title string
    Count int
}

func Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World! ->", r.URL.Path[0:])
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := static.ParseAsset("html/layout.html")
    if err != nil {
        panic(err)
    }
    pd := PageData{"タイトルやで", 10}
    err = tmpl.Execute(w, pd)
    if err != nil {
        panic(err)
    }
}