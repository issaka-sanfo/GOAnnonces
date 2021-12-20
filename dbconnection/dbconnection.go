package dbconnection

import (
	"database/sql"
    "fmt"
    "goanonnonces.leboncoin.fr/messages"
    _ "github.com/lib/pq" // for postgresql
)

// DB properties
const (
    DB_USER     = "newuser"
    DB_PASSWORD = "newpass" //////////////////////////////////////////////////////******************Change thePASSWORD
    DB_NAME     = "annonces"
)

// DB set up
func DBsetup() *sql.DB {
    dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)

    messages.CheckError(err)

    return db
}