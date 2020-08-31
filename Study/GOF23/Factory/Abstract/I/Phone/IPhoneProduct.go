package Phone


type IPhoneProduct interface{
	Open()
	Close()
	Call()
	SendSMS()
}