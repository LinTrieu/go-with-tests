// Pointers & errors
// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors
package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)

	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(5))
		want := Bitcoin(15)
		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		starting := Bitcoin(5)
		wallet := Wallet{starting}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, starting)
	})
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error, but did not want one")
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("want error but did not get one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertBalance(t *testing.T, w Wallet, want Bitcoin) {
	t.Helper()
	got := w.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
