package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example"))
	fmt.Println(buffer.String())
}