package main

import (
	"testing"
	"reflect"
)

func TestConvertSingleRuneToNumeral(t *testing.T) {
	const testInput rune = '3'
	expectedResult, expectedError := 3, error(nil)
	result, err := ConvertRuneToInteger(testInput)
	if (result != expectedResult) || (err != expectedError) {
		t.Errorf("Expected (%v, %v), got (%v, %v)", expectedResult, expectedError, result, err)
	}
}

func TestConvertingNonDigitRuneToNumeralFails(t *testing.T) {
	const testInput rune = 'a'
	_, err := ConvertRuneToInteger(testInput)
	if err == nil  { //Having difficulty with finding a way to effectively compare results
		t.Errorf("Expected non-nil error, got %v", err)
	}
}

func TestConvertRunesToInteger(t *testing.T) {
	testInput := []rune{'3', '1'}
	expected := 31
	result, _ := ConvertRunesToInteger(testInput)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestRuneIsAdjacentToSymbolReturnsFalseWhenNotAdjacent(t *testing.T) {
	testInput := [][]rune{
		{'.', '.', '.', '&'},
		{'.', '6', '5', '.'},
		{'.', '.', '.', '.'},
	}
	expected := false
	result := AdjacentToSymbol(testInput, 1, 1)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestRuneIsAdjacentToSymbolReturnsTrueWhenIsAdjacent(t *testing.T) {
	testInput := [][]rune{
		{'.', '.', '.', '&'},
		{'.', '6', '5', '.'},
		{'.', '.', '.', '.'},
	}
	expected := true
	result := AdjacentToSymbol(testInput, 1, 2)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestRuneIsAdjacentToSymbolIgnoresTheSymbolUnderTest(t *testing.T) {
	testInput := [][]rune{
		{'.', '.', '.', '&'},
		{'.', '6', '5', '.'},
		{'.', '.', '.', '.'},
	}
	expected := false
	result := AdjacentToSymbol(testInput, 0, 3)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestNumberIsAdjacentToSymbolWhenAdjacent(t *testing.T) {
	testInput := [][]rune{
		{'.', '.', '.', '&'},
		{'.', '6', '5', '.'},
		{'.', '.', '.', '.'},
	}
	expected := true
	result := NumberIsAdjacentToSymbol(testInput, 1, 1, 3)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestNumberIsAdjacentToSymbolWhenNotAdjacent(t *testing.T) {
	testInput := [][]rune{
		{'.', '.', '.', '.'},
		{'.', '6', '5', '.'},
		{'.', '.', '.', '.'},
	}
	expected := false
	result := NumberIsAdjacentToSymbol(testInput, 1, 1, 3)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGetNumberIndicesFromRuneSlice(t *testing.T) {
	testInput := []rune{'.','1','3','*','a','6','2','1','.','.'}
	expected := []SliceCoords{{1,3}, {5, 8}}
	result := GetNumberCoordsFromRuneSlice(testInput)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestConvertStringToRuneGrid (t *testing.T) {
	testInput := "Hello.\nIt'sme"
	expected := [][]rune{
		{'H','e','l','l','o','.'},
		{'I','t','\'','s','m','e'},
	}
	result := ConvertStringToRuneGrid(testInput)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}