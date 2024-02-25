package utils

import (
	"bufio"
	"os"
)

func ProcessSingleStringLineFile(inputFile string) string {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

func ProcessStringLinesFile(inputFile string) []string {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	processedFile := make([]string, 0)

	for scanner.Scan() {
		processedFile = append(processedFile, scanner.Text())
	}

	return processedFile
}
