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
		sum += EvaluateGameMatches(game)

		//fmt.Printf("Game: %+v, Value: %d\n", game, EvaluateGame(game))
	}

	fmt.Println(sum)
}