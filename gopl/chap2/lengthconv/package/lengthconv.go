// package lenghthconv performs inches to centimeters conversion

package lengthconv

import "fmt"

type In float64
type Cm float64

func (in In) string() string {
	return fmt.Sprintf("%gin", in)
}

func (cm Cm) string() string {
	return fmt.Sprintf("%gcm", cm)
}

