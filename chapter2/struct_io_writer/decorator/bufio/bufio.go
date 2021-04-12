package main

import (
	"bufio"
	"os"
)

func main() {
	buffer := bufio.NewWriter(os.Stdout) // ある程度の分量ごとにまとめて書き出すbuffo.Writerという構造体もあります
	buffer.WriteString("buffo.Writer ")
	buffer.Flush()
	buffer.WriteString("example\n")
	buffer.Flush()
}
// Flush()を自動で呼び出すためには、バッファサイズ指定のbuffer.NewWriterSize(os.Stdout, bufferSize)関数で、bufio.Writerを作成します。　　　　　　　　　　