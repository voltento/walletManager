package httpQueryModels

// Data structure models for communication between clients and the wallet manager.

type Account struct {
	Id       string  `json:"id"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

type Payment struct {
	Id           int
	From_account string  `json:"from_account_id"`
	Amount       float64 `json:"amount"`
	To_account   string  `json:"to_account_id"`
}

// General response entity for queries which don't have some specific data in a response
type GeneralResponse struct {
	Response string `json:"response"`
	Err      string `json:"err,omitempty"`
}
