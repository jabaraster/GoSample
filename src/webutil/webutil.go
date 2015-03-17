package webutil

import (
    "net/http"
    "../assets"
)

func WriteWithTemplate(w http.ResponseWriter, r *http.Request, templatePath string,  data interface{}) {
    tmpl, err := assets.ParseAsset(templatePath)
    if err != nil {
        http.NotFound(w, r)
        return
    }
    err = tmpl.Execute(w, data)
    if err != nil {
        panic(err)
    }
}
