package main

// ---------------------------------------------------------
// EXERCISE: Try the scopes
//
//  1. Create: printer.go
//
//  2. In printer.go:
//     1. Create a function named hello
//     2. Inside the hello function, print your name
//
//  4. In printer.go:
//     1. Call the bye function from
//        inside the hello function
// ---------------------------------------------------------

import (
	"fmt"
)

func hello() {
	fmt.Println("cyber33")
	bye()
}
