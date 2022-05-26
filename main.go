package main

import (
  "fmt"
  "parser/Components"
)

func main() {
    ParsedSuccess, EmailsDomains := customerimporter.ParseExcelFile("customers.csv");

    if ParsedSuccess == true {
      customerimporter.PrintResults(EmailsDomains)
    } else {
      fmt.Printf("ERROR: Excel file parse failed")
    }
}
