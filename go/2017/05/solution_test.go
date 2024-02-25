package main

import "testing"

func TestProcessJumpsPart1(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{0, 3, 0, 1, -3}, 5},
	}

	for _, test := range tests {
		result := processJumpsPart1(test.input)
		if result != test.expected {
			t.Errorf("%v -> expected %d for part 1 but got %d", test.input, test.expected, result)
		}
	}
}

func TestProcessJumpsPart2(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{0, 3, 0, 1, -3}, 10},
	}

	for _, test := range tests {
		result := processJumpsPart2(test.input)
		if result != test.expected {
			t.Errorf("%v -> expected %d for part 2 but got %d", test.input, test.expected, result)
		}
	}
}
