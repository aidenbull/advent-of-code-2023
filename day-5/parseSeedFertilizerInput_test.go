package main

import (
	"testing"
	"github.com/google/go-cmp/cmp"
)

func TestParseOutSeeds(t *testing.T) {
	input := "seeds: 1 2 10 30"

	expected := []int{1, 2, 10, 30}
	result := ParseSeeds(input)

	if !cmp.Equal(expected, result) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestParseOutSingleSetOfMappings(t *testing.T) {
	input := "seed-to-soil map:\n4 5 2\n1 1 1"

	expected := MapSet{
		[]FertilizerMap{
			{4, 5, 2},
			{1, 1, 1},
		},
	}
	result := ParseSetOfMappings(input)

	if !cmp.Equal(expected, result) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestParseOutSeedsAndMultipleMapSets(t *testing.T) {
	input := "seeds: 4 1 3\n" + 
	"\n" +
	"seed-to-soil map:\n" +
	"4 5 2\n" +
	"1 1 1\n" + 
	"\n" + 
	"soil-to-fertilizer map:\n" +
	"1 4 7\n" +
	"1 9 3"

	expectedSeeds := []int {4, 1, 3}
	expectedMapSet1 := MapSet{[]FertilizerMap{
		{4, 5, 2},
		{1, 1, 1},
	}}
	expectedMapSet2 := MapSet{[]FertilizerMap{
		{1, 4, 7},
		{1, 9, 3},
	}}

	expected := SeedsAndMaps{
		expectedSeeds,
		[]MapSet {
			expectedMapSet1,
			expectedMapSet2,
		},
	}

	result := ParseSeedsAndMaps(input)

	if !cmp.Equal(expected, result) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestApplyMapToInput(t *testing.T) {
	testInput := 1
	testMap := MapSet{[]FertilizerMap{
		{10, 0, 5},
	}}

	expected := 11
	result := ApplyMap(testInput, testMap)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestApplyMapToInputWhenNoMatch(t *testing.T) {
	testInput := 5
	testMap := MapSet{[]FertilizerMap{
		{10, 0, 5},
	}}

	expected := 5
	result := ApplyMap(testInput, testMap)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestApplyMultipleMapsToInput(t *testing.T) {
	testInput := 1
	testMapSets := []MapSet{
		{[]FertilizerMap{
			{10, 0, 5},
		}},
		{[]FertilizerMap{
			{-10, 10, 2},
		}},
	}

	expected := -9
	result := ApplyMultipleMaps(testInput, testMapSets) 

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestApplySeedRangeToMultipleMaps(t *testing.T) {
	testInputStart := 5
	testInputRangeLen := 10
	testMapSets := []MapSet{
		{[]FertilizerMap{
			{72, 0, 10},
		}},
		{[]FertilizerMap{
			{62, 10, 10},
		}},
	}

	expected := 62
	result := ApplySeedRangeToMultipleMaps(testInputStart, testInputRangeLen, testMapSets)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}