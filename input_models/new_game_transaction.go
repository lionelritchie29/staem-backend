package input_models

type NewGameTransaction struct {
	User int
	CardNo string
	CardExp string
	BillingAddress string
	BillingCity string
	PostalCode string
	PhoneNumber string
	Country string
	PaymentMethod int
	Details []NewGameTransactionDetail
}
