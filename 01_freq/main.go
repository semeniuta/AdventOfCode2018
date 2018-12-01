package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func createScanner(filename string) *bufio.Scanner {

	file, err := os.Open(filename)
	checkError(err)

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	return scanner

}

func parseNumber(line string) int {

	sSign := line[0]
	sNumber := string(line[1:len(line)])

	sign := 1
	if sSign == '-' {
		sign = -1
	}

	number, err := strconv.Atoi(sNumber)
	checkError(err)

	return sign * number
}

func readChanges(filename string) []int {

	scanner := createScanner(filename)

	rValidate, err := regexp.Compile("^[+-][0-9]+$")
	checkError(err)

	var changes []int

	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())

		matched := rValidate.MatchString(line)

		if !matched {
			log.Fatalln("Wrong formatting in line", line)
			return nil
		}

		number := parseNumber(line)

		changes = append(changes, number)
	}

	return changes

}

func calculateFinalState(changes []int) int {

	var state int
	for _, change := range changes {
		state += change
	}

	return state
}

type stateMonitor struct {
	seenStates map[int]bool
	done bool
}

func newStateMonitor() *stateMonitor {
	
	monitor := stateMonitor{
		seenStates: make(map[int]bool),
	}
	
	return &monitor
}

func (m *stateMonitor) haveSeen(state int) bool {
	_, exists := m.seenStates[state]
	return exists
}

func (m *stateMonitor) remember(state int) {
	m.seenStates[state] = true
}

func checkForRepeat(changes []int) int {

	var state int

	monitor := newStateMonitor()

	for !monitor.done {

		for _, change := range changes {

			state += change

			if monitor.haveSeen(state) {
				monitor.done = true
				break
			}

			monitor.remember(state)
		}
	}

	return state
}

func main() {

	// Correct output:
	// Final state: 505
	// State seen twice the first time: 72330

	filename := "input.txt"
	changes := readChanges(filename)

	finalState := calculateFinalState(changes)
	fmt.Println("Final state:", finalState)

	stateCyclic := checkForRepeat(changes)
	fmt.Println("State seen twice the first time:", stateCyclic)

}
