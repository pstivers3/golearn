package main

import "fmt"

type twoFloats struct {
	a, b float64
}

// method persScaleMethod with receiver tf
// pointer makes the twoFloats values perisistent
func (tf *twoFloats) persScaleMethod(s float64) {
	tf.a = tf.a * s
	tf.b = tf.b * s
}

// func persScaleFunc with same functionality as method persScaleMethod
// pointer type makes the twoFloats values perisistent
func persScaleFunc(tf *twoFloats, s float64) {
	tf.a = tf.a * s
	tf.b = tf.b * s
}

// method scaleMethod, returns two floats
// becasue the receiver type is not a pointer, twoFloats values are not persistent
// if this method didn't return anything, it would have no effect
func (tf twoFloats) scaleMethod(s float64) (float64, float64) {
	tf.a = tf.a * s
	tf.b = tf.b * s
	return tf.a, tf.b
}

// func scaleFunc with same functionality as ScaleMethod
func scaleFunc(tf twoFloats, s float64) (float64, float64) {
	tf.a = tf.a * s
	tf.b = tf.b * s
	return tf.a, tf.b
}

// returns the two values in the twoFloats struct
func returnTwoFloats(tf twoFloats) (float64, float64) {
	return tf.a, tf.b
}

func main() {
	tf := twoFloats{1,2}
	tf.persScaleMethod(10)
	fmt.Println(returnTwoFloats(tf)) // 10 20, persistent
	persScaleFunc(&tf, 2)
	fmt.Println(returnTwoFloats(tf)) // 20 40, persistent

	tf2 := twoFloats{3, 4}
	fmt.Println(tf2.scaleMethod(10)) // 30 40, not persistent
	fmt.Println(scaleFunc(tf2, 2)) // 6 8, not persistent

	// scales the persistent values held in tf
	// the scaling itself in call to scaleMethod is not persistent
	fmt.Println(tf.scaleMethod(3)) // 60 120 
	fmt.Println(returnTwoFloats(tf)) // 20 40
	fmt.Println(returnTwoFloats(tf2)) // 3 4
}
