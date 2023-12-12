package main

import (
	"regexp"
	"strings"
	"strconv"
	"fmt"
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
		panic("Oh it's happening! Oh my god it's happening! Everybody STAY FUCKING CALM!")
	}
	return gameNum
}

// Takes a single line of game input as a string and returns a
// struct containing the information about that game
func processGame(gameInput string) CubeGame {
	var game CubeGame
	game.GameNum = processGameNum(gameInput)
	inputRunes := []rune(gameInput)
	
	matchPrefix := regexp.MustCompile(`^Game [0-9]+: `)
	prefixStringRunes := []rune(matchPrefix.FindString(gameInput))	
	
	setsRunes := inputRunes[len(prefixStringRunes):]
	game.Sets = processSets(string(setsRunes))

	return game
}

func preProcessGame(input string) string {
	return strings.TrimSpace(strings.ToLower(input))
}

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
	processedInput := preProcessGame(input)

	tokens := strings.Split(processedInput, "\n")
	sum := 0
	for _, token := range tokens {
		game := processGame(token)
		if gameIsValid(game) {
			sum += game.GameNum
		}
	}

	fmt.Println(sum)
}