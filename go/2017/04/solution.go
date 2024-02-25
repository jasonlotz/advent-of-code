package main

import (
	"bufio"
	"fmt"
	"github.com/jasonlotz/advent-of-code/go/utils"
	"os"
	"sort"
	"strings"
)

var INPUT_FILE = "../../../input-files/2017/04/input.txt"

func main() {
	passphrases := utils.ProcessStringLinesFile(INPUT_FILE)

	validBasicPassphrases := 0
	validAnagramPassphrases := 0

	for _, line := range passphrases {
		isValidBasic := isValidBasicPassphrase(line)
		if isValidBasic {
			validBasicPassphrases++
		}
		isValidAnagram := isValidAnagramPassphrase(line)
		if isValidAnagram {
			validAnagramPassphrases++
		}
	}
	fmt.Println("Valid basic passphrases:", validBasicPassphrases)
	fmt.Println("Valid anagram passphrases:", validAnagramPassphrases)
}

func processFile(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	processedFile := make([]string, 0)

	for scanner.Scan() {
		processedFile = append(processedFile, scanner.Text())
	}

	return processedFile
}

func isValidBasicPassphrase(passphrase string) bool {
	seen := make(map[string]bool)
	words := strings.Split(passphrase, " ")

	for _, word := range words {
		if seen[word] {
			return false
		}
		seen[word] = true
	}

	return true
}

func isValidAnagramPassphrase(passphrase string) bool {
	seen := make(map[string]bool)
	words := strings.Split(passphrase, " ")

	for _, word := range words {
		word = sortString(word)
		if seen[word] {
			return false
		}
		seen[word] = true
	}

	return true
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
