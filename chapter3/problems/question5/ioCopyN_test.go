package main

import (
	"testing"
	"os"
)

func TestCopyN(t *testing.T) {
	main()
	fileSize := 3
	info, err := os.Stat("new.txt")
	if err != nil {
		panic(err)
	}

	if int(info.Size()) != fileSize {
		t.Fatal("size was not fileSize.")
	}
}

