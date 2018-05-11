package main

import "fmt"

func main() {
    m := map[string]int {
        "a": 1,
        "b": 2,
        "c": 3,
    }

    for key, value := range m {
        fmt.Println ("key: ", key, ", value: ", value)
    }

    fmt.Println ("\n", "m[\"a\"] : ", m["a"])
    m["a"]++
    fmt.Println ("m[\"a\"]++ : ", m["a"])
}
