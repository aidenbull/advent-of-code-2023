package main

import (
	"strings"
)

func main() {
	file := OpenFile("input")
	fileString := strings.ToLower(ReadFile(file))
	// Trim any trailing and leading whitespace before splitting. It appears the input has one trailing newline
	tokens := strings.Split(strings.TrimSpace(fileString), "\n")
	PrintPartOneSolution(tokens)
	PrintPartTwoSolution(tokens)
}