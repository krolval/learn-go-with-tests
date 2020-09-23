package pointers07

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

// Wallet struct
type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet Deposit method
func (w *Wallet) Deposit(quantity Bitcoin) {
	w.balance += quantity
	fmt.Printf("address of balance in test is %v \n", &w.balance)
}

// Wallet Balance method
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Wallet Withdraw method
func (w *Wallet) Withdraw(quantity Bitcoin) error {
	if quantity > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= quantity
	return nil
}
