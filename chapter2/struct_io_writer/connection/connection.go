package main

import (
	"io"
	"os"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	// connはnet.Connという通信のコネクションを表すインターフェースが返ってくる。詳しくは3.4.3章で解説
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, "GET / HTTP/1.0/\r\nHost: ascii.jp\r\n\r\n")
	io.Copy(os.Stdout, conn)
} 