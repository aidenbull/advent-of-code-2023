package main

import (
	"regexp"
	"strings"
	"strconv"
	"fmt"
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
	winningNumsStrings := strings.Split(strings.TrimSpace(gameComponents[0]), " ")
	yourNumsStrings := strings.Split(strings.TrimSpace(gameComponents[1]), " ")

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