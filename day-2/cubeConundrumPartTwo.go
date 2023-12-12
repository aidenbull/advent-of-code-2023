package main

import (
	"strings"
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func powerSetOfGame(game CubeGame) int {
	redMax := 0
	greenMax := 0
	blueMax := 0

	for _, set := range game.Sets {
		redMax = max(set.Red, redMax)
		greenMax = max(set.Green, greenMax)
		blueMax = max(set.Blue, blueMax)
	}

	return redMax * greenMax * blueMax
}

func PrintPartTwoSolution(input string) {
	processedInput := PreProcessGame(input)

	tokens := strings.Split(processedInput, "\n")
	sum := 0
	for _, token := range tokens {
		game := ProcessGame(token)
		sum += powerSetOfGame(game)
	}

	fmt.Println(sum)
}