package otplessgo

import "fmt"

// Represents an oAuth initialization payload
type OAuthPayload struct {
	Channel     OAuthChannel `json:"channel"`
	RedirectURI string       `json:"redirect_uri"`
}

// OAuth initialization success response
type OAuthResponse struct {
	RequestId string `json:"requestId"`
	Link      string `json:"link"`
}

// Sends a initialization request for the oAuth Channel.
func (o *OTPLessClient) InitiateOAuth(payload *OAuthPayload) (*ApiResponse[OAuthResponse], error) {
	var failureResp ApiErrorResponse
	var successResp OAuthResponse

	resp, err := o.restyClient.R().SetBody(payload).SetError(&failureResp).SetResult(&successResp).Post("/initiate/oauth")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() > 300 {
		return &ApiResponse[OAuthResponse]{
			Response:      nil,
			ErrorResponse: &failureResp,
			RawResponse:   resp.RawResponse,
		}, fmt.Errorf(failureResp.Message)
	}

	return &ApiResponse[OAuthResponse]{
		Response:      &successResp,
		ErrorResponse: &failureResp,
		RawResponse:   resp.RawResponse,
	}, nil
}
