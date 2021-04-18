package main

import (
	"fmt"
)

func main() {
	fmt.Println("start sub()")
	// 終了を受け取るためのチャネル
	done := make(chan bool) // バッファをつけるなら第３引数につける
	go func() {
		fmt.Println("sub() is finished")
		// 終了を通知
		done <- true
	}()
	// 終了を待つ
	<- done // データを読み捨てる場合は代入文が不要
	fmt.Println("all tasks are finished")
}