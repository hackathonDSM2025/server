package quiz

import "time"

type Quiz struct {
	QuizID     int       `json:"quiz_id" db:"quiz_id"`
	HeritageID int       `json:"heritage_id" db:"heritage_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

type Question struct {
	QuestionID    int       `json:"question_id" db:"question_id"`
	QuizID        int       `json:"quiz_id" db:"quiz_id"`
	QuestionText  string    `json:"question_text" db:"question_text"`
	Option1       string    `json:"option1" db:"option1"`
	Option2       string    `json:"option2" db:"option2"`
	Option3       string    `json:"option3" db:"option3"`
	Option4       string    `json:"option4" db:"option4"`
	CorrectAnswer int       `json:"correct_answer" db:"correct_answer"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}

type Badge struct {
	BadgeID     int       `json:"badge_id" db:"badge_id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	ImageURL    string    `json:"image_url" db:"image_url"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type UserBadge struct {
	UserBadgeID int       `json:"user_badge_id" db:"user_badge_id"`
	UserID      int       `json:"user_id" db:"user_id"`
	BadgeID     int       `json:"badge_id" db:"badge_id"`
	EarnedAt    time.Time `json:"earned_at" db:"earned_at"`
}

type UserVisitStatus struct {
	Visited       bool `json:"visited"`
	QuizCompleted bool `json:"quiz_completed"`
	QuizScore     int  `json:"quiz_score"`
}