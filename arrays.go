package main

import "fmt"

func mainFunc() {

	basicDeclartion()
}

func basicDeclartion() {
	var arr = [5]int{}
	var arr2 = [5]int{2, 4}
	var arr3 = [...]int{1, 2, 3}
	arr4 := [5]int{1: 10, 2: 40}

	fmt.Println(arr, arr2, arr3, arr4)
}

// func highestNumber(arr) {

// }
