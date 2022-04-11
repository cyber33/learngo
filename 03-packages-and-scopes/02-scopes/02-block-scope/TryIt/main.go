// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// https://go.dev/play/p/V3QcHXwJbmB

package main

import "fmt"

func nope() { // block scope starts

	// hello and ok are only visible here
	const ok = true
	var hello = "Hello"

	_ = hello
        fmt.Println(hello, ok)
} // block scope ends

func main() { // block scope starts

	// hello and ok are not visible here
	// even if I call the function, they are still locked in the scope
	// of that function
	nope()

	// ERROR:
	// fmt.Println(hello, ok)

} // block scope ends
