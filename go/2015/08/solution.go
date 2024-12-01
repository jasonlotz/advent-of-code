package main

import (
	"fmt"
	"strconv"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2015/08/input.txt"
var sampleInputFile = "../../../input-files/2015/08/input-sample.txt"
var isSampleMode = false

func getInput() []string {
	file := inputFile

	if isSampleMode {
		file = sampleInputFile
	}

	return utils.ProcessStringLinesFile(file)
}

func main() {
	part1()
	part2()
}

func part1() {
	input := getInput()
	difference := 0

	for _, line := range input {
		processedLine := decode(line)

		difference += len(line) - len(processedLine)
	}

	fmt.Println("Part 1: ", difference)
}

func part2() {
	input := getInput()
	difference := 0

	for _, line := range input {
		encodedLine := encode(line)

		difference += len(encodedLine) - len(line)
	}

	fmt.Println("Part 2: ", difference)
}

func decode(line string) []rune {
	var out []rune

	// Remove the surrounding quotes
	line = line[1 : len(line)-1]

	for i := 0; i < len(line); i++ {
		if line[i] == '\\' {
			if line[i+1] == 'x' {
				escInt, err := strconv.ParseInt(line[i+2:i+4], 16, 0)
				if err != nil {
					panic(err)
				}
				out = append(out, rune(escInt))
				i += 3
			} else {
				out = append(out, rune(line[i+1]))
				i++
			}
		} else {
			out = append(out, rune(line[i]))
		}

	}

	return out
}

func encode(line string) string {
	encoded := "\""

	for _, char := range line {
		if char == '"' || char == '\\' {
			encoded += "\\"
		}
		encoded += string(char)
	}

	encoded += "\""

	return encoded
}
