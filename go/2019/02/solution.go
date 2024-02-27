package main

import (
	"fmt"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var INPUT_FILE = "../../../input-files/2019/02/input.txt"

func main() {
	part1()
}

func part1() {
	input := utils.ProcessSingleStringLineFile(INPUT_FILE)
	opCodes := utils.CommaDelimitedStringToIntArray(input)
	opCodes[1] = 12
	opCodes[2] = 2
	processOpCodes(opCodes)
	fmt.Println("Part 1:", opCodes[0])
}

func processOpCodes(opCodes []int) {
	for i := 0; i < len(opCodes); i += 4 {
		switch opCodes[i] {
		case 1:
			opCodes[opCodes[i+3]] = opCodes[opCodes[i+1]] + opCodes[opCodes[i+2]]
		case 2:
			opCodes[opCodes[i+3]] = opCodes[opCodes[i+1]] * opCodes[opCodes[i+2]]
		case 99:
			return
		default:
			panic("Invalid OpCode")
		}
	}
}
