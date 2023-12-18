package main

import (
	"fmt"
)

func PrintPartOneSolution(input string) {
	runeGrid := ConvertStringToRuneGrid(input)
	sum := 0
	for i, line := range runeGrid {
		numberCoords := GetNumberCoordsFromRuneSlice(line)
		for _, num := range numberCoords {
			if NumberIsAdjacentToSymbol(runeGrid, i, num.Start, num.End) {
				foundInt, _ := ConvertRunesToInteger(line[num.Start:num.End])
				sum += foundInt
			}
		}
	}
	fmt.Println(sum)
}