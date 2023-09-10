package facade

import "fmt"

type Account struct {
	Name string
}

func NewAccount(name string) *Account {
	return &Account{
		Name: name,
	}
}

func (a *Account) CheckAccount(name string) error {
	if a.Name != name {
		return fmt.Errorf("account name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}
