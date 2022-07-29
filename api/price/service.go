package price

import (
	"net/http"

	"github.com/c3duan/Crypto-Trade-API/api/model"
	"github.com/c3duan/Crypto-Trade-API/pkg"
)

type PriceService struct {
	engine pkg.RatesEngine
}

func NewPriceService(engine pkg.RatesEngine) pkg.PriceService {
	return &PriceService{engine: engine}
}

func (s *PriceService) GetPrice(baseCurrencyID, quoteCurrencyID string) *model.Response {
	price, err := s.engine.GetPrice(baseCurrencyID, quoteCurrencyID)
	if err != nil {
		return model.NewResponse(http.StatusBadRequest, nil, err.Error())
	}

	return model.NewResponse(http.StatusOK, model.PriceResponse{
		BaseCurrencyID:  baseCurrencyID,
		QuoteCurrencyID: quoteCurrencyID,
		Price:           price,
	}, "")
}
