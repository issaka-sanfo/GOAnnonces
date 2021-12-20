package messages

import "fmt"

// Function for handling messages
func PrintMessage(message string) {
    fmt.Println("")
    fmt.Println(message)
    fmt.Println("")
}

// Function for handling errors
func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}