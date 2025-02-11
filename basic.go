package main

import (
	"fmt"
)

func main() {
	mainFunc()
}

func basic1() {
	var a int
	var b string
	var c float64
	var d bool
	e := 45 // short declaration must be inside a function and initialize the variable

	var x = 10
	x = 12

	y := 5
	y = 7
	// y :=7 // this will give error as y is already declared

	fmt.Println(a, b, c, d, e, x, y)
	// 0  ""  0 false
}

func basic2() {
	// multiple declartion

	var a, b, c, d int = 1, 3, 5, 7
	var x, y = 9, "sahil"
	fmt.Println(a, b, c, d, x, y)

	const pi = 3.14
	fmt.Println(pi)
	fmt.Printf("value of pi is %v", pi)

	var cc int

	fmt.Printf("value of cc is %T", cc)
	// %v is used to print the value of the arguments
	// %T is used to print the type of the arguments

	i := 42
	fmt.Printf("%#v\n", i) // Output: 42 (same for basic types)
	fmt.Printf("%04d\n", i)

	s := []int{1, 2, 3}
	fmt.Printf("%#v\n", s) // Output: []int{1, 2, 3}

	i2 := 42
	fmt.Printf("%v\n", i2) // Output: 42

	s2 := []int{1, 2, 3}
	fmt.Printf("%v\n", s2) // Output: [1 2 3]
}
