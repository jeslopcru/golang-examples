package main

import (
	"fmt"
	"github.com/jeslopcru/golang-examples/01-hello"
	"github.com/jeslopcru/golang-examples/02-calculator"
)

func main() {
	fmt.Println("01 - Hello World")
	fmt.Println(hello.HelloWorld("Jesús"))

	fmt.Println("02 - Calculator")
	fmt.Println(calculator.Add(5, 2))
}
