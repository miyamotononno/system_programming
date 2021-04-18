package main

import (
  "fmt"
  "time"
)

// 新しく作られるgoroutineが呼ぶ関数
func sub() {
  fmt.Println("sub() is running")
  time.Sleep(time.Second)
  fmt.Println("sub() is finished")
}

func main() {
  fmt.Println("start sub()")
  // goroutineを作って関数を実行
  go sub()
  time.Sleep(2*time.Second)
}