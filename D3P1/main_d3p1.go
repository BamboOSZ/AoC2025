package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 3 Part 1")

	// read file
	input, err := os.ReadFile("input_d3p1.txt")
	if err != nil {
		panic(err)
	}

	// split input into lines
	lines := strings.Split(string(input), "\n")

	var total int = 0

	// process each line
	for _, line := range lines {
		line = strings.TrimSpace(line)
		index1 := 0
		index2 := 0

		for index := 0; index < len(line)-1; index++ {
			if line[index] > line[index1] {
				index1 = index
			}
		}

		index2 = index1 + 1

		for index := index2; index < len(line); index++ {
			if line[index] > line[index2] {
				index2 = index
			}
		}

		value := int(line[index1]-'0')*10 + int(line[index2]-'0')

		fmt.Println("Index1: ", index1, "Index2: ", index2, "Value: ", value)

		total += value
	}

	fmt.Println("Total: ", total)

}
