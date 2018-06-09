package main

import (
	"fmt"
	"sync"
	"time"
)

// takes an integer and a pointer to a wait group
func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done() // decrement the waitgroup counter
}

func main() {
	no := 3
	var wg sync.WaitGroup // declare a zero value variable of type WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1) // increment waitgroup counter by 1
		go process(i, &wg)
	}
	// Wait() blocks the Goroutine in which wg is passed, until the counter becomes zero.
	// Therefore main stops here until the counter becomes zero. 
	wg.Wait()
	fmt.Println("All go routines finished executing")
}

