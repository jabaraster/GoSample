package env

import (
    "flag"
    "strings"
    "fmt"
)

var (
    mode string
)

func init() {
    flag.StringVar(&mode, "mode", "debug", "debug or production")
    if flag.Parsed() {
        flag.Parse()
    }
}

func Dump() {
    fmt.Printf("mode -> %s\n", mode)
}

func Mode() (string) {
    return mode
}

func IsProductionMode() bool {
    return strings.EqualFold("production", mode)
}

func IsDebugMode() bool {
    return strings.EqualFold("debug", mode)
}