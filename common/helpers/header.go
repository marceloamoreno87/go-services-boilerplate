package helpers

import "net/http"

func GetAccessToken(r *http.Request) string {

	bearer := r.Header.Get("Authorization")
	return bearer[7:]
}

func GetStripeSignature(r *http.Request) string {
	return r.Header.Get("Stripe-Signature")
}
