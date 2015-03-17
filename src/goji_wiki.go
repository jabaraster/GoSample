package main

import (
    "github.com/zenazn/goji"

    "./assets"
    "./env"
    "./wiki"
)

func main() {
    env.Dump()

    goji.Get("/wiki", wiki.SampleView)
    goji.Post("/wiki/:title", wiki.SaveWiki)

    goji.Get("/css/*", assets.CssHandler)
    goji.Get("/js/*", assets.AssetsHandler)

    goji.Serve()
}
