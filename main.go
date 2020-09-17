package main

import (
	"fmt"
	"os"
)

const (
	screenWidth  = 800
	screenHeight = 800
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
}
