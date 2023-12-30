package main

import (
	"testing"
	"github.com/google/go-cmp/cmp"
)

func TestConvertInputToRanges(t *testing.T) {
	input := []int{0,5,12,3,6,4}

	expected := []Range{{0, 5}, {12, 15}, {6, 10}}
	result := ConvertInputToRanges(input)

	if !cmp.Equal(expected, result) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}