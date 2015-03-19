package wiki

import (
    "os"
    "time"
    "net/http"

    _ "github.com/mattn/go-sqlite3"
    "github.com/naoina/genmai"
    "github.com/zenazn/goji/web"

    "../env"
    "../webutil"
)

var (
    db *genmai.DB
)

type Page struct {
    Title string
}

type Wiki struct {
    Id   int64  `db:"pk" json:"id"`
    Title string `db:"unique" json:"title"`
    Body string `json:"body"`
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
    var err error
    db, err = createDb()
    if err != nil {
        panic(err)
    }
    db.SetLogOutput(os.Stdout)
    if err := db.CreateTableIfNotExists(&Wiki{}); err != nil {
        panic(err)
    }
    db.Insert(&Wiki{ Title: "サンプルページ", Body: "簡単なWikiアプリを作ってみました。これはサンプル用のページです。ページ間のリンクなんかは実装していないので、ほんとに簡単なアプリです。" })
}

func createDb() (*genmai.DB, error) {
    switch env.DbKind() {
        case env.DbKindSQLite3:
            return genmai.New(&genmai.SQLite3Dialect{}, "./wiki.db")
        case env.DbKindPostgreSQL:
            return genmai.New(&genmai.PostgresDialect{}, "host=" + env.PostgresHost() + " dbname=" + env.PostgresDbName() + " user=" + env.PostgresUser() + " password=" + env.PostgresPassword())
    }
    panic(env.DbKind())
}

func New(w http.ResponseWriter, r *http.Request) {
    webutil.WriteTemplateResponse(w, r, "html/wiki/new.html", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
    webutil.WriteTemplateResponse(w, r, "html/wiki/index.html", nil)
}

func IndexData(w http.ResponseWriter, r *http.Request) {
    var result []Wiki
    err := db.Select(&result)
    if  err != nil {
        panic(err)
    }
    webutil.WriteJsonResponse(w, result)
}

func View(c web.C, w http.ResponseWriter, r *http.Request) {
    wiki := getByTitle(c)
    if wiki == nil {
        http.NotFound(w, r)
        return
    }
    webutil.WriteTemplateResponse(w, r, "html/wiki/view.html", wiki)
}

func Edit(c web.C, w http.ResponseWriter, r *http.Request) {
    webutil.WriteTemplateResponse(w, r, "html/wiki/edit.html", &Page{c.URLParams["title"]})
}

func EditData(c web.C, w http.ResponseWriter, r *http.Request) {
    wiki := getByTitle(c)
    if wiki == nil {
        http.NotFound(w ,r)
        return
    }
    webutil.WriteJsonResponse(w, wiki)
}

func Delete(c web.C, w http.ResponseWriter, r *http.Request) {
    wiki := getByTitle(c)
    if wiki == nil {
        http.NotFound(w, r)
        return
    }
    _, err := db.Delete(wiki)
    if err != nil {
        panic(err);
    }
    http.Redirect(w, r, "/wiki/", http.StatusSeeOther)
}

func Save(c web.C, w http.ResponseWriter, r *http.Request) {
    wiki := getByTitle(c)
    if wiki == nil {
        title := c.URLParams["title"]
        body := r.FormValue("body")

        data := Wiki{Title: title, Body: body}
        _, err := db.Insert(&data)
        if err != nil {
            panic(err)
        }
        err2 := webutil.WriteJsonResponse(w, data)
        if err2 != nil {
            panic(err2)
        }
    } else {
        wiki.Body = r.FormValue("body")
        _, err := db.Update(wiki)
        if err != nil {
            panic(err)
        }
    }
}

func getByTitle(c web.C) *Wiki {
    var result []Wiki
    err := db.Select(&result, db.Where("title", "=", c.URLParams["title"]))
    if err != nil {
        panic(err)
    }
    switch (len(result)) {
        case 0:
            return nil
        case 1:
            return &result[0]
    }
    panic(result)
}