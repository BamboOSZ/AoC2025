package main

import (
	"fmt"
	"os"
	"strconv"
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

	table := [][]int{}
	for i := 0; i < len(lines)-1; i++ {
		line := strings.Fields(lines[i])
		tab := []int{}
		for j := 0; j < len(line); j++ {
			val, _ := strconv.Atoi(line[j])
			tab = append(tab, val)
		}

		table = append(table, tab)
	}

	operations := strings.Fields(lines[len(lines)-1])

	for x := 0; x < len(operations); x++ {
		if operations[x] == "*" {
			opTotal := 1
			for y := 0; y < len(table); y++ {
				opTotal *= table[y][x]
			}

			fmt.Println("Operation: ", operations[x], "Total: ", opTotal)

			total += opTotal
		}

		if operations[x] == "+" {
			opTotal := 0
			for y := 0; y < len(table); y++ {
				opTotal += table[y][x]
			}

			fmt.Println("Operation: ", operations[x], "Total: ", opTotal)

			total += opTotal
		}
	}

	fmt.Println("Total: ", total)
}
