package main

type DigitNotFound struct {
	s string
}

func (e DigitNotFound) Error() string {
	return e.s
}