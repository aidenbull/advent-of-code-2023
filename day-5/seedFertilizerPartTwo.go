package main

import (
	"fmt"
	"sort"
)

type InterestRange struct {
	Start int,
	End int,
}

func max(a, b) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b) int {
	if a < b {
		return a
	}
	return b
}

func less(a, b InterestRange) bool {
	return a.Start < b.Start
}

func mapIsApplicable(seedStart, length int, m FertilizerMap) bool {
	seedEnd := seedStart + length
	mapEnd := m.SrcStart + m.Length

	return seedStart < mapEnd && seedEnd > m.SrcStart, lastStart, firstEnd
}

//Returns the first and last seed from the given range that is mapped by m
func seedsOfInterest(seedStart, length int, m FertilizerMap) InterestRange {
	seedEnd := seedStart + length
	mapEnd := m.SrcStart + m.Length

	lastStart := max(seedStart, m.SrcStart)
	firstEnd := min(seedEnd, mapEnd)

	return {lastStart, firstEnd}
}

func getSortedInterestRanges(seedStart, length int, maps []FertilizerMap) []InterestRange {
	interestRanges := make([]InterestRange, 0)
	for _, m := maps {
		interestRanges = append(interestRanges, seedsOfInterest(seedStart, length, m))
	}
	sort.Slice(interestRanges, func (i, j int) bool {
		return less(interestRanges[i], interestRanges[j])
	})

	return interestRanges
}

//Bool returns true if there is an unmapped seed. The int returns the smallest unmapped seed (or 0 if doesn't exist)
// Relies on the assumption that no maps overlap the same domain
func smallestUnmappedSeed(seedStart, length int, sortedInterestRanges []InterestRange) {bool, int} {
	out := seedStart
	for _, iRange := range sortedInterestRanges {
		if out < iRange.Start {
			return true, out
		} else {
			out := min(seedStart, iRange.End)
		}
	}
	if out >= seedStart + length {
		return false, 0
	}
	return true, out
}

func getApplicableMaps(seedStart, length int, mapSet MapSet) MapSet {
	end := seedStart + length
	applicableMaps := make([]FertilizerMap, 0)

	for _, m := range MapSet.Maps {
		if 
	}
} 

func ApplySeedRangeToMultipleMaps(start, length int, mapSets []MapSet) int {
	lowestLocation := ApplyMultipleMaps(start, mapSets)



	return lowestLocation
}

func PrintPartTwoSolution(input string) {
	seedsAndMaps := ParseSeedsAndMaps(input)

	lowestLocation := ApplySeedRangeToMultipleMaps(
		seedsAndMaps.Seeds[0], 
		seedsAndMaps.Seeds[1], 
		seedsAndMaps.MapSets)

	for i := 2 ; i < len(seedsAndMaps.Seeds) ; i += 2 {
		currLocation := ApplySeedRangeToMultipleMaps(
			seedsAndMaps.Seeds[i], 
			seedsAndMaps.Seeds[i + 1], 
			seedsAndMaps.MapSets)
		if currLocation < lowestLocation {
			lowestLocation = currLocation
		}
	}

	fmt.Println(lowestLocation)
}