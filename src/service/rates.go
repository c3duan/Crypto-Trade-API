package service

import (
	"fmt"

	"github.com/c3duan/Crypto-Trade-API/pkg"
	"github.com/shopspring/decimal"
)

var (
	one = decimal.RequireFromString("1")
	mockRates = map[string]decimal.Decimal{
		"BTC-USD": decimal.RequireFromString("20023.4"),
		"ETH-USD": decimal.RequireFromString("2017.3"),
		"SOL-USD": decimal.RequireFromString("43.2"),
	}
)

type constantRatesEngine struct {
	rates map[string]decimal.Decimal 
}

func NewConstantRatesService() pkg.RatesEngine {
	return &constantRatesEngine{rates: mockRates}
}

func (s *constantRatesEngine) GetPrice(baseCurrencyID, quoteCurrencyID string) (decimal.Decimal, error) {
	key := fmt.Sprintf("%s-%s", baseCurrencyID, quoteCurrencyID)
	if rate, found := s.rates[key]; found {
		return rate, nil
	}

	return one, nil
}
