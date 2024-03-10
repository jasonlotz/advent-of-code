package main

import (
	"fmt"
	"sort"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2018/04/input.txt"
var testInputFile = "../../../input-files/2018/04/input-sample.txt"
var isTestMode = false

type Action struct {
	Year    int
	Month   int
	Day     int
	Hour    int
	Minute  int
	GuardId int
	Type    ActionType
}

type ActionType int

const (
	BeginShift ActionType = iota
	FallAsleep
	WakeUp
)

type GuardSleepMap map[int]DateSleepMap

type DateSleepMap map[string]MinutesSlept

type MinutesSlept []int

func getInput() []string {
	file := inputFile

	if isTestMode {
		file = testInputFile
	}

	lines := utils.ProcessStringLinesFile(file)

	sort.Strings(lines)

	return lines
}

func main() {
	input := getInput()
	actions := parseInput(input)

	guardSleep := processActions(actions)
	if isTestMode {
		printGuardSleep(guardSleep)
	}

	sleepiestGuard, sleepiestMinute := findSleepiestGuard(guardSleep)
	fmt.Printf("Sleepiest guard: %d\n", sleepiestGuard)
	fmt.Printf("Sleepiest minute: %d\n", sleepiestMinute)

	fmt.Println("Part 1:", sleepiestGuard*sleepiestMinute)

	sleepiestGuardByMinute, sleepiestOverallMinute := findSleepiestGuardByMinute(guardSleep)
	fmt.Printf("Sleepiest guard by minute: %d\n", sleepiestGuardByMinute)
	fmt.Printf("Sleepiest overall minute: %d\n", sleepiestOverallMinute)

	fmt.Println("Part 2:", sleepiestGuardByMinute*sleepiestOverallMinute)
}

func parseInput(input []string) []Action {
	actions := []Action{}

	for _, line := range input {
		action := Action{}

		_, _ = fmt.Sscanf(line, "[%d-%d-%d %d:%d]", &action.Year, &action.Month, &action.Day, &action.Hour, &action.Minute)

		rawAction := line[19:] // Remove the date from the line
		switch rawAction {
		case "falls asleep":
			action.Type = FallAsleep
		case "wakes up":
			action.Type = WakeUp
		default:
			action.Type = BeginShift
			_, _ = fmt.Sscanf(rawAction, "Guard #%d begins shift", &action.GuardId)
		}

		actions = append(actions, action)
	}

	return actions
}

func processActions(actions []Action) GuardSleepMap {
	guardSleep := GuardSleepMap{}
	lastGuard := 0
	lastDate := ""
	lastFellAsleep := 0

	for _, action := range actions {
		lastDate = fmt.Sprintf("%d-%d-%d", action.Year, action.Month, action.Day)

		if action.Type == BeginShift {
			lastGuard = action.GuardId

			// If the guard isn't in the map, add them
			if _, ok := guardSleep[lastGuard]; !ok {
				guardSleep[lastGuard] = DateSleepMap{}
			}
		}

		// If the date isn't in the guard's map, add it
		if _, ok := guardSleep[lastGuard][lastDate]; !ok {
			guardSleep[lastGuard][lastDate] = make(MinutesSlept, 60)
		}

		if action.Type == FallAsleep {
			lastFellAsleep = action.Minute
		}

		if action.Type == WakeUp {
			for i := lastFellAsleep; i < action.Minute; i++ {
				guardSleep[lastGuard][lastDate][i]++
			}
		}
	}

	return guardSleep
}

func findSleepiestGuard(guardSleep GuardSleepMap) (sleepiestGuard int, sleepiestMinute int) {
	sleepiestGuard = 0
	sleepiestGuardMinutes := 0

	for guard, dates := range guardSleep {
		minutesAsleep := 0
		for _, minutes := range dates {
			for _, minute := range minutes {
				minutesAsleep += minute
			}
		}

		if minutesAsleep > sleepiestGuardMinutes {
			sleepiestGuard = guard
			sleepiestGuardMinutes = minutesAsleep
		}
	}

	sleepiestMinute, _ = findSleepiestMinute(guardSleep[sleepiestGuard])

	return sleepiestGuard, sleepiestMinute
}

func findSleepiestMinute(dates DateSleepMap) (sleepiestMinute int, sleepiestMinuteCount int) {
	minutesAggregate := make(MinutesSlept, 60)

	for _, minutes := range dates {
		for i, minute := range minutes {
			minutesAggregate[i] += minute
		}
	}

	sleepiestMinute = 0
	for i, minute := range minutesAggregate {
		if minute > minutesAggregate[sleepiestMinute] {
			sleepiestMinute = i
		}
	}

	return sleepiestMinute, minutesAggregate[sleepiestMinute]
}

func findSleepiestGuardByMinute(guardSleep GuardSleepMap) (sleepiestGuardByMinute int, sleepiestOverallMinute int) {
	sleepiestGuardByMinute = 0
	sleepiestOverallMinute = 0
	sleepiestMinuteCount := 0

	for guard, dates := range guardSleep {
		minute, count := findSleepiestMinute(dates)

		if count > sleepiestMinuteCount {
			sleepiestGuardByMinute = guard
			sleepiestOverallMinute = minute
			sleepiestMinuteCount = count
		}
	}

	return sleepiestGuardByMinute, sleepiestOverallMinute
}

func printGuardSleep(guardSleep GuardSleepMap) {
	for guard, dates := range guardSleep {
		fmt.Printf("Guard %d\n", guard)
		for date, minutes := range dates {
			fmt.Printf("\t%s: %v\n", date, minutes)
		}
	}
}
