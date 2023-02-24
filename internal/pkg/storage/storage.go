package storage

import (
	"sync"
	"time"
)

type userIncome struct {
	Balance   float64
	UpdatedAt time.Time
}

type IncomesById map[string]userIncome

type Storage struct {
	incomes IncomesById
	mu      sync.Mutex
}

func New() *Storage {
	return &Storage{incomes: make(IncomesById), mu: sync.Mutex{}}
}

func (s *Storage) GetIncomes() IncomesById {
	return s.incomes
}

func (s *Storage) SetBalance(username string, income float64) {
	row, _ := s.incomes[username]

	row.Balance = s.incomes[username].Balance + income
	row.UpdatedAt = time.Now()
	s.incomes[username] = row
}

func (s *Storage) GetMutex() *sync.Mutex {
	return &s.mu
}
