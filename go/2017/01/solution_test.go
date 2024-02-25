package main

import (
	"testing"
)

func TestSumCaptchaNext(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1122", 3},
		{"1111", 4},
		{"1234", 0},
		{"91212129", 9},
	}

	for _, test := range tests {
		result := sumCaptchaNext(test.input)
		if result != test.expected {
			t.Errorf("%s -> expected %d but got %d", test.input, test.expected, result)
		}
	}
}

func TestSumCaptchaHalfway(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1212", 6},
		{"1221", 0},
		{"123425", 4},
		{"123123", 12},
		{"12131415", 4},
	}

	for _, test := range tests {
		result := sumCaptchaHalfway(test.input)
		if result != test.expected {
			t.Errorf("%s -> expected %d but got %d", test.input, test.expected, result)
		}
	}
}
