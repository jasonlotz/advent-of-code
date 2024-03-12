package main

import (
	"fmt"
	"regexp"
	"strconv"
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
	part2()
}

func part1() {
	input := getInput()

	splitInput := strings.Split(input, "\n\n")

	count := 0

	for _, rawPassport := range splitInput {
		passport := parsePassport(rawPassport)
		if passport.hasRequiredFields() {
			count++
		}
	}

	fmt.Println("Part 1:", count)
}

func part2() {
	input := getInput()

	splitInput := strings.Split(input, "\n\n")

	count := 0

	for _, rawPassport := range splitInput {
		passport := parsePassport(rawPassport)
		if passport.hasRequiredFields() && passport.validateValues() {
			count++
		}
	}

	fmt.Println("Part 2:", count)
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

func (p Passport) hasRequiredFields() bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, field := range requiredFields {
		if _, ok := p[field]; !ok {
			return false
		}
	}

	return true
}

func (p Passport) validateValues() bool {
	for field, value := range p {
		if !isFieldValid(field, value) {
			return false
		}
	}

	return true
}

func isFieldValid(field string, value string) bool {
	switch field {
	case "byr":
		return isYearValid(value, 1920, 2002)
	case "iyr":
		return isYearValid(value, 2010, 2020)
	case "eyr":
		return isYearValid(value, 2020, 2030)
	case "hgt":
		return isHeightValid(value)
	case "hcl":
		return isHairColorValid(value)
	case "ecl":
		return isEyeColorValid(value)
	case "pid":
		return isPassportIdValid(value)
	case "cid":
		return true
	default:
		return false
	}
}

func isYearValid(value string, min int, max int) bool {
	pattern := regexp.MustCompile(`^(\d{4})$`)

	if match := pattern.FindStringSubmatch(value); match != nil {
		year, _ := strconv.Atoi(match[1])
		return year >= min && year <= max
	}

	return false
}

func isHeightValid(value string) bool {
	cmPattern := regexp.MustCompile(`^(\d+)cm$`)
	inPattern := regexp.MustCompile(`^(\d+)in$`)

	if cmMatch := cmPattern.FindStringSubmatch(value); cmMatch != nil {
		height, _ := strconv.Atoi(cmMatch[1])
		return height >= 150 && height <= 193
	} else if inMatch := inPattern.FindStringSubmatch(value); inMatch != nil {
		height, _ := strconv.Atoi(inMatch[1])
		return height >= 59 && height <= 76
	}

	return false
}

func isHairColorValid(value string) bool {
	match, _ := regexp.MatchString(`^#[0-9a-f]{6}$`, value)
	return match
}

func isEyeColorValid(value string) bool {
	match, _ := regexp.MatchString(`^(amb|blu|brn|gry|grn|hzl|oth)$`, value)
	return match
}

func isPassportIdValid(value string) bool {
	match, _ := regexp.MatchString(`^\d{9}$`, value)
	return match
}
