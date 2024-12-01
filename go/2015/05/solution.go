package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2015/05/input.txt"
var sampleInputFile = "../../../input-files/2015/05/input-sample.txt"
var isSampleMode = false

func main() {
	flag.BoolVar(&isSampleMode, "sample", false, "Use the sample input")
	flag.Parse()

	file := inputFile
	if isSampleMode {
		file = sampleInputFile
	}

	input := utils.ProcessStringLinesFile(file)
	niceStrings1, niceStrings2 := countNiceStrings(input)

	fmt.Println("Part 1:", niceStrings1)
	fmt.Println("Part 2:", niceStrings2)
}

func countNiceStrings(input []string) (int, int) {
	niceStrings1 := 0
	niceStrings2 := 0

	for _, line := range input {
		if isNiceString(line) {
			niceStrings1++
		}
		if isNiceString2(line) {
			niceStrings2++
		}
	}

	return niceStrings1, niceStrings2
}

func isNiceString(line string) bool {
	return hasThreeVowels(line) && hasDoubleLetter(line) && !hasForbiddenStrings(line)
}

func hasThreeVowels(line string) bool {
	vowels := "aeiou"
	count := 0

	for _, c := range line {
		if strings.ContainsRune(vowels, c) {
			count++
		}
	}

	return count >= 3
}

func hasDoubleLetter(line string) bool {
	for i := 0; i < len(line)-1; i++ {
		if line[i] == line[i+1] {
			return true
		}
	}

	return false
}

func hasForbiddenStrings(line string) bool {
	forbiddenStrings := []string{"ab", "cd", "pq", "xy"}

	for _, s := range forbiddenStrings {
		if strings.Contains(line, s) {
			return true
		}
	}

	return false
}

func isNiceString2(line string) bool {
	return hasPairOfLettersTwice(line) && hasRepeatLetterWithLetterInBetween(line)
}

func hasPairOfLettersTwice(line string) bool {
	for i := 0; i < len(line)-1; i++ {
		pair := line[i : i+2]
		if strings.Count(line, pair) >= 2 {
			return true
		}
	}

	return false
}

func hasRepeatLetterWithLetterInBetween(line string) bool {
	for i := 0; i < len(line)-2; i++ {
		if line[i] == line[i+2] {
			return true
		}
	}

	return false
}
