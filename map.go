package main

import (
	"fmt"
)

func mapFunc() {
	// var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
	// var a = map[int]string{2: "sahil", 1: "demo"}

	// 	Maps are references to hash tables.

	// If two map variables refer to the same hash table,
	// changing the content of one variable affect the content of the othe

	var b = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	delete(b, "year")

	var c = make(map[string]string)
	var d map[string]string

	// d["a"] = "c"

	// c["a"] = "c"
	// fmt.Println(b, c)

	// If key does not exist in the map. So the value is an empty string ('')

	var value, valuePresntOrNot = b["brand"]
	fmt.Println(value, valuePresntOrNot)

	fmt.Println(c == nil)
	fmt.Println(d == nil)

	// iterate over the maps

	for k, v := range b {
		fmt.Printf("%v : %v, ", k, v)
	}

}
