package main

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 1, 1}, 3},
		{[]int{1, 1, -2}, 0},
		{[]int{-1, -2, -3}, -6},
	}

	for _, test := range tests {
		result := sum(test.input)
		if result != test.expected {
			t.Errorf("%v -> expected %d but got %d", test.input, test.expected, result)
		}
	}
}

func TestSeenTwice(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, -1}, 0},
		{[]int{3, 3, 4, -2, -4}, 10},
		{[]int{-6, 3, 8, 5, -6}, 5},
		{[]int{7, 7, -2, -7, -4}, 14},
	}

	for _, test := range tests {
		result := seenTwice(test.input)
		if result != test.expected {
			t.Errorf("%v -> expected %d but got %d", test.input, test.expected, result)
		}
	}
}
