package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filestr, _ := ioutil.ReadFile("input")
	fmt.Println(string(filestr))
}
