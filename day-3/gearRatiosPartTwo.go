package main

import (
	"unicode"
	"fmt"
)

func GetNumberAtIndex(runeSlice []rune, index int) (int, error) {
	startIndex := index
	endIndex := index + 1

	for (startIndex - 1) >= 0 && unicode.IsDigit(runeSlice[startIndex - 1]) {
		startIndex--
	}
	for endIndex < len(runeSlice) && unicode.IsDigit(runeSlice[endIndex]) {
		endIndex++ 
	}

	return ConvertRunesToInteger(runeSlice[startIndex:endIndex])
}

func FindAdjacentNumbers(runeGrid [][]rune, y, x int) []int {
	out := make([]int, 0)
	
	yBorder := []int{-1, 0, 1}
	xBorder := []int{-1, 0, 1}
	readingNumber := false

	for _, j := range yBorder {
		for _, i := range xBorder {
			currY := y + j
			currX := x + i

			if 	(i == 0 && j == 0) {
				//Gotta throw this in case there are numbers on both sides of the middle row
				readingNumber = false
				continue
			}
			if 	(currY < 0 || currY >= len(runeGrid)) ||
				(currX < 0 || currX >= len(runeGrid[currY])) {
				continue
			}

			if unicode.IsDigit(runeGrid[currY][currX]) {
				if !readingNumber {
					newNum, _ := GetNumberAtIndex(runeGrid[currY], currX)
					out = append(out, newNum)
				}
				readingNumber = true
			} else {
				readingNumber = false
			}
		}
		//Moving onto next line so will be reading a new number
		readingNumber = false
	}

	return out
}

func PrintPartTwoSolution(input string) {
	runeGrid := ConvertStringToRuneGrid(input)
	sum := 0

	for j, line := range runeGrid {
		for i, r := range line {
			if r == '*' {
				adjNums := FindAdjacentNumbers(runeGrid, j, i)
				if len(adjNums) == 2 {
					sum += adjNums[0] * adjNums[1]
				}
			}
		}
	}

	fmt.Println(sum)
}