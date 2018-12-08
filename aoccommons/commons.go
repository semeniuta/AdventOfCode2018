package aoccommons

import (
	"bufio"
	"os"
	"log"
	"strings"
	"regexp"
	"io/ioutil"
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

func ReadAll(filename string) string {
	
	file, err := os.Open(filename)
	CheckError(err)
	
	reader := bufio.NewReader(file)
	
	data, err := ioutil.ReadAll(reader)
	CheckError(err)

	return string(data)

}

func RegexParse(expression *regexp.Regexp, s string) map[string]string {

	match := expression.FindStringSubmatch(s)

	result := make(map[string]string)

	for i, name := range expression.SubexpNames() {
		if i > 0 {
			result[name] = match[i]
		}
	}

	return result

}

func AbsInt(b int) int {

	if b < 0 {
		return -b
	}
	return b

}