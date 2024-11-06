package stripe

import "github.com/stripe/stripe-go/v80"

type CreatePrice struct {
	ProductName string
	Amount      int64
	Currency    string
	Interval    string
}

func NewCreatePrice() *CreatePrice {
	return &CreatePrice{}
}

func (sc *CreatePrice) Request() (price *stripe.Price, err error) {

	params := &stripe.PriceParams{
		Currency:   stripe.String(string(sc.Currency)),
		UnitAmount: stripe.Int64(sc.Amount),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(string(sc.Interval)),
		},
		ProductData: &stripe.PriceProductDataParams{Name: stripe.String(sc.ProductName)},
	}

	price, err = NewStripeClient().StripeClient.Prices.New(params)
	if err != nil {
		return
	}
	return

}

type UpdatePrice struct {
	PriceID string
	Amount  int64
}

func NewUpdatePrice() *UpdatePrice {
	return &UpdatePrice{}
}

func (up *UpdatePrice) Request() (price *stripe.Price, err error) {

	params := &stripe.PriceParams{
		UnitAmount: stripe.Int64(up.Amount),
	}

	price, err = NewStripeClient().StripeClient.Prices.Update(up.PriceID, params)
	if err != nil {
		return
	}
	return

}

type GetPrice struct {
	ProductID *string
}

func NewGetPrice(productID *string) *GetPrice {
	return &GetPrice{ProductID: productID}
}

func (gp *GetPrice) Request() (price *stripe.Price, err error) {

	params := &stripe.PriceSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "product:'" + *gp.ProductID + "'",
		},
	}

	iter := NewStripeClient().StripeClient.Prices.Search(params)

	for iter.Next() {
		price = iter.Price()
	}

	if err = iter.Err(); err != nil {
		return
	}

	return

}
