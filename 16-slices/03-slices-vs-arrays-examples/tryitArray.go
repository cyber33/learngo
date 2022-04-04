package main

import (
        "fmt"
)

func main () {

/*
        // Let define and initialize a couple arrays and operate on them
        var books [5]string
        books[0] = "dracula"
        books[1] = "1984"
        books[2] = "island"

*/

        var shopping_list [10]string
        shopping_list[0] = "apples"
        shopping_list[1] = "oranges"
        shopping_list[3] = "paper towels"

        fmt.Printf("Today my shopping list is: %#v\n", shopping_list)

        shopping_list[2] = "milk"

        fmt.Printf("Today my shopping list is: %#v\n", shopping_list)

        shopping_list[1] = "dishwashing liquid"

        fmt.Printf("Today my shopping list is: %#v\n", shopping_list)

}