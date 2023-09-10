package facade

import "fmt"

type Wallet struct {
	balance int
}

func NewWallet() *Wallet {
	return &Wallet{
		balance: 0,
	}
}
func (w *Wallet) CreditBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
	return
}

func (w *Wallet) DebitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("balance is sufficient")
	}
	w.balance -= amount
	fmt.Println("Wallet balance is sufficient")
	return nil
}

type Ledger struct {
}

func (s *Ledger) MakeEntry(accountID, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountID, txnType, amount)
	return
}

type Notification struct {
}

func (n *Notification) sendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

func (n *Notification) sendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}
