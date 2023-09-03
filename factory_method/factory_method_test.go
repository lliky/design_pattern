package factory_method

import (
	"fmt"
	"testing"
)

func Test_GetResult(t *testing.T) {
	operFactory := &AddFactory{}
	oper := operFactory.CreateOperation()
	oper.(*OperatorAdd).NumberA = 1
	oper.(*OperatorAdd).NumberB = 2
	result, _ := oper.GetResult()
	fmt.Println("result = ", result)
}
