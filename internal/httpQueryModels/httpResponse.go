package httpQueryModels

type Account struct {
	Id       string  `json:"id"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

type GeneralResponse struct {
	Response string `json:"response"`
	Err      string `json:"err,omitempty"`
}
