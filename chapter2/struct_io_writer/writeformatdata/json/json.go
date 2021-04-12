package main

import (
	"os"
	"encoding/json"
)

func main() {
	encoder := json.NewEncoder(os.Stdout) // io.Writerの例と組み合わせれば。サーバーにJSONを送ったり、ブラウザにJSONを返すことも簡単にできる
	encoder.SetIndent("", "     ")
	encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello": "world",
	})
}