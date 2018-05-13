// from https://blog.golang.org/go-slices-usage-and-internals

package main

import "fmt"

// in declaration of AppendByte, 'data ...byte' apparently allows func to take
// unlimited amounts of additional data of tupe byte ?? q
// see the call to AppendByte(p, 7, 11, 13) below
func AppendByte(slice []byte, data ...byte) []byte {
    m := len(slice)
    n := m + len(data)
    if n > cap(slice) { // if necessary, reallocate
        // allocate double what's needed, for future growth.
        newSlice := make([]byte, (n+1)*2)
        copy(newSlice, slice)
        slice = newSlice
    }
	fmt.Println(slice)
    slice = slice[0:n]
	fmt.Println(slice)
	fmt.Println(slice[m:n])
    copy(slice[m:n], data)
	fmt.Println(slice)
    return slice
}

func main() {
	p := []byte{2, 3, 5}
	p = AppendByte(p, 7, 11, 13)
	fmt.Println(p)
}
