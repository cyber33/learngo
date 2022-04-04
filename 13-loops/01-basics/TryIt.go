package main

import (
        "fmt"
)

func main() {
        var floatSum float64

        for i := 1; i <= 1000; i++ {
                floatSum += 3.14159
                fmt.Println(floatSum)
        }

        fmt.Println(floatSum)
}