package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
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

		assertBalance(t, wallet, Euro(20))
		assertError(t, err, "not enough Euro")
	})
}

func assertBalance(t *testing.T, wallet Wallet, expected Euro) {
	result := wallet.Balance()
	if expected != result {
		t.Errorf("result %s and expected %s", result, expected)
	}
}

func assertError(t *testing.T, error error, expected string) {
	if error == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if error.Error() != expected {
		t.Errorf("result %s and expected %s", error, expected)
	}
}
