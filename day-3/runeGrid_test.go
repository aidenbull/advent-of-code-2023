package main

import (
	"testing"
)

func TestConvertSingleRuneToNumeral(t *testing.T) {
	const testInput rune = '3'
	expectedResult, expectedError := 3, error(nil)
	result, err := ConvertRuneToInteger(testInput)
	if (result != expectedResult) || (err != expectedError) {
		t.Errorf("Expected (%d, %d), got (%d, %d)", expectedResult, expectedError, result, err)
	}
}

func TestConvertingNonDigitRuneToNumeralFails(t *testing.T) {
	const testInput rune = 'a'
	_, err := ConvertRuneToInteger(testInput)
	if err == nil  { //Having difficulty with finding a way to effectively compare results
		t.Errorf("Expected non-nil error, got %d", err)
	}
}

func TestConvertRunesToInteger(t *testing.T) {
	testInput := []rune{'3', '1'}
	expected := 31
	result, _ := ConvertRunesToInteger(testInput)

	if expected != result {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestFindIndexOfFirstDigitInRuneSlice(t *testing.T) {
	testInput := []rune{'.', '.', '0'}
	expected := 2
	result := FindFirstDigit(testInput)

	if expected != result {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestFindIndexOfFirstDigitFailsIfNoDigitsPresent(t *testing.T) {
	testInput := []rune{'.','&','$','-'}
	expected := -1
	result := FindFirstDigit(testInput)

	//Noticing a bit of duplication here. Might be worth investigating
	if expected != result {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestRuneIsAdjacentToSymbolWhenNotAdjacent(t *testing.T) {
	testInput := [][]rune{
		{'.', '.', '.', '&'},
		{'.', '6', '5', '.'},
		{'.', '.', '.', '.'},
	}
	expected := false
	result := AdjacentToSymbol(testInput, 1, 1)

	if expected != result {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}