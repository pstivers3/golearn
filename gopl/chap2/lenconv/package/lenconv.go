// package lenghthconv performs inches to centimeters conversion

package lenconv

import "fmt"

type In float64
type Cm float64

func (in In) String() string {
	return fmt.Sprintf("%g in", in)
}

func (cm Cm) String() string {
	return fmt.Sprintf("%g cm", cm)
}

