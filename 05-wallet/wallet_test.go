package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(10)

	result := wallet.Balance()

	expected := 10

	if expected != result {
		t.Errorf("result %d and expected %d", result, expected)
	}
}
