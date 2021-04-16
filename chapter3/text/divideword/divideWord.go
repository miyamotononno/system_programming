package main

import (
	"bufio"
	"fmt"
	"strings"
)

var source = `１行目
２行目
３行目
`

func main() {
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan(){ // scanner.Split(bufio.ScanWords)とすると単語区切りになる
		fmt.Printf("%#v\n", scanner.Text())
	} 
}