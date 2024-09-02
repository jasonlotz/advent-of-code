package main

import (
	"fmt"
	"strings"
)

var input = "3113322113"

func main() {
	for i := 0; i < 50; i++ {
		input = updateSequence(input)
		fmt.Println(i+1, len(input))
	}
}

func updateSequence(input string) string {
	var stringBuilder strings.Builder
	var count int
	var currentLetter rune

	for i, c := range input {
		if i == 0 {
			currentLetter = c
			count = 1
			continue
		}

		if c == currentLetter {
			count++
		} else {
			stringBuilder.WriteString(fmt.Sprintf("%d%c", count, currentLetter))
			currentLetter = c
			count = 1
		}
	}

	stringBuilder.WriteString(fmt.Sprintf("%d%c", count, currentLetter))

	return stringBuilder.String()
}
