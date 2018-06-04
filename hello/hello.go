package main

import "fmt"

func main() {
	fmt.Println(HelloWorld("Jesús"))
}

const HELLO = "Hello, "

func HelloWorld(name string) string {
	if name == "" {
		name = "World"
	}
	return HELLO + name
}
