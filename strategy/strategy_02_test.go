package strategy

import (
	"fmt"
	"testing"
)

func TestCashContext_GetResult(t *testing.T) {
	// 策略模式
	// 只有一个结构 CashSuper
	cc := NewCashContext("rebate")
	result := cc.GetResult(500)
	fmt.Println(result)
}
