package proxy

import "testing"

func TestProxy(t *testing.T) {
	mm := NewSchoolGirl("美女")
	proxy := NewProxy(mm)
	proxy.GiveFlowers()
	proxy.GiveDolls()
	proxy.GiveChocolate()
}
