package main

import "fmt"

func main() {
	var data []int
	fmt.Println(data == nil)
	fmt.Println("---------------------")
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	f := x
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
	fmt.Println("z: ", z)
	fmt.Println("d: ", d)
	fmt.Println("e: ", e)
	fmt.Println("f: ", f)
	fmt.Println("---------------------")
	x = []int{1, 2, 3, 4}
	y = x[:2]
	z = x[1:]
	fmt.Println("x: ", x)
	x[1] = 20
	y[0] = 10
	z[1] = 30
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
	fmt.Println("z: ", z)
	fmt.Println("---------------------")
	x1 := [...]int{5, 6, 7, 8}
	y1 := make([]int, 2)
	num := copy(y1, x1[1:])
	fmt.Println(y1, num)
}
