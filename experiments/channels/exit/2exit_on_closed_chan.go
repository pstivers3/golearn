// from golangbot.com

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
	// for over range of ch
	// exits when the chan in the goroutine closes
	for  v := range ch {
        fmt.Println("Received ", v)
    }
	fmt.Println("channel closed, exited for loop")
}
