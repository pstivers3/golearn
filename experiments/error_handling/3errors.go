// from tour.golang.org

package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{ time.Now(), "it didn't work", }
}

func main() {
	fmt.Println(MyError{ time.Now(), "it didn't work", })
	fmt.Println(&MyError{ time.Now(), "it didn't work", })
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

