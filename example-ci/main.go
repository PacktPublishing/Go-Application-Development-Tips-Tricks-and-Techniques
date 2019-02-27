package main

import (
	"fmt"
	"io"
	"os"
)

// PrintHello prints hello
func PrintHello(out io.Writer) error {
	_, err := fmt.Fprintln(out, "Hello World")
	return err
}

func main() {
	err := PrintHello(os.Stdout)
	if err != nil {
		panic(err)
	}
}
