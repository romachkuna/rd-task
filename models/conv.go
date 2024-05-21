package models

type ConversionResponse struct {
	BaseCurrency   int `json:"base_currency"`
	TargetCurrency int `json:"target_currency"`
	Amount         int `json:"amount"`
}
