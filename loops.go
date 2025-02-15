package main

import (
	"fmt"
)

func loopMain() {
	// basicLoops()
	// basicPattern(15)
	// printOnlyEven()
	var arr = []int{2, 4, 43, 12, 43, 4}
	rangeLoop(arr)
	// var isNumPresent = findNumber(arr, 473)
	// fmt.Println("Nos is present", isNumPresent)
}

func basicLoops() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even", i)
		}
	}
}

func printOnlyEven() {
	for i := 1; i < 20; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd ", i)

	}
}

func findNumber(arr []int, num int) bool {
	var l = len(arr) - 1

	for l > -1 {
		if arr[l] == num {
			return true
		}
		l--
	}
	return false
}

func basicPattern(maxHeight int) {
	for i := 1; i <= maxHeight; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := maxHeight - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func rangeLoop(arr []int) {
	for index, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
