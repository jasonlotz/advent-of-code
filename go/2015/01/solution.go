package main

import (
	"fmt"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2015/01/input.txt"

func main() {
	part1()
}

func part1() {
	input := utils.ProcessSingleStringLineFile(inputFile)

	floor := 0
	basementReachedOn := 0

	for i, c := range input {
		if c == '(' {
			floor++
		} else {
			floor--
		}

		if floor == -1 && basementReachedOn == 0 {
			basementReachedOn = i + 1
		}
	}

	fmt.Println("Part 1:", floor)
	fmt.Println("Part 2:", basementReachedOn)
}
