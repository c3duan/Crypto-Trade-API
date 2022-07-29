package user

import (
	"encoding/json"
	"net/http"

	"github.com/c3duan/Crypto-Trade-API/api/model"
	"github.com/c3duan/Crypto-Trade-API/pkg"
	"github.com/gorilla/mux"
)

type UserController struct {
	service pkg.UserService
}

func NewUserController(service pkg.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) GetUserCurrencies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userID := vars["user_id"]
	
	response := c.service.GetUserCurrencies(userID)
	response.Send(w)
}

func (c *UserController) GetUserBalances(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userID := vars["user_id"]
	
	response := c.service.GetUserBalances(userID)
	response.Send(w)
}

func (c *UserController) GetUserBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := r.URL.Query()
	userID := vars.Get("user_id")
	currencyID := vars.Get("currency_id")
	
	if currencyID == "" {
		model.NewResponse(http.StatusBadRequest, nil, "invalid currency id").Send(w)
		return
	}

	response := c.service.GetUserBalance(userID, currencyID)
	response.Send(w)
}

func (c *UserController) Buy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var request model.UserBuyRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		model.NewResponse(http.StatusBadRequest, "", err.Error()).Send(w)
		return
	}

	if request.CurrencyID == "" || request.CurrencyID == "USD" {
		model.NewResponse(http.StatusBadRequest, nil, "invalid currency id").Send(w)
		return
	}

	if request.BuyAmount.IsZero() || request.PriceInUSD.IsZero() {
		model.NewResponse(http.StatusBadRequest, nil, "invalid decimal fields").Send(w)
		return
	}

	response := c.service.Buy(request)
	response.Send(w)
}

func (c *UserController) Sell(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var request model.UserSellRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		model.NewResponse(http.StatusBadRequest, nil, err.Error()).Send(w)
		return
	}

	if request.CurrencyID == "" || request.CurrencyID == "USD" {
		model.NewResponse(http.StatusBadRequest, nil, "invalid currency id").Send(w)
		return
	}

	if request.SellAmount.IsZero() || request.PriceInUSD.IsZero() {
		model.NewResponse(http.StatusBadRequest, nil, "invalid decimal fields").Send(w)
		return
	}

	response := c.service.Sell(request)
	response.Send(w)
}
