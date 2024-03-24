package main

import (
	"fmt"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2016/07/input.txt"
var testInputFile = "../../../input-files/2016/07/input-sample.txt"
var isTestMode = false

func getInput() []string {
	file := inputFile

	if isTestMode {
		file = testInputFile
	}

	return utils.ProcessStringLinesFile(file)
}

func main() {
	input := getInput()

	tlsCount := 0
	slsCount := 0

	for _, line := range input {
		fmt.Println("Line:", line)
		supernetSequence, hypernetSequence := parseInput(line)
		fmt.Println("Supernet:", supernetSequence)
		fmt.Println("Hypernet:", hypernetSequence)

		if isTls(line) {
			tlsCount++
		}

		if isSsl(line) {
			slsCount++
		}
	}

	fmt.Println("Part 1:", tlsCount)
	fmt.Println("Part 2:", slsCount)
}

func parseInput(ip string) (supernetSequence string, hypernetSequence string) {
	isHypernet := false

	for i := 0; i < len(ip); i++ {
		if ip[i] == '[' {
			supernetSequence += " "
			isHypernet = true
		} else if ip[i] == ']' {
			hypernetSequence += " "
			isHypernet = false
		} else {
			if isHypernet {
				hypernetSequence += string(ip[i])
			} else {
				supernetSequence += string(ip[i])
			}
		}
	}

	return supernetSequence, hypernetSequence
}

func hasAbba(sequence string) bool {
	for i := 0; i < len(sequence)-3; i++ {
		if sequence[i] == sequence[i+3] && sequence[i+1] == sequence[i+2] && sequence[i] != sequence[i+1] {
			return true
		}
	}

	return false
}

func isTls(ip string) bool {
	supernetSequence, hypernetSequence := parseInput(ip)

	if hasAbba(supernetSequence) && !hasAbba(hypernetSequence) {
		return true
	}

	return false
}

func getAbas(sequence string) []string {
	abas := []string{}

	for i := 0; i < len(sequence)-2; i++ {
		if sequence[i] == sequence[i+2] && sequence[i] != sequence[i+1] {
			abas = append(abas, sequence[i:i+3])
		}
	}

	return abas
}

func getBab(aba string) string {
	return string(aba[1]) + string(aba[0]) + string(aba[1])
}

func isSsl(ip string) bool {
	supernetSequence, hypernetSequence := parseInput(ip)

	supernetAbas := []string{}
	hypernetAbas := []string{}

	supernetAbas = append(supernetAbas, getAbas(supernetSequence)...)
	hypernetAbas = append(hypernetAbas, getAbas(hypernetSequence)...)

	for _, aba := range supernetAbas {
		bab := getBab(aba)

		for _, hypernetAba := range hypernetAbas {
			if bab == hypernetAba {
				return true
			}
		}
	}

	return false
}
