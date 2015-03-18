package wiki

import (
    "testing"
    "fmt"
)

func TestCnv(t *testing.T) {
    data := &Wiki{Title:"Title",Body:"Body"}
    fmt.Println(data)
    ary := []string{"a",}
    fmt.Println(len(ary))
}