package main

import (
	"fmt"
)

func GetLowestLocationFromSeedRange(start, length int, mapSets []MapSet) int {
	lowestLocation := ApplyMultipleMaps(start, mapSets)
	for i := start + 1 ; i < start + length ; i++ {
		currLocation := ApplyMultipleMaps(i, mapSets)
		if currLocation < lowestLocation {
			lowestLocation = currLocation
		}
	}

	return lowestLocation
}

func PrintPartTwoSolution(input string) {
	seedsAndMaps := ParseSeedsAndMaps(input)

	lowestLocation := GetLowestLocationFromSeedRange(
		seedsAndMaps.Seeds[0], 
		seedsAndMaps.Seeds[1], 
		seedsAndMaps.MapSets)

	for i := 2 ; i < len(seedsAndMaps.Seeds) ; i += 2 {
		currLocation := GetLowestLocationFromSeedRange(
			seedsAndMaps.Seeds[i], 
			seedsAndMaps.Seeds[i + 1], 
			seedsAndMaps.MapSets)
		if currLocation < lowestLocation {
			lowestLocation = currLocation
		}
	}

	fmt.Println(lowestLocation)
}