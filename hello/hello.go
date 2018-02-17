package main

import (
  "fmt"
  "github.com/gravytates/go_example/stringutil"
)

func main() {
  input := "!oG ,olleH\n"
  reverseInput := stringutil.Reverse(input)
  fmt.Printf(input)
  fmt.Printf(reverseInput)
}
