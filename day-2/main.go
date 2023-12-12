package main

import (
	"github.com/aidenbull/fileio"
)

func main() {
	file := fileio.OpenFile("input")
	fileString := fileio.ReadFile(file)
	PrintPartOneSolution(fileString)
}
