package decorator

import "testing"

func TestShow(t *testing.T) {
	person := NewPerson("小燕")
	th := NewTShirts()
	uw := NewUnderWear()
	c := NewColor()
	h := NewHat()

	h.SetShow(person)
	uw.SetShow(h)
	c.SetShow(uw)
	th.SetShow(c)
	th.Show()
}
