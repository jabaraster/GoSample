package wiki

import (
    "fmt"
    "os"
    "time"
    "net/http"
    "encoding/json"

    _ "github.com/mattn/go-sqlite3"
    "github.com/naoina/genmai"
    "github.com/zenazn/goji/web"

    "../webutil"
)

var (
    db *genmai.DB
)

type Wiki struct {
    Id   int64  `db:"pk" json:"id"`
    Title string `db:"unique" json:"title"`
    Body string `json:"title"`
    Created time.Time `json:"created"`
    Updated time.Time `json:"updated"`
}

func (wiki *Wiki) BeforeInsert() error {
    n := time.Now()
    wiki.Created = n
    wiki.Updated = n
    return nil
}

func (wiki *Wiki) BeforeUpdate() error {
    n := time.Now()
    wiki.Updated = n
    return nil
}

func init() {
    db, err := genmai.New(&genmai.SQLite3Dialect{}, ":memory:")
    if err != nil {
        panic(err)
    }
    db.SetLogOutput(os.Stdout)
    if err := db.CreateTableIfNotExists(&Wiki{}); err != nil {
        panic(err)
    }
}

func Cnv(wiki *Wiki) (string, error) {
    _, e := db.Insert(wiki)
    if e != nil {
        panic(e)
    }
    b, err := json.Marshal(wiki)
    return string(b), err
}

func SampleView(w http.ResponseWriter, r *http.Request) {
    webutil.WriteWithTemplate(w, r, "html/wiki/sample.html", nil)
}

func SaveWiki(c web.C, w http.ResponseWriter, r *http.Request) {
    title := r.Form.Get("Title")
    body := r.Form.Get("Body")
    fmt.Printf("Title: %s, Body: %s\n", title, body)

    data := &Wiki{Title: title, Body: body}
    db.Insert(data)
    outjson, err := json.Marshal(data)
    if err != nil {
        fmt.Println(err)
        panic(err)
    }
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, string(outjson))
}

