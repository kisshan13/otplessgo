package otplessgo

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func getRequestId(client *resty.Client, path string, payload interface{}) (*ApiResponse[ApiSuccessWithRequestID], error) {
	var failureResp ApiErrorResponse
	var successResp ApiSuccessWithRequestID

	resp, err := client.R().SetBody(payload).SetError(&failureResp).SetResult(&successResp).Post(path)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == 400 {
		return &ApiResponse[ApiSuccessWithRequestID]{
			Response:      nil,
			ErrorResponse: &failureResp,
			RawResponse:   resp.RawResponse,
		}, fmt.Errorf(failureResp.Message)
	}

	return &ApiResponse[ApiSuccessWithRequestID]{
		Response:      &successResp,
		ErrorResponse: &failureResp,
		RawResponse:   resp.RawResponse,
	}, nil
}
