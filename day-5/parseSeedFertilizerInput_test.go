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
