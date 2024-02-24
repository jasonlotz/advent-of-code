package main

import (
	"testing"
)

func TestSumCaptchaNext(t *testing.T) {
	input := "1122"
	expected := 3
	result := sumCaptchaNext(input)
	testResult(input, result, expected, t)

	input = "1111"
	expected = 4
	result = sumCaptchaNext(input)
	testResult(input, result, expected, t)

	input = "1234"
	expected = 0
	result = sumCaptchaNext(input)
	testResult(input, result, expected, t)

	input = "91212129"
	expected = 9
	result = sumCaptchaNext(input)
	testResult(input, result, expected, t)
}

func TestSumCaptchaHalfway(t *testing.T) {
	input := "1212"
	expected := 6
	result := sumCaptchaHalfway(input)
	testResult(input, result, expected, t)

	input = "1221"
	expected = 0
	result = sumCaptchaHalfway(input)
	testResult(input, result, expected, t)

	input = "123425"
	expected = 4
	result = sumCaptchaHalfway(input)
	testResult(input, result, expected, t)

	input = "123123"
	expected = 12
	result = sumCaptchaHalfway(input)
	testResult(input, result, expected, t)

	input = "12131415"
	expected = 4
	result = sumCaptchaHalfway(input)
	testResult(input, result, expected, t)
}

func testResult(input string, result int, expected int, t *testing.T) {
	if result != expected {
		t.Error(input, "-> expected result of", expected, "but got", result)
	}
}
