package main

import (
    "github.com/zenazn/goji"

    "./assets"
    "./env"
    "./wiki"
    "net/http"
)

func main() {
    env.Dump()

    goji.Get("/wiki", http.RedirectHandler("/wiki/", http.StatusMovedPermanently))
    goji.Get("/wiki/", wiki.Index)
    goji.Get("/wiki/index", wiki.Index)
    goji.Get("/wiki/index/data", wiki.IndexData)
    goji.Get("/wiki/new", wiki.New)
    goji.Get("/wiki/:title", wiki.View)
    goji.Get("/wiki/:title/edit", wiki.Edit)
    goji.Get("/wiki/:title/edit/data", wiki.EditData)
    goji.Post("/wiki/:title/delete", wiki.Delete)
    goji.Post("/wiki/:title", wiki.Save)

    goji.Get("/css/*", assets.CssHandler)
    goji.Get("/js/*", assets.AssetsHandler)

    goji.Serve()
}
