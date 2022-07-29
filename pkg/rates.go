package pkg

import "github.com/shopspring/decimal"

type RatesEngine interface {
	GetPrice(baseCurrencyID, quoteCurrencyID string) (decimal.Decimal, error)
}