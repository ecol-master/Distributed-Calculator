package storage

import (
	"context"
	"database/sql"

	"distributed_calculator/internal/expression"

	_ "github.com/mattn/go-sqlite3"
)

func NewDBStorage(filename string) *DBStorage {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		panic(err)
	}

	err = db.PingContext(context.TODO())
	if err != nil {
		panic(err)
	}
	storage := &DBStorage{db: db}
	err = storage.createTables()

	if err != nil {
		panic(err)
	}
	return storage
}

func (s *DBStorage) createTables() error {
	const (
		expressionTable = `
		CREATE TABLE IF NOT EXISTS expressions(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			expression TEXT NOT NULL,
			user_id INTEGER NOT NULL,

			FOREIGN KEY (user_id) REFERENCES expressions(id)
		);
		`
	)

	if _, err := s.db.ExecContext(context.TODO(), expressionTable); err != nil {
		return err
	}

	return nil
}

func (s *DBStorage) insertExpression(expr *expression.Expression) error {
	var q = `
	INSERT INTO expressions (expression, user_id) values($1, $2)
	`
	if _, err := s.db.ExecContext(context.TODO(), q, expr.Expression, expr.UserID); err != nil {
		return err
	}

	return nil
}
