// from golangbot.go
// given an int, calculates the sqare of the digits plus the cube of the digits

package main

import (
    "fmt"
)

func calcSquares(number int, sqch chan int) {
    sum := 0
    for number != 0 {
        digit := number % 10
        sum += digit * digit
        number /= 10
    }
    sqch <- sum
}

func calcCubes(number int, cubech chan int) {
    sum := 0
    for number != 0 {
        digit := number % 10
        sum += digit * digit * digit
        number /= 10
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
    fmt.Println("Final output", squares + cubes)
}
