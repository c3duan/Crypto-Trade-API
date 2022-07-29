package model

import (
	"encoding/json"
	"net/http"

	"github.com/c3duan/Crypto-Trade-API/src/store/schema"
	"github.com/shopspring/decimal"
)

type UserCurrenciesResponse struct {
	UserID      string            `json:"user_id"`
	Currencies  []schema.Currency `json:"currencies"`
}

type UserBalancesResponse struct {
	UserID      string            `json:"user_id"`
	Balances    []schema.Balance  `json:"balances"`
}

func (r UserBalancesResponse) Equal(other UserBalancesResponse) bool {
	if r.UserID != other.UserID {
		return false
	}

	if len(r.Balances) != len(other.Balances) {
		return false
	}

	for i, balance := range r.Balances {
		otherBalance := other.Balances[i]
		if balance.CurrencyID != otherBalance.CurrencyID || !balance.Balance.Equal(otherBalance.Balance) {
			return false
		}
	}

	return true
}

type UserBalanceResponse struct {
	UserID      string            `json:"user_id"`
	CurrencyID  string            `json:"currency_id"`
	Balance     decimal.Decimal   `json:"balances"`
}

func (r UserBalanceResponse) Equal(other UserBalanceResponse) bool {
	return r.UserID == other.UserID && r.Balance.Equal(other.Balance)
}

type PriceResponse struct {
	BaseCurrencyID   string          `json:"base_currency_id"`
	QuoteCurrencyID  string          `json:"quote_currency_id"`
	Price            decimal.Decimal `json:"price"`
}

type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func (res *Response) Send(w http.ResponseWriter) {
	w.WriteHeader(res.Status)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		w.Write([]byte("Error When Encode Response"))
	}
}

func NewResponse(code int, data interface{}, message string) *Response {
	return &Response{
		Status:  code,
		Message: message,
		Data:    data,
	}
}