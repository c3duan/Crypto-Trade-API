package auth

import (
	"net/http"

	"github.com/c3duan/Crypto-Trade-API/api/model"
	"github.com/c3duan/Crypto-Trade-API/pkg"
)

// Dummy authenticator that always return true if the api key is not empty
type auth struct {
}

func NewAuth() pkg.Auth {
	return &auth{}
}

func (a *auth) HasReadAccess(apiKey string) bool {
	return len(apiKey) > 0
}

func (a *auth) HasWriteAccess(apiKey string) bool {
	return len(apiKey) > 0
}


func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		auth := NewAuth()
		apiKey := r.Header.Get("X-API-Key")

		switch r.Method {
		case http.MethodGet:
			if !auth.HasReadAccess(apiKey) {
				model.NewResponse(http.StatusUnauthorized, nil, "API key has not write access").Send(w)
				return
			}
		case http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete:
			if !auth.HasWriteAccess(apiKey) {
				model.NewResponse(http.StatusUnauthorized, nil, "API key has not write access").Send(w)
				return
			}
		}
		
		next.ServeHTTP(w, r)
	})
}
