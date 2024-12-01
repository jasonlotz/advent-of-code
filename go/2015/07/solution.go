package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2015/07/input.txt"
var sampleInputFile = "../../../input-files/2015/07/input-sample.txt"
var isSampleMode = false

type Instruction struct {
	action string
	inputA string
	inputB string
}

type Wire struct {
	instructions []Instruction
	value        uint16
	evaluated    bool
}

type WireMap map[string]*Wire

func getInput() []string {
	file := inputFile

	if isSampleMode {
		file = sampleInputFile
	}

	return utils.ProcessStringLinesFile(file)
}

func main() {
	input := getInput()

	wireMap := parseInstructions(input)

	for _, wire := range wireMap {
		wire.signalValue(&wireMap)
	}

	fmt.Println("Part 1:", wireMap["a"].value)

	wireMap["b"].value = wireMap["a"].value

	for label, wire := range wireMap {
		if label == "b" {
			continue
		}
		wire.evaluated = false
		wire.value = 0
	}

	for _, wire := range wireMap {
		wire.signalValue(&wireMap)
	}

	fmt.Println("Part 2:", wireMap["a"].value)
}

func parseInstructions(input []string) WireMap {
	wireMap := make(map[string]*Wire)

	for _, line := range input {
		words := strings.Split(line, " ")
		wireLabel := ""

		instruction := Instruction{}

		if findWord("NOT", words) {
			instruction.action = "NOT"
			instruction.inputA = words[1]
			wireLabel = words[3]
		} else if findWord("AND", words) {
			instruction.action = "AND"
			instruction.inputA = words[0]
			instruction.inputB = words[2]
			wireLabel = words[4]
		} else if findWord("OR", words) {
			instruction.action = "OR"
			instruction.inputA = words[0]
			instruction.inputB = words[2]
			wireLabel = words[4]
		} else if findWord("LSHIFT", words) {
			instruction.action = "LSHIFT"
			instruction.inputA = words[0]
			instruction.inputB = words[2]
			wireLabel = words[4]
		} else if findWord("RSHIFT", words) {
			instruction.action = "RSHIFT"
			instruction.inputA = words[0]
			instruction.inputB = words[2]
			wireLabel = words[4]
		} else {
			instruction.action = "ASSIGN"
			instruction.inputA = words[0]
			wireLabel = words[2]
		}

		if _, ok := wireMap[wireLabel]; !ok {
			wireMap[wireLabel] = &Wire{}
		}

		wire := *wireMap[wireLabel]
		wire.instructions = append(wire.instructions, instruction)
		wireMap[wireLabel] = &wire
	}

	return wireMap
}

func findWord(word string, words []string) bool {
	for _, w := range words {
		if w == word {
			return true
		}
	}

	return false
}

func (w *Wire) signalValue(wireMap *WireMap) {
	if w.evaluated {
		return
	}

	for _, inst := range w.instructions {
		w.value = processInstruction(inst, wireMap)
	}

	w.evaluated = true
}

func processInstruction(inst Instruction, wireMap *WireMap) uint16 {
	if inst.action == "ASSIGN" {
		return getValue(inst.inputA, wireMap)
	} else if inst.action == "AND" {
		return getValue(inst.inputA, wireMap) & getValue(inst.inputB, wireMap)
	} else if inst.action == "OR" {
		return getValue(inst.inputA, wireMap) | getValue(inst.inputB, wireMap)
	} else if inst.action == "LSHIFT" {
		return getValue(inst.inputA, wireMap) << getValue(inst.inputB, wireMap)
	} else if inst.action == "RSHIFT" {
		return getValue(inst.inputA, wireMap) >> getValue(inst.inputB, wireMap)
	} else if inst.action == "NOT" {
		return ^getValue(inst.inputA, wireMap)
	}

	panic("Invalid instruction")
}

func getValue(input string, wireMap *WireMap) uint16 {
	if val, err := strconv.Atoi(input); err == nil {
		return uint16(val)
	}

	wire := (*wireMap)[input]
	wire.signalValue(wireMap)
	return wire.value
}
