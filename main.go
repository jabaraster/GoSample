package main

import (
    "net/http"
    "./app_web"
    "./assets"
    "fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello!!!")
}

func main() {
//    http.HandleFunc("/", app_web.Handler)
    http.HandleFunc("/view", app_web.ViewHandler)
    http.HandleFunc("/chat", app_web.ChatHnaler)
    http.HandleFunc("/s", handler)

    http.HandleFunc("/css/", assets.CssHandler)
    http.HandleFunc("/js/", assets.AssetsHandler)

    http.ListenAndServe(":9090", nil)
}
