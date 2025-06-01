package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// // Set properties of the predefined Logger, including
	// // the log entry prefix and a flag to disable printing
	// // the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}
	// Request greeting messages for the names.
	messages2, err := greetings.Hellos(names)

	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages2)

	// Get a greeting message and print it.
	message, err := greetings.Hello("Gladys")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
	// Request a greeting
	// messageError, err2 := greetings.Hello("")
	// fmt.Println(messageError, err2)

	// If an error was returned, print it to the console and
	// exit the program.
	// if err2 != nil {
	// 	log.Fatal(err2)
	// }
}
