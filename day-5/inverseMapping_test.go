//Tests for helping with part 2

package main

import (
	"testing"
	"github.com/google/go-cmp/cmp"
)

func TestCreateInverseMappingForSingleMapping(t *testing.T) {
	input := MapSet{[]FertilizerMap{
		{4, 5, 2},
	}}

	expected := MapSet{[]FertilizerMap{
		{5, 4, 2},
	}}

	result := InvertMapSet(input)

	if (!cmp.Equal(expected, result)) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestCreateInverseMappingForMultipleMappings(t *testing.T) {
	input := MapSet{[]FertilizerMap{
		{4, 5, 2},
		{10, 20, 5},
	}}

	expected := MapSet{[]FertilizerMap{
		{5, 4, 2},
		{20, 10, 5},
	}}

	result := InvertMapSet(input)

	if (!cmp.Equal(expected, result)) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestApplyingMappingsToSeedAndApplyingInverseGetsOriginalSeed(t *testing.T) {
	input_seed_forward := 5
	input_mapping := MapSet{[]FertilizerMap{
		{4, 5, 2},
	}}
	expected_output_forward := 4

	result_forward := ApplyMap(input_seed_forward, input_mapping)

	if (result_forward != expected_output_forward) {
		t.Errorf("Expected %v, got %v", expected_output_forward, result_forward)
	}

	input_seed_backward := 4
	inverse_mapping := InvertMapSet(input_mapping)
	expected_output_backward := 5

	result_backward := ApplyMap(input_seed_backward, inverse_mapping)

	if (result_backward != expected_output_backward) {
		t.Errorf("Expected %v, got %v", expected_output_backward, result_backward)
	}
}

func TestCreateInverseOfMultipleSetsOfMappings (t *testing.T) {
	input := []MapSet{
		{[]FertilizerMap{
			{4, 5, 2},
		}},
		{[]FertilizerMap{
			{10, 20, 5},
		}},
	}

	expected := []MapSet{
		{[]FertilizerMap{
			{20, 10, 5},
		}},
		{[]FertilizerMap{
			{5, 4, 2},
		}},
	}

	result := InvertMultipleMapSets(input)

	if (!cmp.Equal(expected, result)) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}