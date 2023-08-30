package strategy

import "math"

type CashSuper interface {
	AcceptCash(float64) float64
}

// normal
type CashNormal struct {
}

func NewCashNormal() *CashNormal {
	return &CashNormal{}
}
func (c *CashNormal) AcceptCash(money float64) float64 {
	return money
}

// discount
type CashDiscount struct {
	discount float64
}

func NewCashDiscount(discount float64) *CashDiscount {
	return &CashDiscount{discount: discount}
}
func (c *CashDiscount) AcceptCash(money float64) float64 {
	return c.discount * money
}

// rebate
type CashRebate struct {
	base   float64
	rebate float64
}

func NewCashRebate(base, rebate float64) *CashRebate {
	return &CashRebate{
		base:   base,
		rebate: rebate,
	}
}
func (c *CashRebate) AcceptCash(money float64) float64 {
	return money - math.Floor(money/c.base)*c.rebate
}

type CashFactory struct {
}

func (c *CashFactory) CreateCashAccept(t string) CashSuper {
	var cs CashSuper
	switch t {
	case "normal":
		cs = NewCashNormal()
	case "discount":
		cs = NewCashDiscount(0.8)
	case "rebate":
		cs = NewCashRebate(600, 100)
	}
	return cs
}
