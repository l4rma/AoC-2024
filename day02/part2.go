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
	var prevNum int = 0
	var ascending bool = true
	isLastChance := false

	// check report
	first, err := strconv.Atoi(report[0])
	check(err)
	second, err := strconv.Atoi(report[1])
	check(err)
	third, err := strconv.Atoi(report[2])
	check(err)

	if first == second {
		if second == third {
			return false
		} else {
			if second > third {
				ascending = false
			}
		}
	} else {
		if first > second {
			ascending = false
		}
	}

	for _, e := range report {
		num, err := strconv.Atoi(e)
		check(err)

		if prevNum != 0 {
			if ascending {
				if prevNum > num {
					if isLastChance {
						return false
					} else {
						isLastChance = true
					}
				}
			} else {
				if prevNum < num {
					if isLastChance {
						return false
					} else {
						isLastChance = true
					}
				}
			}
			if num > prevNum+3 || num < prevNum-3 || prevNum == num {
				if isLastChance {
					return false
				} else {
					isLastChance = true
				}
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
