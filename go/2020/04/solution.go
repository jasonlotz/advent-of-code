package main

import (
	"fmt"
	"strings"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2020/04/input.txt"
var testInputFile = "../../../input-files/2020/04/input-sample.txt"
var isTestMode = false

type Passport map[string]string

func getInput() string {
	file := inputFile

	if isTestMode {
		file = testInputFile
	}

	return utils.ProcessFullFile(file)
}

func main() {
	part1()
}

func part1() {
	input := getInput()

	splitInput := strings.Split(input, "\n\n")

	count := 0

	for _, rawPassport := range splitInput {
		passport := parsePassport(rawPassport)
		if passport.isValid() {
			count++
		}
	}

	fmt.Println("Part 1:", count)
}

func parsePassport(rawPassport string) Passport {
	passport := make(Passport)

	splitPassport := strings.Fields(rawPassport)

	for _, field := range splitPassport {
		pair := strings.Split(field, ":")
		passport[pair[0]] = pair[1]
	}

	return passport
}

func (p Passport) isValid() bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, field := range requiredFields {
		if _, ok := p[field]; !ok {
			return false
		}
	}

	return true
}
