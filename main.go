package main

import (
	"fmt"
	"os"
)

var (
	Version string = "0.0.0"
)

func main() {
	fmt.Fprintf(os.Stdout, "%s\n", Version)
}
