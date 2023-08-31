package proxy

import "fmt"

type IGiveGift interface {
	GiveDolls()
	GiveFlowers()
	GiveChocolate()
}
type SchoolGirl struct {
	Name string
}

func NewSchoolGirl(name string) SchoolGirl {
	return SchoolGirl{Name: name}
}

type Pursuit struct {
	MM SchoolGirl
}

func NewPursuit(mm SchoolGirl) Pursuit {
	return Pursuit{MM: mm}
}

func (p *Pursuit) GiveDolls() {
	fmt.Printf("%v 送你洋娃娃\n", p.MM)
}

func (p *Pursuit) GiveFlowers() {
	fmt.Printf("%v 送你鲜花\n", p.MM)
}

func (p *Pursuit) GiveChocolate() {
	fmt.Printf("%v 送你巧克力\n", p.MM)
}

type Proxy struct {
	GG Pursuit
}

func NewProxy(mm SchoolGirl) Proxy {
	return Proxy{GG: NewPursuit(mm)}
}
func (p *Proxy) GiveDolls() {
	p.GG.GiveDolls()
}
func (p *Proxy) GiveFlowers() {
	p.GG.GiveFlowers()
}

func (p *Proxy) GiveChocolate() {
	p.GG.GiveChocolate()
}
