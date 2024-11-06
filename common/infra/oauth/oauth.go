package oauth

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
)

var GoogleOAuthConfig = &oauth2.Config{
	RedirectURL:  os.Getenv("GOOGLE_OAUTH2_CLIENT_CALLBACK_URL"),
	ClientID:     os.Getenv("GOOGLE_OAUTH2_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_OAUTH2_CLIENT_SECRET"),
	Scopes: []string{
		"openid",
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	},
	Endpoint: google.Endpoint,
}

var AppleOAuthConfig = &oauth2.Config{
	RedirectURL:  os.Getenv("APPLE_OAUTH2_CLIENT_CALLBACK_URL"),
	ClientID:     os.Getenv("APPLE_OAUTH2_CLIENT_ID"),
	ClientSecret: os.Getenv("APPLE_OAUTH2_CLIENT_SECRET"),
	Scopes: []string{
		"name",
		"email",
	},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://appleid.apple.com/auth/authorize",
		TokenURL: "https://appleid.apple.com/auth/token",
	},
}

var FacebookOAuthConfig = &oauth2.Config{
	RedirectURL:  os.Getenv("FACEBOOK_OAUTH2_CLIENT_CALLBACK_URL"),
	ClientID:     os.Getenv("FACEBOOK_OAUTH2_CLIENT_ID"),
	ClientSecret: os.Getenv("FACEBOOK_OAUTH2_CLIENT_SECRET"),
	Scopes: []string{
		"public_profile",
		"email",
	},
	Endpoint: facebook.Endpoint,
}

type GoogleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
}

type AppleUser struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type FacebookUser struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture struct {
		Data struct {
			URL string `json:"url"`
		} `json:"data"`
	} `json:"picture"`
}
