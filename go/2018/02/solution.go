package main

import (
	"fmt"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var INPUT_FILE = "../../../input-files/2018/02/input.txt"

func main() {
	part1()
	part2()
}

func part1() {
	values := utils.ProcessStringLinesFile(INPUT_FILE)
	found2Count := 0
	found3Count := 0

	for _, value := range values {
		letterCount := countLetters(value)

		found2 := false
		found3 := false

		for _, count := range letterCount {
			if !found2 && count == 2 {
				found2 = true
			}
			if !found3 && count == 3 {
				found3 = true
			}

			if found2 && found3 {
				break
			}
		}

		if found2 {
			found2Count++
		}

		if found3 {
			found3Count++
		}
	}

	fmt.Println("Part 1: ", found2Count*found3Count)
}

func part2() {
	values := utils.ProcessStringLinesFile(INPUT_FILE)
	oneOffMatch := findOneOffMatch(values)
	fmt.Println("Part 2: ", oneOffMatch)
}

func countLetters(word string) map[string]int {
	letterCount := make(map[string]int)
	for _, letter := range word {
		letterCount[string(letter)]++
	}
	return letterCount
}

func findOneOffMatch(words []string) string {
	seen := make(map[string]bool)

	for _, word := range words {
		for i := 0; i < len(word); i++ {
			// slice indices are inclusive:exclusive and can exceed the length of the array
			newWord := word[:i] + "?" + word[i+1:]
			if seen[newWord] {
				return newWord
			}
			seen[newWord] = true
		}
	}

	return ""
}
