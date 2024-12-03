package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	safe   int = 0
	unsafe int = 0
)

func main() {
	input, err := os.ReadFile("input.txt")
	check(err)

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		report := strings.Split(line, " ")
		if len(report) > 1 {
			if isSafe(report) {
				safe++
			} else {
				unsafe++
			}
		}
	}
	fmt.Println("Unsafe: ", unsafe)
	fmt.Println("Safe: ", safe)
}

func isSafe(report []string) bool {
	fmt.Println(report)
	var prevNum int = 0
	var ascending bool = true

	// check report
	first, err := strconv.Atoi(report[0])
	check(err)
	second, err := strconv.Atoi(report[1])
	check(err)
	if first > second {
		ascending = false
	}

	for _, e := range report {
		num, err := strconv.Atoi(e)
		check(err)

		if prevNum != 0 {
			if prevNum == num {
				return false
			}
			if ascending {
				if prevNum > num {
					return false
				}
			} else {
				if prevNum < num {
					return false
				}
			}
		}
		if prevNum != 0 {
			if num > prevNum+3 || num < prevNum-3 {
				return false
			}
		}
		prevNum = num
	}
	return true
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
