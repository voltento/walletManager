package httpQueryModels

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

type GeneralResponse struct {
	Response string `json:"response"`
	Err      string `json:"err,omitempty"`
}
