package main

import "fmt"

func main() {
	var s string = "Hello there"
	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]
	fmt.Println(s, s2, s3, s4)
	fmt.Println("---------------------")
	s = "Hello 🌥"
	// convert string to byte array
	var bs []byte = []byte(s)
	// convert string to rune array
	var rs []rune = []rune(s)
	fmt.Println(bs)
	fmt.Println(rs)
}
