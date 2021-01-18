package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var common map[[2]int]bool

func visitedLocations(first []string) [][2]int {
	var xcur, ycur int
	var firstVisited [][2]int
	for _, inst := range first {
		dir := string(inst[0])
		mag, _ := strconv.Atoi(string(inst[1]))
		var xm, ym int

		if dir == "R" {
			xm = 1
		} else if dir == "L" {
			xm = -1
		} else if dir == "U" {
			ym = 1
		} else if dir == "D" {
			ym = -1
		}

		for i := 0; i < mag; i++ {
			xcur += xm
			ycur += ym
			firstVisited = append(firstVisited, [2]int{xcur, ycur})
		}
	}
	return firstVisited
}

func main() {
	// File opening
	file, err := os.Open("3/input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	f := true

	var first []string
	var second []string
	for scanner.Scan() {
		if f {
			first = strings.Split(scanner.Text(), ",")
			f = false
		} else {
			second = strings.Split(scanner.Text(), ",")
		}
	}

	firstVisited := visitedLocations(first)
	secondVisited := visitedLocations(second)

	fmt.Println(firstVisited)
	fmt.Println(secondVisited)

}
