package auth

import (
	"context"
	"database/sql"

	"hackathon-dsm-server/internal/pkg/utils/errors"
)

type Repository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	GetUserByID(ctx context.Context, userID int) (*User, error)
}

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) Repository {
	return &MySQLRepository{db: db}
}

func (r *MySQLRepository) CreateUser(ctx context.Context, user *User) error {
	query := `INSERT INTO users (username, password) VALUES (?, ?)`
	result, err := r.db.ExecContext(ctx, query, user.Username, user.Password)
	if err != nil {
		return errors.Conflict("Username already exists")
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return errors.InternalServerError("Failed to get user ID")
	}

	user.UserID = int(userID)
	return nil
}

func (r *MySQLRepository) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	query := `SELECT user_id, username, password, created_at FROM users WHERE username = ?`
	row := r.db.QueryRowContext(ctx, query, username)

	user := &User{}
	err := row.Scan(&user.UserID, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NotFound("User not found")
		}
		return nil, errors.InternalServerError("Database error")
	}

	return user, nil
}

func (r *MySQLRepository) GetUserByID(ctx context.Context, userID int) (*User, error) {
	query := `SELECT user_id, username, password, created_at FROM users WHERE user_id = ?`
	row := r.db.QueryRowContext(ctx, query, userID)

	user := &User{}
	err := row.Scan(&user.UserID, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NotFound("User not found")
		}
		return nil, errors.InternalServerError("Database error")
	}

	return user, nil
}