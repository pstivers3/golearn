package main

import "fmt"

func functions1() []func() {
	arr := []int{1, 2, 3, 4}
	result := make([]func(), 0)

	for i := range arr {
		result = append(result, func() { fmt.Printf("index - %d, value - %d\n", i, arr[i]) })
	}

	return result
}

// not as written in http://keshavabharadwaj.com/2016/03/31/closure_golang/
// tried to accomodate 'func(index,  val int)' in line 23
func functions2() []func(x int, y int) {
	arr := []int{1, 2, 3, 4}
	result := make([]func(x int, y int), 0)

	for i := range arr {
		result = append(result, func(index,  val int) { fmt.Printf("index - %d, value - %d\n", i, val) })
	}

	return result
}

// as written. Doesn't work.
// Line 37, cannot use func literal (type func(int, int)) as type func() in append
/*
func functions3() []func() {
	arr := []int{1, 2, 3, 4}
	result := make([]func(), 0)

	for i := range arr {
		result = append(result, func(index, val int) { fmt.Printf("index - %d, value - %d\n", index, val) })
	}

	return result
}
*/

func main() {
	fns1 := functions1()
	for f := range fns1 {
		fns1[f]()
	}

	fmt.Println()
	fns2:= functions2()
	for f := range fns2{
		fns2[f](1,2)
	}

	/*
	fns3 := functions3()
	for f := range fns3 {
		fns3[f]()
	}
	*/
}
