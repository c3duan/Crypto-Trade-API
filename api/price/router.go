package price

import (
	"github.com/c3duan/Crypto-Trade-API/pkg"
	"github.com/gorilla/mux"
)

func New(rt *mux.Router, engine pkg.RatesEngine) {
	route := rt.PathPrefix("/prices").Subrouter()

	service := NewPriceService(engine)
	controller := NewPriceController(service)

	route.HandleFunc("", controller.GetPrice).Methods("GET")
}