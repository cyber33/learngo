// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Print the literals
//
//  1. Print a few integer literals
//
//  2. Print a few float literals
//
//  3. Print true and false bool constants
//
//  4. Print your name using a string literal
//
//  5. Print a non-english sentence using a string literal
//
// ---------------------------------------------------------

import (
	"fmt"
)

func main() {
	// Use fmt.Println()
	fmt.Println("Let's try english and convert this sentence using Google's translate.")
        fmt.Println("\n\n\n\n")
	// Arabic not visible but runs and the cursor goes in the right direction when you move over the text i.e. right to left.
	fmt.Println("لنجرب اللغة الإنجليزية ونحول هذه الجملة باستخدام ترجمة Google.")
        // aStr := "لنجرب اللغة الإنجليزية ونحول هذه الجملة باستخدام ترجمة Google."
	// This is to confirm I used the copy 'button' in the Google Translate app and not just the mouse - I'm trusting Google at this point the copy is the true UTF-8 Arabic text
        aStr := "لنجرب اللغة الإنجليزية ونحول هذه الجملة باستخدام ترجمة Google."
	fmt.Printf("% x\n", aStr)
}
