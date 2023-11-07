package main

import (
  "strings"
  "os"
)

// Joining args in
func main(){
  fmt.Println(strings.Join(os.Args[1:], " "))
}
