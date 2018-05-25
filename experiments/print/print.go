package main

import (
    "fmt"
    "os"
)

func main() {
  fmt.Printf("%v\n", os.Stderr)

  fmt.Printf("\r*** first line ***") // how come only prints the last line ?? q 
  fmt.Printf("\r*** second line ***")
  fmt.Printf("\r*** third line ***")
}

