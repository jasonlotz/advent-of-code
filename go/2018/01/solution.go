package main

import (
	"fmt"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var INPUT_FILE = "../../../input-files/2018/01/input.txt"

func main() {
	part1()
	part2()
}

func part1() {
	inputs := utils.ProcessIntLinesFile(INPUT_FILE)
	sum := sum(inputs)
	fmt.Println("Part 1:", sum)
}

func part2() {
	inputs := utils.ProcessIntLinesFile(INPUT_FILE)
	result := seenTwice(inputs)
	fmt.Println("Part 2:", result)
}

func sum(inputs []int) int {
	sum := 0
	for _, i := range inputs {
		sum += i
	}

	return sum
}

func seenTwice(inputs []int) int {
	seen := make(map[int]bool)
	sum := 0
	i := 0

	for {
		seen[sum] = true
		newSum := sum + inputs[i]
		if seen[newSum] {
			return newSum
		}

		sum = newSum
		i = (i + 1) % len(inputs)
	}
}
