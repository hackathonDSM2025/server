package auth

import "time"

type User struct {
	UserID    int       `json:"user_id" db:"user_id"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"`
	Nickname  string    `json:"nickname" db:"nickname"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}