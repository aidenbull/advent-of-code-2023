package main

import (
)

type ActiveDuplication struct {
	DuplicateQueueSize int
	NumberOfDuplicates int
}

func calculateCurrentNumberOfCards(activeDuplications []ActiveDuplication) int {
	currentMultiples := 1

	for _, duplication := range activeDuplications {
		currentMultiples += duplication.NumberOfDuplicates
	}

	return currentMultiples
}

func CalculatePartTwo(games []ScratchGame) int {
	sum := 0
	activeDuplications := make([]ActiveDuplication, 0) 

	for _, game := range games {
		currentMultiples := calculateCurrentNumberOfCards(activeDuplications)

		//Determine number of cards at this stage, and update the active duplications
		nextActiveDuplications := make([]ActiveDuplication, 0)
		for _, duplication := range activeDuplications {
			duplication.DuplicateQueueSize--
			if duplication.DuplicateQueueSize > 0 {
				nextActiveDuplications = append(nextActiveDuplications, duplication)
			}
		}
		activeDuplications = nextActiveDuplications

		sum += currentMultiples
		
		matches := EvaluateGameMatches(game)
		if matches > 0 {
			activeDuplications = append(activeDuplications, ActiveDuplication{matches, currentMultiples}) 
		}
	}
	
	return sum
}