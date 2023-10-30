package main

import "fmt"

type person struct {
	age  int
	name string
}

func main() {
	p := person{}
	i := 2
	s := "こんにちは"
	fmt.Println(i, s, p)
	modifyFails(i, s, p)
	fmt.Println(i, s, p)
	fmt.Println("---------------")
	m := map[int]string{
		1: "1番目",
		2: "2番目",
	}
	modMap(m)
	fmt.Println(m)
	s1 := []int{1, 2, 3}
	modSlice(s1)
	fmt.Println(s1)
}

func modifyFails(i int, s string, p person) {
	i *= 2
	s = "さようなら"
	p.name = "Bob"
	fmt.Println(i, s, p)
}

func modMap(m map[int]string) {
	m[2] = "こんにちは"
	m[3] = "さようなら"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}
