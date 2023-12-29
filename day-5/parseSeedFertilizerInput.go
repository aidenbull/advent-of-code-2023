package main

import (
	"regexp"
	"strings"
	"strconv"
)

type FertilizerMap struct {
	DestStart int
	SourceStart int
	Length int
}

type MapSet struct {
	Maps []FertilizerMap
}

func ParseSeeds(input string) []int {
	matchNumbers := regexp.MustCompile(`\d+`)
	seedStrings := matchNumbers.FindAllString(input, -1)

	out := make([]int, 0)
	for _, seedNumStr := range seedStrings {
		//Swallowing the error because this code is expected to be run 'in production' once. YAGNI
		seedNum, _ := strconv.Atoi(seedNumStr) 
		out = append(out, seedNum)
	}

	return out
}

func ParseSetOfMappings(input string) MapSet {
	matchMappingPrefix := regexp.MustCompile(`^.+-to-.+ map:\n`)
	prefix := matchMappingPrefix.FindString(input)
	mapsStr := strings.TrimPrefix(input, prefix)
	
	maps := strings.Split(mapsStr, "\n")
	out := make([]FertilizerMap, 0)

	for _, m := range maps {
		matchNumbers := regexp.MustCompile(`\d+`)
		numbers := matchNumbers.FindAllString(m, -1)

		destStart, _ := strconv.Atoi(numbers[0])
		srcStart, _ := strconv.Atoi(numbers[1])
		length, _ := strconv.Atoi(numbers[2])
		
		out = append(out, FertilizerMap{destStart, srcStart, length})
	}

	return MapSet{out}
}