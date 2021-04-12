package main

import (
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("os.File example\n")) // Write()が受け取るのは文字列ではなくてバイト列なので変換を行う
	file.Close()
}