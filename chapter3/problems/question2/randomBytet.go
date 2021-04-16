package main

import (
	crand "crypto/rand"
	"os"
	"io"
)

func main() {
	newFile, err := os.Create("new.txt") 
	if err != nil {
		panic(err)
	}

	_, err = io.CopyN(newFile, crand.Reader, 1024)

	if err != nil {
		panic(err)
	}
}