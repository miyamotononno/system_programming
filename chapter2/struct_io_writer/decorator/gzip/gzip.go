package main

import (
	"compress/gzip"
	"os"
	"io"
)

// かきこまれたgzipを圧縮して、あらかじめ渡されていたos.Fileに中継する
func main() {
	file, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}
	writer := gzip.NewWriter(file)
	writer.Header.Name = "test.txt"
	io.WriteString(writer, "gzip.Writer example\n")
	writer.Close()
}