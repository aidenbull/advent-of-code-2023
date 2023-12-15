package main

import (
	"testing"
)

func TestConvertSingleRuneToNumeral(t *testing.T) {
	const testInput1 rune = '3'
	expectedResult1, expectedError1 := 3, error(nil)
	result1, error1 := ConvertRuneToInteger(testInput1)
	if (result1 != expectedResult1) || (error1 != expectedError1) {
		t.Errorf("Expected (%d, %d), got (%d, %d)", expectedResult1, expectedError1, result1, error1)
	}
}

// func TestConvertRunesToInteger(t *testing.T) {
// 	testInput := []rune{'3', '1'}
// 	expected := 31
// 	result, _ := ConvertRunesToInteger(testInput)

// 	if expected != result {
// 		t.Errorf("Expected %d, got %d", expected, result)
// 	}
// }
