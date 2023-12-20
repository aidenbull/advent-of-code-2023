package main

import (
	"testing"
)

func TestReadScratchCardGameFromString(t *testing.T) {
	testInput := "Card   1: 43 19 57 | 34 68 13 38"

	expected := ScratchGame{WinningNums:{43, 19, 57}, YourNums:{34, 68, 13, 38}}
	result := ScratchGameFromString(testInput)

	if expected != result {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}