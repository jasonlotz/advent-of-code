package main

import "testing"

func TestIsValidate(t *testing.T) {
	tests := []struct {
		line                string
		expectedIsValid     bool
		expectedInvalidChar string
	}{
		{"{([(<{}[<>[]}>{[]{[(<()>", false, "}"},
		{"[[<[([]))<([[{}[[()]]]", false, ")"},
		{"[{[{({}]{}}([{[{{{}}([]", false, "]"},
		{"[<(<(<(<{}))><([]([]()", false, ")"},
		{"<{([([[(<>()){}]>(<<{{", false, ">"},
	}

	for _, test := range tests {
		isValid, invalidChar, _ := isValidLine(test.line)
		if isValid != test.expectedIsValid || invalidChar != test.expectedInvalidChar {
			t.Error("Expected", test.expectedIsValid, test.expectedInvalidChar, "got", isValid, invalidChar)
		}
	}
}
