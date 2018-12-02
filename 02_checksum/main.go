package main

import (
	"fmt"
	"github.com/semeniuta/AdventOfCode2018/aoccommons"
	"strings"
)

func readLines(filename string) []string {

	scanner := aoccommons.CreateScanner(filename)

	var lines []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lines = append(lines, line)
	}

	return lines

}

func analyzeLine(line string) (bool, bool) {

	freqs := make(map[rune]int)

	for _, r := range line {

		val, exists := freqs[r]
		if exists {
			freqs[r] = val + 1
		} else {
			freqs[r] = 1
		}

	}

	var appear2 bool
	var appear3 bool

	for _, v := range freqs {

		if v == 2 {
			appear2 = true
			continue
		}

		if v == 3 {
			appear3 = true
			continue
		}

		if appear2 && appear3 {
			break
		}

	}

	return appear2, appear3
}

func computeChecksum(lines []string) int {
	
	var n2 int
	var n3 int

	for _, line := range lines {
		
		appear2, appear3 := analyzeLine(line)
		
		if appear2 {
			n2++
		}

		if appear3 {
			n3++
		}
	}

	return n2 * n3
}

func main() {

	// Correct checksum: 7350

	filename := "input.txt"
	lines := readLines(filename)

	checksum := computeChecksum(lines)
	fmt.Println("Checksum:", checksum)

}
