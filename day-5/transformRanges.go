package main

//Start is inclusive, End is exclusive
type SeedRange struct {
	Start int;
	End int;
}

func GetOutputRanges(seeds SeedRange, mapSet MapSet) []SeedRange {
	return []SeedRange{}
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

func mergeOverlappingRange(r1, r2 SeedRange) SeedRange {
	return SeedRange{min(r1.Start, r2.Start), max(r1.End, r2.End)}
}

func InsertAndMergeRange(new_range SeedRange, existing_ranges []SeedRange) []SeedRange {
	for i := 0; i<len(existing_ranges); i++ {
		if rangesOverlap(new_range, existing_ranges[i]){
			merged_range := mergeOverlappingRange(new_range, existing_ranges[i])
			existing_ranges[i] = existing_ranges[len(existing_ranges)-1]
			return InsertAndMergeRange(
				merged_range, 
				existing_ranges[:len(existing_ranges)-1],
			)
		}
	}
	return append(existing_ranges, new_range)
}