package decorator

import "fmt"

type IShow interface {
	Show()
}

type Person struct {
	name string
}

func NewPerson(name string) *Person {
	return &Person{name: name}
}
func (p *Person) Show() {
	fmt.Println("装扮的", p.name)
}

type TShirts struct {
	show IShow
}

func NewTShirts() *TShirts {
	return &TShirts{}
}
func (t *TShirts) SetShow(show IShow) {
	t.show = show
}
func (t *TShirts) Show() {
	fmt.Printf("大T恤 ")
	t.show.Show()
}

type UnderWear struct {
	show IShow
}

func NewUnderWear() *UnderWear {
	return &UnderWear{}
}
func (t *UnderWear) SetShow(show IShow) {
	t.show = show
}
func (t *UnderWear) Show() {
	fmt.Printf("内裤 ")
	t.show.Show()
}

type Color struct {
	show IShow
}

func NewColor() *Color {
	return &Color{}
}
func (t *Color) SetShow(show IShow) {
	t.show = show
}
func (t *Color) Show() {
	fmt.Printf("红色 ")
	t.show.Show()
}

type Hat struct {
	show IShow
}

func NewHat() *Hat {
	return &Hat{}
}
func (t *Hat) SetShow(show IShow) {
	t.show = show
}
func (t *Hat) Show() {
	fmt.Printf("帽子 ")
	t.show.Show()
}
