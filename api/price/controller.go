package price

import (
	"net/http"

	"github.com/c3duan/Crypto-Trade-API/api/model"
	"github.com/c3duan/Crypto-Trade-API/pkg"
)

type PriceController struct {
	service pkg.PriceService
}

func NewPriceController(service pkg.PriceService) *PriceController {
	return &PriceController{service: service}
}

func (c *PriceController) GetPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := r.URL.Query()
	baseCurrencyID := vars.Get("base_currency_id")
	quoteCurrencyID := vars.Get("quote_currency_id")

	if baseCurrencyID == "" || quoteCurrencyID == "" {
		model.NewResponse(http.StatusBadRequest, nil, "invalid currency ids").Send(w)
		return
	}

	response := c.service.GetPrice(baseCurrencyID, quoteCurrencyID)
	response.Send(w)
}
