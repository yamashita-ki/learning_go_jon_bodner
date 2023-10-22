package main

import "fmt"

func main() {
	var nilMap map[string]int
	fmt.Println(len(nilMap))
	fmt.Println(nilMap["abc"])
	fmt.Println("---------------------")
	totalWins := map[string]int{}
	fmt.Println(totalWins == nil)
	fmt.Println(totalWins["abc"])
	totalWins["abc"] = 3
	fmt.Println(totalWins)
	fmt.Println(totalWins["abc"])
	fmt.Println("---------------------")
	teams := map[string][]string{
		"writers":   {"natsume", "mori", "kunikida"},
		"knight":    {"takeda", "uesugi"},
		"musicians": {"label", "list"},
	}
	fmt.Println(teams)
	fmt.Println("---------------------")
	totalWins = map[string]int{}
	totalWins["writers"] = 1
	totalWins["knight"] = 2
	fmt.Println(totalWins["writers"])
	fmt.Println(totalWins["musicians"])
	totalWins["musicians"]++
	fmt.Println(totalWins["musicians"])
	fmt.Println("---------------------")
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)
	v, ok = m["world"]
	fmt.Println(v, ok)
	// if it does not exit, ok returns false
	v, ok = m["goodbye"]
	fmt.Println(v, ok)
	fmt.Println("---------------------")
	m = map[string]int{
		"hello": 5,
		"world": 0,
	}
	fmt.Println(m)
	delete(m, "hello")
	fmt.Println(m)
	fmt.Println("---------------------")
	intSet := map[int]bool{}
	vals := []int{1, 2, 3, 4, 5, 6, 2}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[100])
	if intSet[100] {
		fmt.Println("100 exists")
	}
	if intSet[1] {
		fmt.Println("1 exists")
	}
}
