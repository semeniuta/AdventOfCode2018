package main

import (
	"fmt"
	"github.com/semeniuta/AdventOfCode2018/aoccommons"
)

func areActive(a byte, b byte) bool {

	diff := int(a) - int(b)
	absDiff := aoccommons.AbsInt(diff)
	return absDiff == 32

}

func strech(polymer string) string {

	var streched []byte

	var i int
	for i < len(polymer)-1 {

		a := polymer[i]
		b := polymer[i+1]

		if areActive(a, b) {
			i += 2
			continue
		}

		streched = append(streched, a)
		i++

	}

	streched = append(streched, polymer[len(polymer)-1])

	return string(streched)

}

func strechLoop(polymer string) string {

	currentPolymer := polymer
	currentLen := len(currentPolymer)

	for true {

		newPolymer := strech(currentPolymer)
		newLen := len(newPolymer)

		currentPolymer = newPolymer

		if newLen == currentLen {
			break
		}

		currentLen = newLen
	}

	return currentPolymer

}

func removeChar(s string, c byte) string {

	c2 := c + 32

	var res []byte
	for i := 0; i < len(s); i++ {

		currentChar := s[i]
		if currentChar == c || currentChar == c2 {
			continue
		}

		res = append(res, currentChar)

	}

	return string(res)

}

func detectProblematicChar(s string) (byte, int) {

	scores := make(map[byte]int)

	var c byte
	for c = 'A'; c <= 'Z'; c++ {

		sModified := removeChar(s, c)
		streched := strechLoop(sModified)
		scores[c] = len(streched)

	}

	minScore := len(s)
	var minScoreChar byte
	for c, score := range scores {
		if score < minScore {
			minScore = score
			minScoreChar = c
		}
	}

	return minScoreChar, minScore
}

func main() {

	// Correct 1st: 10250
	// Correct 2nd: 6188

	s := aoccommons.ReadAll("input.txt")

	streched := strechLoop(s)
	fmt.Println("Length of the stretched polymer:", len(streched))

	minScoreChar, minScore := detectProblematicChar(s)
	fmt.Printf("Problematic element: %c; resulting length: %d\n", minScoreChar, minScore)

}
