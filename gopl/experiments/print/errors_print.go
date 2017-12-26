 package main

 import (
     "fmt"
     "os"
 )

 func main() {
     _, err := os.Open("textfile3") // file doesn't exist
     if err != nil {
         fmt.Println("os.Stderr: ", os.Stderr)
         fmt.Printf("os.Stderr: %v\n", os.Stderr)
         fmt.Println("err: ", err)
         fmt.Printf("err: %v\n", err)
     }
 }
