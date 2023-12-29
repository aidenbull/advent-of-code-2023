package main

import (
	"regexp"
	"strings"
	"strconv"
)

type FertilizerMap struct {
	DestStart int
	SrcStart int
	Length int
}

type MapSet struct {
	Maps []FertilizerMap
}

type SeedsAndMaps struct {
	Seeds []int
	MapSets []MapSet
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

func ParseSeedsAndMaps(input string) SeedsAndMaps {
	preProcessedInput := strings.TrimSpace(input)
	majorComponentTokens := strings.Split(preProcessedInput, "\n\n")

	seeds := ParseSeeds(majorComponentTokens[0])

	mapSets := make([]MapSet, 0)
	for _, mapSetStr := range majorComponentTokens[1:] {
		mapSets = append(mapSets, ParseSetOfMappings(mapSetStr))
	}

	return SeedsAndMaps{seeds, mapSets}
}

func ApplyMap(input int, mapSet MapSet) int {
	for _, m := range mapSet.Maps {
		if input >= m.SrcStart && input < (m.SrcStart + m.Length) {
			return input + (m.DestStart - m.SrcStart)
		}
	}
	return input
}

func ApplyMultipleMaps(input int, mapSets []MapSet) int {
	out := input
	
	for _, mapSet := range mapSets {
		out = ApplyMap(out, mapSet)
	}

	return out
} 