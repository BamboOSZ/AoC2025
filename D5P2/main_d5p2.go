package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 5 Part 1")

	input, err := os.ReadFile("input_d5p2.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	total := 0
	ranges := [][]int{}

	for _, line := range lines {
		split := strings.Split(line, "-")
		start, _ := strconv.Atoi(split[0])
		end, _ := strconv.Atoi(split[1])

		ranges = append(ranges, []int{start, end})

	}

	slices.SortFunc(ranges, func(a, b []int) int {
		if a[0] == b[0] {
			return 0
		}
		if a[0] < b[0] {
			return -1
		} else {
			return 1
		}
	})

	mergedRanges := [][]int{}

	for i := 0; i < len(ranges); i++ {

		start := ranges[i][0]
		end := ranges[i][1]

		newRange := true

		for i := 0; i < len(mergedRanges); i++ {
			ranStart := mergedRanges[i][0]
			ranEnd := mergedRanges[i][1]
			if start >= ranStart && end <= ranEnd {
				fmt.Println("Range already contained in another range ")
				newRange = false
				break
			}

			if start < ranStart && end >= ranStart-1 {
				mergedRanges[i][0] = start
				newRange = false
				break
			}

			if start <= ranEnd+1 && end > ranEnd {
				mergedRanges[i][1] = end
				newRange = false
				break
			}
		}

		if newRange {
			mergedRanges = append(mergedRanges, []int{start, end})
		}
	}

	for _, ran := range mergedRanges {
		total += ran[1] - ran[0] + 1
	}

	fmt.Println("Total: ", total)
}
