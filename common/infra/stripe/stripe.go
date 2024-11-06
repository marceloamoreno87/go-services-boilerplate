package stripe

import (
	"os"

	"github.com/stripe/stripe-go/v80/client"
)

type StripeClient struct {
	StripeClient *client.API
}

func NewStripeClient() *StripeClient {
	stripe := &client.API{}
	stripe.Init(os.Getenv("STRIPE_CLIENT_SECRET"), nil)
	return &StripeClient{
		StripeClient: stripe,
	}
}

type StripeSignatureSecret struct {
	SignatureSecret string `json:"signature_secret"`
}

func NewStripeSignatureSecret() *StripeSignatureSecret {
	return &StripeSignatureSecret{
		SignatureSecret: os.Getenv("STRIPE_SIGNATURE_SECRET"),
	}
}
