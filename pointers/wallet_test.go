package pointers

import (
	"testing"
)

type Transaction int

const (
	Deposit Transaction = iota
	Withdraw
)

func Test_wallet(t *testing.T) {
	tests := []struct {
		name   string
		wallet Wallet
		amount rupees
		action Transaction
		want   rupees
	}{
		{
			"Deposit",
			Wallet{amount: 0},
			rupees(10),
			Deposit,
			rupees(10),
		}, {
			"Withdraw",
			Wallet{amount: rupees(50)},
			rupees(20),
			Withdraw,
			rupees(30),
		}, {
			"Overdraft",
			Wallet{amount: rupees(20)},
			rupees(30),
			Withdraw,
			rupees(20),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			switch tt.action {
			case Deposit:
				tt.wallet.Deposit(tt.amount)
			case Withdraw:
				err = tt.wallet.Withdraw(tt.amount)
			}
			if err != nil {
				assertError(t, err)
			}
			if got := tt.wallet.Balance(); got != tt.want {
				t.Errorf("Wallet starting contents: %#v\tgot: %v\twant: %v", tt.wallet, got, tt.want)
			}
		})
	}
}

func assertError(t *testing.T, got error) {
	t.Helper()
	if got == nil {
		t.Fatalf("Wanted error %v but got none", ErrInsufficientFunds)
	}
	if got != ErrInsufficientFunds {
		t.Errorf("Expected error %v\tgot: %v", ErrInsufficientFunds, got)
	}
}
