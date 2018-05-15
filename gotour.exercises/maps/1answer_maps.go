package main

import (
    "golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    set := strings.Fields(s)
    var st_map map[string]int = make(map[string]int)
    for i := range set {
        count := 0
        for j := range set {
            if set[i] == set[j] {
                count = count + 1
            }
        }
        st_map[set[i]] = count
    }
    return st_map
}

func main() {
    wc.Test(WordCount)
}
