package otplessgo

import (
	"fmt"
	"os"
	"testing"
)

func TestOAuthIntegration(t *testing.T) {

	fmt.Println(os.Getenv("OTPLESS_CLIENT_ID"))

	client := NewOTPLessClient(os.Getenv("OTPLESS_CLIENT_ID"), os.Getenv("OTPLESS_CLIENT_SECRET"), VersionOne)

	oauthRes, err := client.InitiateOAuth(&OAuthPayload{
		Channel:     OAuthGoogle,
		RedirectURI: "http://localhost:3000",
	})

	if err != nil {

		if oauthRes.ErrorResponse != nil {
			t.Errorf("OAuth Initiation Failed : %v \nStatus Code %v \nReason: %v", err.Error(), oauthRes.RawResponse.StatusCode, oauthRes.ErrorResponse.Description)
			return
		}

		t.Errorf("OAuth Initiation Failed : %v", err.Error())
		return
	}

	if oauthRes.RawResponse.StatusCode == 200 {
		t.Logf("OAuth Initiation Success %v", oauthRes.Response.Link)
	}
}
