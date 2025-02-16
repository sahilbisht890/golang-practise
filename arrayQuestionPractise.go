package main

import (
	"fmt"
	"sort"
)

func arrayQuestionMain() {
	var arr = []int{42, 12, 89, 43, 12, 89}

	// fmt.Println("Array is: ", twoSum([]int{2, 11, 8, 44, 12, 3, 13, 7, 15}, 26))
	// fmt.Println("Array is: ", twoSumIndex([]int{2, 11, 8, 44, 12, 3, 13, 7, 15}, 26))
	// fmt.Println("Merged Sorted Array", mergeSortedArray([]int{4, 12, 13, 20}, []int{1, 10, 14, 18, 19}))
	fmt.Println("majority Element", majorityElement([]int{1, 2, 1, 2, 2, 1, 1, 1, 1}))
	secondLargest(arr)
}

func sortedArray(arr []int) []int {
	temp := arr[:]
	sort.Ints(temp)
	return temp
}

func secondLargest(arr []int) {
	var largest int = -1
	var secLargest int = -1
	for _, val := range arr {
		if val > largest {
			secLargest = largest
			largest = val
		} else if val > secLargest && val != largest {
			secLargest = val
		}
	}
	if largest == secLargest || secLargest == -1 {
		fmt.Println("No second largest")
	} else {
		fmt.Println("Second largest Number", secLargest)
	}
}

// just the elements
func twoSum(arr []int, sum int) []int {
	var new_arr = sortedArray(arr)
	var digits []int
	lef := 0
	rig := len(arr) - 1
	for lef < rig {
		sum_temp := new_arr[lef] + new_arr[rig]
		if sum_temp == sum {
			digits = []int{new_arr[lef], new_arr[rig]}
			return digits
		}

		if sum_temp > sum {
			rig--
		} else {
			lef++
		}
	}
	return []int{-1, -1}
}

// return the index

func twoSumIndex(arr []int, sum int) []int {
	lef := 0
	var store = make(map[int]int)

	for lef < len(arr) {
		value_needed := sum - arr[lef]
		_, value_present := store[value_needed]
		if value_present {
			return []int{store[value_needed], lef}
		}
		store[arr[lef]] = lef
		lef++

	}
	return []int{-1, -1}
}

func mergeSortedArray(arr1 []int, arr2 []int) []int {
	var mergedArray = make([]int, len(arr1)+len(arr2))
	var i, j, k = 0, 0, 0

	for i < len(arr1) && j < len(arr2) {
		var temp int
		if arr1[i] > arr2[j] {
			temp = arr2[j]
			j++
		} else {
			temp = arr1[i]
			i++
		}
		mergedArray[k] = temp
		k++
	}

	for i < len(arr1) {
		mergedArray[k] = arr1[i]
		k++
		i++
	}

	for j < len(arr2) {
		mergedArray[k] = arr2[j]
		k++
		j++
	}

	return mergedArray
}

// Given an array nums of size n, return the majority element.
// The majority element is the element that appears more than ⌊n / 2⌋ times.
// You may assume that the majority element always exists in the array.

func majorityElement(arr []int) (result int) {
	result = arr[0]
	var count = 1
	for i := 1; i < len(arr); i++ {
		if count == 0 {
			result = arr[i]
			count = 1
			continue
		}
		if arr[i] == result {
			count++
		} else {
			count--
		}
	}

	return
}
