package stripe

import (
	"strconv"

	"github.com/stripe/stripe-go/v80"
)

type CustomerAddress struct {
	City       string
	Country    string
	Line1      string
	Line2      string
	PostalCode string
	State      string
}

type CreateCustomer struct {
	Email       string
	Name        string
	Description string
	Phone       string
	Address     CustomerAddress
	Metadata    map[string]string
}

func NewCreateCustomer(
	email string,
	name string,
	description string,
	phone string,
	address CustomerAddress,
	metadata map[string]string,
) *CreateCustomer {
	return &CreateCustomer{
		Address:     address,
		Metadata:    metadata,
		Email:       email,
		Name:        name,
		Description: description,
		Phone:       phone,
	}
}

func (sc *CreateCustomer) Request() (customer *stripe.Customer, err error) {

	params := &stripe.CustomerParams{
		Email:       stripe.String(sc.Email),
		Name:        stripe.String(sc.Name),
		Description: stripe.String(sc.Description),
		Phone:       stripe.String(sc.Phone),
		Address: &stripe.AddressParams{
			City:       stripe.String(sc.Address.City),
			Country:    stripe.String(sc.Address.Country),
			Line1:      stripe.String(sc.Address.Line1),
			Line2:      stripe.String(sc.Address.Line2),
			PostalCode: stripe.String(sc.Address.PostalCode),
			State:      stripe.String(sc.Address.State),
		},
		Metadata: sc.Metadata,
	}

	customer, err = NewStripeClient().StripeClient.Customers.New(params)
	if err != nil {
		return
	}
	return

}

type UpdateCustomer struct {
	CustomerID  string
	Email       string
	Name        string
	Description string
	Phone       string
	Address     CustomerAddress
	Metadata    map[string]string
}

func NewUpdateCustomer(
	customerID string,
	email string,
	name string,
	description string,
	phone string,
	address CustomerAddress,
	metadata map[string]string,
) *UpdateCustomer {
	return &UpdateCustomer{
		CustomerID:  customerID,
		Email:       email,
		Name:        name,
		Description: description,
		Phone:       phone,
		Address:     address,
		Metadata:    metadata,
	}
}

func (uc *UpdateCustomer) Request() (customer *stripe.Customer, err error) {

	params := &stripe.CustomerParams{
		Email:       stripe.String(uc.Email),
		Name:        stripe.String(uc.Name),
		Description: stripe.String(uc.Description),
		Phone:       stripe.String(uc.Phone),
		Address: &stripe.AddressParams{
			City:       stripe.String(uc.Address.City),
			Country:    stripe.String(uc.Address.Country),
			Line1:      stripe.String(uc.Address.Line1),
			Line2:      stripe.String(uc.Address.Line2),
			PostalCode: stripe.String(uc.Address.PostalCode),
			State:      stripe.String(uc.Address.State),
		},
		Metadata: uc.Metadata,
	}

	customer, err = NewStripeClient().StripeClient.Customers.Update(uc.CustomerID, params)
	if err != nil {
		return
	}
	return

}

type DeleteCustomer struct {
	CustomerID string
}

func NewDeleteCustomer(customerID string) *DeleteCustomer {
	return &DeleteCustomer{
		CustomerID: customerID,
	}
}

func (dc *DeleteCustomer) Request() (err error) {
	params := &stripe.CustomerParams{}

	_, err = NewStripeClient().StripeClient.Customers.Del(dc.CustomerID, params)
	if err != nil {
		return
	}
	return

}

type GetCustomerByUserID struct {
	UserID int
}

func NewGetCustomer(userID int) *GetCustomerByUserID {
	return &GetCustomerByUserID{
		UserID: userID,
	}
}

func (gc *GetCustomerByUserID) Request() (customer *stripe.Customer, err error) {

	params := &stripe.CustomerSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "metadata['userId']:'" + strconv.Itoa(gc.UserID) + "'",
		},
	}
	iter := NewStripeClient().StripeClient.Customers.Search(params)

	for iter.Next() {
		customer = iter.Customer()
	}

	if err = iter.Err(); err != nil {
		return
	}

	return

}
