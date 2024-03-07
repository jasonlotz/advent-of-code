package main

import (
	"fmt"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var INPUT_FILE = "../../../input-files/2021/01/input.txt"

func main() {
	part1()
	part2()
}

func part1() {
	input := utils.ProcessIntLinesFile(INPUT_FILE)
	count := countIncreases(input)

	fmt.Println("Part 1:", count)
}

func part2() {
	input := utils.ProcessIntLinesFile(INPUT_FILE)

	count := countIncreasesWithSlidingWindow(input)

	fmt.Println("Part 2:", count)
}

func countIncreases(input []int) int {
	count := 0
	previous := 0

	for _, num := range input {
		if previous != 0 && previous < num {
			count++
		}

		previous = num
	}

	return count
}

func countIncreasesWithSlidingWindow(input []int) int {
	count := 0
	previous := 0

	for i := 2; i < len(input); i++ {
		current := input[i] + input[i-1] + input[i-2]
		if previous != 0 && previous < current {
			count++
		}

		previous = current
	}

	return count
}
