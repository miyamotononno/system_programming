package main

import (
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "http.ResponseWriter sample")
}

// 実行したらlocalhost:8080にアクセス
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}