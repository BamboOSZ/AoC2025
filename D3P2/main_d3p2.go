package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 3 Part 2")

	input, err := os.ReadFile("input_d3p2.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	var total int = 0
	var wordLength = 12

	for _, line := range lines {
		line = strings.TrimSpace(line)
		indexes := []int{}
		for index := 0; index < wordLength; index++ {
			indexes = append(indexes, 0)
		}

		for index := 0; index < len(indexes); index++ {
			for cursor := indexes[index]; cursor < len(line)-wordLength+index+1; cursor++ {
				valAtCursor := int(line[cursor] - '0')
				currMax := int(line[indexes[index]] - '0')
				if valAtCursor > currMax {
					indexes[index] = cursor
				}
			}

			if index+1 < len(indexes) {
				indexes[index+1] = indexes[index] + 1
			}

		}

		var number string = ""
		for index := 0; index < wordLength; index++ {
			number += string(line[indexes[index]])
		}

		value, _ := strconv.Atoi(number)

		fmt.Println("Value: ", value)

		total += value
	}

	fmt.Println("Total: ", total)

}
