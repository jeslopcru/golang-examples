package main

import (
	"fmt"
	"github.com/jeslopcru/golang-examples/07-http-score-store"
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) ObtainPlayerScore(name string) int {
	return 123
}

func main() {

	fmt.Println("Server Running in http://localhost:5000")
	server := &http_score_store.PlayerServer{Store: &InMemoryPlayerStore{}}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

//func consoleMain() {
//	fmt.Println("01 - Hello")
//	fmt.Println(hello.HelloWorld("Jesús"))
//	fmt.Println("")
//	fmt.Println("02 - Calculator")
//	sum := calculator.Add(5, 2)
//	fmt.Println(fmt.Sprintf("Add 5 plus 2 is %v", sum))
//	numberList := []int{1, 2, 3}
//	sumMultiple := calculator.AddMultiple(numberList)
//	fmt.Println(fmt.Sprintf("AddMultiple %v is %v", numberList, sumMultiple))
//	fmt.Println("")
//	fmt.Println("03 - Geometry")
//	rectangle := geometry.Rectangle{10.0, 5.0}
//	fmt.Println(fmt.Sprintf("Rectangle%v with Area: %f", rectangle, rectangle.Area()))
//	fmt.Println("")
//	fmt.Println("04 - More Geometry")
//	circle := more_geometry.Circle{5}
//	fmt.Println(fmt.Sprintf("circle%v with Area: %f", circle, circle.Area()))
//	fmt.Println("")
//	fmt.Println("05 - Wallet")
//	myWallet := wallet.Wallet{}
//	myWallet.Balance()
//	fmt.Println(fmt.Sprintf("my balance is %v ", myWallet.Balance()))
//	myWallet.Deposit(wallet.Euro(10))
//	fmt.Println(fmt.Sprintf("if I deposit 10, my balance now is %v ", myWallet.Balance()))
//	myWallet.Withdraw(wallet.Euro(3))
//	fmt.Println(fmt.Sprintf("if I withdraw 3, my balance now is %v ", myWallet.Balance()))
//	fmt.Println("")
//}
