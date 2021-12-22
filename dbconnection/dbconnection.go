package dbconnection

import (
	"database/sql"
    "fmt"
    _ "github.com/lib/pq" // for postgresql
    "log"
	"os"
)

// DB properties
var host = os.Getenv("pgHost")
var port = os.Getenv("pgPort")
var user = os.Getenv("pgUser")
var password = os.Getenv("pgPassword")
var dbname = os.Getenv("pgDbName")



// DB set up
// func DBsetup() *sql.DB {
//     dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_USER, DB_PORT, DB_PASSWORD, DB_NAME)
//     db, err := sql.Open("postgres", dbinfo)

//     messages.CheckError(err)

//     return db
// }

func DBsetup() (*sql.DB) {
	var err error
	dbInformation := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",  host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbInformation)
	if err != nil {
		log.Fatal("This is the error: ", err)
		fmt.Printf("Cannot connect to %s database", dbInformation)
		return nil
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	log.Println("Database Connection established")
	return db
}