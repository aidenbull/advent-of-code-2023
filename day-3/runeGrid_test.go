package main

import (
	"testing"
)

func TestConvertRuneToNumeral(t *testing.T) {
	const testInput1 rune = '3'
	expectedResult1, expectedError1 := 3, nil
	result1, error1 := ConvertRuneToInteger() {
		
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