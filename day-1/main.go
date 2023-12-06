package main

import (
	"strings"
)

func main() {
	file := OpenFile("input")
	fileString := ReadFile(file)
	// Trim trailing and leading whitespace and split by newlines
	tokens := strings.Split(strings.TrimSpace(fileString), "\n")
	PrintPartOneSolution(tokens)
}