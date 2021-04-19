package main

import (
	"net"
)

func client() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	// connを使った読み書き
}

func server() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	// 一度で終了しないためにAccept()を何度も繰り返し呼ぶ
	for {
		conn, net := ln.Accept()
		if err != nil {
			// handle error
			panic(err)
		}
		// 1リクエスト処理中に他のリクエストのAccceptが行えるように
		// goroutineを使って非同期にレスポンスを処理する
		go func() {
			// connを使った読みかき
		}()
	}
}