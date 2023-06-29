package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("My Error")
	fmt.Println(err)
	fmt.Errorf("error: %s", msg)
}
