package main

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		p        PasswordPolicy
		expected bool
	}{
		{PasswordPolicy{1, 3, 'a', "abcde"}, true},
		{PasswordPolicy{1, 3, 'b', "cdefg"}, false},
		{PasswordPolicy{2, 9, 'c', "ccccccccc"}, true},
	}
	for _, test := range tests {
		if test.p.validate() != test.expected {
			t.Error("Expected", test.expected, "got", !test.expected)
		}
	}
}

func TestValidate2(t *testing.T) {
	tests := []struct {
		p        PasswordPolicy
		expected bool
	}{
		{PasswordPolicy{1, 3, 'a', "abcde"}, true},
		{PasswordPolicy{1, 3, 'b', "cdefg"}, false},
		{PasswordPolicy{2, 9, 'c', "ccccccccc"}, false},
	}
	for _, test := range tests {
		if test.p.validate2() != test.expected {
			t.Error("Expected", test.expected, "got", !test.expected)
		}
	}
}
