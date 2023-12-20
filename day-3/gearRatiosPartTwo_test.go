package main

import (
	"testing"
	"github.com/google/go-cmp/cmp"
)

func TestGetNumberAtSliceIndex(t *testing.T) {
	testInputSlice := []rune {'.','3','5','4','9','.','.'}
	testInputIndex := 4
	
	expected := 3549
	result, _ := GetNumberAtIndex(testInputSlice, testInputIndex)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFindAdjacentNumbers(t *testing.T) {
	testInput := [][]rune {
		{'5','2','.','.','.'},
		{'.','.','*','.','.'},
		{'.','1','0','3','.'},
	}
	
	expected := []int{52, 103}
	result := FindAdjacentNumbers(testInput, 1, 2)

	if !cmp.Equal(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}