package main

import (
	"os"
	"fmt"
	"time"
)

func main() {
	fmt.Fprintf(os.Stdout, "write with os.Stdout at %v", time.Now())
}