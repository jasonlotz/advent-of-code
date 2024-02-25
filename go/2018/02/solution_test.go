package main

import (
	"github.com/jasonlotz/advent-of-code/go/utils"
	"testing"
)

func TestCountLetters(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{"abcdef", map[string]int{"a": 1, "b": 1, "c": 1, "d": 1, "e": 1, "f": 1}},
		{"bababc", map[string]int{"a": 2, "b": 3, "c": 1}},
		{"abbcde", map[string]int{"a": 1, "b": 2, "c": 1, "d": 1, "e": 1}},
		{"abcccd", map[string]int{"a": 1, "b": 1, "c": 3, "d": 1}},
		{"aabcdd", map[string]int{"a": 2, "b": 1, "c": 1, "d": 2}},
		{"abcdee", map[string]int{"a": 1, "b": 1, "c": 1, "d": 1, "e": 2}},
		{"ababab", map[string]int{"a": 3, "b": 3}},
	}

	for _, test := range tests {
		result := countLetters(test.input)
		if !utils.CompareMaps(result, test.expected) {
			t.Errorf("%s -> expected %v but got %v", test.input, test.expected, result)
		}
	}
}
