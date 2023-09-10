package template_method

import "fmt"

type Email struct {
	Otp
}

func (e *Email) GenRandomOPT(i int) string {
	randomOTP := "1234"
	fmt.Printf("EMAIL: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (e *Email) SaveOTPCache(s string) {
	fmt.Printf("EMAIL: saving otp: %s to cache\n", s)
}

func (e *Email) GetMessage(s string) string {
	return "EMAIL OTP for login is " + s
}

func (e *Email) SendNotification(s string) error {
	fmt.Printf("EMAIL: sending email: %s\n", s)
	return nil
}
