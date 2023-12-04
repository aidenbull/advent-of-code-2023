package main

import (
	"fmt"
	"./fileio"
)

func main() {
	file := fileio.OpenFile("input")
	file_bytes := fileio.ReadFile(file)
	fmt.Println(file_bytes)
}