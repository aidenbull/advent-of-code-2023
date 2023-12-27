package main

import (
	"strings"
	"fmt"
)

type ActiveDuplication struct {
	RemainingQueueSize int
	NumberOfDuplicates int
}

func calculateNumberOfCardsAtThisStep(activeDuplications []ActiveDuplication) int {
	numCardsAtThisStep := 1

	for _, duplication := range activeDuplications {
		numCardsAtThisStep += duplication.NumberOfDuplicates
	}

	return numCardsAtThisStep
}

func decrementAndRemoveActiveDuplications(activeDuplications []ActiveDuplication) []ActiveDuplication {
	nextActiveDuplications := make([]ActiveDuplication, 0)
	for _, duplication := range activeDuplications {
		duplication.RemainingQueueSize--
		if duplication.RemainingQueueSize > 0 {
			nextActiveDuplications = append(nextActiveDuplications, duplication)
		}
	}
	return nextActiveDuplications
}

func CalculatePartTwo(games []ScratchGame) int {
	sum := 0
	activeDuplications := make([]ActiveDuplication, 0) 

	for _, game := range games {
		numCardsAtThisStep := calculateNumberOfCardsAtThisStep(activeDuplications)
		sum += numCardsAtThisStep

		activeDuplications = decrementAndRemoveActiveDuplications(activeDuplications)

		matches := EvaluateGameMatches(game)
		if matches > 0 {
			activeDuplications = append(activeDuplications, ActiveDuplication{matches, numCardsAtThisStep}) 
		}
	}
	
	return sum
}

func PrintPartTwoSolution(input string) {
	tokens := strings.Split(input, "\n")
	
	games := make([]ScratchGame, 0)
	for _, gameStr := range tokens {
		games = append(games, ScratchGameFromString(gameStr))
	}

	fmt.Println(CalculatePartTwo(games))
}