package main

import (
	"fmt"
	"strings"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2022/05/input.txt"
var sampleInputFile = "../../../input-files/2022/05/input-sample.txt"
var isSampleMode = false

type instruction struct {
	qty, from, to int
}

func (i instruction) String() string {
	return fmt.Sprintf("Move %d from %d to %d", i.qty, i.from, i.to)
}

func getInput() (rawState string, rawInstructions string) {
	file := inputFile

	if isSampleMode {
		file = sampleInputFile
	}

	input := strings.TrimRight(utils.ProcessFullFile(file), "\n")
	splitInput := strings.Split(input, "\n\n")

	return splitInput[0], splitInput[1]
}

func main() {
	part1()
	part2()
}

func part1() {
	rawState, rawInstructions := getInput()

	state := parseInitialState(rawState)
	instructions := parseInstructions(rawInstructions)

	followInstructions(&state, instructions)

	result := readTopOfStacks(state)

	fmt.Println("Part 1:", result)
}

func part2() {
	rawState, rawInstructions := getInput()

	state := parseInitialState(rawState)
	instructions := parseInstructions(rawInstructions)

	followUpdatedInstructions(&state, instructions)

	result := readTopOfStacks(state)

	fmt.Println("Part 2:", result)
}

func parseInitialState(rawState string) []utils.Stack {
	state := rawState

	formatted := [][]string{}
	for _, row := range strings.Split(state, "\n") {
		formatted = append(formatted, strings.Split(row, ""))
	}
	oRows, oCols := len(formatted), len(formatted[0])

	actual := []utils.Stack{}
	for c := 0; c < oCols-1; c++ {
		if formatted[oRows-1][c] != " " {
			// Found a column with values, move up the column and add to the actual stack
			stack := []string{}
			for r := oRows - 2; r >= 0; r-- {
				char := formatted[r][c]
				if char != " " {
					stack = append(stack, char)
				}
			}
			actual = append(actual, stack)
		}
	}

	return actual
}

func parseInstructions(rawInstructions string) []instruction {
	instructions := []instruction{}

	for _, row := range strings.Split(rawInstructions, "\n") {
		inst := instruction{}
		_, err := fmt.Sscanf(row, "move %d from %d to %d", &inst.qty, &inst.from, &inst.to)
		if err != nil {
			panic(err)
		}
		// Make the instructions 0-indexed
		inst.from--
		inst.to--
		instructions = append(instructions, inst)
	}

	return instructions
}

func followInstructions(state *[]utils.Stack, instructions []instruction) {
	for _, inst := range instructions {
		for i := 0; i < inst.qty; i++ {
			val, ok := (*state)[inst.from].Pop()
			if !ok {
				panic("No value to move")
			}

			(*state)[inst.to].Push(val)
		}
	}
}

func followUpdatedInstructions(state *[]utils.Stack, instructions []instruction) {
	for _, inst := range instructions {
		vals, ok := (*state)[inst.from].PopMultiple(inst.qty)
		if !ok {
			panic("No value to move")
		}

		for _, val := range vals {
			(*state)[inst.to].Push(val)
		}
	}
}

func readTopOfStacks(state []utils.Stack) string {
	result := ""
	for _, stack := range state {
		val, ok := stack.Peek()
		if !ok {
			panic("No value to read from stack")
		}
		result += val
	}
	return result
}
