package main

import (
	"flag"
	"fmt"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/YEAR/DAY/input.txt"
var testInputFile = "../../../input-files/YEAR/DAY/input-sample.txt"
var isTestMode = false

func getInput() []string {
	flag.BoolVar(&isTestMode, "test", false, "Use the test input")
	flag.Parse()

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
