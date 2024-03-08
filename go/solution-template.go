package main

import (
	"fmt"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/YEAR/DAY/input.txt"
var testInputFile = "../../../input-files/YEAR/DAY/input-sample.txt"
var isTestMode = true

func getInput() []string {
	file := inputFile

	if isTestMode {
		file = testInputFile
	}

	return utils.ProcessStringLinesFile(file)
}

func main() {
	part1()
}

func part1() {
	input := getInput()
	fmt.Println(input)
}
