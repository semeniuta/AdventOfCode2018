package aoccommons

import (
	"bufio"
	"os"
	"log"
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