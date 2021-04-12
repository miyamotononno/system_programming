package main

import (
	"os"
	"io"
)

func main() {
	file, err := os.Open("file.go")
	if err != nil {
		panic(err)
	}
	defer file.Close() // 現在のスコープが終了したら、その後ろに書かれている行保処理を実行する
	io.Copy(os.Stdout, file)
}