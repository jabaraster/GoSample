package wiki

import (
    "testing"
    "fmt"
)

func TestCnv(t *testing.T) {
    s, e := Cnv(&Wiki{Title:"Title",Body:"Body"})
    if e != nil {
        fmt.Errorf("%s\n", e)
        t.Fail()
    }
    fmt.Println(s)
}