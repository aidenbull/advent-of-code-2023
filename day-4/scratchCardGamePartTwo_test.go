package main

import (
	"testing"
)

func TestPartTwoOneStartingCard(t *testing.T) {
	testInput := []ScratchGame {
		{[]int{1,2,3}, []int{1, 4, 5, 6}},
	}

	expected := 1
	result := CalculatePartTwo(testInput)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}


func TestPartTwoTwoStartingCards(t *testing.T) {
	testInput := []ScratchGame {
		{[]int{1,2,3}, []int{1, 4, 5, 6}},
		{[]int{4,5,6}, []int{1, 4, 5, 6}},
	}

	expected := 3
	result := CalculatePartTwo(testInput)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}