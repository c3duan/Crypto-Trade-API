package model

import "github.com/shopspring/decimal"

type UserBuyRequest struct {
	UserID      string           `json:"user_id"`
	CurrencyID  string           `json:"currency_id"`
	BuyAmount   decimal.Decimal  `json:"buy_amount"`
	PriceInUSD  decimal.Decimal  `json:"price_usd"`
}

type UserSellRequest struct {
	UserID      string           `json:"user_id"`
	CurrencyID  string           `json:"currency_id"`
	SellAmount  decimal.Decimal  `json:"sell_amount"`
	PriceInUSD  decimal.Decimal  `json:"price_usd"`
}
