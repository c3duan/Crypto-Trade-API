package user

import (
	"net/http"

	"github.com/c3duan/Crypto-Trade-API/api/model"
	"github.com/c3duan/Crypto-Trade-API/pkg"
	"github.com/c3duan/Crypto-Trade-API/src/store/schema"
)

const USD string = "USD"

type UserService struct {
	store pkg.Store
}

func NewUserService(store pkg.Store) pkg.UserService {
	return &UserService{store: store}
}

func (s *UserService) IsCurrencyAllowed(currencies []schema.Currency, currencyID string) bool {	
	for _, currency := range currencies {
		if currency.ID == currencyID {
			return true
		}
	}

	return false
}

func (s *UserService) GetUserCurrencies(userID string) *model.Response {
	currencies, err := s.store.GetCurrencies(userID)
	if err != nil {
		return model.NewResponse(http.StatusBadRequest, nil, err.Error())
	}

	data := model.UserCurrenciesResponse{
		UserID:     userID,
		Currencies: currencies,
	}
	
	return model.NewResponse(http.StatusOK, data, "")
}

func (s *UserService) GetUserBalances(userID string) *model.Response {
	balances, err := s.store.GetBalances(userID)
	if err != nil {
		return model.NewResponse(http.StatusBadRequest, nil, err.Error())
	} 

	data := model.UserBalancesResponse{
		UserID:   userID,
		Balances: balances,
	}
	
	return model.NewResponse(http.StatusOK, data, "")
}

func (s *UserService) GetUserBalance(userID, currencyID string) *model.Response {
	balance, err := s.store.GetBalance(userID, currencyID)
	if err != nil {
		return model.NewResponse(http.StatusBadRequest, nil, err.Error())
	} 
	
	data := model.UserBalanceResponse{
		UserID:     userID,
		CurrencyID: currencyID,
		Balance:    balance,
	}

	return model.NewResponse(http.StatusOK, data, "")
}

func (s *UserService) Buy(request model.UserBuyRequest) *model.Response {
	currencies, err := s.store.GetCurrencies(request.UserID)
	if err != nil {
		return model.NewResponse(http.StatusBadRequest, nil, err.Error())
	}

	if !s.IsCurrencyAllowed(currencies, request.CurrencyID) {
		return model.NewResponse(http.StatusBadRequest, nil, "user cannot transact currency")
	}

	usdBalance, err := s.store.GetBalance(request.UserID, USD)
	if err != nil {
		return model.NewResponse(http.StatusBadRequest, nil, err.Error())
	}

	amountInUSD := request.BuyAmount.Mul(request.PriceInUSD)
	if usdBalance.LessThan(amountInUSD) {
		return model.NewResponse(http.StatusBadRequest, nil, "insufficient fund to buy")
	}

	currencyBalance, err := s.store.GetBalance(request.UserID, request.CurrencyID)
	if err != nil {
		return model.NewResponse(http.StatusBadRequest, nil, err.Error())
	}

	err = s.store.UpdateBalance(request.UserID, USD, usdBalance.Sub(amountInUSD))
	if err != nil {
		return model.NewResponse(http.StatusBadRequest, nil, err.Error())
	}

	err = s.store.UpdateBalance(request.UserID, request.CurrencyID, currencyBalance.Add(request.BuyAmount))
	if err != nil {
		// assumption: best effort revert
		_ = s.store.UpdateBalance(request.UserID, USD, usdBalance)
		return model.NewResponse(http.StatusBadRequest, nil, err.Error())
	}

	return model.NewResponse(http.StatusOK, nil, "buy success")
}

func (s *UserService) Sell(request model.UserSellRequest) *model.Response {
	currencies, err := s.store.GetCurrencies(request.UserID)
	if err != nil {
		return model.NewResponse(http.StatusBadRequest, nil, err.Error())
	}

	if !s.IsCurrencyAllowed(currencies, request.CurrencyID) {
		return model.NewResponse(http.StatusBadRequest, nil, "user cannot transact currency")
	}

	currencyBalance, err := s.store.GetBalance(request.UserID, request.CurrencyID)
	if err != nil {
		return model.NewResponse(http.StatusBadRequest, nil, err.Error())
	}

	if currencyBalance.LessThan(request.SellAmount) {
		return model.NewResponse(http.StatusBadRequest, nil, "insufficient fund to sell")
	}

	usdBalance, err := s.store.GetBalance(request.UserID, USD)
	if err != nil {
		return model.NewResponse(http.StatusBadRequest, nil, err.Error())
	}

	err = s.store.UpdateBalance(request.UserID, request.CurrencyID, currencyBalance.Sub(request.SellAmount))
	if err != nil {
		return model.NewResponse(http.StatusBadRequest, nil, err.Error())
	}

	amountInUSD := request.SellAmount.Mul(request.PriceInUSD)
	err = s.store.UpdateBalance(request.UserID, USD, usdBalance.Add(amountInUSD))
	if err != nil {
		// assumption: best effort revert
		_ = s.store.UpdateBalance(request.UserID, request.CurrencyID, currencyBalance)
		return model.NewResponse(http.StatusBadRequest, nil, err.Error())
	}

	return model.NewResponse(http.StatusOK, nil, "sell success")
}
