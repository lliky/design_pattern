package strategy

type CashContext struct {
	cs CashSuper
}

func NewCashContext(t string) *CashContext {
	var cs CashSuper
	switch t {
	case "normal":
		cs = NewCashNormal()
	case "discount":
		cs = NewCashDiscount(0.8)
	case "rebate":
		cs = NewCashRebate(600, 100)
	}
	return &CashContext{cs: cs}
}
func (cs *CashContext) GetResult(money float64) float64 {
	return cs.cs.AcceptCash(money)
}
