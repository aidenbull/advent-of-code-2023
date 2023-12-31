package main

import (
	"unicode"
	"strings"
)

type BadConversion struct {
	s string
}

func (e *BadConversion) Error() string {
	return e.s
}

func ConvertRuneToInteger(r rune) (int, error) {
	isDigit := unicode.IsDigit(r)
	if !isDigit {
		return 0, &BadConversion{"Input is not a digit!"}
	}
	return int(r - '0'), nil
}

//Does not support negative powers
func intPow(x, y int) int {
	total := 1
	for i := 0; i<y; i++ {
		total *= x
	}
	return total
}

//Assumes base10 number (at least for the moment)
func ConvertRunesToInteger(runes []rune) (int, error) {
	result := 0
	for i, r := range runes {
		out, err := ConvertRuneToInteger(r)
		if (err != nil) {
			return 0, err
		}
		result += out * intPow(10, len(runes) - (i + 1))
	}

	return result, nil
}

func isNotableSymbol(r rune) bool {
	return r != '.' && !unicode.IsDigit(r)
}

func AdjacentToSymbol(runeGrid [][]rune, y, x int) bool {
	yBorder := []int{-1, 0, 1}
	xBorder := []int{-1, 0, 1}
	foundSymbol := false
	for _, j := range yBorder {
		for _, i := range xBorder {
			//Ignore the symbol being checked
			if i == 0 && j == 0 {
				continue
			}
			currY := y + j
			currX := x + i
			//Ignore the symbol if the current y coord is outside of the slice range
			if currY < 0 || currY >= len(runeGrid) {
				continue
			}
			//Same for x (need to do this second because accessing an index)
			if currX < 0 || currX >= len(runeGrid[currY]) {
				continue
			}
			foundSymbol = foundSymbol || isNotableSymbol(runeGrid[currY][currX])
		}
	}
	return foundSymbol
}

func NumberIsAdjacentToSymbol(runeGrid [][]rune, lineIndex, xStart, xEnd int) bool {
	out := false
	for i := xStart; i < xEnd; i++ {
		out = AdjacentToSymbol(runeGrid, lineIndex, i) || out
	}
	return out
}

type SliceCoords struct {
	Start int
	End int
}

func GetNumberCoordsFromRuneSlice(runes []rune) []SliceCoords {
	out := make([]SliceCoords, 0)
	currStartIndex := 0
	readingNumber := false
	for i, r := range runes {
		if (unicode.IsDigit(r)) {
			if !readingNumber {
				currStartIndex = i
				readingNumber = true
			}
		} else {
			if readingNumber {
				out = append(out, SliceCoords{currStartIndex, i})
				readingNumber = false
			}
		}
	}
	if readingNumber {
		out = append(out, SliceCoords{currStartIndex, len(runes)})
	}
	return out
}

func ConvertStringToRuneGrid(input string) [][]rune {
	out := make([][]rune, 0)
	lines := strings.Split(input, "\n")
	for _, str := range lines {
		out = append(out, []rune(str))
	}
	return out
}