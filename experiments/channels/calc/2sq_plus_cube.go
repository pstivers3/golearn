// from golangbot.go
// given an int, calculates the sqare of the digits plus the cube of the digits

package main

import (
    "fmt"
)

func digits(number int, dch chan int) {
    for number != 0 {
        digit := number % 10
        dch <- digit
        number /= 10
    }
    close(dch)
}
func calcSquares(number int, sqch chan int) {
    sum := 0
    dch := make(chan int)
    go digits(number, dch)
    for digit := range dch {
        sum += digit * digit
    }
    sqch <- sum
}

func calcCubes(number int, cubech chan int) {
    sum := 0
    dch := make(chan int)
    go digits(number, dch)
    for digit := range dch {
        sum += digit * digit * digit
    }
    cubech <- sum
}

func main() {
    number := 123
    sqch := make(chan int)
    cubech := make(chan int)
    go calcSquares(number, sqch)
    go calcCubes(number, cubech)
    squares, cubes := <-sqch, <-cubech
    fmt.Println("Final output", squares+cubes)
}
