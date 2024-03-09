package main

import (
	"fmt"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2019/02/input.txt"

func main() {
	part1()
	part2()
}

func part1() {
	input := utils.ProcessSingleStringLineFile(inputFile)
	opCodes := utils.CommaDelimitedStringToIntArray(input)
	opCodes[1] = 12
	opCodes[2] = 2
	processOpCodes(opCodes)
	fmt.Println("Part 1:", opCodes[0])
}

func part2() {
	input := utils.ProcessSingleStringLineFile(inputFile)
	opCodes := utils.CommaDelimitedStringToIntArray(input)
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			opCodesCopy := make([]int, len(opCodes))
			copy(opCodesCopy, opCodes)
			opCodesCopy[1] = noun
			opCodesCopy[2] = verb
			processOpCodes(opCodesCopy)
			if opCodesCopy[0] == 19690720 {
				fmt.Println("Part 2:", 100*noun+verb)
				return
			}
		}
	}
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
