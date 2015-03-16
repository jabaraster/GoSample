package main

import (
    "net/http"
    "./app_web"
    "./assets"
    "fmt"
    "./env"
    "github.com/zenazn/goji/bind"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello!!!")
}

func main() {
    env.Dump()

    http.HandleFunc("/view", app_web.ViewHandler)
    http.HandleFunc("/chat", app_web.ChatHnaler)
    http.HandleFunc("/s", handler)

    http.HandleFunc("/css/", assets.CssHandler)
    http.HandleFunc("/js/", assets.AssetsHandler)

    http.ListenAndServe(bind.Default().Addr().Network(), nil)
}
