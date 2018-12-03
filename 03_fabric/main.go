package main

import (
	"fmt"
	"strings"
	"strconv"
	"github.com/semeniuta/AdventOfCode2018/aoccommons"
)

type patch struct {
	id int
	left int
	top int
	width int
	height int
}

func parseNumbers(s string, sep string) []int {

	sNums := strings.Split(s, sep)

	var res []int

	for _, sNum := range sNums {
		num, err := strconv.Atoi(sNum)
		aoccommons.CheckError(err)
		res = append(res, num)
	}

	return res
}

func parseLine(line string) patch {

	// #120 @ 773,696: 15x23

	p := patch{}

	elements := strings.Split(line, " ")

	sID := elements[0]
	sCorner := elements[2]
	sSides := elements[3]

	id, err := strconv.Atoi(sID[1:len(sID)])
	aoccommons.CheckError(err)
	p.id = id

	corner := parseNumbers(sCorner[:len(sCorner) - 1], ",")
	sides := parseNumbers(sSides, "x")

	p.left = corner[0]
	p.top = corner[1]
	p.width = sides[0]
	p.height = sides[1]

	return p

}

type coordinate struct {
	row int
	col int
}

type coordinateInfo struct {
	nClaims int
	claimIDs []int
}

func buildMap(patches []patch) map[coordinate]int {

	m := make(map[coordinate]int)

	for _, p := range patches {
		
		for i := p.top; i < p.top + p.height; i++ {
			for j := p.left; j < p.left + p.width; j++ {

				c := coordinate{i, j}

				val, ok := m[c]
				if ok {
					m[c] = val + 1
				} else {
					m[c] = 1
				}

			}
		}

	}

	return m

}


func main() {

	// Correct (1st): 112418

	filename := "input.txt"
	lines := aoccommons.ReadLines(filename)

	var patches []patch
	for _, line := range lines {
		patches = append(patches, parseLine(line))
	}

	m := buildMap(patches)

	var count int
	for _, v := range m {
		if v >= 2 {
			count++
		}
	}

	fmt.Println(count)
	

}