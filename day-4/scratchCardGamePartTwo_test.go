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

func TestPartTwoThreeStartingCards(t *testing.T) {
	testInput := []ScratchGame {
		{[]int{1,2,3}, []int{1, 2, 5, 6}},
		{[]int{4,5,6}, []int{1, 4, 5, 6}},
		{[]int{7,8,9}, []int{1, 7}},
	}

	expected := 7
	result := CalculatePartTwo(testInput)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPartTwoWithDroppedDuplicates(t *testing.T) {
	testInput := []ScratchGame {
		{[]int{1,2}, []int{1, 2}}, // 2 matches, 1 card 
		{[]int{4,5,6}, []int{4, 5, 6}}, // 3 matches, 2 cards
		{[]int{7}, []int{7}}, // 1 match, 4 cards
		{[]int{7}, []int{1}}, // 0 matches, 7 cards
		{[]int{1}, []int{2}}, // 0 matches, 3 cards
	} //TOTAL : 17 cards

	expected := 17
	result := CalculatePartTwo(testInput)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}