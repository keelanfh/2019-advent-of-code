package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func run(noun, verb int, arr []int) int {
	// Copying to create a new slice
	iarr := append([]int{}, arr...)

	iarr[1] = noun
	iarr[2] = verb

	for i := 0; i <= len(iarr); i += 4 {
		op, a1, a2, res := iarr[i], iarr[i+1], iarr[i+2], iarr[i+3]

		switch op {
		// addition
		case 1:
			iarr[res] = iarr[a1] + iarr[a2]
		// multiplication
		case 2:
			iarr[res] = iarr[a1] * iarr[a2]
		// end program
		case 99:
			return iarr[0]
		}
	}
	return 0
}

func main() {
	file, err := ioutil.ReadFile("2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := string(file)
	arr := strings.Split(data, ",")
	iarr := make([]int, len(arr))
	for i, s := range arr {
		iarr[i], _ = strconv.Atoi(s)
	}

	fmt.Println(run(12, 2, iarr))

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			if run(noun, verb, iarr) == 19690720 {
				fmt.Println(100*noun + verb)
				break
			}
		}
	}
}
