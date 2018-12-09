package main

import (
	"fmt"
	"github.com/semeniuta/AdventOfCode2018/aoccommons"
	"log"
	"sort"
	"strconv"
)

type action int

const (
	beginShift = iota
	fallAsleep
	wakeUp
)

type event struct {
	year        int
	month       int
	day         int
	hours       int
	minutes     int
	guardID     int
	guardAction action
	source      string
}

func processLines(lines []string) (map[int]*event, []int) {

	reMain := aoccommons.CompileRegex("^\\[(?P<year>\\d\\d\\d\\d)-(?P<month>\\d\\d)-(?P<day>\\d\\d) (?P<hours>\\d\\d):(?P<minutes>\\d\\d)\\] (?P<action>.*)$")

	reBegin := aoccommons.CompileRegex("^Guard #(?P<guardID>\\d+) begins shift$")

	events := make(map[int]*event)
	var sortedKeys []int

	for _, line := range lines {

		parseMap := aoccommons.RegexParse(reMain, line)

		eventEntry := event{}
		eventEntry.source = line
		eventEntry.year, _ = strconv.Atoi(parseMap["year"])
		eventEntry.month, _ = strconv.Atoi(parseMap["month"])
		eventEntry.day, _ = strconv.Atoi(parseMap["day"])
		eventEntry.hours, _ = strconv.Atoi(parseMap["hours"])
		eventEntry.minutes, _ = strconv.Atoi(parseMap["minutes"])

		matched := reBegin.MatchString(parseMap["action"])
		if matched {
			beginParseMap := aoccommons.RegexParse(reBegin, parseMap["action"])
			eventEntry.guardID, _ = strconv.Atoi(beginParseMap["guardID"])
			eventEntry.guardAction = beginShift
		} else if parseMap["action"] == "falls asleep" {
			eventEntry.guardAction = fallAsleep
		} else if parseMap["action"] == "wakes up" {
			eventEntry.guardAction = wakeUp
		} else {
			log.Fatalln("Wrong line:", parseMap["action"])
		}

		mapKey, _ := strconv.Atoi(parseMap["year"] + parseMap["month"] + parseMap["day"] + parseMap["hours"] + parseMap["minutes"])

		events[mapKey] = &eventEntry
		sortedKeys = append(sortedKeys, mapKey)

	}

	sort.Ints(sortedKeys)

	var currentGuardID int
	for _, mapKey := range sortedKeys {

		event := events[mapKey]

		switch event.guardAction {

		case beginShift:
			currentGuardID = event.guardID

		case fallAsleep:
			event.guardID = currentGuardID

		case wakeUp:
			event.guardID = currentGuardID

		}
	}

	return events, sortedKeys

}

type guardStats struct {
	totalSleep int
	durations  []int
}

func prepareStats(events map[int]*event, sortedKeys []int) map[int]*guardStats {

	stats := make(map[int]*guardStats)

	var currentAction action
	var currentGuard int
	var lastFallsAsleep int
	var sleepDuration int
	var sleepDataReady bool
	var sleepMinutes []int

	for _, k := range sortedKeys {

		event := events[k]
		currentAction = event.guardAction
		currentGuard = event.guardID

		if currentAction == fallAsleep {
			lastFallsAsleep = event.minutes
			sleepDataReady = false
		}

		if currentAction == wakeUp {

			sleepDuration = event.minutes - lastFallsAsleep

			fmt.Println(currentGuard, "has slept for", sleepDuration, "from", lastFallsAsleep, "to", event.minutes)

			var minutesList []int
			for minute := lastFallsAsleep; minute < event.minutes; minute++ {
				minutesList = append(minutesList, minute)
			}
			sleepMinutes = minutesList

			sleepDataReady = true
		}

		entry, ok := stats[currentGuard]

		if !ok {

			s := guardStats{}
			s.durations = make([]int, 60, 60)

			stats[currentGuard] = &s
			entry = stats[currentGuard]
		}

		if event.guardAction == wakeUp && sleepDataReady {

			entry.totalSleep += sleepDuration
			fmt.Println("Total sleep duration for", currentGuard, "is now", entry.totalSleep)

			for _, minute := range sleepMinutes {
				entry.durations[minute]++
			}
		}
	}

	return stats

}

func detectSleepy(stats map[int]*guardStats) (int, int) {

	var maxMinutes int
	var sleepyID int
	for gID, gStats := range stats {

		if gStats.totalSleep > maxMinutes {
			maxMinutes = gStats.totalSleep
			sleepyID = gID
		}
	}

	sleepyMinutes := stats[sleepyID].durations
	var mostFrequentValue int
	var mostFrequentIdx int
	for i := 0; i < 60; i++ {

		if sleepyMinutes[i] > mostFrequentValue {
			mostFrequentValue = sleepyMinutes[i]
			mostFrequentIdx = i
		}

	}

	return sleepyID, mostFrequentIdx

}

func main() {

	lines := aoccommons.ReadLines("input.txt")
	events, sortedKeys := processLines(lines)

	stats := prepareStats(events, sortedKeys)

	for gID, s := range stats {
		fmt.Println(gID, s.totalSleep)
	}

	sleepyID, mostFrequentIdx := detectSleepy(stats)

	for i := 0; i < 60; i++ {
		fmt.Println(i, ":", stats[sleepyID].durations[i])
	}

	fmt.Println(sleepyID, mostFrequentIdx, sleepyID * mostFrequentIdx)
	
	// 409 51 20859

}
