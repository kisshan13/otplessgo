package otplessgo

import "net/http"

// Represents a authentication channel / method
type Channel string

// Represents the OTP Length (4 or 6)
type OtpLength int

// Represents the metadata for a request
type Metadata map[string]string

// Represents authentication channel for oAuth authentication
type OAuthChannel string

const (
	ChannelWhatsapp  = Channel("WHATSAPP")
	ChannelSMS       = Channel("SMS")
	ChannelVoiceCall = Channel("VOICE_CALL")
	ChannelViber     = Channel("VIBER")
	ChannelEmail     = Channel("EMAIL")
)

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

const (
	DigitFour = OtpLength(4)
	DigitSix  = OtpLength(6)
)

// API Response for request made to OTPLESS API
type ApiResponse[T any] struct {
	Response      *T
	ErrorResponse *ApiErrorResponse
	RawResponse   *http.Response
}

// API Error response for request made to OTPLESS API
type ApiErrorResponse struct {
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
	ErrorCode   string `json:"errorCode,omitempty"`
}

type ApiSuccessWithRequestID struct {
	RequestID string `json:"requestId"`
}

type VerifySuccessResponse struct {
	RequestID     string `json:"requestId"`
	IsOTPVerified bool   `json:"isOTPVerified"`
	Message       string `json:"message"`
}

type CodeVerifyPayload struct {
	Code string `json:"code"`
}

type CodeVerificationResponse struct {
	RequestId string `json:"requestId"`
	Message   string `json:"message"`
}

type UserDetails struct {
	Token       string         `json:"token"`
	Status      string         `json:"status"`
	CompletedAt int            `json"completedAt"`
	Identity    []UserIdentity `json:"identities"`
	Network     UserNetwork    `json:"network"`
	DeviceInfo  DeviceInfo     `json:"deviceInfo"`
}

type UserIdentity struct {
	IdentityType      string   `json:"identityType"`
	IdentityValue     string   `json:"identityValue"`
	Channel           string   `json:"channel"`
	Methods           []string `json:"methods"`
	Verified          bool     `json:"verified"`
	VerifiedTimeStamp int      `json:"verifiedTimeStamp"`
}

type UserNetwork struct {
	IP string `json:"ip"`
}

type DeviceInfo struct {
	UserAgent        string  `json:"userAgent"`
	Platform         string  `json:"platform"`
	Vendor           string  `json:"vendor"`
	Language         string  `json:"language"`
	CookieEnabled    bool    `json:"cookieEnabled"`
	ScreenWidth      int     `json:"screenWidth"`
	ScreenHeight     int     `json:"screenHeight"`
	ScreenColorDepth int     `json:"screenColorDepth"`
	DevicePixelRatio float64 `json:"devicePixelRatio"`
}

type RequestStatusResponse struct {
	Token       string `json:"token,omitempty"`
	Status      string `json:"status,omitempty"`
	CompletedAt int    `json:"completedAt,omitempty"`
}

type RequestIdentity struct {
	Type              string   `json:"identityType,omitempty"`
	Value             string   `json:"identityValue,omitempty"`
	Channel           string   `json:"identityChannel,omitempty"`
	Methods           []string `json:"identityMethods,omitempty"`
	Verified          bool     `json:"verified,omitempty"`
	VerifiedTimeStamp int      `json:"verifiedTimeStamp,omitempty"`
}
