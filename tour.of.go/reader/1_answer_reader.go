package main

import (
	"golang.org/x/tour/reader"
	"fmt"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(s []byte) (n int, err error) {
    s = s[:cap(s)];
    for i := range s {
        s[i] = 'A'
		fmt.Println(s[i]) // prints 65, which is asci char A,  many times
    }
	fmt.Println(cap(s)) // prints 2048 many times, after a bunch of 65 from the for loop
    return cap(s), nil
}

func main() {
	reader.Validate(MyReader{})
}
