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

func ConvertRunesToInteger(runes []rune) (int, error) {

	return 0, nil
}
