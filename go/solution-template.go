package main

import (
	"fmt"
	"github.com/jasonlotz/advent-of-code/go/utils"
)

var INPUT_FILE = "../../../input-files/{year}/{day}/input.txt"

func main() {
	part1()
}

func part1() {
	input := utils.ProcessStringLinesFile(INPUT_FILE)
	fmt.Println(input)
}
