package main

import "fmt"

func stuctMain() {
	// A struct (short for structure) is used to create a
	// collection of members of different data types, into a single variable.
	type person struct {
		age  int
		name string
	}

	var user person = person{
		age:  24,
		name: "sahil",
	}

	var user2 person
	user2.age = 21
	user2.name = "demo"

	fmt.Println(user, user2)
}
