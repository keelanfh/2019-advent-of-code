package main

import (
	"fmt"
	"strconv"
)

func checkNumber(i int) (bool, bool) {
	s := strconv.Itoa(i)
	cm := make(map[int]int)
	var p int
	// Loop through the string
	for _, r := range s {
		j, _ := strconv.Atoi(string(r))
		cm[j]++
		if j < p {
			return false, false
		}
		p = j
	}
	var p1 bool
	for _, num := range cm {
		if num == 2 {
			return true, true
		}
		if num >= 2 {
			p1 = true
		}
	}
	return p1, false
}

func main() {
	var p1total, p2total int
	for i := 123257; i <= 647015; i++ {
		p1, p2 := checkNumber(i)
		if p1 {
			p1total++
		}
		if p2 {
			p2total++
		}
	}
	fmt.Println(p1total)
	fmt.Println(p2total)
}
