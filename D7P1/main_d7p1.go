package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 7 Part 1")

	input, err := os.ReadFile("input_d7p1.txt")
	if err != nil {
		panic(err)
	}

	table := [][]byte{}

	lines := strings.Split(string(input), "\n")

	total := 0

	for _, line := range lines {
		table = append(table, []byte(line))
	}

	for y := 0; y < len(table); y++ {
		for x := 0; x < len(table[y]); x++ {
			if table[y][x] == 'S' {
				table[y+1][x] = '|'
			}

			if y > 0 && table[y][x] == '.' && table[y-1][x] == '|' {
				table[y][x] = '|'
			}

			if table[y][x] == '^' && table[y-1][x] == '|' {
				table[y][x-1] = '|'
				table[y][x+1] = '|'
				total++
			}
		}
	}

	for _, line := range table {
		fmt.Println(string(line))
	}

	fmt.Println("Total: ", total)
}
