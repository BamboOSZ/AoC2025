package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 7 Part 1")

	input, err := os.ReadFile("input_d7p2_s.txt")
	if err != nil {
		panic(err)
	}

	table := [][]byte{}
	paths := [][]int{}

	lines := strings.Split(string(input), "\n")

	total := 0

	for _, line := range lines {
		table = append(table, []byte(line))
		paths = append(paths, make([]int, len(line)))
	}

	for y := 0; y < len(table); y++ {
		for x := 0; x < len(table[y]); x++ {
			if table[y][x] == 'S' {
				table[y+1][x] = '|'
				paths[y+1][x] += 1
			}

			if y > 0 && table[y][x] == '.' && table[y-1][x] == '|' {
				table[y][x] = '|'
				paths[y][x] += paths[y-1][x]
			}

			if table[y][x] == '^' && table[y-1][x] == '|' {
				paths[y][x] = paths[y-1][x]
				table[y][x-1] = '|'
				table[y][x+1] = '|'

				paths[y][x-1] += paths[y][x]
				paths[y][x+1] += paths[y][x]
			}
		}
	}

	for y := 0; y < len(table); y++ {
		for x := 0; x < len(table[y]); x++ {
			if table[y][x] == '.' {
				fmt.Printf(" ")
				continue
			} else {
				if table[y][x] != '|' {
					fmt.Printf("%c", table[y][x])
					continue
				}
				fmt.Printf("%d", paths[y][x])
			}
		}
		fmt.Println()
	}

	for i := 0; i < len(paths[len(paths)-1]); i++ {
		total += paths[len(paths)-1][i]
	}

	fmt.Println("Total: ", total)
}
