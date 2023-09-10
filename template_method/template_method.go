package template_method

type IOtp interface {
	GenRandomOPT(int) string
	SaveOTPCache(string)
	GetMessage(string) string
	SendNotification(string) error
}

type Otp struct {
	O IOtp
}

func (o *Otp) GetAndSendOTP(length int) error {
	opt := o.O.GenRandomOPT(length)
	o.O.SaveOTPCache(opt)
	message := o.O.GetMessage(opt)
	return o.O.SendNotification(message)
}
