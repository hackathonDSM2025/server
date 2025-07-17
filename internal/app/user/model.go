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

type UserReview struct {
	ReviewID          int       `json:"review_id" db:"review_id"`
	HeritageID        int       `json:"heritage_id" db:"heritage_id"`
	HeritageName      string    `json:"heritage_name" db:"heritage_name"`
	HeritageImageURL  string    `json:"heritage_image_url" db:"heritage_image_url"`
	Rating            int       `json:"rating" db:"rating"`
	ReviewText        string    `json:"review_text" db:"review_text"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}

type UserBadge struct {
	Name              string    `json:"name" db:"name"`
	ImageURL          string    `json:"image_url" db:"image_url"`
	EarnedAt          time.Time `json:"earned_at" db:"earned_at"`
	HeritageName      string    `json:"heritage_name" db:"heritage_name"`
	HeritageImageURL  string    `json:"heritage_image_url" db:"heritage_image_url"`
}

