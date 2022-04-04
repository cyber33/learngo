package main

import (
        "fmt"
)

func main() {
        type Company struct {
                Name          string
                EmployeeCount int
                NonProfit     bool
	  }

        company1 := Company {
                Name:          "Acme", 
                EmployeeCount: 200, 
                NonProfit:     false,
        }

        fmt.Printf("company1: %+v\n", company1)
}