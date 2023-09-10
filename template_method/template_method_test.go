package template_method

import (
	"fmt"
	"testing"
)

func TestOtp_GetAndSendOTP(t *testing.T) {
	smsOTP := &Sms{}
	o := Otp{
		O: smsOTP,
	}
	o.GetAndSendOTP(4)

	fmt.Println("")
	emailOTP := &Email{}
	o = Otp{
		O: emailOTP,
	}
	o.GetAndSendOTP(4)
}
