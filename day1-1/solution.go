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
	for _, str := range strArray {
		num, err := strconv.Atoi(str)
		if err != nil {
			continue
		}

		numArray = append(numArray, num)
	}

	sort.Ints(numArray)
	fmt.Println(numArray[0])
}
