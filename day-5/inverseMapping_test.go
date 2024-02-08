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

	result := InvertMapping(input)

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

	result := InvertMapping(input)

	if (!cmp.Equal(expected, result)) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}