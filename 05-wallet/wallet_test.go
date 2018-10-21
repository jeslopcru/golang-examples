package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, expected Euro) {
		result := wallet.Balance()
		if expected != result {
			t.Errorf("result %s and expected %s", result, expected)
		}
	}

	t.Run("Balance", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Euro(10))
		assertBalance(t, wallet, Euro(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Euro(20)}
		wallet.Withdraw(Euro(10))
		assertBalance(t, wallet, Euro(10))
	})

	t.Run("Withdraw without enough funds", func(t *testing.T) {
		wallet := Wallet{balance: Euro(20)}
		err := wallet.Withdraw(Euro(100))

		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	})
}
