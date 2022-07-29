package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/c3duan/Crypto-Trade-API/api/model"
	"github.com/c3duan/Crypto-Trade-API/src/store/schema"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const ADDRS string = "http://0.0.0.0:8080"

func TestGetUserCurrencies(t *testing.T) {
	type currenciesResponse struct {
		Status  int                           `json:"status"`
		Data    model.UserCurrenciesResponse  `json:"data"`
	}

	testCases := []struct {
		name         string
		userID       string
		statusCode   int
		err          error
		expected     model.UserCurrenciesResponse
	}{
		{
			name: "get currencies for existing user 1",
			userID: "1",
			statusCode: http.StatusOK,
			err: nil,
			expected: model.UserCurrenciesResponse{
				UserID:     "1",
				Currencies: []schema.Currency{
					{
						Name: "Bitcoin",
						ID:   "BTC",
					},
					{
						Name: "Ethereum",
						ID:   "ETH",
					},
					{
						Name: "Solana",
						ID:   "SOL",
					},
					{
						Name: "US Dollar",
						ID:   "USD",
					},
				},
			},
		},
		{
			name: "get currencies for existing user 2",
			userID: "2",
			statusCode: http.StatusOK,
			err: nil,
			expected: model.UserCurrenciesResponse{
				UserID:     "2",
				Currencies: []schema.Currency{
					{
						Name: "Bitcoin",
						ID:   "BTC",
					},
					{
						Name: "US Dollar",
						ID:   "USD",
					},
				},
			},
		},
		{
			name: "user not found",
			userID: "4",
			statusCode: http.StatusBadRequest,
			err: fmt.Errorf("user not found"),
			expected: model.UserCurrenciesResponse{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", fmt.Sprintf("%s/users/currencies/%s", ADDRS, tc.userID), nil)
			require.NoError(t, err)

			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-API-Key", "test-key")

			client := http.DefaultClient
			res, err := client.Do(req)
			require.NoError(t, err)
			defer res.Body.Close()
			
			assert.Equal(t, tc.statusCode, res.StatusCode)
			data, err := ioutil.ReadAll(res.Body)
			require.NoError(t, err)

			if tc.err == nil {
				var actual currenciesResponse
				err = json.Unmarshal(data, &actual)
				require.NoError(t, err)
				assert.Equal(t, tc.expected, actual.Data)
			} else {
				var actual model.Response
				err = json.Unmarshal(data, &actual)
				require.NoError(t, err)
				assert.Equal(t, tc.err.Error(), actual.Message)
			}
		})
	}
}

func TestGetUserBalancies(t *testing.T) {
	type balancesResponse struct {
		Status  int                        `json:"status"`
		Data    model.UserBalancesResponse `json:"data"`
	}

	testCases := []struct {
		name         string
		userID       string
		statusCode   int
		err          error
		expected     model.UserBalancesResponse
	}{
		{
			name: "get balances for existing user 1",
			userID: "1",
			statusCode: http.StatusOK,
			err: nil,
			expected: model.UserBalancesResponse{
				UserID:   "1",
				Balances: []schema.Balance{
					{
						CurrencyID: "BTC",
						Balance: decimal.RequireFromString("22.34"),
					},
					{
						CurrencyID: "ETH",
						Balance: decimal.RequireFromString("243.141"),
					},
					{
						CurrencyID: "SOL",
						Balance: decimal.RequireFromString("2022.72"),
					},
					{
						CurrencyID: "USD",
						Balance: decimal.RequireFromString("4013418.82"),
					},
				},
			},
		},
		{
			name: "get balances for existing user 3",
			userID: "3",
			statusCode: http.StatusOK,
			err: nil,
			expected: model.UserBalancesResponse{
				UserID:   "3",
				Balances: []schema.Balance{
					{
						CurrencyID: "ETH",
						Balance: decimal.RequireFromString("203.141"),
					},
					{
						CurrencyID: "USD",
						Balance: decimal.RequireFromString("432529.82"),
					},
				},
			},
		},
		{
			name: "user not found",
			userID: "5",
			statusCode: http.StatusBadRequest,
			err: fmt.Errorf("user not found"),
			expected: model.UserBalancesResponse{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", fmt.Sprintf("%s/users/balances/%s", ADDRS, tc.userID), nil)
			require.NoError(t, err)

			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-API-Key", "test-key")

			client := http.DefaultClient
			res, err := client.Do(req)
			require.NoError(t, err)
			defer res.Body.Close()
			
			assert.Equal(t, tc.statusCode, res.StatusCode)
			data, err := ioutil.ReadAll(res.Body)
			require.NoError(t, err)

			if tc.err == nil {
				var actual balancesResponse
				json.Unmarshal(data, &actual)
				assert.True(t, tc.expected.Equal(actual.Data))
			} else {
				var actual model.Response
				json.Unmarshal(data, &actual)
				assert.Equal(t, tc.err.Error(), actual.Message)
			}
		})
	}
}

type priceResponse struct {
	Status  int                 `json:"status"`
	Data    model.PriceResponse `json:"data"`
}

type balanceResponse struct {
	Status  int                       `json:"status"`
	Data    model.UserBalanceResponse `json:"data"`
}

func getPrice(t *testing.T, client *http.Client, currencyID string) decimal.Decimal {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/prices", ADDRS), nil)
	require.NoError(t, err)
	q := req.URL.Query()
	q.Add("base_currency_id", currencyID)
	q.Add("quote_currency_id", "USD")
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	require.NoError(t, err)

	defer res.Body.Close()
	
	assert.Equal(t, http.StatusOK, res.StatusCode)
	data, err := ioutil.ReadAll(res.Body)
	require.NoError(t, err)

	var priceRes priceResponse
	err = json.Unmarshal(data, &priceRes)
	require.NoError(t, err)

	return priceRes.Data.Price
}

func checkBalances(t *testing.T, client *http.Client, userID, currencyID string, 
	expectedCurrencyBalance, expectedUSDBalance decimal.Decimal) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users/balance", ADDRS), nil)
	require.NoError(t, err)
	q := req.URL.Query()
	q.Add("user_id", userID)
	q.Add("currency_id", currencyID)
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", "test-key")

	res, err := client.Do(req)
	require.NoError(t, err)

	defer res.Body.Close()
		
	assert.Equal(t, http.StatusOK, res.StatusCode)
	data, err := ioutil.ReadAll(res.Body)
	require.NoError(t, err)

	var balanceRes balanceResponse
	err = json.Unmarshal(data, &balanceRes)
	require.NoError(t, err)

	assert.True(t, expectedCurrencyBalance.Equal(balanceRes.Data.Balance))

	req, err = http.NewRequest("GET", fmt.Sprintf("%s/users/balance", ADDRS), nil)
	require.NoError(t, err)
	q = req.URL.Query()
	q.Add("user_id", userID)
	q.Add("currency_id", "USD")
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", "test-key")

	res, err = client.Do(req)
	require.NoError(t, err)

	defer res.Body.Close()
		
	assert.Equal(t, http.StatusOK, res.StatusCode)
	data, err = ioutil.ReadAll(res.Body)
	require.NoError(t, err)

	balanceRes = balanceResponse{}
	err = json.Unmarshal(data, &balanceRes)
	require.NoError(t, err)

	assert.True(t, expectedUSDBalance.Equal(balanceRes.Data.Balance))
}

func TestBuy(t *testing.T) {
	testCases := []struct {
		name                     string
		userID                   string
		currencyID               string
		amount                   decimal.Decimal
		statusCode               int
		message                  string
		expectedCurrencyBalance  decimal.Decimal
		expectedUSDBalance       decimal.Decimal
	}{
		{
			name:                    "successfully buy transactable currency",
			userID:                  "1",
			currencyID:              "BTC",
			amount:                  decimal.RequireFromString("0.75"),
			statusCode:              http.StatusOK,
			message:                 "buy success",
			expectedCurrencyBalance: decimal.RequireFromString("23.09"),
			expectedUSDBalance:      decimal.RequireFromString("3998401.27"),
		},
		{
			name:                    "not transactable currency",
			userID:                  "2",
			currencyID:              "SOL",
			amount:                  decimal.RequireFromString("13.4"),
			statusCode:              http.StatusBadRequest,
			message:                 "user cannot transact currency",
			expectedCurrencyBalance: decimal.Zero,
			expectedUSDBalance:      decimal.RequireFromString("713418.82"),
		},
		{
			name:                    "insufficient USD fund",
			userID:                  "2",
			currencyID:              "BTC",
			amount:                  decimal.RequireFromString("66"),
			statusCode:              http.StatusBadRequest,
			message:                 "insufficient fund to buy",
			expectedCurrencyBalance: decimal.RequireFromString("202.34"),
			expectedUSDBalance:      decimal.RequireFromString("713418.82"),
		},
		{
			name:                    "zero amount",
			userID:                  "1",
			currencyID:              "BTC",
			amount:                  decimal.Zero,
			statusCode:              http.StatusBadRequest,
			message:                 "invalid decimal fields",
			expectedCurrencyBalance: decimal.RequireFromString("23.09"),
			expectedUSDBalance:      decimal.RequireFromString("3998401.27"),
		},
		{
			name:                    "cannot buy USD with USD",
			userID:                  "1",
			currencyID:              "USD",
			amount:                  decimal.RequireFromString("0.75"),
			statusCode:              http.StatusBadRequest,
			message:                 "invalid currency id",
			expectedCurrencyBalance: decimal.RequireFromString("3998401.27"),
			expectedUSDBalance:      decimal.RequireFromString("3998401.27"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {	
			client := http.DefaultClient
		
			priceInUSD := getPrice(t, client, tc.currencyID)
			buyReq := model.UserBuyRequest{
				UserID:     tc.userID,
				CurrencyID: tc.currencyID,
				BuyAmount:  tc.amount,
				PriceInUSD: priceInUSD,
			}
			data, err := json.Marshal(buyReq)
			require.NoError(t, err)
			req, err := http.NewRequest("POST", fmt.Sprintf("%s/users/buy", ADDRS), bytes.NewBuffer(data))
			require.NoError(t, err)

			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-API-Key", "test-key")

			res, err := client.Do(req)
			require.NoError(t, err)
			defer res.Body.Close()
			
			assert.Equal(t, tc.statusCode, res.StatusCode)
			data, err = ioutil.ReadAll(res.Body)
			require.NoError(t, err)

			var actual model.Response
			json.Unmarshal(data, &actual)
			assert.Equal(t, tc.message, actual.Message)

			checkBalances(t, client, tc.userID, tc.currencyID, tc.expectedCurrencyBalance, tc.expectedUSDBalance)
		})
	}
}

func TestSell(t *testing.T) {
	testCases := []struct {
		name                     string
		userID                   string
		currencyID               string
		amount                   decimal.Decimal
		err                      error
		statusCode               int
		message                  string
		expectedCurrencyBalance  decimal.Decimal
		expectedUSDBalance       decimal.Decimal
	}{
		{
			name:                    "successfully sell transactable currency",
			userID:                  "3",
			currencyID:              "ETH",
			amount:                  decimal.RequireFromString("32"),
			statusCode:              http.StatusOK,
			message:                 "sell success",
			expectedCurrencyBalance: decimal.RequireFromString("171.141"),
			expectedUSDBalance:      decimal.RequireFromString("497083.42"),
		},
		{
			name:                    "not transactable currency",
			userID:                  "2",
			currencyID:              "SOL",
			amount:                  decimal.RequireFromString("13.4"),
			statusCode:              http.StatusBadRequest,
			message:                 "user cannot transact currency",
			expectedCurrencyBalance: decimal.Zero,
			expectedUSDBalance:      decimal.RequireFromString("713418.82"),
		},
		{
			name:                    "insufficient BTC fund",
			userID:                  "2",
			currencyID:              "BTC",
			amount:                  decimal.RequireFromString("420"),
			statusCode:              http.StatusBadRequest,
			message:                 "insufficient fund to sell",
			expectedCurrencyBalance: decimal.RequireFromString("202.34"),
			expectedUSDBalance:      decimal.RequireFromString("713418.82"),
		},
		{
			name:                    "zero amount",
			userID:                  "3",
			currencyID:              "ETH",
			amount:                  decimal.Zero,
			statusCode:              http.StatusBadRequest,
			message:                 "invalid decimal fields",
			expectedCurrencyBalance: decimal.RequireFromString("171.141"),
			expectedUSDBalance:      decimal.RequireFromString("497083.42"),
		},
		{
			name:                    "cannot sell USD with USD",
			userID:                  "2",
			currencyID:              "USD",
			amount:                  decimal.RequireFromString("0.75"),
			statusCode:              http.StatusBadRequest,
			message:                 "invalid currency id",
			expectedCurrencyBalance: decimal.RequireFromString("713418.82"),
			expectedUSDBalance:      decimal.RequireFromString("713418.82"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {	
			client := http.DefaultClient
		
			priceInUSD := getPrice(t, client, tc.currencyID)
			buyReq := model.UserSellRequest{
				UserID:     tc.userID,
				CurrencyID: tc.currencyID,
				SellAmount: tc.amount,
				PriceInUSD: priceInUSD,
			}
			data, err := json.Marshal(buyReq)
			require.NoError(t, err)
			req, err := http.NewRequest("POST", fmt.Sprintf("%s/users/sell", ADDRS), bytes.NewBuffer(data))
			require.NoError(t, err)

			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-API-Key", "test-key")

			res, err := client.Do(req)
			require.NoError(t, err)
			defer res.Body.Close()
			
			assert.Equal(t, tc.statusCode, res.StatusCode)
			data, err = ioutil.ReadAll(res.Body)
			require.NoError(t, err)

			expectedMessage := tc.message
			if tc.err != nil {
				expectedMessage = tc.err.Error()
			}

			var actual model.Response
			json.Unmarshal(data, &actual)
			assert.Equal(t, expectedMessage, actual.Message)

			checkBalances(t, client, tc.userID, tc.currencyID, tc.expectedCurrencyBalance, tc.expectedUSDBalance)
		})
	}
}
