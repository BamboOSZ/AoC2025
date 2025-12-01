package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	dialMax = 99
	dialMin = 0
)

func main() {

	var dial int = 50
	var total int = 0

	// read file
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// split input into lines
	lines := strings.Split(string(input), "\n")

	// process each line
	for _, line := range lines {
		line = strings.TrimSpace(line) // remove \r, \n, and other whitespace
		// split line into instructions
		direction := line[0] // gets first byte/character
		distance := line[1:] // remaining string as distance
		dist, err := strconv.Atoi(distance)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Direction: %c, Distance: %d\n", direction, dist)
		if err != nil {
			continue // skip invalid instructions
		}

		switch direction {
		case 'L':
			dial = TurnLeft(dial, dist)
			if dial == 0 {
				total++
			}
		case 'R':
			dial = TurnRight(dial, dist)
			if dial == 0 {
				total++
			}
		}

		fmt.Printf("Dial: %d, Total: %d\n", dial, total)
	}

	fmt.Printf("Total number of 0 is %d\n", total)
}

func TurnRight(dial, distance int) int {
	var newDial int = (dial + distance) % (dialMax + 1)
	return newDial
}

func TurnLeft(dial, distance int) int {
	var newDial int = (dial - distance)
	if newDial < 0 {
		newDial += (dialMax + 1)
	}
	newDial = newDial % (dialMax + 1)
	return newDial
}

