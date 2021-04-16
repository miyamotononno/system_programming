package main

import (
	"archive/zip"
	"strings"
	"io"
	"os"
)

func main() {
	file, err := os.Create("new.zip")
	if err != nil {
		panic(err)
	}
	zipWriter := zip.NewWriter(file) // zipファイル書き込み用の構造体
	defer zipWriter.Close()

	// この構造体自体はio.Writerではないが、Create()メソッドを呼ぶと、
	// 個別のファイルに書き込むためのio.Writerが返ってくる
	text, err := zipWriter.Create("newFile.txt")
	if err != nil {
		panic(err)
	}

	reader := strings.NewReader("write to zip file")
	io.Copy(text, reader)
}