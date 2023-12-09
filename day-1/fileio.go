package main

import (
	"os"
	"io"
)

func OpenFile(filepath string) *os.File {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err.Error())
	}
	return file
}

func ReadFile(file *os.File) string {
	out := make([]byte, 0)
	data := make([]byte, 100) 
	for {
		numRead, readErr := file.Read(data)
		if readErr == io.EOF{
			break
		}
		out = append(out, data[0:numRead]...) // Only append up to the number of bytes read for the last copy
	}
	return string(out)
}