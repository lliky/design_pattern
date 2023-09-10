package facade

import "fmt"

type SecurityCode struct {
	code int
}

func NewSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		code: code,
	}
}

func (s *SecurityCode) CheckCode(code int) error {
	if s.code != code {
		return fmt.Errorf("security Code is incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}
