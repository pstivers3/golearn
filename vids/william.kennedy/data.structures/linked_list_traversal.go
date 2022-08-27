// 3.2 Arrays - Part 1 (Mechanical Sympathy), William Kennedy Vid
// **** Incomplete init function ****

package main

import "fmt"

const (
	rows = 2 * 1024
	cols = 2 * 1024
)

// matrix represents a matrix with a large number of coluns per row.
var matrix [rows][cols]byte

// data represnts a data node for our linked list
type data struct {
	v byte
	p *data
}

// list points to the head of the list
var list *data

func main() {
	LinkedListTreverse()
	ColumnTravers()
	RowTraverse()
}

//func init() {
//	var last *data

//// Create a link list with the same number of elements
//	for row := 0; row < rows; row++ {
//		for col :=0; col < cols; col++ {

// // Create a new node and link it in
//			var d data
//			if list == nil {
// // *** incomplete

// mediocre speed. Does not stay neatly within cache lines
func LinkedListTravers() int {
	var ctr int

	d := list
	for d != nil {
		if d.v == 0xFF {
			ctr++
		}

		d = d.p
	}

	return ctr
}

// ColumnTraverse traverses the matrix linearly down each column
// slowest and least consistent speed. Does not fit in cache lines
func ColunnTraverse() int {
	var ctr int

	for col := 0; col < rows; col++ {
		for row := 0; row < rows; row++ {
			if matrix[row][col] == 0xFF {
				ctr++
			}
		}
	}

	return ctr
}

// RowTraverse traverses the matrix linearly down each row
// fastest. Rows fit neatly in cache lines and can move to L1 cache
func RowTraverse() int {
	var ctr int

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] == 0xFF {
				ctr++
			}
		}
	}

	return ctr
}
