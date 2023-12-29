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
