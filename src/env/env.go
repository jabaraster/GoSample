package env

import (
    "os"
    "strings"
    "fmt"
)

const (
    KeyMode = "WIKI_MODE"
    KeyDbKind = "WIKI_DB_KIND"
    KeyPostgresHost     = "WIKI_POSTGRES_HOST"
    KeyPostgresDbName   = "WIKI_POSTGRES_DB_NAME"
    KeyPostgresUser     = "WIKI_POSTGRES_USER"
    KeyPostgresPassword = "WIKI_POSTGRES_PASSWORD"

    ModeProduction = "production"
    ModeDebug      = "debug"

    DbKindSQLite3    = "SQLite3"
    DbKindPostgreSQL = "PostgreSQL"
)

var (
    mode   string
    dbKind string
    postgresHost     string
    postgresDbName   string
    postgresUser     string
    postgresPassword string
)

func init() {
    mode = getEnv(KeyMode, ModeDebug)
    dbKind = getEnv(KeyDbKind, DbKindSQLite3)

    postgresHost = getEnv(KeyPostgresHost, "")
    postgresDbName = getEnv(KeyPostgresDbName, "")
    postgresUser = getEnv(KeyPostgresUser, "")
    postgresPassword = getEnv(KeyPostgresPassword, "")
}

func Dump() {
    dump("mode", mode)
    dump("dbKind", dbKind)

    dump("postgresHost", postgresHost)
    dump("postgresDbName", postgresDbName)
    dump("postgresUser", postgresUser)
    dump("postgresPassword", postgresPassword)
}

func Mode() string {
    return mode
}

func IsProductionMode() bool {
    return strings.EqualFold(ModeProduction, mode)
}

func IsDebugMode() bool {
    return strings.EqualFold(ModeDebug, mode)
}

func DbKind() string {
    return dbKind
}

func PostgresHost() string {
    return postgresHost
}

func PostgresDbName() string {
    return postgresDbName
}

func PostgresUser() string {
    return postgresUser
}

func PostgresPassword() string {
    return postgresPassword
}

func dump(key string, val string) {
    fmt.Printf("%s -> %s\n", key, val)
}

func getEnv(key string, defaultValue string) string {
    ret := os.Getenv(key)
    if ret == "" {
        return defaultValue
    }
    return ret
}
