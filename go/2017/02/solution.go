package main

import (
	"fmt"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var INPUT_FILE = "../../../input-files/2017/02/input.txt"

func main() {
	values := utils.ProcessIntsLinesFile(INPUT_FILE)
	part1Score, part2Score := calcScores(values)

	fmt.Println("Part 1 score:", part1Score)
	fmt.Println("Part 2 score:", part2Score)
}

func calcScores(values [][]int) (part1Score, part2Score int) {
	part1Score = 0
	part2Score = 0

	for _, value := range values {
		min, max := findMinMax(value)
		part1Score += max - min

		i, j := findDivisible(value)
		part2Score += i / j
	}

	return part1Score, part2Score
}

func findMinMax(numbers []int) (min, max int) {
	min = numbers[0]
	max = numbers[0]

	for _, n := range numbers {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	return min, max
}

func findDivisible(numbers []int) (int, int) {
	for i, n := range numbers {
		for j, m := range numbers {
			if i == j {
				continue
			}

			if n%m == 0 {
				return n, m
			}
		}
	}

	return 0, 0
}
