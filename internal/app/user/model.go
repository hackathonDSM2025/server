package user

import "time"

type User struct {
	UserID    int       `json:"user_id" db:"user_id"`
	Email     string    `json:"email" db:"email"`
	Nickname  string    `json:"nickname" db:"nickname"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type VisitedPlace struct {
	Name        string    `json:"name" db:"name"`
	VisitedAt   time.Time `json:"visited_at" db:"visited_at"`
	Completed   bool      `json:"completed" db:"quiz_completed"`
}

type UserBadgeInfo struct {
	Name      string    `json:"name" db:"name"`
	ImageURL  string    `json:"image_url" db:"image_url"`
	EarnedAt  time.Time `json:"earned_at" db:"earned_at"`
}