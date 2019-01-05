package main

import (
	"fmt"
	"math"
)

var gridSize = 300

type grid struct {
	serialNumber int
}

type coordinate struct {
	x int
	y int
}

func (g *grid) powerLevel(coord coordinate) int {

	rackID := coord.x + 10

	power := rackID * coord.y

	power += g.serialNumber

	power *= rackID

	power = hundredsDigit(power)

	power -= 5

	return power
}

func (g *grid) findHighestPowerSquare() (coordinate, int) {

	cache := make(map[coordinate]int)

	var highestPower int
	var highestPowerSquare coordinate

	for y := 1; y <= gridSize-2; y++ {
		for x := 1; x <= gridSize-2; x++ {

			var totalPower int

			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {

					cell := coordinate{x + j, y + i}
					power, ok := cache[cell]
					if !ok {
						power = g.powerLevel(cell)
						cache[cell] = power
					}

					totalPower += power
					
				}
			}

			if totalPower > highestPower {
				highestPower = totalPower
				highestPowerSquare = coordinate{x, y}
			}
			
		}

	}

	return highestPowerSquare, highestPower

}

func hundredsDigit(number int) int {

	nDigits := calculateNumDigits(number)

	if nDigits < 2 {
		return 0
	}

	remainder := number % 1000

	return remainder / 100

}

func calculateNumDigits(number int) int {

	i := -1

	var done bool
	for !done {

		i++

		a := int(math.Pow10(i))
		b := int(math.Pow10(i+1) - 1)

		if a <= number && number <= b {
			done = true
		}
	}

	return i + 1

}

func main() {

	// Correct (1st): 245, 14

	g := grid{serialNumber: 9810}

	hiCoord, _ := g.findHighestPowerSquare()
	fmt.Println("Coordinate of the 3x3 square with highest power:", hiCoord)

}
