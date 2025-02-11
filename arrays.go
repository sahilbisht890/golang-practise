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

	// Slices are similar to arrays, but are more powerful and flexible.
	// Like arrays, slices are also used to store multiple values of the same type in a single variable.
	// However, unlike arrays, the length of a slice can grow and shrink as you see fit.

	// slice_name := []datatype{values}

	myslice1 := []int{1}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))

	temp := [6]int{10, 11, 12, 13, 14, 15}
	myslice2 := temp[2:4]
	fmt.Println(myslice2)

	myslice3 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))

}

func sliceAppend() {
	myslice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	// slice3 = append(slice1, slice2...)
	// Note: The '...' after slice2 is necessary when appending the elements of one slice to another.

}

// func highestNumber(arr) {

// }
