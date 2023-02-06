package storage

import "sync"

type IncomesById map[string]float64

type Storage struct {
	Incomes IncomesById
	Mu      sync.Mutex
}
