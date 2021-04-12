package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Create("nultiwriter.txt")
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, os.Stdout) // io.MultiWriter()は複数のio.Writerを受け取り、それら全てに対し、書き込まれた内容を同時に書き込むデコレータです
	io.WriteString(writer, "io.MultiWriter example\n")
}