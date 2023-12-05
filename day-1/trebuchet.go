package main

import (
	"fmt"
)

func main() {
	file := OpenFile("input")
	file_bytes := ReadFile(file)
	fmt.Println(file_bytes)
}