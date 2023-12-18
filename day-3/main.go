package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	filestr, _ := ioutil.ReadFile("input")
	preProcessedInput := strings.TrimSpace(string(filestr))
	PrintPartOneSolution(preProcessedInput)
}
