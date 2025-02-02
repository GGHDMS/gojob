package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("no money")

func NewAccount(owner string) *Account {
	return &Account{owner: owner, balance: 0}
}

func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a *Account) GetBalance() int {
	return a.balance
}

func (a *Account) Withdraw(amount int) error {
	if amount > a.balance {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a *Account) Owner() string {
	return a.owner
}

func (a *Account) String() string {
	return fmt.Sprintf("Account{owner: %s, balance: %d}", a.owner, a.balance)
}
