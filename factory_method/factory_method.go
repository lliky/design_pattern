package factory_method

import "errors"

type Operation interface {
	GetResult() (float64, error)
}

type OperatorNum struct {
	NumberA float64
	NumberB float64
}
type OperatorAdd struct {
	OperatorNum
}

func NewOperatorAdd() *OperatorAdd {
	return &OperatorAdd{}
}

func (add *OperatorAdd) GetResult() (float64, error) {
	return add.NumberA + add.NumberB, nil
}

type OperatorSub struct {
	OperatorNum
}

func NewOperatorSub() *OperatorSub {
	return &OperatorSub{}
}
func (sub *OperatorSub) GetResult() (float64, error) {
	return sub.NumberA - sub.NumberB, nil
}

type OperatorMul struct {
	OperatorNum
}

func NewOperatorMul() *OperatorMul {
	return &OperatorMul{}
}

func (mul *OperatorMul) GetResult() (float64, error) {
	return mul.NumberA * mul.NumberB, nil
}

type OperatorDiv struct {
	OperatorNum
}

func NewOperatorDiv() *OperatorDiv {
	return &OperatorDiv{}
}

func (div *OperatorDiv) GetResult() (float64, error) {
	if div.NumberB == 0 {
		return 0, errors.New("divisor cannot be equal to 0")
	}
	return div.NumberA / div.NumberB, nil
}

type IFactory interface {
	CreateOperation() Operation
}

type AddFactory struct {
}

func (a AddFactory) CreateOperation() Operation {
	return NewOperatorAdd()
}

type SubFactory struct {
}

func (a SubFactory) CreateOperation() Operation {
	return NewOperatorSub()
}

type MulFactory struct {
}

func (a MulFactory) CreateOperation() Operation {
	return NewOperatorMul()
}

type DivFactory struct {
}

func (a DivFactory) CreateOperation() Operation {
	return NewOperatorDiv()
}
