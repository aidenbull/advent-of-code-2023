package main

import(
	"fmt"
)

// 'zero' and '0' never actually appear in the input but it lets me do a little index trick
var keywords = [...]string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
} 

func getIntFromWordOrNumeral(input string) (int, error) {
	for i, str := range keywords {
		if (input == str) {
			return i % 10, nil
		}
	}
	return -1, DigitNotFound{"getIntFromWordOrNumeral: Token '" + input + "'didn't match any keywords"}
}

func findAllKeywords(input string) []string {
	runes := []rune(input)
	foundKeywords := make([]string, 0)
	for i, _ := range runes {
		for _, keyword := range keywords {
			keywordRunes := []rune(keyword)
			if i+len(keywordRunes) > len(runes) {
				continue
			}
 			if keyword == string(runes[i:i+len(keywordRunes)]) {
				foundKeywords = append(foundKeywords, keyword)
				break
			}
		}
	}
	
	return foundKeywords
}

func GetFirstAndLastDigitsPartTwo(input string) (int, int, error) {
	matches := findAllKeywords(input)
	if len(matches) == 0 {
		return -1, -1, DigitNotFound{"GetFirstDigitPartTwo: Failed to find digit match in string: " + input}
	}
	// Swallowing errors at my own peril
	firstDigit, _ := getIntFromWordOrNumeral(matches[0])
	lastDigit, _ := getIntFromWordOrNumeral(matches[len(matches) - 1])
	return firstDigit, lastDigit, nil
}

func PrintPartTwoSolution(tokens []string) {
	sum := 0
	for _, str := range tokens {
		val_1, val_2, err := GetFirstAndLastDigitsPartTwo(str)

		if err != nil {
			panic("Bad input: \n" + err.Error())
		}
		sum += val_1 * 10 + val_2
	}
	fmt.Println("Solution for part 2:", sum)
}