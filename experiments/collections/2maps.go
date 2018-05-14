package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"k1": 7, "k2": 14}
	m["k3"] = 21
	delete(m, "k2")
	fmt.Println(m)
	v, prsnt := m["k1"]
	fmt.Println(v, prsnt, "\n")

	v, prsnt = m["k2"]
	fmt.Println(v, prsnt)
	_, prsnt = m["k2"]
	fmt.Println(prsnt)
	v = m["k3"]
	fmt.Println(v, m["k3"], "\n")

	for k, v := range m {
		fmt.Printf("%s -> %d\n", k, v)
	}

	for k := range m {
        fmt.Println("key:", k)
	}
}
