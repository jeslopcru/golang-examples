package wallet

import (
	"errors"
	"fmt"
)

type Euro int

func (eur Euro) String() string {
	return fmt.Sprintf("%d EUR", eur)
}

type Wallet struct {
	balance Euro
}

func (wallet *Wallet) Deposit(amount Euro) {
	wallet.balance += amount
}
func (wallet *Wallet) Balance() Euro {
	return wallet.balance
}
func (wallet *Wallet) Withdraw(amount Euro) error {
	if amount > wallet.balance {
		return errors.New("not enough Euro")
	}
	wallet.balance -= amount
	return nil
}
