package main

import "os"

func main() {
	f, _ := os.Create("somefile")
	f.WriteString("Hello World")
}
