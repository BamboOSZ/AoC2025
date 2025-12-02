package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 2 Part 2")

	input, err := os.ReadFile("input_d2p2.txt")
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
		for index := start; index <= end; index++ {
			num := strconv.Itoa(index)
			numLen := len(num)
			for chunkSize := 1; chunkSize <= numLen/2; chunkSize++ {
				if numLen%chunkSize != 0 {
					continue
				}

				chunks := []string{}
				for iterator := 0; iterator < numLen; iterator += chunkSize {
					chunks = append(chunks, num[iterator:iterator+chunkSize])
				}
				allEqual := true
				first := chunks[0]
				for _, chunk := range chunks {
					if chunk != first {
						allEqual = false
						break
					}
				}

				if allEqual {
					fmt.Println("Repeated number found: ", num)
					results = append(results, index)
					result += index
					break
				}
			}
		}
	}

	fmt.Println(results)
	fmt.Println("Result: ", result)
}
