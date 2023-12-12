package main

import (
	"strings"
	"fmt"
)

func gameIsValid(game CubeGame) bool {
	const redMax = 12
	const greenMax = 13
	const blueMax = 14
	for _, set := range game.Sets {
		if set.Red > redMax {
			return false
		}
		if set.Green > greenMax {
			return false
		}
		if set.Blue > blueMax {
			return false
		}
	}
	return true
}

func PrintPartOneSolution(input string) {
	processedInput := PreProcessGame(input)

	tokens := strings.Split(processedInput, "\n")
	sum := 0
	for _, token := range tokens {
		game := ProcessGame(token)
		if gameIsValid(game) {
			sum += game.GameNum
		}
	}

	fmt.Println(sum)
}