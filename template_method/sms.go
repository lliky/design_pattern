package template_method

import "fmt"

type Sms struct {
	Otp
}

func (S *Sms) GenRandomOPT(i int) string {
	randomOTP := "1234"
	fmt.Printf("SMS: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (S *Sms) SaveOTPCache(s string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", s)
}

func (S *Sms) GetMessage(s string) string {
	return "SMS OTP for login is " + s
}

func (S *Sms) SendNotification(s string) error {
	fmt.Printf("SMS: sending sms: %s\n", s)
	return nil
}
