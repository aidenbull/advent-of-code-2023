package main

import (
	"regexp"
	"strings"
	"strconv"
)

type ScratchGame struct {
	WinningNums []int
	YourNums []int
}

func ScratchGameFromString(input string) ScratchGame {
	//Remove prefix
	matchPrefix := regexp.MustCompile(`Card +[0-9]+:`)
	gamePrefix := matchPrefix.FindString(input)
	game := strings.TrimPrefix(input, gamePrefix)
	
	//Split winning nums from your nums
	gameComponents := strings.Split(game, "|")
	
	//Get tokens from components
	matchSpaces := regexp.MustCompile(`\s+`)
	winningNumsStrings := matchSpaces.Split(strings.TrimSpace(gameComponents[0]), -1)
	yourNumsStrings := matchSpaces.Split(strings.TrimSpace(gameComponents[1]), -1)

	var out ScratchGame
	for _, numStr := range winningNumsStrings {
		num, _ := strconv.Atoi(numStr)
		out.WinningNums = append(out.WinningNums, num)
	}
	for _, numStr := range yourNumsStrings {
		num, _ := strconv.Atoi(numStr)
		out.YourNums = append(out.YourNums, num)
	}

	return out
}

func intPow(x, y int) int {
	total := 1
	for i := 0; i < y; i++ {
		total *= x
	}
	return total
}

func EvaluateGame(game ScratchGame) int {
	numMatches := 0

	for _, winningNum := range game.WinningNums {
		for _, yourNum := range game.YourNums {
			if winningNum == yourNum {
				numMatches++
				break
			}
		}
	}

	if numMatches == 0 {
		return 0
	}

	return intPow(2, numMatches - 1)
}