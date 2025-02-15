package main

import "fmt"

func funcMain() {
	fmt.Println(myFunction(12, "sahil"))

	// you can store the 2 value
	a, b := myFunction(5, "hi")
	fmt.Println(a, b)
	fmt.Println(factorial(5))
}

func sum(a int, b int) int {
	return a + b
}

func sum2(a int, b int) (result int) {
	result = a + b
	return
}

// Note: When you are working with multiple parameters, the function call must have the
// number of arguments as there are parameters, and the arguments must be passed in the same order

func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

// recursion

func factorial(num int) int {
	if num < 2 {
		return 1
	}
	return num * factorial(num-1)
}
