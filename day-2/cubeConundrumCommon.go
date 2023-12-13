package main

import (
	"regexp"
	"strconv"
	"strings"
)

type CubeSet struct {
	Red int
	Green int
	Blue int
}

type CubeGame struct {
	GameNum int
	Sets []CubeSet
}

// Takes a single set within a game as a string and returns
// a struct describing the cubes drawn in that set
func processSet(setInput string) CubeSet {
	var set CubeSet
	matchColourCount := regexp.MustCompile(`[0-9]+ (red|green|blue)`)
	colourStrs := matchColourCount.FindAllString(setInput, -1)

	if len(colourStrs) > 3 {
		panic("Somehow there are more than 3 colours in this set: " + setInput)
	}

	matchCount := regexp.MustCompile(`[0-9]+`)
	matchColour := regexp.MustCompile(`(red|green|blue)`)
	for _, colourStr := range colourStrs {
		thisColour := matchColour.FindString(colourStr)
		thisCount, err := strconv.Atoi(matchCount.FindString(colourStr))

		if err != nil {
			panic("Failed to convert string to integer: " + err.Error())
		}

		switch thisColour {
		case "red":
			set.Red = thisCount
		case "green":
			set.Green = thisCount
		case "blue":
			set.Blue = thisCount
		}
	}

	return set
}

func processSets(setsInput string) []CubeSet {
	setStrings := strings.Split(setsInput, ";")
	sets := make([]CubeSet, 0)
	for _, set := range setStrings {
		sets = append(sets, processSet(set))
	}
	return sets
}

func processGameNum(gameInput string) int {
	firstNumMatch := regexp.MustCompile(`[0-9]+`)
	gameNumStr := firstNumMatch.FindString(gameInput)
	gameNum, err := strconv.Atoi(gameNumStr)
	if err != nil {
		panic("Oh my god! Okay it's happening! Everybody stay calm. Stay calm. STAY FUCKING CALM!")
	}
	return gameNum
}

// Takes a single line of game input as a string and returns a
// struct containing the information about that game
func ProcessGame(gameInput string) CubeGame {
	var game CubeGame
	game.GameNum = processGameNum(gameInput)
	inputRunes := []rune(gameInput)
	
	matchPrefix := regexp.MustCompile(`^Game [0-9]+: `)
	prefixStringRunes := []rune(matchPrefix.FindString(gameInput))	
	
	setsRunes := inputRunes[len(prefixStringRunes):]
	game.Sets = processSets(string(setsRunes))

	return game
}

func PreProcessGame(input string) string {
	return strings.TrimSpace(strings.ToLower(input))
}