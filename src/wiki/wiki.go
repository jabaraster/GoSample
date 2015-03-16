package wiki

import (
    "os"
    "time"
    _ "github.com/mattn/go-sqlite3"
    "github.com/naoina/genmai"
)

var (
    db *genmai.DB
)

type Wiki struct {
    Id   int64  `db:"pk"`
    Title string `db:"unique"`
    Body string `default:""`
    Created time.Time
    Updated time.Time
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

func Init() {
    db, err := genmai.New(&genmai.SQLite3Dialect{}, ":memory:")
    if err != nil {
        panic(err)
    }
    db.SetLogOutput(os.Stdout)
    if err := db.CreateTableIfNotExists(&Wiki{}); err != nil {
        panic(err)
    }
    db.Insert(&Wiki{Title: "Title", Body: "Body"})
}