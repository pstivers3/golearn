// converts centemeters to inches

package main

import (
	"fmt"
	"os"
	"strconv"
	"../lenconv/package"
)

func main() {
	for _, arg := range os.Args[1:] {
		c, err := strconv.ParseFloat (arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		cm := lenconv.Cm(c)
		fmt.Printf("%.2f = %s\n", cm, lenconv.CmToIn(cm))
	}
}
