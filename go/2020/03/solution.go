package main

import (
	"fmt"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2020/03/input.txt"
var testInputFile = "../../../input-files/2020/03/input-sample.txt"
var isTestMode = false

func getInput() []string {
	file := inputFile

	if isTestMode {
		file = testInputFile
	}

	return utils.ProcessStringLinesFile(file)
}

func main() {
	part1()
	part2()
}

func part1() {
	input := getInput()
	treeCount := 0

	treeCount = countTrees(input, 3, 1)

	fmt.Println("Part 1:", treeCount)
}

func part2() {
	input := getInput()
	treeCount := 0

	treeCount = countTrees(input, 1, 1)
	treeCount *= countTrees(input, 3, 1)
	treeCount *= countTrees(input, 5, 1)
	treeCount *= countTrees(input, 7, 1)
	treeCount *= countTrees(input, 1, 2)

	fmt.Println("Part 2:", treeCount)
}

func countTrees(input []string, right, down int) int {
	treeCount := 0
	width := len(input[0])
	x := 0
	y := 0

	for y < len(input) {
		if input[y][x%width] == '#' {
			treeCount++
		}

		x += right
		y += down
	}

	return treeCount
}
