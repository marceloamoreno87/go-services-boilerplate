package stripe

import (
	"fmt"

	"github.com/stripe/stripe-go/v80"
)

type CreatePaymentMethod struct {
	MethodType string
	Token      string
	TaxID      string
	Name       string
	Email      string
	City       string
	State      string
	Country    string
	Line1      string
	Line2      string
	PostalCode string
}

func NewCreatePaymentMethod(
	methodType string,
) *CreatePaymentMethod {
	return &CreatePaymentMethod{
		MethodType: methodType,
	}
}

func (sc *CreatePaymentMethod) SetBillet(
	name string,
	email string,
	taxID string,
	city string,
	state string,
	country string,
	line1 string,
	line2 string,
	postalCode string,
) *CreatePaymentMethod {
	sc.TaxID = taxID
	sc.Name = name
	sc.Email = email
	sc.City = city
	sc.State = state
	sc.Country = country
	sc.Line1 = line1
	sc.Line2 = line2
	sc.PostalCode = postalCode
	return sc
}

func (sc *CreatePaymentMethod) SetToken(token string) *CreatePaymentMethod {
	sc.Token = token
	return sc
}
func (sc *CreatePaymentMethod) SetName(name string) *CreatePaymentMethod {
	sc.Name = name
	return sc
}

func (sc *CreatePaymentMethod) Request() (paymentMethod *stripe.PaymentMethod, err error) {
	var params *stripe.PaymentMethodParams

	switch sc.MethodType {
	case "card":
		params = &stripe.PaymentMethodParams{
			Type: stripe.String(sc.MethodType),
			Card: &stripe.PaymentMethodCardParams{
				Token: stripe.String(sc.Token),
			},
		}
	case "pix":
		params = &stripe.PaymentMethodParams{
			Type: stripe.String(sc.MethodType),
			Pix:  &stripe.PaymentMethodPixParams{},
		}
	case "boleto":
		params = &stripe.PaymentMethodParams{
			Type: stripe.String(sc.MethodType),
			Boleto: &stripe.PaymentMethodBoletoParams{
				TaxID: stripe.String(sc.TaxID),
			},
			BillingDetails: &stripe.PaymentMethodBillingDetailsParams{
				Name:  stripe.String(sc.Name),
				Email: stripe.String(sc.Email),
				Address: &stripe.AddressParams{
					City:       stripe.String(sc.City),
					State:      stripe.String(sc.State),
					Country:    stripe.String(sc.Country),
					Line1:      stripe.String(sc.Line1),
					Line2:      stripe.String(sc.Line2),
					PostalCode: stripe.String(sc.PostalCode),
				},
			},
		}
	default:
		return nil, fmt.Errorf("unsupported payment method type: %s", sc.MethodType)
	}

	paymentMethod, err = NewStripeClient().StripeClient.PaymentMethods.New(params)
	if err != nil {
		return
	}
	return
}

type UpdatePaymentMethod struct {
	PaymentMethodID string
	OrderId         string
}

func NewUpdatePaymentMethod() *UpdatePaymentMethod {
	return &UpdatePaymentMethod{}
}

func (upm *UpdatePaymentMethod) Request() (paymentMethod *stripe.PaymentMethod, err error) {
	params := &stripe.PaymentMethodParams{
		Metadata: map[string]string{
			"order_id": upm.OrderId,
		},
	}

	paymentMethod, err = NewStripeClient().StripeClient.PaymentMethods.Update(upm.PaymentMethodID, params)
	if err != nil {
		return
	}
	return
}

type GetPaymentMethod struct {
	customer *stripe.Customer
}

func NewGetPaymentMethod(customer *stripe.Customer) *GetPaymentMethod {
	return &GetPaymentMethod{
		customer: customer,
	}
}

func (gpm *GetPaymentMethod) Request() (paymentMethod *stripe.PaymentMethodList, err error) {
	params := &stripe.PaymentMethodListParams{
		Customer: stripe.String(gpm.customer.ID),
	}
	iter := NewStripeClient().StripeClient.PaymentMethods.List(params)

	var paymentMethods stripe.PaymentMethodList
	for iter.Next() {
		paymentMethods.Data = append(paymentMethods.Data, iter.PaymentMethod())
	}

	if err := iter.Err(); err != nil {
		return nil, err
	}

	return &paymentMethods, nil
}

type AttachPaymentMethod struct {
	CustomerID      string
	PaymentMethodID string
}

func NewAttachPaymentMethod(
	customerID string,
	paymentMethodID string,
) *AttachPaymentMethod {
	return &AttachPaymentMethod{
		CustomerID:      customerID,
		PaymentMethodID: paymentMethodID,
	}
}

func (apm *AttachPaymentMethod) Request() (paymentMethod *stripe.PaymentMethod, err error) {
	params := &stripe.PaymentMethodAttachParams{
		Customer: stripe.String(apm.CustomerID),
	}

	paymentMethod, err = NewStripeClient().StripeClient.PaymentMethods.Attach(apm.PaymentMethodID, params)
	if err != nil {
		return
	}
	return
}
