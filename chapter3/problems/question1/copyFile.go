package main

import (
	"os"
	"io"
)


func main() {
	file, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 引数が1以外の場合、new.txtに埋め込む挙動
	if len(os.Args) != 2 {
		newFile, err := os.Create("new.txt")
		if err != nil {
			panic(err)
		}
		io.Copy(newFile, file)
	} else {
		fileName := os.Args[1]
		newFile, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		io.Copy(newFile, file)
	}
}