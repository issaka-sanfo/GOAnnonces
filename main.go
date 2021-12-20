package main // The main package


// Import the necessary packages
import (
    "fmt" // for printing messages and text
    "log" // for logging errors and printing messaging
    "net/http" // a Go HTTP package for handling HTTP requesting when creating GO APIs
	"goanonnonces.leboncoin.fr/restapi" // Import the restapi package to get APIs
    "goanonnonces.leboncoin.fr/sqlschema"
)

// Main function
func main() {
    sqlschema.CreateTables()
	restapi.Route()
    // serve the app
    fmt.Println("Server at 80")
    log.Fatal(http.ListenAndServe(":80", restapi.Router))
}