package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 6 Part 1")

	input, err := os.ReadFile("input_d6p1.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	total := 0

	for _, line := range lines {
		_ = line
	}

	fmt.Println("Total: ", total)
}

