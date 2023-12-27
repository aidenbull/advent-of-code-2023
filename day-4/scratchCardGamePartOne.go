package main

import (
	"strings"
	"fmt"
)

func PrintPartOneSolution(input string) {
	tokens := strings.Split(input, "\n")
	sum := 0

	for _, gameStr := range tokens {
		game := ScratchGameFromString(gameStr)
		sum += CalculateValue(EvaluateGameMatches(game))
	}

	fmt.Println(sum)
}