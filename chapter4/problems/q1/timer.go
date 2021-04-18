package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start timer")
	timer := time.After(3 * time.Second) 
	for {
		select {
			case <- timer:
				fmt.Println("the end!")
				return
			default:
				fmt.Println("still waiting")
		}
	}
}