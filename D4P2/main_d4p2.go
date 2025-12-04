package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 4 Part 2")

	input, err := os.ReadFile("input_d4p2.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	var total int = 0

	table := [][]byte{}

	for _, line := range lines {
		table = append(table, []byte(line))
	}

	untilChange := true
	for untilChange {
		lastTotal := total
		for y := 0; y < len(table); y++ {
			for x := 0; x < len(table[y]); x++ {
				if table[y][x] == '@' {
					neighbours := 0
					if y-1 >= 0 && table[y-1][x] == '@' {
						neighbours++
					}
					if y+1 < len(table) && table[y+1][x] == '@' {
						neighbours++
					}
					if x-1 >= 0 && table[y][x-1] == '@' {
						neighbours++
					}
					if x+1 < len(table[y]) && table[y][x+1] == '@' {
						neighbours++
					}
					if y-1 >= 0 && x-1 >= 0 && table[y-1][x-1] == '@' {
						neighbours++
					}
					if y-1 >= 0 && x+1 < len(table[y]) && table[y-1][x+1] == '@' {
						neighbours++
					}
					if y+1 < len(table) && x-1 >= 0 && table[y+1][x-1] == '@' {
						neighbours++
					}
					if y+1 < len(table) && x+1 < len(table[y]) && table[y+1][x+1] == '@' {
						neighbours++
					}
					if neighbours < 4 {
						total++
						table[y][x] = 'x'
					}
				}

			}
		}

		fmt.Println("Total: ", total)

		if lastTotal == total {
			untilChange = false
		}
	}

	fmt.Println("Total: ", total)
}
