package main

import (
	"regexp"
	"strconv"
)

func ParseSeeds(input string) []int {
	matchNumbers := regexp.MustCompile(`\d+`)
	seedStrings := matchNumbers.FindAllString(input, -1)

	out := make([]int, 0)
	for _, seedNumStr := range seedStrings {
		//Swallowing the error because this code is expected to be run once. YAGNI
		seedNum, _ := strconv.Atoi(seedNumStr) 
		out = append(out, seedNum)
	}

	return out
}