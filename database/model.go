package database

type Account struct {
	Id       string
	Currency string
	Balance  float64
}

type PaymentDirection int

const (
	Incoming PaymentDirection = iota
	Outgoing
)

type Payment struct {
	AccountId string
	Amount    float64
	ToAccount string
	Direction PaymentDirection
}
