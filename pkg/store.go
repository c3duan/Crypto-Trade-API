package pkg

import (
	"github.com/c3duan/Crypto-Trade-API/src/store/schema"
	"github.com/shopspring/decimal"
)

type Store interface {
	GetCurrencies(userID string) ([]schema.Currency, error)
	GetBalances(userID string) ([]schema.Balance, error)
	GetBalance(userID, currencyID string) (decimal.Decimal, error)
	UpdateBalance(userID, currencyID string, newBalance decimal.Decimal) error
}
