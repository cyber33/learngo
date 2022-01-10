// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// package main is a special package
// it allows Go to create an executable file
package main

/*
This is a multi-line comment.

import keyword makes another package available
  for this .go "file".

import "fmt" lets you access fmt package's functionality
  here in this file.
*/
import "fmt"

// "func main" is special.
//
// Go has to know where to start
//
// func main creates a starting point for Go
//
// After compiling the code,
// Go runtime will first run this function
func main() {
	// after: import "fmt"
	// Println function of "fmt" package becomes available

	// Look at what it looks like by typing in the console:
	//   go doc -src fmt Println

	/*
		C:\Windows\System32>cd C:\Users\mark_\Data\Google\Interview Prep\Golang\learngo

		C:\Users\mark_\Data\Google\Interview Prep\Golang\learngo>go doc -src fmt Println
		package fmt // import "fmt"

		// Println formats using the default formats for its operands and writes to standard output.
		// Spaces are always added between operands and a newline is appended.
		// It returns the number of bytes written and any write error encountered.
		func Println(a ...interface{}) (n int, err error) {
				return Fprintln(os.Stdout, a...)
		}

		C:\Users\mark_\Data\Google\Interview Prep\Golang\learngo>

	*/

	// Println is just an exported function from
	//   "fmt" package

	// Exported = First Letter is uppercase
	fmt.Println("Hello Gopher Boyee!")

	// Go cannot call Println function by itself.
	// That's why you need to call it here.
	// It only calls `func main` automatically.

	// -----

	// Go supports Unicode characters in string literals
	// And also in source-code: KÖSTEBEK!
	//
	// Because: Literal ~= Source Code

	// EXERCISE: Remove the comments from below --> //
	fmt.Println("Merhaba Köstebek!")
}
