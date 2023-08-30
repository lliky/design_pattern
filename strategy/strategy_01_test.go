package strategy

import (
	"fmt"
	"testing"
)

func TestCashFactory_CreateCashAccept(t *testing.T) {
	// 简单工程模式
	//两个结构体 CashFactory  CashSuper
	cf := &CashFactory{}
	cs := cf.CreateCashAccept("discount")
	result := cs.AcceptCash(600)
	fmt.Println(result)
}
