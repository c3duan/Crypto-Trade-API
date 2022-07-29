package schema

import "github.com/shopspring/decimal"

type User struct {
	ID         string                      `json:"id"`
	Currencies []Currency	               `json:"currencies"`
	Balances   map[string]decimal.Decimal  `json:"balances"`   // currency_id -> balance
}
