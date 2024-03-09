package main

import (
	"fmt"
	"sort"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2021/10/input.txt"

var pointsPart1 = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var pointsPart2 = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

var openToClose = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

var closeToOpen = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
}

func main() {
	input := utils.ProcessStringLinesFile(inputFile)
	part1Score := 0
	part2Scores := make([]int, 0)

	for _, line := range input {
		isValid, invalidChar, remainingChars := isValidLine(line)
		if !isValid {
			part1Score += pointsPart1[invalidChar]
		} else if len(remainingChars) != 0 {
			part2LineScore := 0
			for len(remainingChars) > 0 {
				part2LineScore *= 5
				remainingChar, _ := remainingChars.Pop()
				part2LineScore += pointsPart2[openToClose[remainingChar]]
			}

			part2Scores = append(part2Scores, part2LineScore)
		}
	}

	fmt.Println("Part 1 Score:", part1Score)

	sort.Ints(part2Scores)
	middle := len(part2Scores) / 2
	fmt.Println("Part 2 Score:", part2Scores[middle])
}

func isValidLine(input string) (isValid bool, invalidChars string, remainingChars utils.Stack) {
	s := utils.Stack{}

	for _, c := range input {
		if _, ok := openToClose[string(c)]; ok {
			s.Push(string(c))
		} else {
			open, _ := s.Pop()
			if closeToOpen[string(c)] != open {
				return false, string(c), s
			}
		}
	}

	return true, "", s
}
