package main

import (
	"fmt"
	"github.com/jeslopcru/golang-examples/01-hello"
	"github.com/jeslopcru/golang-examples/02-calculator"
	"github.com/jeslopcru/golang-examples/03-geometry"
	"github.com/jeslopcru/golang-examples/04-more-geometry"
)

func main() {
	fmt.Println("01 - Hello")
	fmt.Println(hello.HelloWorld("Jesús"))
	fmt.Println("")

	fmt.Println("02 - Calculator")
	sum := calculator.Add(5, 2)
	fmt.Println(fmt.Sprintf("Add 5 plus 2 is %v", sum))
	numberList := []int{1, 2, 3}
	sumMultiple := calculator.AddMultiple(numberList)
	fmt.Println(fmt.Sprintf("AddMultiple %v is %v", numberList, sumMultiple))
	fmt.Println("")

	fmt.Println("03 - Geometry")
	rectangle := geometry.Rectangle{10.0, 5.0}
	fmt.Println(fmt.Sprintf("Rectangle%v with Area: %f", rectangle, rectangle.Area()))
	fmt.Println("")

	fmt.Println("04 - More Geometry")
	circle := more_geometry.Circle{5}
	fmt.Println(fmt.Sprintf("circle%v with Area: %f", circle, circle.Area()))
	fmt.Println("")

}
