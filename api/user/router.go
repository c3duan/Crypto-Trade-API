package user

import (
	"github.com/c3duan/Crypto-Trade-API/pkg"
	"github.com/c3duan/Crypto-Trade-API/src/middleware/auth"
	"github.com/gorilla/mux"
)

func New(rt *mux.Router, store pkg.Store) {
	route := rt.PathPrefix("/users").Subrouter()
	route.Use(auth.CheckAuth)

	service := NewUserService(store)
	controller := NewUserController(service)

	route.HandleFunc("/currencies/{user_id}", controller.GetUserCurrencies).Methods("GET")
	route.HandleFunc("/balances/{user_id}", controller.GetUserBalances).Methods("GET")
	route.HandleFunc("/balance", controller.GetUserBalance).Methods("GET")
	route.HandleFunc("/buy", controller.Buy).Methods("POST")
	route.HandleFunc("/sell", controller.Sell).Methods("POST")
}