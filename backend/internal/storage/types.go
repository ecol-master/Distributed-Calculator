package storage

import (
	"database/sql"
)

// type storage interface {
// 	GetExpression()
// 	GetAllExpressions()

// 	CreateExpression()
// 	CreateUser()

// 	UpdateExpression()
// }

// type Cache struct {
// 	mutex sync.Mutex
// 	data  CacheData
// }

// type Storage struct {
// 	filepath string
// 	ch       *Cache
// 	interval time.Duration
// 	stop     chan struct{}
// }

type Storage struct {
	db *sql.DB
}
