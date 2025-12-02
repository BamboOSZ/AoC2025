package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 2 Part 1")

	input, err := os.ReadFile("input_d2p1.txt")
	if err != nil {
		panic(err)
	}

	results := []int{}
	var result int = 0

	ranges := strings.Split(string(input), ",")
	for _, r := range ranges {
		fmt.Println("Range: ", r)
		r2 := strings.Split(r, "-")
		start, _ := strconv.Atoi(r2[0])
		end, _ := strconv.Atoi(r2[1])
		fmt.Println("Start: ", start, "End: ", end)
		for i := start; i <= end; i++ {
			num := strconv.Itoa(i)
			numLen := len(num)
			if numLen%2 == 0 {
				if num[:numLen/2] == num[numLen/2:] {
					fmt.Println("Repeated number found: ", num)
					results = append(results, i)
					result += i
				}
			}
		}
	}

	fmt.Println(results)
	fmt.Println("Result: ", result)
}
