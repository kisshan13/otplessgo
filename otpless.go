package otplessgo

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Version string

const baseUrl = "https://auth.otpless.app/auth"

const (
	VersionOne = Version("v1")
)

// Represents an OTPLESS Client
type OTPLessClient struct {
	clientId     string
	clientSecret string
	baseUrl      string
	version      Version
	restyClient  *resty.Client
}

// Creates a new OTPLESS Client
func NewOTPLessClient(clientId string, clientSecret string, version Version) *OTPLessClient {

	restyClient := resty.New()

	restyClient.SetHeaders(map[string]string{
		"clientId":     clientId,
		"clientSecret": clientSecret,
		"Content-Type": "application/json",
	})

	restyClient.SetBaseURL(baseUrl + "/" + string(version))

	return &OTPLessClient{clientId: clientId, clientSecret: clientSecret, baseUrl: baseUrl, version: version, restyClient: restyClient}
}

// Checks the authentication status for the given RequestID
func (o *OTPLessClient) RequestStatus(requestId string) (*ApiResponse[RequestStatusResponse], error) {
	var failureResponse ApiErrorResponse
	var successResponse RequestStatusResponse

	resp, err := o.restyClient.R().SetBody(map[string]string{
		"requestId": requestId,
	}).SetError(&failureResponse).SetResult(&successResponse).Post("/status")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() > 300 {
		return &ApiResponse[RequestStatusResponse]{
			Response:      nil,
			ErrorResponse: &failureResponse,
			RawResponse:   resp.RawResponse,
		}, fmt.Errorf(failureResponse.Message)
	}

	return &ApiResponse[RequestStatusResponse]{
		Response:      &successResponse,
		ErrorResponse: nil,
		RawResponse:   resp.RawResponse,
	}, nil
}
