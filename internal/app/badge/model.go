package badge

import "time"

type Badge struct {
	BadgeID     int       `json:"badge_id" db:"badge_id"`
	HeritageID  int       `json:"heritage_id" db:"heritage_id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	ImageURL    string    `json:"image_url" db:"image_url"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type UserBadgeInfo struct {
	Name         string    `json:"name" db:"name"`
	ImageURL     string    `json:"image_url" db:"image_url"`
	EarnedAt     time.Time `json:"earned_at" db:"earned_at"`
	HeritageName string    `json:"heritage_name" db:"heritage_name"`
}

type BadgeWithHeritage struct {
	BadgeID      int       `json:"badge_id" db:"badge_id"`
	HeritageID   int       `json:"heritage_id" db:"heritage_id"`
	Name         string    `json:"name" db:"name"`
	Description  string    `json:"description" db:"description"`
	ImageURL     string    `json:"image_url" db:"image_url"`
	HeritageName string    `json:"heritage_name" db:"heritage_name"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}