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
