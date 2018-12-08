package main

import (
	"fmt"
	"strings"
	"strconv"
	"github.com/semeniuta/AdventOfCode2018/aoccommons"
)

type point struct {
	x int
	y int
}

type boundaryElement struct {
	value int
	points []point
}

type boundary struct {
	xMin boundaryElement
	xMax boundaryElement
	yMin boundaryElement
	yMax boundaryElement
}

func (be *boundaryElement) addPoints(indices []int, points []point) {
	for _, idx := range indices {
		be.points = append(be.points, points[idx])
	}
}

type elementInfo struct {
	value int
	index int
}

func readCoordinates(filename string) []point {

	lines := aoccommons.ReadLines(filename)
	
	var points []point
	for _, line := range lines {

		substrings := strings.Split(line, ",")
		
		x, err := strconv.Atoi(strings.TrimSpace(substrings[0]))
		aoccommons.CheckError(err)
		
		y, err := strconv.Atoi(strings.TrimSpace(substrings[1]))
		aoccommons.CheckError(err)

		p := point{x, y}
		points = append(points, p)

	}

	return points

}

func determineBoundary(points []point) boundary {

	memoryX := make(map[int][]int)
	memoryY := make(map[int][]int)

	xMax := elementInfo{points[0].x, 0}
	xMin := elementInfo{points[0].x, 0}
	yMax := elementInfo{points[0].y, 0}
	yMin := elementInfo{points[0].y, 0}

	for i := 0; i < len(points); i++ {

		p := points[i]

		memoryX[p.x] = append(memoryX[p.x], i)
		memoryY[p.y] = append(memoryY[p.y], i)

		if p.x > xMax.value {
			xMax.index = i
			xMax.value = p.x
		}

		if p.x < xMin.value {
			xMin.index = i
			xMin.value = p.x
		}

		if p.y > yMax.value {
			yMax.index = i
			yMax.value = p.y
		}

		if p.y < yMin.value {
			yMin.index = i
			yMin.value = p.y
		}

	}

	b := boundary{}
	
	b.xMax.value = xMax.value
	b.xMax.addPoints(memoryX[xMax.value], points)

	b.xMin.value = xMin.value
	b.xMin.addPoints(memoryX[xMin.value], points)

	b.yMax.value = yMax.value
	b.yMax.addPoints(memoryY[yMax.value], points)

	b.yMin.value = yMin.value
	b.yMin.addPoints(memoryY[yMin.value], points)

	return b

}

func manhattanDist(p1 point, p2 point) int {

	diffX := aoccommons.AbsInt(p1.x - p2.x)
	diffY := aoccommons.AbsInt(p1.y - p2.y)

	return diffX + diffY

}

func getNonBoundaryPoints(b boundary, points []point) map[int]bool {

	boundaryPoints := make(map[point]bool)
	nonBoundaryPoints := make(map[int]bool)

	for _, pt := range b.xMax.points {
		boundaryPoints[pt] = true
	}

	for _, pt := range b.xMin.points {
		boundaryPoints[pt] = true
	}

	for _, pt := range b.yMax.points {
		boundaryPoints[pt] = true
	}

	for _, pt := range b.yMin.points {
		boundaryPoints[pt] = true
	}

	for idx, pt := range points {
		
		_, ok := boundaryPoints[pt]
		if !ok {
			nonBoundaryPoints[idx] = true
		}
	}

	return nonBoundaryPoints

}

func findClosest(b boundary, points []point) map[point]int {

	nPoints := len(points)

	closest := make(map[point]int)

	for i := b.xMin.value; i <= b.xMax.value; i++ {

		for j := b.yMin.value; j <= b.yMax.value; j++ {

			pt0 := point{i, j}

			distances := make([]int, nPoints, nPoints)

			minDist := b.xMax.value + b.yMax.value
			for pointIndex, pt1 := range points {
				
				dist := manhattanDist(pt0, pt1)
				distances[pointIndex] = dist
				
				if dist < minDist {
					minDist = dist
				}
			}

			nTies := 0
			indexWithMinDist := 0
			for pointIndex := range points {
				if distances[pointIndex] == minDist {
					nTies++
					indexWithMinDist = pointIndex
				}
			}

			if nTies == 1 {
				closest[pt0] = indexWithMinDist
			} else {
				closest[pt0] = -1
			}

		}

	}

	return closest

}

func findLargest(closest map[point]int, nonBoundaryPoints map[int]bool) (int, int) {

	sizes := make(map[int]int)

	for _, idx := range(closest) {

		_, isNonBoundaryPoint := nonBoundaryPoints[idx]

		if idx != -1 && isNonBoundaryPoint {
			
			val, ok := sizes[idx]
			if ok {
				sizes[idx] = val + 1
			} else {
				sizes[idx] = 1
			}

		} 
	}

	maxIdx := 0
	maxSz := sizes[0]
	for idx, sz := range(sizes) {
		if sz > maxSz {
			maxIdx = idx
			maxSz = sz
		}
	}

	return maxIdx, maxSz

}


func main() {

	points := readCoordinates("small.txt")
	b := determineBoundary(points)
	nonBoundaryPoints := getNonBoundaryPoints(b, points)
	closest := findClosest(b, points)

	fmt.Println("Boundary:", b)
	fmt.Println("Non-boundary points:")
	for idx := range nonBoundaryPoints {
		fmt.Println(points[idx])
	}

	largestIdx, largestSize := findLargest(closest, nonBoundaryPoints)
	fmt.Println("Largest:", points[largestIdx], largestSize)

}