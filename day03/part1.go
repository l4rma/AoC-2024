package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	check(err)

	lines := strings.Split(string(input), "\n")

	pattern := `mul\(\d{1,3},\d{1,3}\)`
	regex, err := regexp.Compile(pattern)
	check(err)
	var matches []string

	for _, line := range lines {
		lineMatches := regex.FindAllString(line, -1)
		matches = append(matches, lineMatches...)
	}

	newPattern := `\d{1,3},\d{1,3}`
	newRegex, err := regexp.Compile(newPattern)
	check(err)
	var math []string
	for _, line := range matches {
		lineMatches := newRegex.FindAllString(line, -1)
		math = append(math, lineMatches...)
	}

	total := 0
	for _, line := range math {
		xy := strings.Split(line, ",")

		x, err := strconv.Atoi(xy[0])
		check(err)

		y, err := strconv.Atoi(xy[1])
		check(err)

		sum := x * y
		total += sum
	}
	fmt.Println("Total:", total)

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
