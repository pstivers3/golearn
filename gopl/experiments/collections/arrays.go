package main

import "fmt"

func main() {
    var s [10]string
    s[0] = "one"
    s[1] = "two"
    s[2] = "three"
    s[3] = "four"

    fmt.Println("s[0]: \"one\"")
    fmt.Println("s[1]: \"two\"")
    fmt.Println("s[2]: \"three\"")
    fmt.Println("s[3]: \"four\"\n")

    fmt.Println("s[0:]: ", s[0:])
    fmt.Println("s[:2]: ", s[:2] )
    fmt.Println("s[0:3]: ", s[0:3], "\n")

    fmt.Println("lenth of s is ", len(s), "\n")
}
