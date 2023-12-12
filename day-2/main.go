package main

import (
	"fmt"
	"strings"

	"github.com/aidenbull/fileio"
)

func main() {
	file := fileio.OpenFile("input")
	fileString := strings.TrimSpace(strings.ToLower(fileio.ReadFile(file)))
	tokens := strings.Split(fileString, "\n")
	for _, token := range tokens {
		fmt.Println(ProcessGame(token))
	}
}
