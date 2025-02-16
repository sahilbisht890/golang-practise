package main

import (
	"fmt"
)

func stringFunc() {
	fmt.Println(romanToInt("III"))
	fmt.Println(isAnagram("anagram", "nagaram"))
}

func romanToInt(s string) int {
	var romanValues = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var result = 0
	var prev = 0
	for i := len(s) - 1; i > -1; i-- {
		var value, _ = romanValues[string(s[i])]
		if value < prev {
			result = result - value
		} else {
			result = result + value
		}
		prev = value
	}

	return result
}

func isAnagram(s string, t string) bool {
	var arr = [256]int{}

	for i := 0; i < len(s); i++ {
		var temp = s[i]
		arr[temp] += 1
	}

	for i := 0; i < len(t); i++ {
		var temp = t[i]
		arr[temp] -= 1
		if arr[temp] < 0 {
			return false
		}
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			return false
		}
	}

	return true

}

func lengthOfLastWord(s string) int {
	var l = len(s) - 1
	var r int

	for l > -1 {
		if string(s[l]) == " " {
			l--
			continue
		} else {
			r = l
			for r > -1 && string(s[r]) != " " {
				r--
			}
			break
		}
	}

	return l - r
}
