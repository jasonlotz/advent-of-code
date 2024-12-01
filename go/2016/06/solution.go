package main

import (
	"fmt"
	"math"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2016/06/input.txt"
var sampleInputFile = "../../../input-files/2016/06/input-sample.txt"
var isSampleMode = false

type LetterPositions []LetterMap
type LetterMap map[string]int

func getInput() []string {
	file := inputFile

	if isSampleMode {
		file = sampleInputFile
	}

	return utils.ProcessStringLinesFile(file)
}

func main() {
	input := getInput()

	letterPositions := countLetters(input)

	messageMostCommon, messageLeastCommon := findMessage(letterPositions)

	fmt.Println("Part 1:", messageMostCommon)
	fmt.Println("Part 2:", messageLeastCommon)
}

func countLetters(input []string) LetterPositions {
	letterPositions := make(LetterPositions, len(input[0]))

	for i := range letterPositions {
		letterCount := make(LetterMap)
		letterPositions[i] = letterCount
	}

	for _, line := range input {
		for i, letter := range line {
			letterPositions[i][string(letter)]++
		}
	}

	return letterPositions
}

func findMessage(letterPositions LetterPositions) (messageMostCommon string, messageLeastCommon string) {
	for _, letterCount := range letterPositions {
		maxLetter, minLetter := "", ""
		maxCount := 0
		minCount := math.MaxInt

		for letter, count := range letterCount {
			if count > maxCount {
				maxLetter = letter
				maxCount = count
			}

			if count < minCount {
				minLetter = letter
				minCount = count
			}
		}
		messageMostCommon += maxLetter
		messageLeastCommon += minLetter
	}

	return messageMostCommon, messageLeastCommon
}
