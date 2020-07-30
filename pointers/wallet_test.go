package bitcoin

import "testing"

func TestWallet(t *testing.T) {

	assertError := func(t *testing.T, got, want error) {
		t.Helper()
		if got == nil {
			t.Error("Expected error, but got none")
		}
		if want != got {
			t.Errorf("Thrown: %q, want: %q", got, want)
		}
	}

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		_ = wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Throw error when deposit negative amount", func(t *testing.T) {
		wallet := Wallet{}
		err := wallet.Deposit(Bitcoin(-10))
		assertBalance(t, wallet, Bitcoin(0))
		assertError(t, err, errorInsufficientBalance)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		_ = wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw when amount greater balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(30))
		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, errorInsufficientBalance)
	})

}
