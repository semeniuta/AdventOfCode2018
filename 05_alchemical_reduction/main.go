package main

import (
	"github.com/semeniuta/AdventOfCode2018/aoccommons"
	"fmt"
)

func absInt(b int) int {
	
	if b < 0 {
		return -b
	}
	return b

}

func areActive(a byte, b byte) bool {
	
	diff := int(a) - int(b)
	absDiff := absInt(diff)
	return absDiff == 32

}

func strech(polymer string) string {

	var streched []byte

	var i int
	for i < len(polymer) - 1 {

		a := polymer[i]
		b := polymer[i+1]

		if areActive(a, b) {
			i += 2
			continue
		}

		streched = append(streched, a)
		i++

	}

	streched = append(streched, polymer[len(polymer) - 1])

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

func main() {

	// Correct 1st: 10250

	s := aoccommons.ReadAll("input.txt")

	streched := strechLoop(s)
	fmt.Println(len(streched))

	//s := "dabAcCaCBAcCcaDA"
	//fmt.Println(s)
	//fmt.Println(strechLoop(s))
	
}