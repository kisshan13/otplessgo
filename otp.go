package otplessgo

import (
	"fmt"
)

// Struct represents the payload for sending OTP to a phone number.
type PhoneOTPPayload struct {
	Phone     string    `json:"phone"`
	Expiry    int       `json:"expiry"`
	OtpLength OtpLength `json:"otpLength"`
	Metadata  Metadata  `json:"metadata,omitempty"`
	Channels  []Channel `json:"channels"`
}

// Struct represents the payload for sending OTP to an email address.
type EmailOTPPayload struct {
	Email     string    `json:"email"`
	Expiry    int       `json:"expiry"`
	OtpLength OtpLength `json:"otpLength"`
	Metadata  Metadata  `json:"metadata,omitempty"`
}

// Struct represents the payload for verifying the otp.
type VerifyOTPPayload struct {
	RequestId string `json:"requestId"`
	Otp       string `json:"otp"`
}

type emailPayloadForApi struct {
	Email     string    `json:"email"`
	Expiry    int       `json:"expiry"`
	OtpLength OtpLength `json:"otpLength"`
	Metadata  Metadata  `json:"metadata,omitempty"`
	Channel   []Channel `json:"channel"`
}

// Sends an OTP to a Phone Number
func (o *OTPLessClient) SendPhoneOTP(payload *PhoneOTPPayload) (*ApiResponse[ApiSuccessWithRequestID], error) {
	return getRequestId(o.restyClient, "/initiate/otp", payload)
}

// Sends an OTP to an email address
func (o *OTPLessClient) SendEmailOTP(payload *EmailOTPPayload) (*ApiResponse[ApiSuccessWithRequestID], error) {

	apiPayload := &emailPayloadForApi{
		Email:     payload.Email,
		Expiry:    payload.Expiry,
		OtpLength: payload.OtpLength,
		Metadata:  payload.Metadata,
		Channel:   []Channel{ChannelEmail},
	}

	return getRequestId(o.restyClient, "/initiate/otp", apiPayload)
}

// Verify the OTP for a given RequestID
func (o *OTPLessClient) VerifyOTP(payload *VerifyOTPPayload) (*ApiResponse[VerifySuccessResponse], error) {
	var failureResponse ApiErrorResponse
	var successResponse VerifySuccessResponse

	resp, err := o.restyClient.R().SetBody(payload).SetError(failureResponse).SetResult(successResponse).Post("/verify/otp")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == 400 {
		return &ApiResponse[VerifySuccessResponse]{
			Response:      nil,
			ErrorResponse: &failureResponse,
			RawResponse:   resp.RawResponse,
		}, fmt.Errorf(failureResponse.Message)
	}

	return &ApiResponse[VerifySuccessResponse]{
		Response:      &successResponse,
		ErrorResponse: nil,
		RawResponse:   resp.RawResponse,
	}, err
}
