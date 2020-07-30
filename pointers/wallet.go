package bitcoin

import (
	"errors"
	"fmt"
)

var errorInsufficientBalance = errors.New("Expected error")

// Stringer todo
type Stringer interface {
	String() string
}

// Bitcoin todo
type Bitcoin int

// Wallet todo
type Wallet struct {
	balance Bitcoin
}

// Deposit todo
func (w *Wallet) Deposit(amount Bitcoin) error {
	if amount < 0 {
		return errorInsufficientBalance
	}
	w.balance += amount
	return nil
}

// Balance todo
func (w *Wallet) Balance() (balance Bitcoin) {
	return w.balance
}

// Withdraw todo
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errorInsufficientBalance
	}
	if amount < 0 {
		return errorInsufficientBalance
	}
	w.balance -= amount
	return nil
}

// ToString todo
func (b Bitcoin) String() string {
	return fmt.Sprintf("%dBTC", b)
}
