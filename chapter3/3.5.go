package main

import "fmt"

func main() {
	type person struct {
		name string
		age  int
		pet  string
	}
	var fred person
	fmt.Println(fred)
	bob := person{
		"bob",
		20,
		"cat",
	}
	fmt.Println(bob)
	fmt.Println(bob.name)
	fmt.Println("---------------------")
	pet := struct {
		name string
		kind string
	}{
		name: "doggy G",
		kind: "dog",
	}
	fmt.Println(pet)
	fmt.Println("---------------------")
}
