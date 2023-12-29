package main 

import (
	"io/ioutil"
)

func main() {
	fileStr, _ := ioutil.ReadFile("input")
	PrintPartOneSolution(string(fileStr))
	PrintPartTwoSolution(string(fileStr))
}