package storage

import (
	ex "distributed_calculator/internal/expression"
)

type CacheData = map[string]ex.Expression

// Cache constructor
func NewCache() *Cache {
	return &Cache{data: make(CacheData)}
}

func (c *Cache) Update(expr ex.Expression) {
	c.mutex.Lock()
	c.data[expr.ExpressionID] = expr
	c.mutex.Unlock()
}

// returns expression from cache.data
func (c *Cache) GetExpressionByID(expessionID string) (ex.Expression, bool) {
	value, found := c.data[expessionID]
	return value, found
}

func (c *Cache) AddExpression(expr ex.Expression) {
	c.mutex.Lock()
	c.data[expr.ExpressionID] = expr
	c.mutex.Unlock()
}

// returns private field cache.data
func (c *Cache) Data() CacheData {
	return c.data
}

func (c *Cache) SetData(newData CacheData) {
	c.data = newData
}
