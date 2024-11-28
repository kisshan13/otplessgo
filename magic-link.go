package otplessgo

import "fmt"

// Struct represents the payload for sending magic link to a phone number.
type PhoneMagicLinkPayload struct {
	RedirectURI string    `json:"redirectURI"`
	Expiry      int       `json:"expiry"`
	PhoneNumber string    `json:"phoneNumber"`
	Channels    []Channel `json:"channels"`
	Metadata    Metadata  `json:"metadata,omitempty"`
}

// Struct represents the payload for sending magic link to an email.
type EmailMagicLinkPayload struct {
	RedirectURI string   `json:"redirectURI"`
	Expiry      int      `json:"expiry"`
	Email       string   `json:"email"`
	Metadata    Metadata `json:"metadata,omitempty"`
}

type emailMagicLinkPayloadForApi struct {
	RedirectURI string    `json:"redirectURI"`
	Expiry      int       `json:"expiry"`
	Email       string    `json:"email"`
	Metadata    Metadata  `json:"metadata,omitempty"`
	Channel     []Channel `json:"channels"`
}

// Sends the Magic link to a phone number
func (o *OTPLessClient) SendPhoneMagicLink(payload *PhoneMagicLinkPayload) (*ApiResponse[ApiSuccessWithRequestID], error) {
	return getRequestId(o.restyClient, "/initiate/magiclink", payload)
}

// Sends the Magic link to an email number
func (o *OTPLessClient) SendEmailMagicLink(payload *EmailMagicLinkPayload) (*ApiResponse[ApiSuccessWithRequestID], error) {
	apiPayload := emailMagicLinkPayloadForApi{
		RedirectURI: payload.RedirectURI,
		Expiry:      payload.Expiry,
		Email:       payload.Email,
		Metadata:    payload.Metadata,
		Channel:     []Channel{ChannelEmail},
	}

	return getRequestId(o.restyClient, "/initiate/magiclink", apiPayload)
}

// Verify the code and returns the user details object.
func (o *OTPLessClient) VerifyCode(code string) (*ApiResponse[UserDetails], error) {
	apiPayload := CodeVerifyPayload{
		Code: code,
	}

	var failureResp ApiErrorResponse
	var successResp UserDetails

	resp, err := o.restyClient.R().SetBody(apiPayload).SetError(&failureResp).SetResult(&successResp).Post("/verify/code")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() > 300 {
		return nil, fmt.Errorf(failureResp.Message)
	}

	return &ApiResponse[UserDetails]{
		Response:      &successResp,
		ErrorResponse: &failureResp,
		RawResponse:   resp.RawResponse,
	}, nil
}
