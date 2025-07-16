package user

import "time"

type User struct {
	UserID    int       `json:"user_id" db:"user_id"`
	Username  string    `json:"username" db:"username"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type VisitedPlace struct {
	Name        string    `json:"name" db:"name"`
	VisitedAt   time.Time `json:"visited_at" db:"visited_at"`
	Completed   bool      `json:"completed" db:"quiz_completed"`
}

