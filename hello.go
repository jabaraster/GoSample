package main

import (
    "net/http"
    "./app_web"
    "fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello!!!")
}

func main() {
    http.HandleFunc("/", app_web.Handler)
    http.HandleFunc("/view", app_web.ViewHandler)
    http.HandleFunc("/s", handler)
    http.ListenAndServe(":9090", nil)
}
