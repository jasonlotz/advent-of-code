package main

import (
	"testing"
)

func TestIsValidBasicPassphrase(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"aa bb cc dd ee", true},
		{"aa bb cc dd aa", false},
		{"aa bb cc dd aaa", true},
	}

	for _, test := range tests {
		result := isValidBasicPassphrase(test.input)
		if result != test.expected {
			t.Errorf("%s -> expected %t but got %t", test.input, test.expected, result)
		}
	}
}

func TestIsValidAnagramPassphrase(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abcde fghij", true},
		{"abcde xyz ecdab", false},
		{"a ab abc abd abf abj", true},
		{"iiii oiii ooii oooi oooo", true},
		{"oiii ioii iioi iiio", false},
	}

	for _, test := range tests {
		result := isValidAnagramPassphrase(test.input)
		if result != test.expected {
			t.Errorf("%s -> expected %t but got %t", test.input, test.expected, result)
		}
	}
}
