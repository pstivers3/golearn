package main

import (
    "fmt"
)

func producer(chnl chan int) {
    for i := 0; i < 3; i++ {
        chnl <- i
    }
    close(chnl)
}

func main() {
    ch := make(chan int)
    go producer(ch)
	// continuous for loop
	// exits when the chan in the goroutine closes
    for {
        v, ok := <-ch
        fmt.Println("Received ", v, ok)
        if ok == false {
			fmt.Println("channel closed, exiting for loop")
            break
        }
    }
}
