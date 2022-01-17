package main

import (
	"fmt"
	"log"
	"os"	
)

func main() {
	// From: https://zetcode.com/golang/writefile/ 

	f, err := os.Create("Arabic.txt")

 	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

        aStr := "لنجرب اللغة الإنجليزية ونحول هذه الجملة باستخدام ترجمة Google."

	_, err2 := f.WriteString(aStr)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}
