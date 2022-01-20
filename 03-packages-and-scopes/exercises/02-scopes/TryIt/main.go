package main

// ---------------------------------------------------------
// EXERCISE: Try the scopes
//
//  1. Create: main.go
//
//  3. In main.go:
//     1. Create the usual func main
//     2. Call your function just by using its name: hello
//     3. Create a function named bye
//     4. Inside the bye function, print "bye bye"
//
// ---------------------------------------------------------

import (
	"fmt"
)

func main() {
	hello()
}

func bye() {
	fmt.Println("bye bye")
}

