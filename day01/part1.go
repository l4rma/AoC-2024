package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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
	sort.Slice(leftList, func(i, j int) bool {
		return leftList[i] < leftList[j]
	})

	sort.Slice(rightList, func(i, j int) bool {
		return rightList[i] < rightList[j]
	})

	var res []int
	for n, leftValue := range leftList {
		diff := leftValue - rightList[n]
		if diff < 0 {
			diff = -diff
		}
		res = append(res, diff)
	}

	var sum int
	for _, e := range res {
		sum += e
	}

	fmt.Println(sum)
}
