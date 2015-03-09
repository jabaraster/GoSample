package static
import (
    "html/template"
)

var ns = template.New("complex")

func ParseAsset(path string) (*template.Template, error) {
    src, err := Asset(path)
    if err != nil {
        return nil, err
    }
    return ns.Parse(string(src))
}
