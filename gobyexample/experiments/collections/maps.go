package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"k1": 7, "k2": 14}
	m["k3"] = 21
	delete(m, "k2")
	fmt.Println(m)
	v, prs := m["k1"]
	fmt.Println(v, prs)
	v, prs = m["k2"]
	fmt.Println(v, prs)
	_, prs = m["k2"]
	fmt.Println(prs)
	v = m["k3"]
	fmt.Println(v, m["k3"])

	for k, v := range m {
		fmt.Printf("%s -> %d\n", k, v)
	}

	for k := range m {
        fmt.Println("key:", k)
	}
}
