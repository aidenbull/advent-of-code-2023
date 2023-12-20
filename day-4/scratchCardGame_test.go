package main

import (
	"testing"
	"github.com/google/go-cmp/cmp"
)

func TestReadScratchCardGameFromString(t *testing.T) {
	testInput := "Card   1: 43 19 57 | 34 68 13 38"

	expected := ScratchGame{[]int{43, 19, 57}, []int{34, 68, 13, 38}}
	result := ScratchGameFromString(testInput)

	if !cmp.Equal(expected, result) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestEvaluateGameMatchesWithValueZero(t *testing.T) {
	testInput := ScratchGame{[]int{1, 2, 3}, []int{4, 5, 6, 7}}

	expected := 0
	result := EvaluateGameMatches(testInput)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestEvaluateGameMatchesWithPositiveValue(t *testing.T) {
	testInput := ScratchGame{[]int{1, 2, 3}, []int{3, 6, 7, 1}}

	expected := 2
	result := EvaluateGameMatches(testInput)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

//Breaking some testing rules (no overlapping tests) because this is quick to write.
// Just doing this for AdventOfCode because I know the problem domain isn't changing
func TestCalculateGameValueGivenMatches(t *testing.T) {
	testInput1 := 0
	testInput2 := 10

	expected1 := 0 
	expected2 := 512

	result1 := CalculateValue(testInput1)
	result2 := CalculateValue(testInput2)
	
	if 	expected1 != result1 {
		t.Errorf("Expected %v, got %v", expected1, result1)
	}
	
	if 	expected2 != result2 {
		t.Errorf("Expected %v, got %v", expected2, result2)
	}
}