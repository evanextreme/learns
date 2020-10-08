// This is a runable go file where I code along with https://learnxinyminutes.com/docs/go/

package main

import (
	"fmt"
)

func main() { // main method runs first, functions not allowed to be called outside definitions
	fmt.Println("Hello World!")
	whatarenums()
}

// variable types declared after the names. x int not int x.

func whatarenums() {
	var x int // variable type definition
	x = 3     // variable value declaration
	y := 4    // variable declaration that infers type
	fmt.Println("X ", x, "Y ", y)
	sum, prod := testreturns(x, y)
	fmt.Println("Sum ", sum, "Prod ", prod)
}

func testreturns(x int, y int) (sum, prod int) {
	return x + y, x * y
}

func learntotype(){
	t1a := "This is a string"
	t1b := `This is also a string, and auto escapes stuff like quotes`
	g := 'Σ' // This is a rune. Its an alias for INT32 and I need to look up what that means tbh. Holds a unicode point.
	var i uint = 32 // has unsigned ints
	var pi float32 = 22. / 7
	n := byte('\n') // conversion / typecasting happens this way.
	var a [4]int // size of arrays are immutable
	a2 := [...]int{2, 3, 4, 5} // initialize array at declaration

	//Slices are arrays but with dynamic sizing

	s2 := []int{2, 3, 4, 5}

}