package main

import "fmt"

// takes a send only chan
func sendData(sendch chan<- int) {
    sendch <- 10 // data can be sent to the chan
}

func main() {
    chnl := make(chan int)
    go sendData(chnl) // chnl get's converted to send only during the func call
    fmt.Println(<-chnl) // chnl is bidirectional again outside of the func
}
