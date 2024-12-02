package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var leftList []int
	var rightList []int

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		set := strings.Split(line, "   ")
		if len(set) > 1 {
			x, _ := strconv.Atoi(set[0])
			y, _ := strconv.Atoi(set[1])

			leftList = append(leftList, x)
			rightList = append(rightList, y)
		}
	}

	var simScore int

	for _, leftVal := range leftList {
		var hits int
		for _, rightVal := range rightList {
			if leftVal == rightVal {
				hits += 1
			}
		}
		simScore += leftVal * hits
	}

	fmt.Println(simScore)
}
