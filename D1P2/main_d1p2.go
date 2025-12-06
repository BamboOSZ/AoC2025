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

	input, err := os.ReadFile("input_d1p2.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		direction := line[0]
		distance := line[1:]
		dist, err := strconv.Atoi(distance)
		if err != nil {
			panic(err)
		}
		fmt.Printf("move Dial: %d, in Direction: %c, by Distance: %d\n", dial, direction, dist)

		zerosCrossed := 0

		switch direction {
		case 'R':
			zerosCrossed = (dial + dist) / (dialMax + 1)
			dial = (dial + dist) % (dialMax + 1)

		case 'L':
			if dist >= dial {
				if dial != 0 {
					zerosCrossed++
				}

				zerosCrossed += (dist - dial) / (dialMax + 1)
			}
			dial = ((dial-dist)%(dialMax+1) + (dialMax + 1)) % (dialMax + 1)
		}

		fmt.Printf("Zeros crossed: %d\n", zerosCrossed)

		total += zerosCrossed

		fmt.Printf("Dial after move: %d, Total zeros: %d\n", dial, total)
	}

	fmt.Printf("Final Total number of zeros crossed is %d\n", total)
}
