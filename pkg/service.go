package pkg

import "github.com/c3duan/Crypto-Trade-API/api/model"

type UserService interface {
	GetUserCurrencies(userID string) *model.Response
	GetUserBalances(userID string) *model.Response
	GetUserBalance(userID, currencyID string) *model.Response
	Buy(request model.UserBuyRequest) *model.Response 
	Sell(request model.UserSellRequest) *model.Response 
}

type PriceService interface {
	GetPrice(baseCurrencyID, quoteCurrencyID string) *model.Response
}
