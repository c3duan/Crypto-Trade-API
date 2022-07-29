package store

import (
	"github.com/c3duan/Crypto-Trade-API/src/store/schema"
	"github.com/shopspring/decimal"
)

var mockUserData = []*schema.User{
	{
		ID: "1",
		Currencies: []schema.Currency{
			{
				Name: "Bitcoin",
				ID:   "BTC",
			},
			{
				Name: "Ethereum",
				ID:   "ETH",
			},
			{
				Name: "Solana",
				ID:   "SOL",
			},
			{
				Name: "US Dollar",
				ID:   "USD",
			},
		},
		Balances: map[string]decimal.Decimal{
			"BTC": decimal.RequireFromString("22.34"),
			"ETH": decimal.RequireFromString("243.141"),
			"SOL": decimal.RequireFromString("2022.72"),
			"USD": decimal.RequireFromString("4013418.82"),
		},
	},
	{
		ID: "2",
		Currencies: []schema.Currency{
			{
				Name: "Bitcoin",
				ID:   "BTC",
			},
			{
				Name: "US Dollar",
				ID:   "USD",
			},
		},
		Balances: map[string]decimal.Decimal{
			"BTC": decimal.RequireFromString("202.34"),
			"USD": decimal.RequireFromString("713418.82"),
		},
	},
	{
		ID: "3",
		Currencies: []schema.Currency{
			{
				Name: "Ethereum",
				ID:   "ETH",
			},
			{
				Name: "US Dollar",
				ID:   "USD",
			},
		},
		Balances: map[string]decimal.Decimal{
			"ETH": decimal.RequireFromString("203.141"),
			"USD": decimal.RequireFromString("432529.82"),
		},
	},
}