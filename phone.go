package otplessgo

// Initiate a phone call approval for given phone number
func (o *OTPLessClient) PhoneCallApproval(phone string) (*ApiResponse[ApiSuccessWithRequestID], error) {
	return getRequestId(o.restyClient, "/initiate/voice_approval", map[string]string{"phoneNumber": phone})
}
