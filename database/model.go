package database

type Account struct {
	Id       string  `json:"id"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

type PaymentDirection int

const (
	Incoming PaymentDirection = iota
	Outgoing
)

type Payment struct {
	AccountId string           `json:"id"`
	Amount    float64          `json:"amount"`
	ToAccount string           `json:"id"`
	Direction PaymentDirection `json:"direction"`
}
