// methods and functions that take a value argument
// funcs that take a value argument must take a value when called 
// methods that take a value argument can take either a value or a pointer when called 

package main

import "fmt"

type twoFloats struct {
	a, b float64
}

func (tf twoFloats) persScaleMethod(s float64) (float64, float64) {
	tf.a = tf.a * s
	tf.b = tf.b * s
	return tf.a, tf.b
}

func persScaleFunc(tf twoFloats, s float64) (float64, float64) {
	tf.a = tf.a * s
	tf.b = tf.b * s
	return tf.a, tf.b
}

func returnTwoFloats(tf twoFloats) (float64, float64) {
	return tf.a, tf.b
}

func main() {
	// no changes to values in twoFloats strict because reciever to method is not a type pointer
	// and tf argument to func is not a type pointer
	tf := &twoFloats{1,2}
	tf.persScaleMethod(2)
	fmt.Println(returnTwoFloats(*tf)) // 1 2

	tf1 := twoFloats{1,2}
	tf1.persScaleMethod(3)
	p := &tf1
	fmt.Println(p.persScaleMethod(2)) // 2, 4 interpreted as *p because receiver is a value
	fmt.Println(returnTwoFloats(tf1)) // 1, 2 

	fmt.Println(persScaleFunc(tf1, 3)) // 3 6 
}
