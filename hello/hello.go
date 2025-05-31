package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    message, _ := greetings.Hello("Gladys")
    fmt.Println(message)

    // Request a greeting
    messageError, err2 := greetings.Hello("")
    fmt.Println(messageError, err2)

    // If an error was returned, print it to the console and
    // exit the program.
    if err2 != nil {
        log.Fatal(err2)
    }
}