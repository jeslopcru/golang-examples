package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Balance", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Euro(10))
		result := wallet.Balance()
		expected := Euro(10)
		if expected != result {
			t.Errorf("result %s and expected %s", result, expected)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Euro(20)}
		wallet.Withdraw(Euro(10))
		result := wallet.Balance()
		expected := Euro(10)
		if expected != result {
			t.Errorf("result %s and expected %s", result, expected)
		}
	})
}
