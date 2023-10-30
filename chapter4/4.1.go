package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
	fmt.Println("---------------------")
	x = 10
	if x > 5 {
		a, x := 20, 5
		fmt.Println(a, x)
	}
	fmt.Println("---------------------") 
}
