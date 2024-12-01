package main

import (
	"flag"
	"fmt"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/yyyy/dd/input.txt"
var sampleInputFile = "../../../input-files/yyyy/dd/input-sample.txt"
var isSampleMode = false

func getInput() []string {
	flag.BoolVar(&isSampleMode, "sample", false, "Use the sample input")
	flag.Parse()

	file := inputFile
	if isSampleMode {
		file = sampleInputFile
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
