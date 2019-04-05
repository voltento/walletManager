package httpModel

// Data structure models for communication between clients and the wallet manager.

// General response entity for queries which don't have some specific data in a response
type GeneralResponse struct {
	Response string `json:"response"`
	Err      string `json:"err,omitempty"`
}

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

type ChangeBalanceRequest struct {
	Id     string  `json:"id"`
	Amount float64 `json:"change_amount"`
}

type SendMoneyRequest struct {
	FromAccId string  `json:"from_account"`
	ToAccId   string  `json:"to_account"`
	Amount    float64 `json:"change_amount"`
}
