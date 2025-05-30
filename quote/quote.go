package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	// Get the Go quote
	fmt.Println(quote.Go())
	
	// Get other quotes
	fmt.Println(quote.Glass())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
}