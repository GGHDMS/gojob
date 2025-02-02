package main

import (
	"fmt"
	"gojob/accounts"
	"log"
)

func main4() {
	account := accounts.NewAccount("test")
	fmt.Println(account)
	account.Deposit(10000)
	fmt.Println(account)
	err := account.Withdraw(1000)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account)

	account.ChangeOwner("seokmun")
	fmt.Println(account.Owner())
}
