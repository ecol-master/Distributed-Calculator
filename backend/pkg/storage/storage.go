package storage

import (
	cache "distributed_calculator/pkg/cache"
	ex "distributed_calculator/pkg/expression"
)

type Storage struct {
	filepath string
	ch       *cache.Cache
}

func NewStorage(fp string) *Storage {
	return &Storage{
		filepath: fp,
		ch:       cache.NewCache(),
	}
}

func GetExpressionByID() (ex.Expression, error) {

	return ex.Expression{}, nil
}

// function open file and add new data to cache
func (s *Storage) loadFileData() {

}

// function add new expression to cache and to data file
// every 5 seconds
func (s *Storage) AddExpression(ex.Expression) {

}
