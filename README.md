# Crypto-Trade-API

## Endpoints

1. GET users/currencies/{user_id} 
    
    Currency list: For a given user, get the currencies they can transact
    ```zsh
    curl 0.0.0.0:8080/users/currencies/1 -H "X-API-Key: key"
    ```

2. POST users/buy

    For a given user, buy a selected currency.
    ```zsh
    curl -X POST 0.0.0.0:8080/users/buy -H "Content-Type: application/json" -H "X-API-Key: key" -d '{"user_id": "1", "currency_id": "BTC", "buy_amount": "1", "price_usd": "22030"}'
    ```

3. POST users/sell

    For a given user, sell a selected currency.
    ```zsh
    curl -X POST 0.0.0.0:8080/users/sell -H "Content-Type: application/json" -H "X-API-Key: key" -d '{"user_id": "1", "currency_id": "BTC", "sell_amount": "1", "price_usd": "22030"}'
    ```

4. GET users/balances/{user_id}

    For a given user, get a list of balances
    ```zsh
    curl 0.0.0.0:8080/users/balances/1 -H "X-API-Key: key"
    ```

5. GET users/balance

    For a given user and currency, get the currency balance
    ```zsh
    curl 0.0.0.0:8080/users/balance/\?user_id=1\&currency_id=BTC -H "X-API-Key: key"
    ```

6. GET prices

    For a given pair of base and quote currency, get the conversion rate
    ```zsh
    curl 0.0.0.0:8080/prices\?base_currency_id=BTC\&quote_currency_id=USD
    ```


## Assumptions
- Buy endpoint only supports USD -> Crypto (No Crypto -> Crypto conversion)
- Sell endpoint only supports Crypto -> USD (No Crypto -> Crypto conversion)
- Authentication to access user info and perform user buy and sell is done via API Key. The current authentication only check if the API key is non-empty.
- We assume buys and sells are performed at a pre-determined price (quote) and all transaction will succeed if sufficient funds are present ignoring market volatility.


## Future Improvements
- Use persistent database to store Users, Currencies, Balances, and Transactions.
- Use Transactional Database to reliably rollback all changes if transactions fail.
- Add actual Authentication Middleware that uses either JWT or API Key.
- Add actual Rates Engine that polls real-time price data from various Oracle (CoinGecko, CoinMarketCap, Coinbase, ChainLink, etc.)
- Place user buys and sells as actual orders to different CEXs and DEXs (Coinbase, FTX, Binance, Uniswap, Curve, etc.)
- Add Risk monitoring for outstanding orders/credits


## Installation Steps
1. Clone the repo:
```zsh
git clone https://github.com/c3duan/Crypto-Trade-API.git
```
2. Build the app:
```zsh
go build
```

3. Run the app:
```zsh
go run main.go
```


## Testing Steps

1. Run the app:
```zsh
go run main.go
```

2. Run the tests:
```zsh
cd test
go test
```

> Note: since all data are stored in memory, we need to re-start the app after each test run to revert to initial data state.
