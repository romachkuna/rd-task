package models

type ConversionResponse struct {
	BaseCurrency       string  `json:"base_currency"`
	BaseCurrencyAmount float64 `json:"base_currency_amount"`
	TargetCurrency     string  `json:"target_currency"`
	ConvertedAmount    float64 `json:"amount"`
}
