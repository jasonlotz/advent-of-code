package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var INPUT_FILE = "../../../input-files/2016/04/input.txt"

type Room struct {
	name     string
	sectorId int
	checksum string
}

func main() {
	input := utils.ProcessStringLinesFile(INPUT_FILE)
	rooms := processInput(input)
	invalidSectorIdSum := 0

	for _, room := range rooms {
		if room.isValid() {
			invalidSectorIdSum += room.sectorId
			if strings.Contains(room.decrypt(), "north") {
				fmt.Println(room.decrypt(), room.sectorId)
			}
		}
	}

	fmt.Println("Part 1 (invalidSectorIdSum):", invalidSectorIdSum)
}

func processInput(input []string) []Room {
	rooms := make([]Room, len(input))

	for i, line := range input {
		rooms[i] = createRoom(line)
	}

	return rooms
}

func createRoom(s string) Room {
	// name
	re := regexp.MustCompile(`[a-z]+-`)
	matches := re.FindAllString(s, -1)
	name := ""
	for _, match := range matches {
		name += strings.Replace(match, "-", "", -1)
	}

	// sectorId
	re = regexp.MustCompile(`\d+`)
	matches = re.FindAllString(s, -1)
	sectorId, err := strconv.Atoi(matches[0])
	if err != nil {
		panic(err)
	}

	// checksum
	checksum := ""
	re = regexp.MustCompile(`\[[a-z]+\]`)
	matches = re.FindAllString(s, -1)
	for _, match := range matches {
		checksum += match
	}
	checksum = checksum[1 : len(checksum)-1]

	return Room{name, sectorId, checksum}
}

func (r Room) countChars() map[rune]int {
	counts := make(map[rune]int)

	for _, char := range r.name {
		counts[char]++
	}

	return counts
}

func (r Room) isValid() bool {
	counts := r.countChars()

	sortedCharCountArr := sortCharCountMap(counts)

	// take the top 5 of the sorted list and compare to the checksum
	for i, char := range r.checksum {
		if char != sortedCharCountArr[i].char {
			return false
		}
	}

	return true
}

type CharCount struct {
	char  rune
	count int
}

func sortCharCountMap(m map[rune]int) []CharCount {
	var charCountArr []CharCount
	for k, v := range m {
		charCountArr = append(charCountArr, CharCount{k, v})
	}

	sort.Slice(charCountArr, func(i, j int) bool {
		if charCountArr[i].count == charCountArr[j].count {
			return charCountArr[i].char < charCountArr[j].char // sort by rune in ascending order in case of tie
		}
		return charCountArr[i].count > charCountArr[j].count // sort by value in descending order
	})

	return charCountArr
}

func (r Room) decrypt() string {
	decrypted := ""
	for _, char := range r.name {
		decrypted += string(shiftChar(char, r.sectorId))
	}
	return decrypted
}

func shiftChar(c rune, n int) rune {
	// convert to 0-25
	c -= 'a'
	// shift
	c = (c+rune(n))%26 + 'a'
	return c
}
