package main

import (
	"fmt"
)

type marble struct {
	value int
	prev *marble
	next *marble
}

func (m *marble) isLast() bool {
	return m.next.value == 0
}

func printCircle(zeroMarble *marble) {

	if zeroMarble.isLast() {
		fmt.Printf("%d\n", zeroMarble.value)
		return
	}

	m := zeroMarble
	for !m.isLast() {
		fmt.Printf("%d ", m.value)
		m = m.next
	}
	fmt.Printf("%d ", m.value)

	fmt.Printf("\n")

}

func startCircle() *marble {

	m := &marble{}
	m.value = 0
	m.next = m
	m.prev = m

	return m

}

func (m *marble) insert(val int) (*marble, int) {

	if val % 23 != 0 {

		m1 := m.next
		m2 := m.next.next
		
		newCurrent := &marble{
			value: val,
			prev: m1,
			next: m2,
		}

		m1.next = newCurrent
		m2.prev = newCurrent

		return newCurrent, 0
	}

	target := m
	for i := 0; i < 7; i++ {
		target = target.prev
	}

	newCurrent := target.next
	newCurrent.prev = target.prev

	newCurrent.prev.next = newCurrent
	newCurrent.next.prev = newCurrent

	score := val + target.value

	return newCurrent, score

}

func showAround(m *marble) {

	fmt.Printf("%d -> ", m.prev.value)
	fmt.Printf("(%d) -> ", m.value)
	fmt.Printf("%d\n", m.next.value)

}

type playerWheel struct {
	n int
	current int
}

func (pw *playerWheel) next() {

	nextID := pw.current + 1
	if nextID > pw.n {
		nextID = 1
	}

	pw.current = nextID

}

func play(nPLayers int, lastValue int) (map[int]int, int) {

	zeroMarble := startCircle()
	currentMarble := zeroMarble

	var score int
	
	pw := playerWheel{
		n: nPLayers,
		current: 1,
	}

	scores := make(map[int]int)
	for id := 1; id <= nPLayers; id++ {
		scores[id] = 0
	}
	
	for i:= 1; i <= lastValue; i++ {
		
		currentMarble, score = currentMarble.insert(i)
		scores[pw.current] += score

		pw.next()

	}

	maxScore := 0
	for _, s := range scores {
		if s > maxScore {
			maxScore = s
		}
	}
	
	return scores, maxScore

}

func main() {
	
	_, maxScore := play(477, 70851) // Correct: 374690
	fmt.Println("Max score:", maxScore)
}