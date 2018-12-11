package main

import (
	"fmt"
	"github.com/semeniuta/AdventOfCode2018/aoccommons"
	"strconv"
	"strings"
)

type node struct {
	nChildren int
	nMeta     int
	children  []*node
	meta      []int
}

func (nd *node) printTree(indent string) {
	
	fmt.Printf("%s[%d, %d], (%v)\n", indent, nd.nChildren, nd.nMeta, nd.meta)

	for _, child := range nd.children {
		child.printTree(indent + "  ")
	}
}

func (nd *node) sumMeta() int {

	var s int
	for _, num := range nd.meta {
		s += num
	}

	for _, child := range nd.children {
		s += child.sumMeta()
	}

	return s

}

func fillTree(numbers []int, start int) (*node, int) {

	nd := &node{}

	nd.nChildren = numbers[start]
	nd.nMeta = numbers[start+1]

	pointer := start + 2

	for i := 0; i < nd.nChildren; i++ {
		child, s := fillTree(numbers, pointer)
		pointer = s
		nd.children = append(nd.children, child)
	}

	for i := pointer; i < pointer+nd.nMeta; i++ {
		nd.meta = append(nd.meta, numbers[i])
	}

	pointer += nd.nMeta

	return nd, pointer
}

func readNumbers(filename string) []int {

	data := aoccommons.ReadAll(filename)

	var numbers []int

	for _, sNum := range strings.Split(data, " ") {
		num, err := strconv.Atoi(sNum)
		aoccommons.CheckError(err)

		numbers = append(numbers, num)
	}

	return numbers

}

func main() {

	// Correct (1st): 46962

	numbers := readNumbers("input.txt")

	t, _ := fillTree(numbers, 0)

	//fmt.Println(numbers)	
	//t.printTree("")

	fmt.Println("Sum:",  t.sumMeta())

}
