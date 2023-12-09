package main

import(
	"unicode"
	"fmt"
)

func GetFirstDigit(input string) (int, error) {
	for _, char := range input {
		if unicode.IsDigit(char) {
			return int(char) - '0', nil
		}
	}
	return -1, DigitNotFound{"GetFirstDigit: Failed to find digit in string: " + input}
}

func GetLastDigit(input string) (int, error) {
	//Referenced from https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
	runes := []rune(input)
	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return GetFirstDigit(string(runes))
}

func PrintPartOneSolution(tokens []string) {
	sum := 0
	for _, str := range tokens {
		val_1, err_1 := GetFirstDigit(str)
		if err_1 != nil {
			panic("Bad input: \n" + err_1.Error())
		}
		val_2, err_2 := GetLastDigit(str)
		if err_2 != nil {
			panic("Somehow errored out in reverse but not forward: \n" + err_2.Error())
		}
		sum += val_1 * 10 + val_2
	}
	fmt.Println("Solution for part 1:", sum)
}