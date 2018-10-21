package wallet

import "fmt"

type Wallet struct {
	balance int
}

func (wallet Wallet) Deposit(amount int) {
	fmt.Printf("In Deposit, address of balance is: %v \n", &wallet.balance)
	wallet.balance += amount
}
func (wallet Wallet) Balance() int {
	return wallet.balance
}
