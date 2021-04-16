package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	var buffer bytes.buffer
	reader := bytes.NewBufferString("Example of io.TeeReader\n")
	TeeReader := io.TeeReader(reader, &buffer)
	// データを読み捨てる
	_, _ = ioutil.ReadAll(TeeReader)

	// けどバッファに残っている
	fmt.Println(buffer.String())
}