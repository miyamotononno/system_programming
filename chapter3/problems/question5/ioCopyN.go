package main

import (
	"io"
	"os"
)

func copyN(dest io.Writer, src io.Reader, length int) {
	if length <= 0 {
		return
	}
	buffer := make([]byte, length)
	// 決まったバイトだけ読む。指定したバッファのサイズ分まで読み込めない場合はerrが返る
	_, err := io.ReadFull(src,  buffer)
	if err != nil {
		panic(err)
	}
	dest.Write(buffer)
}

func main() {
	src, err := os.Open("original.txt")
	if err != nil {
		panic(err)
	}
	defer src.Close()
	dest, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	copyN(dest, src, 3)	
}
