package main

import (
	"strings"
	"github.com/aidenbull/fileio"
)

func main() {
	file := fileio.OpenFile("input")
	fileString := strings.ToLower(fileio.ReadFile(file))
	// Trim any trailing and leading whitespace before splitting. It appears the input has one trailing newline
	tokens := strings.Split(strings.TrimSpace(fileString), "\n")
	PrintPartOneSolution(tokens)
	PrintPartTwoSolution(tokens)
}