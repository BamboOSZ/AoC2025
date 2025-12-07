package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 5 Part 1")

	input, err := os.ReadFile("input_d5p1.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	total := 0
	ranges := [][]int{}

	readRanges := true

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			readRanges = false
			continue
		}
		if readRanges {
			split := strings.Split(line, "-")
			start, _ := strconv.Atoi(split[0])
			end, _ := strconv.Atoi(split[1])
			ranges = append(ranges, []int{start, end})
		} else {
			id, _ := strconv.Atoi(line)

			for _, ran := range ranges {
				if id >= ran[0] && id <= ran[1] {
					total++
					break
				}
			}
		}
	}

	fmt.Println("Total: ", total)
}
