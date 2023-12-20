package main 

import (
	"strings"
	"io/ioutil"
)

func main() {
	filestr, _ := ioutil.ReadFile("input")
	preProcessedInput := strings.TrimSpace(string(filestr))
	PrintPartOneSolution(preProcessedInput)
}