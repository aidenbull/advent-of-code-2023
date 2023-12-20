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

func TestEvaluateGameValueWithValueZero(t *testing.T) {
	testInput := ScratchGame{[]int{1, 2, 3}, []int{4, 5, 6, 7}}

	expected := 0
	result := EvaluateGame(testInput)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestEvaluateGameValueWithPositiveValue(t *testing.T) {
	testInput := ScratchGame{[]int{1, 2, 3}, []int{3, 6, 7, 1}}

	expected := 2
	result := EvaluateGame(testInput)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}