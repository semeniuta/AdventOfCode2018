package main

import (
	"log"
	"fmt"
	"github.com/semeniuta/AdventOfCode2018/aoccommons"
)

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

func compareLines(line1 string, line2 string) []bool {

	nChars := len(line1)
	if nChars != len(line2) {
		log.Fatalln("String of different length are supplied to compareLines", line1, line2)
	}
	
	diff := make([]bool, nChars, nChars)

	for i := 0; i < nChars; i++ {
		if line1[i] == line2[i] {
			diff[i] = true
		} 
	}

	return diff

}

func moreThanOneDifferent(diff []bool) bool {
	
	var nDifferent int
	var moreThanOne bool
	
	for _, match := range diff {
		if !match {
			nDifferent++
		}
		if nDifferent > 1 {
			moreThanOne = true
			break
		}
	}

	return moreThanOne

}

func getCommonChars(line string, matches []bool) string {

	characters := make([]byte, len(line), len(line))

	for charIdx, match := range matches {
		if match {
			characters[charIdx] = line[charIdx]
		}
	}

	return string(characters)
}

func findSimilar(lines []string) string {

	type pair struct {
		a int
		b int
	}

	alreadyCompared := make(map[pair]bool)

	for i, line1 := range(lines) {

		for j, line2 := range(lines) {

			if i == j {
				continue
			}

			if alreadyCompared[pair{i, j}] {
				continue
			}
			
			diff := compareLines(line1, line2)
			oneCharDifference := !moreThanOneDifferent(diff)

			if oneCharDifference {
				return getCommonChars(line1, diff)
			}

			alreadyCompared[pair{i, j}] = true
			alreadyCompared[pair{j ,i}] = true
		}
	}

	return ""

}

func main() {

	// Correct checksum: 7350
	// Correct common characters: wmlnjevbfodamyiqpucrhsukg

	filename := "input.txt"
	lines := aoccommons.ReadLines(filename)

	checksum := computeChecksum(lines)
	fmt.Println("Checksum:", checksum)

	similar := findSimilar(lines)
	fmt.Println(similar)
	

}
