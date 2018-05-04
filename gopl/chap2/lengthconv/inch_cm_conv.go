// converts between inches and cm

package main

import (
	"fmt"
	"../lengthconv/package"
)

func main() {
	var in lengthconv.In = 1
	fmt.Println(lengthconv.InToCm(in))
}
