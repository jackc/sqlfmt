package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello, world")
	io.Copy(os.Stdout, os.Stdin)
}
