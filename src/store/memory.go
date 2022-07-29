package store

import (
	"fmt"
	"sync"

	"github.com/shopspring/decimal"

	"github.com/c3duan/Crypto-Trade-API/src/store/schema"
	"github.com/c3duan/Crypto-Trade-API/pkg"
)

type memoryStore struct {
	sync.RWMutex
	users   map[string]*schema.User
}

func NewMemoryStore() pkg.Store {
	users := make(map[string]*schema.User)
	for _, user := range mockUserData {
		users[user.ID] = user
	}
	return &memoryStore{
		users: users,
	}
}

func (s *memoryStore) GetCurrencies(userID string) ([]schema.Currency, error) {
	s.RLock()
	defer s.RUnlock()
	if user, found := s.users[userID]; found {
		return user.Currencies, nil
	}
	return nil, fmt.Errorf("user not found")
}

func (s *memoryStore) GetBalances(userID string) ([]schema.Balance, error) {
	s.RLock()
	defer s.RUnlock()

	if user, found := s.users[userID]; found {
		var balances []schema.Balance
		for _, currency := range user.Currencies {
			balances = append(balances, schema.Balance{
				CurrencyID: currency.ID,
				Balance:    user.Balances[currency.ID],
			})
		}
		return balances, nil
	}
	
	return nil, fmt.Errorf("user not found")
}

func (s *memoryStore) GetBalance(userID, currencyID string) (decimal.Decimal, error) {
	s.RLock()
	defer s.RUnlock()
	
	if user, found := s.users[userID]; found {
		return user.Balances[currencyID], nil
	}
	
	return decimal.Zero, fmt.Errorf("user not found")
}

func (s *memoryStore) UpdateBalance(userID, currencyID string, newBalance decimal.Decimal) error {
	s.Lock()
	defer s.Unlock()
	
	if user, found := s.users[userID]; found {
		user.Balances[currencyID] = newBalance
		return nil
	}
	
	return fmt.Errorf("user not found")
}
