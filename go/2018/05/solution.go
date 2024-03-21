package main

import (
	"fmt"
	"strings"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2018/05/input.txt"
var testInputFile = "../../../input-files/2018/05/input-sample.txt"
var isTestMode = true

func getInput() string {
	file := inputFile

	if isTestMode {
		file = testInputFile
	}

	return utils.ProcessSingleStringLineFile(file)
}

func main() {
	part1()
	part2()
}

func part1() {
	input := getInput()

	reactedPolymer := reactPolymer(input)
	if isTestMode {
		fmt.Println("Final polymer:", reactedPolymer)
	}
	fmt.Println("Part 1:", len(reactedPolymer))
}

func part2() {
	input := getInput()

	shortestPolymer := len(input)
	for i := 65; i <= 90; i++ {
		letterToReplace := fmt.Sprintf("%c", i)
		reactedPolymer := reactPolymer(strings.Replace(input, letterToReplace, "", -1))
		reactedPolymer = reactPolymer(strings.Replace(reactedPolymer, strings.ToLower(letterToReplace), "", -1))

		length := len(reactedPolymer)
		fmt.Println("Replaced", letterToReplace, "Length:", length)

		if length < shortestPolymer {
			shortestPolymer = len(reactedPolymer)
		}

		if isTestMode {
			fmt.Println("Final polymer:", reactedPolymer)
		}
	}

	fmt.Println("Part 2:", shortestPolymer)
}

func reactPolymer(polymer string) string {
	reacted := true

	for reacted {
		reacted = false
		for i := 0; i < len(polymer)-1; i++ {
			if polymer[i] != polymer[i+1] && strings.ToLower(string(polymer[i])) == strings.ToLower(string(polymer[i+1])) {
				reacted = true
				polymer = polymer[:i] + polymer[i+2:]
				break
			}
		}
	}

	return polymer
}
