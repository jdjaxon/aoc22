package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)
	strArray := strings.Split(string(data), "\n")
	numArray := []int{}
	totalCal := 0
	for _, str := range strArray {
		num, err := strconv.Atoi(str)
		totalCal += num
		if err != nil {
			numArray = append(numArray, totalCal)
			totalCal = 0
			continue
		}
	}

	// part one
	sort.Sort(sort.Reverse(sort.IntSlice(numArray)))
	fmt.Println("Part one: ", numArray[0])

	// part two
	topThreeCal := 0
	for _, num := range numArray[:3] {
		topThreeCal += num
	}
	fmt.Println("Part two: ", topThreeCal)

}
