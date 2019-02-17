// +build !darwin,!windows,!linux

package main

import "fmt"

func SayHello() {
	fmt.Println("Hello, other OS!")
}
