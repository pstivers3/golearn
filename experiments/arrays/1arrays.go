package main

import "fmt"

func main() {
    var s [5] int
    s[0] = 0
    s[1] = 1
    s[2] = 2
    s[3] = 3
    s[4] = 4

    fmt.Println("s: ", s)

    fmt.Println("s[2:]: ", s[2:])
    fmt.Println("s[:2]: ", s[:2] )
    fmt.Println("s[1:3]: ", s[1:3], "\n")

    fmt.Println("lenth of s is ", len(s), "\n")

	a := [4]int{7,8,9,10}
	a[3] = 1
	fmt.Println(a)
}
