package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

func ProcessIntLinesFile(inputFile string) []int {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	processedFile := make([]int, 0)

	for scanner.Scan() {
		intVal, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		processedFile = append(processedFile, intVal)
	}

	return processedFile
}
func ProcessIntsLinesFile(inputFile string) [][]int {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	processedFile := make([][]int, 0)

	for scanner.Scan() {
		stringVals := strings.Fields(scanner.Text())
		intVals := make([]int, len(stringVals))

		for i, s := range stringVals {
			intVal, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}

			intVals[i] = intVal
		}
		processedFile = append(processedFile, intVals)
	}

	return processedFile
}