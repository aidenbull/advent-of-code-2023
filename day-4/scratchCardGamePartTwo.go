package main

import (
)

func CalculatePartTwo(games []ScratchGame) int {
	sum := 0
	matches := make([]int, 0)

	for _, game := range games {
		sum += 1 + len(matches)
		matches = append(matches, EvaluateGameMatches(game))
	}
	
	return sum
}