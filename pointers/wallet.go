package pointers

import "fmt"

type rupees int

var ErrInsufficientFunds = fmt.Errorf("tried to withdraw an amount larger than the balance")

type Stringer interface {
	String() string
}

type Wallet struct {
	amount rupees
}

func (r rupees) String() string {
	return fmt.Sprintf("%d RUPEE", r)
}

// Pointers on type to dereference the object give so we do not operate with local copy
func (w *Wallet) Deposit(i rupees) {
	w.amount += i
}

// Struct pointers are automatically dereferenced
func (w *Wallet) Balance() rupees {
	return w.amount
}

func (w *Wallet) Withdraw(i rupees) error {
	if i > w.amount {
		return ErrInsufficientFunds
	}
	w.amount -= i
	return nil
}
