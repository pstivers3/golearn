package main

import "fmt"

func main() {
    fmt.Println("*** Arrays ***\n")
    var s [10]string
    s[0] = "one"
    s[1] = "two"
    s[2] = "three"
    s[3] = "four"

    fmt.Println("s[0] = \"one\"")
    fmt.Println("s[1] = \"two\"")
    fmt.Println("s[2] = \"three\"")
    fmt.Println("s[3] = \"four\"\n")

    fmt.Println("s[0:] = ", s[0:])
    fmt.Println("s[:2] = ", s[:2] )
    fmt.Println("s[0:3] = ", s[0:3], "\n")

    fmt.Println("lenth of s is ", len(s), "\n")

    fmt.Println("*** maps ***\n")

    m := map[string]int {
        "a": 1,
        "b": 2,
        "c": 3,
    }

    for key, value := range m {
        fmt.Println ("key: ", key, ", value: ", value)
    }
}
