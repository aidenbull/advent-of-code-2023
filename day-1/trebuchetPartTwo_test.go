package main

import (
	"testing"
	"reflect"
)

func TestFindAllKeywords(t *testing.T) {
	testStr := "helloworld3zero1fortysevenineight"
	expected := []string{"3", "zero", "1", "seven", "nine", "eight"}
	result := findAllKeywords(testStr)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}