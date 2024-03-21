package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

var input = "ugkcyxxp"
var testInput = "abc3231929"
var isTestMode = false

func main() {
	part1()
	part2()
}

func part1() {
	if isTestMode {
		input = testInput
	}

	password := generatePassword(input)
	fmt.Println("Part 1:", password)
}

func part2() {
	if isTestMode {
		input = testInput
	}

	password := generatePassword2(input)
	fmt.Println("Part 2:", password)
}

func generateMd5(input string) string {
	h := md5.New()
	io.WriteString(h, input)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func generatePassword(prefix string) string {
	password := ""
	for i := 0; len(password) < 8; i++ {
		hash := generateMd5(prefix + fmt.Sprint(i))
		if hash[:5] == "00000" {
			fmt.Println("Found a valid hash:", i, hash)
			password += string(hash[5])
		}
	}
	return password
}

func generatePassword2(prefix string) string {
	password := "________"
	for i := 0; strings.Contains(password, "_"); i++ {
		hash := generateMd5(prefix + fmt.Sprint(i))
		if hash[:5] == "00000" {
			position := int(hash[5]) - 48 // convert ascii to int
			if position < 8 && password[position] == '_' {
				password = password[:position] + string(hash[6]) + password[position+1:]
				fmt.Println("Found a valid hash:", i, hash, password)
			}
		}
	}
	return password
}
