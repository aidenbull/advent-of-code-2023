package main

import (
	"fmt"
)

func PrintPartOneSolution(input string) {
	seedsAndMaps := ParseSeedsAndMaps(input)

	lowestLocation := ApplyMultipleMaps(seedsAndMaps.Seeds[0], seedsAndMaps.MapSets)
	for _, seed := range seedsAndMaps.Seeds[1:] {
		currLocation := ApplyMultipleMaps(seed, seedsAndMaps.MapSets)
		if currLocation < lowestLocation {
			lowestLocation = currLocation
		}
	}

	fmt.Println(lowestLocation)
}