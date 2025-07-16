package heritage

import "time"

type Heritage struct {
	HeritageID  int     `json:"heritage_id" db:"heritage_id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	ImageURL    string  `json:"image_url" db:"image_url"`
	QRCode      string  `json:"qr_code" db:"qr_code"`
	StoryText   string  `json:"story_text" db:"story_text"`
	Latitude    float64 `json:"latitude" db:"latitude"`
	Longitude   float64 `json:"longitude" db:"longitude"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type UserVisit struct {
	VisitID       int       `json:"visit_id" db:"visit_id"`
	UserID        int       `json:"user_id" db:"user_id"`
	HeritageID    int       `json:"heritage_id" db:"heritage_id"`
	QuizCompleted bool      `json:"quiz_completed" db:"quiz_completed"`
	QuizScore     int       `json:"quiz_score" db:"quiz_score"`
	VisitedAt     time.Time `json:"visited_at" db:"visited_at"`
}

type Badge struct {
	BadgeID     int       `json:"badge_id" db:"badge_id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	ImageURL    string    `json:"image_url" db:"image_url"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type HeritageReview struct {
	ReviewID   int       `json:"review_id" db:"review_id"`
	UserID     int       `json:"user_id" db:"user_id"`
	HeritageID int       `json:"heritage_id" db:"heritage_id"`
	Rating     int       `json:"rating" db:"rating"`
	ReviewText string    `json:"review_text" db:"review_text"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}