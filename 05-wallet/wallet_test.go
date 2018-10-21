package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(10)

	result := wallet.Balance()
	fmt.Printf("In tests, address of balance is: %v \n", &wallet.balance)

	expected := 10

	if expected != result {
		t.Errorf("result %d and expected %d", result, expected)
	}
}
