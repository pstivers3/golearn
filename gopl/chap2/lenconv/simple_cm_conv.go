// converts between inches and cm

package main

import (
	"fmt"
	"../lenconv/package"
)

func main() {
	var in lenconv.In = 1
	fmt.Println(lenconv.InToCm(in))
}
