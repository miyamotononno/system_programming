package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	header := bytes.NewBufferString("---------HEADER----------\n")
	content := bytes.NewBufferString("Example of io.MultiReader\n")
	footer := bytes.NewBufferString("---------FOOTER----------\n");
	
	reader := io.MultiReader(header, content, footer) // 引数で渡されたio.Readerの全ての入力が繋がっているかのような動作をする
	// 全てのreaderをつなげた出力
	io.Copy(os.Stdout, reader);
}