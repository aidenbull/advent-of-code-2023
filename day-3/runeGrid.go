package main

import (
	"unicode"
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
