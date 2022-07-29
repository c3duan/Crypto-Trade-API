package api

import (
	"net/http"

	"github.com/c3duan/Crypto-Trade-API/api/price"
	"github.com/c3duan/Crypto-Trade-API/api/user"
	"github.com/c3duan/Crypto-Trade-API/src/service"
	"github.com/c3duan/Crypto-Trade-API/src/store"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func New() http.Handler {
	mainRouter := mux.NewRouter()

	store := store.NewMemoryStore()
	engine := service.NewConstantRatesService()
	
	user.New(mainRouter, store)
	price.New(mainRouter, engine)

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT"},
		AllowedHeaders: []string{"Content-Type", "X-CSRF-Token"},
	}).Handler(mainRouter)

	return corsMiddleware
}
