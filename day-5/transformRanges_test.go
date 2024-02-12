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

func TestConvertRangeWhenMappingsAndInputMeetAtRangeBorders(t *testing.T) {
	input_seed_range := SeedRange{0, 100}
	input_mappings := MapSet{[]FertilizerMap{
		{100, 0, 20},
		{130, 30, 20},
		{150, 50, 20},
		{180, 80, 20},
	}}

	expected := []SeedRange{
		{20, 30},
		{70, 80},
		{100, 120},
		{130, 170},
		{180, 200},
	}

	result := SortRangesByIncreasingStart(GetOutputRanges(input_seed_range, input_mappings))

	if (!cmp.Equal(expected, result)) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestConvertMultipleInputRangesToOutputRanges(t *testing.T) {
	input_ranges := []SeedRange{
		{0, 50},
		{50, 100},
	}
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

	result := SortRangesByIncreasingStart(
		TransformInputRangesToOutputRanges(
			input_ranges, 
			input_mappings,
		))

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

func TestInsertAndMergeRangeWhenNoIntersectButStillMergable(t *testing.T) {
	input_new_range := SeedRange{0, 5}
	input_existing_ranges := []SeedRange{
		{5, 10},
	}

	expected := []SeedRange{
		{0, 10},
	}
	result := InsertAndMergeRange(input_new_range, input_existing_ranges)

	if (!cmp.Equal(expected, result)) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestProcessPartTwoSeeds(t *testing.T) {
	input := []int {
		10, 5,
		20, 6,
		0, 2,
	}
	
	expected := []SeedRange{
		{10, 15},
		{20, 26},
		{0, 2},
	}
	result := ProcessPartTwoSeeds(input)

	if (!cmp.Equal(expected, result)) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}