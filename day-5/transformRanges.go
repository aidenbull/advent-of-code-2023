package main

//Start is inclusive, End is exclusive
type SeedRange struct {
	Start int;
	End int;
}

func shiftRange(r SeedRange, shift int) SeedRange {
	return SeedRange{r.Start + shift, r.End + shift}
}

func rangeIntersect(r1, r2 SeedRange) SeedRange {
	return SeedRange{max(r1.Start, r2.Start), min(r1.End, r2.End)}
}

func getOutputRangesRecursive(seeds SeedRange, mapSet MapSet, curr_ranges []SeedRange) []SeedRange {
	if seeds.Start == seeds.End {
		return curr_ranges
	}

	for i := 0; i<len(mapSet.Maps); i++ {
		curr_map := mapSet.Maps[i]
		curr_map_input_range := SeedRange{curr_map.SrcStart, curr_map.SrcStart + curr_map.Length}
		if (rangesOverlap(seeds, curr_map_input_range)) {
			//Intersect of overlap
			intersect := rangeIntersect(seeds, curr_map_input_range)
			curr_ranges = InsertAndMergeRange(
				shiftRange(intersect, (curr_map.DestStart - curr_map.SrcStart)), 
				curr_ranges,
			)

			//Minor optimization. Don't need to search the maps up to this point again
			remainingMaps := MapSet{mapSet.Maps[i+1:]} 

			//Left 
			left := SeedRange{seeds.Start, intersect.Start}
			curr_ranges = getOutputRangesRecursive(
				left,
				remainingMaps,
				curr_ranges,
			) 

			//Right
			right := SeedRange{intersect.End, seeds.End}
			curr_ranges = getOutputRangesRecursive(
				right,
				remainingMaps,
				curr_ranges,
			)

			return curr_ranges
		}
	}

	return InsertAndMergeRange(seeds, curr_ranges)
}

func GetOutputRanges(seeds SeedRange, mapSet MapSet) []SeedRange {
	return getOutputRangesRecursive(seeds, mapSet, []SeedRange{})
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func rangesOverlap(r1, r2 SeedRange) bool {
	return r1.Start < r2.End && r2.Start < r1.End
}

func rangeUnion(r1, r2 SeedRange) SeedRange {
	return SeedRange{min(r1.Start, r2.Start), max(r1.End, r2.End)}
}

func InsertAndMergeRange(new_range SeedRange, existing_ranges []SeedRange) []SeedRange {
	for i := 0; i<len(existing_ranges); i++ {
		if rangesOverlap(new_range, existing_ranges[i]){
			merged_range := rangeUnion(new_range, existing_ranges[i])
			existing_ranges[i] = existing_ranges[len(existing_ranges)-1]
			return InsertAndMergeRange(
				merged_range, 
				existing_ranges[:len(existing_ranges)-1],
			)
		}
	}
	return append(existing_ranges, new_range)
}