package utils

import (
	"strconv"
	"strings"
)

func CommaDelimitedStringToIntArray(input string) []int {
	splitInput := strings.Split(input, ",")
	intSlice := make([]int, len(splitInput))

	for i, v := range splitInput {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		intSlice[i] = num
	}

	return intSlice
}
