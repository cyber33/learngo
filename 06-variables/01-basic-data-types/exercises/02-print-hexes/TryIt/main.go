// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// THIS EXERCISE IS OPTIONAL

// ---------------------------------------------------------
// EXERCISE: Print hexes
//
//  1. Print 0 to 9 in hexadecimal
//
//  2. Print 10 to 15 in hexadecimal
//
//  3. Print 17 in hexadecimal
//
//  4. Print 25 in hexadecimal
//
//  5. Print 50 in hexadecimal
//
//  6. Print 100 in hexadecimal
//
// EXPECTED OUTPUT
//  0 1 2 3 4 5 6 7 8 9
//  10 11 12 13 14 15    <***** I think that is wrong because 10 hex is not the same as 10 decimal which seems to have  been assumed: should be a b c d e f - see note on bottom of this file...
//  17
//  25
//  50
//  100
//
// NOTES
//  https://stackoverflow.com/questions/910309/how-to-turn-hexadecimal-into-decimal-using-brain
//
// https://simple.wikipedia.org/wiki/Hexadecimal_numeral_system
//
// ---------------------------------------------------------

func main() {
	// EXAMPLES:

	// I'm going to print 10 in hexadecimal
	// fmt.Println(0xa)

	// I'm going to print 16 in hexadecimal
	// 0x10
	//   ^^-----  1 * 0 = 0
	//   |
	//   +------ 16 * 1 = 16
	//                  = 16
	// fmt.Println(0x10)

	// I'm going to print 150 in hexadecimal
	// 0x96
	//   ^^-----  1 * 6 = 6
	//   |
	//   +------ 16 * 9 = 144
	//                  = 150
	// fmt.Println(0x96)

	// COMMENT-OUT ALL THE CODE ABOVE, THEN,
	// ADD YOUR OWN SOLUTIONS BELOW

        for x := 0; x < 10; x++ {
		fmt.Printf("%x ", x)
	}

	fmt.Println("")

	for x := 10 ; x < 16; x++ {
		fmt.Printf("%x ", x)
	}

	fmt.Println("")

        // I'm going to print 17 in hexadecimal
        // 0x10
        //   ^^-----  1 * 1 =  1
        //   |
        //   +------ 16 * 1 = 16
        //                  = 17
        fmt.Println(0x11)


        // I'm going to print 25 in hexadecimal
        // 0x10
        //   ^^-----  1 * 9 =  9 
        //   |
        //   +------ 16 * 1 = 16
        //                  = 25 
        fmt.Println(0x19)

        // I'm going to print 50 in hexadecimal
        // 0x10
        //   ^^-----  1 * 2 =  2 
        //   |
        //   +------ 16 * 3 = 48 
        //                  = 50 
        fmt.Println(0x32)

        // I'm going to print 100 in hexadecimal
        // 0x10
        //   ^^-----  1 * 4 =   4 
        //   |
        //   +------ 16 * 6 =  96 
        //                  = 100 
        fmt.Println(0x64)


	// No actually what he wanted me to do is to actually print the hex number so that the output would be decimal...
	// From the solution...

// package main

// import "fmt"

// func main() {
        fmt.Println(0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9)
        fmt.Println(0xa, 0xb, 0xc, 0xd, 0xe, 0xf)
        fmt.Println(0x11) // 17
        fmt.Println(0x19) // 25
        fmt.Println(0x32) // 50
        fmt.Println(0x64) // 100
// }


}

/*

Moral of the story: READ THE README!!!!!

mark_@LAPTOP-3D73FF3E /cygdrive/c/Users/mark_/Data/Google/Interview Prep/Golang/learngo/06-variables/01-basic-data-types/exercises
$ cat README.md
1. **[Print the literals](https://github.com/inancgumus/learngo/tree/master/06-variables/01-basic-data-types/exercises/01-print-the-literals)**

    Print a few values using the literals

2. **[Print hexes](https://github.com/inancgumus/learngo/tree/master/06-variables/01-basic-data-types/exercises/02-print-hexes)** (optional exercise)

    Print numbers in hexadecimal   <*********************************************************************** This was the clue!

mark_@LAPTOP-3D73FF3E /cygdrive/c/Users/mark_/Data/Google/Interview Prep/Golang/learngo/06-variables/01-basic-data-types/exercises
$



*/
