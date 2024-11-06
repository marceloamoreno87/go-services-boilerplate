package stripe

import "github.com/stripe/stripe-go/v80"

type CreateSubscription struct {
	CustomerID      string
	PriceID         string
	PaymentMethodID string
}

func NewCreateSubscription(
	customerID string,
	priceID string,
	paymentMethodID string,
) *CreateSubscription {
	return &CreateSubscription{
		CustomerID:      customerID,
		PriceID:         priceID,
		PaymentMethodID: paymentMethodID,
	}
}

func (sc *CreateSubscription) Request() (subscription *stripe.Subscription, err error) {

	params := &stripe.SubscriptionParams{
		Customer: stripe.String(sc.CustomerID),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price: stripe.String(sc.PriceID),
			},
		},
		DefaultPaymentMethod: stripe.String(sc.PaymentMethodID),
	}

	subscription, err = NewStripeClient().StripeClient.Subscriptions.New(params)
	if err != nil {
		return
	}
	return

}

type UpdateSubscription struct {
	SubscriptionID string
	PriceID        string
}

func NewUpdateSubscription() *UpdateSubscription {
	return &UpdateSubscription{}
}

func (us *UpdateSubscription) Request() (subscription *stripe.Subscription, err error) {

	params := &stripe.SubscriptionParams{
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price: stripe.String(us.PriceID),
			},
		},
	}

	subscription, err = NewStripeClient().StripeClient.Subscriptions.Update(us.SubscriptionID, params)
	if err != nil {
		return
	}
	return

}

type GetSubscription struct {
	CustomerID string
}

func NewGetSubscriptions(customerID string) *GetSubscription {
	return &GetSubscription{CustomerID: customerID}
}

func (gs *GetSubscription) Request() (subscription *stripe.Subscription, err error) {

	params := &stripe.SubscriptionListParams{
		Customer: stripe.String(gs.CustomerID),
	}

	iter := NewStripeClient().StripeClient.Subscriptions.List(params)
	for iter.Next() {
		subscription = iter.Subscription()
	}
	if iter.Err() != nil {
		err = iter.Err()
		return
	}
	return

}
