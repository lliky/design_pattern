package facade

import "fmt"
func NewAPI() API {
	return &apiIml{
		a : NewAModuleAPI(),
		b : NewBModuleAPI(),
	}
}
//API is facade interface of facade package
type API interface {
	Test() string
}
//facade implement
type apiIml struct {
	a AModuleAPI
	b BModuleAPI
}

func (a *apiIml) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet,bRet)
}

func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}
//
type AModuleAPI interface {
	TestA() string
}
type aModuleImpl struct {

}
func (*aModuleImpl)TestA() string{
	return "A module running"
}
func NewBModuleAPI() BModuleAPI {
	return &bModelImpl{}
}
type BModuleAPI interface {
	TestB() string
}
type bModelImpl struct {
}

func (*bModelImpl)TestB()string  {
	return "B module running"
}
