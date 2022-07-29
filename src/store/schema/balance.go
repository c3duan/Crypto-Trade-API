package schema

import "github.com/shopspring/decimal"

type Balance struct {
	CurrencyID  string           `json:"currency_id"`
	Balance     decimal.Decimal  `json:"balance"`
}