package main

import (
	"fmt"
	"time"
)

func PrintPartOneSolution(input string) {
	fmt.Println("\nPart one.")
	parse_time_start := time.Now()

	seedsAndMaps := ParseSeedsAndMaps(input)

	fmt.Printf("Time to parse: %s\n", time.Since(parse_time_start))
	calculation_time_start := time.Now()

	lowestLocation := ApplyMultipleMaps(seedsAndMaps.Seeds[0], seedsAndMaps.MapSets)
	for _, seed := range seedsAndMaps.Seeds[1:] {
		currLocation := ApplyMultipleMaps(seed, seedsAndMaps.MapSets)
		if currLocation < lowestLocation {
			lowestLocation = currLocation
		}
	}

	fmt.Printf("Time to calculate: %s\n", time.Since(calculation_time_start))
	fmt.Printf("Part one answer: %d\n", lowestLocation)
}