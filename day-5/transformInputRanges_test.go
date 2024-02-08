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
		{3, 4},
		{4, 6},
		{7, 9},
	}

	result := GetOutputRanges(input_seed_range, input_mapping)

	if (!cmp.Equal(expected, result)) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}