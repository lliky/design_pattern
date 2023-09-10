package facade

import (
	"fmt"
)

type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
	notification *Notification
	ledger       *Ledger
}

func NewWalletFacade(accountID string, code int) *WalletFacade {
	fmt.Println("starting create account")
	wf := &WalletFacade{
		account:      NewAccount(accountID),
		wallet:       NewWallet(),
		securityCode: NewSecurityCode(code),
		notification: &Notification{},
		ledger:       &Ledger{},
	}
	fmt.Println("account created")
	return wf
}

func (wf *WalletFacade) addMoneyToWallet(accountID string, code, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := wf.account.CheckAccount(accountID)
	if err != nil {
		return err
	}
	err = wf.securityCode.CheckCode(code)
	if err != nil {
		return err
	}
	wf.wallet.CreditBalance(amount)
	wf.notification.sendWalletCreditNotification()
	wf.ledger.MakeEntry(accountID, "credit", amount)
	return nil
}

func (wf *WalletFacade) deductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := wf.account.CheckAccount(accountID)
	if err != nil {
		return err
	}

	err = wf.securityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	err = wf.wallet.DebitBalance(amount)
	if err != nil {
		return err
	}
	wf.notification.sendWalletDebitNotification()
	wf.ledger.MakeEntry(accountID, "debit", amount)
	return nil
}
