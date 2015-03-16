package main

import (
    "github.com/zenazn/goji"

    "./app_web"
    "./assets"
    "./env"
)

func main() {
    env.Dump()
    goji.Get("/wiki", app_web.ViewHandler)

    goji.Get("/css/*", assets.CssHandler)
    goji.Get("/js/*", assets.AssetsHandler)

    goji.Serve()
}