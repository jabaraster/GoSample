package env

import (
    "flag"
    "strings"
    "fmt"
)

var (
    mode string
    webPort int
)

func init() {
    flag.StringVar(&mode, "mode", "debug", "debug or production")
    flag.IntVar(&webPort, "webPort", 8081, "http port")
    flag.Parse()
}

func Dump() {
    fmt.Printf("mode -> %s\n", mode)
    fmt.Printf("webPort -> %d\n", webPort)
}

func Mode() (string) {
    return mode
}

func IsProductionMode() (bool) {
    return strings.EqualFold("production", mode)
}

func WebPort() (int) {
    return webPort
}

