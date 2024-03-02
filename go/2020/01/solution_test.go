package main

import "testing"

func TestFindSumEntries2(t *testing.T) {
	tests := []struct {
		input     []int
		target    int
		expected1 int
		expected2 int
	}{
		{[]int{1721, 979, 366, 299, 675, 1456}, 2020, 299, 1721},
	}

	for _, test := range tests {
		num1, num2 := findSumEntries2(test.input, test.target)
		if num1 != test.expected1 || num2 != test.expected2 {
			t.Errorf("findSum(%v, %v) = %v, %v, want %v, %v", test.input, test.target, num1, num2, test.expected1, test.expected2)
		}
	}
}

func TestFindSumEntries3(t *testing.T) {
	tests := []struct {
		input     []int
		target    int
		expected1 int
		expected2 int
		expected3 int
	}{
		{[]int{1721, 979, 366, 299, 675, 1456}, 2020, 979, 675, 366},
	}

	for _, test := range tests {
		num1, num2, num3 := findSumEntries3(test.input, test.target)
		if num1 != test.expected1 || num2 != test.expected2 || num3 != test.expected3 {
			t.Errorf("findSum(%v, %v) = %v, %v, %v, want %v, %v, %v", test.input, test.target, num1, num2, num3, test.expected1, test.expected2, test.expected3)
		}
	}
}
