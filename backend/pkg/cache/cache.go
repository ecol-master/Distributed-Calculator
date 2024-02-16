package cache

import (
	ex "distributed_calculator/pkg/expression"
	"sync"
)

type Cache struct {
	mutex sync.Mutex
	data  map[string]ex.Expression
}

func NewCache() *Cache {
	return &Cache{data: make(map[string]ex.Expression)}
}

func (c *Cache) Update(expr ex.Expression) {
	c.mutex.Lock()
	c.data[expr.ExpressionID] = expr
	c.mutex.Unlock()
}

func (c *Cache) GetAllData() map[string]ex.Expression {
	return c.data
}

func (c *Cache) GetExpressionByID(expessionID string) (ex.Expression, bool) {
	value, found := c.data[expessionID]
	return value, found
}

func (c *Cache) AddExpression(expr ex.Expression) {
	c.mutex.Lock()
	c.data[expr.ExpressionID] = expr
	c.mutex.Unlock()
}

func (c *Cache) Data() map[string]ex.Expression {
	return c.data
}

func (c *Cache) SetData(newData map[string]ex.Expression) {
	c.data = newData
}
