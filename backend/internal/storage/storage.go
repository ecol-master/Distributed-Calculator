package storage

import (
	"context"
	"database/sql"
	"distributed_calculator/internal/expression"
	"distributed_calculator/internal/users"

	_ "github.com/mattn/go-sqlite3"
)

func NewStorage(filename string) *Storage {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		panic(err)
	}

	err = db.PingContext(context.TODO())
	if err != nil {
		panic(err)
	}
	storage := &Storage{db: db}
	err = storage.createTables()

	if err != nil {
		panic(err)
	}
	return storage
}

func (s *Storage) createTables() error {
	const (
		expressionsTable = `
		CREATE TABLE IF NOT EXISTS expressions(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			expression TEXT NOT NULL,
			user_id INTEGER NOT NULL,
			result INTEGER DEFAULT 0,
			stage INTEGER DEFAULT 0,
			FOREIGN KEY (user_id) REFERENCES users(id)
		);
		`
		usersTable = `
		CREATE TABLE IF NOT EXISTS users(
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				login TEXT NOT NULL,
				password TEXT NOT NULL
		);
		`
	)

	if _, err := s.db.ExecContext(context.TODO(), expressionsTable); err != nil {
		return err
	}

	if _, err := s.db.ExecContext(context.TODO(), usersTable); err != nil {
		return err
	}

	return nil
}

// INSERT INTO TABLES
func (s *Storage) InsertExpression(ctx context.Context, expression string, userID int) (int64, error) {
	var q = `
	INSERT INTO expressions (expression, user_id) values($1, $2)
	`

	res, err := s.db.ExecContext(ctx, q, expression, userID)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (s *Storage) InsertUser(ctx context.Context, login, password string) (int64, error) {
	var q = `
	INSERT INTO users (login, password) values($1, $2)
	`
	res, err := s.db.ExecContext(ctx, q, login, password)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// SELECT FROM TABLES
func (s *Storage) SelectExpressionByID(ctx context.Context, expressionID int) (expression.Expression, error) {
	e := expression.Expression{}
	var q = `
		SELECT id, expression, user_id, result, stage FROM expressions WHERE id = $1
	`
	err := s.db.QueryRowContext(ctx, q, expressionID).Scan(&e.ID, &e.Expression, &e.UserID, &e.Result, &e.Stage)
	return e, err
}

func (s *Storage) SelectExpressionsByUserID(ctx context.Context, userID int) ([]expression.Expression, error) {
	var expressions []expression.Expression
	var q = `
	SELECT id, expression, user_id, result, stage FROM expressions WHERE user_id = $1
	`
	rows, err := s.db.QueryContext(ctx, q, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		e := expression.Expression{}
		err := rows.Scan(&e.ID, &e.Expression, &e.UserID, &e.Result, &e.Stage)
		if err != nil {
			return nil, err
		}
		expressions = append(expressions, e)
	}

	return expressions, nil
}

func (s *Storage) SelectUserByID(ctx context.Context, userID int) (users.User, error) {
	u := users.User{}
	var q = `
	SELECT id, login, password FROM users WHERE id = $1
	`
	err := s.db.QueryRowContext(ctx, q, userID).Scan(&u.ID, &u.Login, &u.Password)
	return u, err
}

// UPDATE

func (s *Storage) UpdateExpression(ctx context.Context, expr expression.Expression) error {
	var q = `
	UPDATE users SET result = $1, stage = $2 WHERE id = $3
	`
	if _, err := s.db.QueryContext(ctx, q, expr.Result, expr.Stage, expr.ID); err != nil {
		return err
	}

	return nil
}
