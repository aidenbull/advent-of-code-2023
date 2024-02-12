package main

import (
	"testing"
	"github.com/google/go-cmp/cmp"
)

func TestConvertInputRangeToOutputRangeForSingleMap(t *testing.T) {
	input_seed_range := SeedRange{3, 9}
	input_mapping := MapSet{[]FertilizerMap{
		{4, 5, 2},
	}}

	expected := []SeedRange{
		{3, 6},
		{7, 9},
	}

	result := GetOutputRanges(input_seed_range, input_mapping)

	if (!cmp.Equal(expected, result)) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestConvertRangeWithTwoDistinctMappingsToCreateFiveRanges(t *testing.T) {
	input_seed_range := SeedRange{0, 100}
	input_mappings := MapSet{[]FertilizerMap{
		{120, 20, 20},
		{160, 60, 20},
	}}

	expected := []SeedRange{
		{0, 20},
		{40, 60},
		{80, 100},
		{120, 140},
		{160, 180},
	}

	result := SortRangesByIncreasingStart(GetOutputRanges(input_seed_range, input_mappings))

	if (!cmp.Equal(expected, result)) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestInsertAndMergeSuccessfullyMergesTwoRanges(t *testing.T) {
	input_new_range := SeedRange{5, 9}
	input_existing_ranges := []SeedRange{
		{7, 13},
	}

	expected := []SeedRange{
		{5, 13},
	}
	result := InsertAndMergeRange(input_new_range, input_existing_ranges)

	if (!cmp.Equal(expected, result)) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestInsertAndMergeRangeWorksWhenTwoMergesWouldOccur(t *testing.T) {
	input_new_range := SeedRange{5, 9}
	input_existing_ranges := []SeedRange{
		{2, 6},
		{8, 13},
	}

	expected := []SeedRange{
		{2, 13},
	}
	result := InsertAndMergeRange(input_new_range, input_existing_ranges)

	if (!cmp.Equal(expected, result)) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}