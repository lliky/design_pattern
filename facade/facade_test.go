package facade

import (
	"fmt"
	"log"
	"testing"
)

func TestNewWalletFacade(t *testing.T) {
	facade := NewWalletFacade("abc", 123)
	err := facade.addMoneyToWallet("abc", 123, 10)
	if err != nil {
		log.Fatalf("Errors: %s\n", err.Error())
	}
	fmt.Println()
	err = facade.deductMoneyFromWallet("abc", 123, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
