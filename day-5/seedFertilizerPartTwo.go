package main

import (
	"fmt"
)

func PrintPartTwoSolution(input string) {
	seedsAndMaps := ParseSeedsAndMaps(input)
	seedRanges := ProcessPartTwoSeeds(seedsAndMaps.Seeds)
	curr_input_ranges := seedRanges
	for _, mapSet := range(seedsAndMaps.MapSets) {
		curr_input_ranges = TransformInputRangesToOutputRanges(curr_input_ranges, mapSet)
	}
	lowestLocation := SortRangesByIncreasingStart(curr_input_ranges)[0].Start
	fmt.Println(lowestLocation)
}