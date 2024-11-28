# OTPLess Go SDK

![OTPLess Icon](/otpless.svg)

## Overview

The OTPLess Go SDK is a simple, easy-to-use client for interacting with OTPLess API for authentication. This SDK allows you to integrate OTP (One-Time Password), Magic Link, OAuth, and other authentication methods into your Go applications. You can easily send OTPs to phone numbers, email addresses, or handle phone call approvals, among other features.

## Installation

To install the OTPLess Go SDK, you can use the following Go command:

```bash
go get github.com/kisshan13/otplessgo
```

## Usage

### 1. Create a New OTPLess Client

To initialize a new OTPLess client, you can use the `NewOTPLessClient` function.

```go
import (
    "github.com/kisshan13/otplessgo"
)

func main() {
    client := otplessgo.NewOTPLessClient("<client-id>", "<client-secret>", otplessgo.VersionOne)
}
```

### 2. Sending Magic Links

### Send Magic Link to Phone Number

```go
import (
    "github.com/kisshan13/otplessgo"
)

func main() {
    client := otplessgo.NewOTPLessClient("<client-id>", "<client-secret>", otplessgo.VersionOne)

    payload := &otplessgo.PhoneMagicLinkPayload{
        RedirectURI: "https://yourapp.com/redirect",
        Expiry: 600,
        PhoneNumber: "1234567890",
        Channels: []otplessgo.Channel{otplessgo.ChannelSMS},
    }

    response, err := client.SendPhoneMagicLink(payload)
}

```

<br />

### Send Magic Link to Email Address

```go
import (
    "github.com/kisshan13/otplessgo"
)

func main() {
    client := otplessgo.NewOTPLessClient("<client-id>", "<client-secret>", otplessgo.VersionOne)

    payload := &otplessgo.EmailMagicLinkPayload{
        RedirectURI: "https://yourapp.com/redirect",
        Expiry: 600,
        Email: "example@email.com"
    }

    response, err := client.SendEmailMagicLink(payload)
}
```

### 3. Sending OTP

### Send OTP to Phone Number

```go
import (
    "github.com/kisshan13/otplessgo"
)

func main() {
    client := otplessgo.NewOTPLessClient("<client-id>", "<client-secret>", otplessgo.VersionOne)

    payload := &otplessgo.PhoneOTPPayload{
        Phone: "",
        OtpLength: otplessgo.DigitFour,
        Expiry: 600,
        Channels: []otplessgo.Channel{otplessgo.ChannelSMS}
    }

    response, err := client.SendPhoneOTP(payload)
}
```

<br />

### Send OTP to an Email address

```go
import (
    "github.com/kisshan13/otplessgo"
)

func main() {
    client := otplessgo.NewOTPLessClient("<client-id>", "<client-secret>", otplessgo.VersionOne)

    payload := &otplessgo.EmailOTPPayload{
        Email: "",
        OtpLength: otplessgo.DigitFour,
        Expiry: 600,
    }

    response, err := client.SendEmailOTP(payload)
}
```

<br />

### Verify OTP

```go
import (
    "github.com/kisshan13/otplessgo"
)

func main() {
    client := otplessgo.NewOTPLessClient("<client-id>", "<client-secret>", otplessgo.VersionOne)

    payload := &otplessgo.VerifyOTPPayload{
        RequestId: "<request-id>",
        Otp      : "1234",
    }

    response, err := client.VerifyOTP(payload)
}
```

### 4. Sending OTP Links

### Send OTP Link to a Phone Number

```go
import (
    "github.com/kisshan13/otplessgo"
)

func main() {
    client := otplessgo.NewOTPLessClient("<client-id>", "<client-secret>", otplessgo.VersionOne)

    payload := &otplessgo.PhoneOTPLinkPayload{
        Phone: "",
        RedirectURI: "http://localhost:8080/handle",
        OtpLength: otplessgo.DigitFour,
        Expiry: 600,
        Channels: []otplessgo.Channel{otplessgo.ChannelSMS}
    }

    response, err := client.SendPhoneOTPLink(payload)
}
```

<br />

### Send OTP Link to an Email address

```go
import (
    "github.com/kisshan13/otplessgo"
)

func main() {
    client := otplessgo.NewOTPLessClient("<client-id>", "<client-secret>", otplessgo.VersionOne)

    payload := &otplessgo.EmailOTPLinkPayload{
        Email: "example@email.com",
        OtpLength: otplessgo.DigitFour
        Expiry: 600,
        RedirectURI: "http://localhost:8080",
        Channel: []otplessgo.Channel{otplessgo.ChannelEmail}
    }

    response, err := client.SendEmailOTPLink(payload)
}

```

### 5. OAuth Authentication

### Handling OAuth Authentications

```go
import (
    "github.com/kisshan13/otplessgo"
)

func main() {
    client := otplessgo.NewOTPLessClient("<client-id>", "<client-secret>", otplessgo.VersionOne)

    payload := &otplessgo.OAuthPayload{
        Channel: otplessgo.OAuthGoogle,
        RedirectURI: "http://localhost:8080/callback/google"
    }

    response, err := client.InitiateOAuth(payload)
}
```

### Verifying Codes

```go
import (
    "github.com/kisshan13/otplessgo"
)

func main() {
    client := otplessgo.NewOTPLessClient("<client-id>", "<client-secret>", otplessgo.VersionOne)

    response, err := client.VerifyCode("<code from oauth redirect or magic links>")
}
```

## API Response Structure

All responses are returned in a generic `ApiResponse[T]` structure, which contains the following fields:

`Response`: The response data of type T.

`ErrorResponse`: In case of an error, contains the error details.

`RawResponse`: The raw HTTP response object.

Example :

```go
type ApiResponse[T any] struct {
    Response      *T
    ErrorResponse *ApiErrorResponse
    RawResponse   *http.Response
}
```

### Error Handling

In case of an error, you can access the error details through the `ErrorResponse` field of the `ApiResponse` structure.

Example :

```go
type ApiErrorResponse struct {
    Message     string `json:"message,omitempty"`
    Description string `json:"description,omitempty"`
    ErrorCode   string `json:"errorCode,omitempty"`
}
```

### Constants

The SDK defines several constants for authentication channels and OTP lengths:

```go
const (
	ChannelWhatsapp  = Channel("WHATSAPP")
	ChannelSMS       = Channel("SMS")
	ChannelVoiceCall = Channel("VOICE_CALL")
	ChannelViber     = Channel("VIBER")
	ChannelEmail     = Channel("EMAIL")
)
```

```go
const (
	OAuthWhatsapp  = OAuthChannel("WHATSAPP")
	OAuthTwitter   = OAuthChannel("TWITTER")
	OAuthGoogle    = OAuthChannel("GOOGLE")
	OAuthApple     = OAuthChannel("APPLE")
	OAuthLinkedin  = OAuthChannel("LINKEDIN")
	OAuthMicrosoft = OAuthChannel("MICROSOFT")
	OAuthFacebook  = OAuthChannel("FACEBOOK")
	OAuthInstagram = OAuthChannel("INSTAGRAM")
	OAuthLine      = OAuthChannel("LINE")
	OAuthSlack     = OAuthChannel("SLACK")
	OAuthDropbox   = OAuthChannel("DROPBOX")
	OAuthGithub    = OAuthChannel("GITHUB")
	OAuthBitbucket = OAuthChannel("BITBUCKET")
	OAuthAtlassian = OAuthChannel("ATLASSIAN")
	OAuthLinear    = OAuthChannel("LINEAR")
	OAuthGitlab    = OAuthChannel("GITLAB")
	OAuthTitktok   = OAuthChannel("TIKTOK")
	OAuthTwitch    = OAuthChannel("TWITCH")
	OAuthTelegram  = OAuthChannel("TELEGRAM")
	OAuthHubspot   = OAuthChannel("HUBSPOT")
	OAuthNotion    = OAuthChannel("NOTION")
	OAuthBox       = OAuthChannel("BOX")
	OAuthXero      = OAuthChannel("XERO")
)
```

```go
const (
	DigitFour = OtpLength(4)
	DigitSix  = OtpLength(6)
)
```

## Contribution Guide

We welcome contributions to enhance the OTPLess Go SDK. To contribute, please follow these steps:

1. **Fork the Repository**: Create a fork of this repository on your GitHub account.

2. **Clone the Fork**: Clone your fork to your local machine using:
   ```bash
   git clone https://github.com/kisshan13/otplessgo.git
   ```

3. **Create a Branch** : Create a new branch for your feature or bug fix:
    ```bash
    git checkout -b feature-name
    ```
4. **Make Changes** : Implement your feature or fix the bug. Ensure your code is well-documented and adheres to the existing coding style.

5. **Commit Changes** : Commit your changes with a meaningful commit message:
    ```bash
    git commit -m "Add a brief description of your change"
    ```
6. **Push Changes**: Push your branch to your fork:
    ```bash
    git push origin feature-name
    ```

7. **Create a pull request** : Open a pull request (PR) to the main repository. Ensure your PR includes a detailed description of your changes.