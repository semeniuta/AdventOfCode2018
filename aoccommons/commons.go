package aoccommons

import (
	"bufio"
	"os"
	"log"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func CreateScanner(filename string) *bufio.Scanner {

	file, err := os.Open(filename)
	CheckError(err)

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	return scanner

}

func ReadLines(filename string) []string {

	scanner := CreateScanner(filename)

	var lines []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lines = append(lines, line)
	}

	return lines

}