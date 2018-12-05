package main

import (
	"fmt"
	"strings"
	"strconv"
	"github.com/semeniuta/AdventOfCode2018/aoccommons"
)

type patch struct {
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

func parseLine(line string) (int, patch) {

	// Example of a line:
	// #120 @ 773,696: 15x23

	p := patch{}

	elements := strings.Split(line, " ")

	sID := elements[0]
	sCorner := elements[2]
	sSides := elements[3]

	id, err := strconv.Atoi(sID[1:len(sID)])
	aoccommons.CheckError(err)

	corner := parseNumbers(sCorner[:len(sCorner) - 1], ",")
	sides := parseNumbers(sSides, "x")

	p.left = corner[0]
	p.top = corner[1]
	p.width = sides[0]
	p.height = sides[1]

	return id, p

}

type coordinate struct {
	row int
	col int
}

type coordinateInfo struct {
	nClaims int
	claimIDs []int
}

func buildMap(patches map[int]patch) map[coordinate]*coordinateInfo {

	m := make(map[coordinate]*coordinateInfo)

	for id, p := range patches {
		
		for i := p.top; i < p.top + p.height; i++ {
			for j := p.left; j < p.left + p.width; j++ {

				c := coordinate{i, j}

				val, ok := m[c]
				if ok {
					m[c].nClaims = val.nClaims + 1
					m[c].claimIDs = append(m[c].claimIDs, id)
				} else {
					var info coordinateInfo
					info.nClaims = 1
					info.claimIDs = append(info.claimIDs, id)
					m[c] = &info
				}

			}
		}

	}

	return m

}

func findSingleClaim(patches map[int]patch, candidates map[int]bool, m map[coordinate]*coordinateInfo) int {

	for id := range candidates {

		p := patches[id]
		ok := true
		
		for i := p.top; i < p.top + p.height; i++ {
			for j := p.left; j < p.left + p.width; j++ {

				if m[coordinate{i, j}].nClaims > 1 {
					ok = false
					break
				}
				
			}
		}

		if ok {
			return id
		}

	}

	return -1
}

func main() {

	// Correct (1st): 112418
	// Correct (2nd): 560

	filename := "input.txt"
	lines := aoccommons.ReadLines(filename)

	patches := make(map[int]patch)

	for _, line := range lines {
		id, p := parseLine(line)
		patches[id] = p
	}

	m := buildMap(patches)

	var count int
	var candidates = make(map[int]bool)
	for _, v := range m {
		
		if v.nClaims >= 2 {
			count++
		}

		if v.nClaims == 1 {
			candidates[v.claimIDs[0]] = true
		}

	}

	singleClaimID := findSingleClaim(patches, candidates, m)

	fmt.Println("Number of square inches of fabric within two or more claims:", count)
	fmt.Println("Claim ID with no overlaps:", singleClaimID)
	

}