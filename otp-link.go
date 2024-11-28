package otplessgo

// Struct represents the payload for sending OTP Link to a phone number.
type PhoneOTPLinkPayload struct {
	RedirectURI string    `json:"redirectURI"`
	Expiry      int       `json:"expiry"`
	OtpLength   OtpLength `json:"otpLength"`
	PhoneNumber string    `json:"phoneNumber"`
	Channel     []Channel `json:"channel"`
}

// Struct represents the payload for sending OTP Link to an email address.
type EmailOTPLinkPayload struct {
	RedirectURI string    `json:"redirectURI"`
	Expiry      int       `json:"expiry"`
	OtpLength   OtpLength `json:"otpLength"`
	Email       string    `json:"Email"`
	Channel     []Channel `json:"channel"`
}

// Sends OTP Link to a phone number
func (o *OTPLessClient) SendPhoneOTPLink(payload *PhoneOTPLinkPayload) (*ApiResponse[ApiSuccessWithRequestID], error) {
	return getRequestId(o.restyClient, "/initiate/otplink", payload)
}

// Sends OTP Link to an Email Address
func (o *OTPLessClient) SendEmailOTPLink(payload *EmailOTPLinkPayload) (*ApiResponse[ApiSuccessWithRequestID], error) {
	return getRequestId(o.restyClient, "/initiate/otplink", payload)
}
