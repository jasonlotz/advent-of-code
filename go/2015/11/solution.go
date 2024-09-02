package main

import (
	"fmt"
	"strings"
)

var input = "vzbxkghb"

func main() {
	fmt.Println("Part 1:", nextPassword(input))
	fmt.Println("Part 2:", nextPassword(nextPassword(input)))
}

func nextPassword(input string) string {
	for {
		input = increment(input)
		if isValid(input) {
			return input
		}
	}
}

func increment(input string) string {
	bytes := []byte(input)
	for i := len(bytes) - 1; i >= 0; i-- {
		if bytes[i] == 'z' {
			bytes[i] = 'a'
		} else {
			bytes[i]++
			break
		}
	}
	return string(bytes)
}

func isValid(input string) bool {
	return hasIncreasingStraight(input) && !strings.ContainsAny(input, "iol") && hasTwoPairs(input)
}

func hasIncreasingStraight(input string) bool {
	for i := 0; i < len(input)-2; i++ {
		if input[i]+1 == input[i+1] && input[i]+2 == input[i+2] {
			return true
		}
	}
	return false
}

func hasTwoPairs(input string) bool {
	pairs := 0
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			pairs++
			i++
		}
	}
	return pairs >= 2
}
