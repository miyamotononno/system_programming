package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	for {
		buffer := make([]byte, 5) // 5byteのバッファを作る
		size, err := os.Stdin.Read(buffer)
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		fmt.Printf("size=%d input='%s'\n", size, string(buffer))
	}
}