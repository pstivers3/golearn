// methods and functions that take a pointer argument
// funcs that take a pointer argument must take a pointer pointer when called
// methods that take a pointer argument can take either a pointer of value when called 
//   because the call argument is interpreted as a pointer if the receiver type is a pointer

package main

import "fmt"

type twoFloats struct {
	a, b float64
}

func (tf *twoFloats) persScaleMethod(s float64) (float64, float64) {
	tf.a = tf.a * s
	tf.b = tf.b * s
	return tf.a, tf.b
}

func persScaleFunc(tf *twoFloats, s float64) (float64, float64) {
	tf.a = tf.a * s
	tf.b = tf.b * s
	return tf.a, tf.b
}

func returnTwoFloats(tf twoFloats) (float64, float64) {
	return tf.a, tf.b
}

func main() {
	tf := &twoFloats{1,2}
	tf.persScaleMethod(2)
	fmt.Println(returnTwoFloats(*tf)) // 4 8

	tf1 := twoFloats{1,2}
	tf1.persScaleMethod(3) // tf is interpreted as a pointer to tf (&tf) because the receiver is a pointer
	p := &tf1
	p.persScaleMethod(2)
	fmt.Println(returnTwoFloats(tf1)) // 6 12 
	fmt.Println(persScaleFunc(&tf1, 2)) // 12 24

	fmt.Println(persScaleFunc(&tf1, 2)) // 24 48
	// fmt.Println(persScaleFunc(tf1, 2)) // throws an exception as expected. 
}
