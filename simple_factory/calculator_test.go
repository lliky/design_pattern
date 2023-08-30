package simplefactory

import (
	"fmt"
	"testing"
)

func TestCreateOperate(t *testing.T) {
	of := OperateFactory{}
	operate := of.CreateOperate("+")
	operate.(*OperatorAdd).NumberA = 1
	operate.(*OperatorAdd).NumberB = 1
	result, _ := operate.GetResult()
	if result != 2 {
		t.Fatal("error")
	}
	fmt.Println(result)
}
