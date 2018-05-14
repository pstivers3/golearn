package main

import "fmt"

func main() {
    m := map[string]int {
        "a": 1,
        "b": 2,
        "c": 3,
    }

    for key, value := range m {
        fmt.Printf ("key: %s, value: %d\n", key, value)
    }

    fmt.Println ("\nm[\"a\"] : ", m["a"])
    m["a"]++
    fmt.Println ("m[\"a\"]++ : ", m["a"])
}
