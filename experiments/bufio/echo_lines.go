// experiment with bufio 

package main

import (
    "bufio"
    "fmt"
    "os"
)

// echo lines typed to stdin after the program is run
// ctrl-d at the beginning of line to exit
func main() {
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
       fmt.Println(input.Text())
    }
}
