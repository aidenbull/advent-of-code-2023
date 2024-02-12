package main

import (
	"fmt"
	"time"
)

func PrintPartTwoSolution(input string) {
	fmt.Println("\nPart two.")
	parse_time_start := time.Now()

	seedsAndMaps := ParseSeedsAndMaps(input)
	seedRanges := ProcessPartTwoSeeds(seedsAndMaps.Seeds)

	fmt.Printf("Time to parse: %s\n", time.Since(parse_time_start))
	calculation_time_start := time.Now()

	curr_input_ranges := seedRanges
	for _, mapSet := range(seedsAndMaps.MapSets) {
		curr_input_ranges = TransformInputRangesToOutputRanges(curr_input_ranges, mapSet)
	}
	lowestLocation := SortRangesByIncreasingStart(curr_input_ranges)[0].Start

	fmt.Printf("Time to calculate: %s\n", time.Since(calculation_time_start))
	fmt.Printf("Part two solution: %d\n", lowestLocation)
}