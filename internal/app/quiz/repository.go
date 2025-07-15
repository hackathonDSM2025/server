package quiz

import (
	"context"
	"database/sql"

	"hackathon-dsm-server/internal/pkg/utils/errors"
)

type Repository interface {
	GetQuizByHeritageID(ctx context.Context, heritageID int) (*Quiz, error)
	GetQuestionsByQuizID(ctx context.Context, quizID int) ([]Question, error)
	UpdateUserVisitScore(ctx context.Context, userID, heritageID, score int) error
	GetBadgeByCondition(ctx context.Context, conditionType string) (*Badge, error)
	CreateUserBadge(ctx context.Context, userID, badgeID int) error
	CheckUserBadgeExists(ctx context.Context, userID, badgeID int) (bool, error)
}

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) Repository {
	return &MySQLRepository{db: db}
}

func (r *MySQLRepository) GetQuizByHeritageID(ctx context.Context, heritageID int) (*Quiz, error) {
	query := `SELECT quiz_id, heritage_id, created_at FROM quiz WHERE heritage_id = ?`
	row := r.db.QueryRowContext(ctx, query, heritageID)

	quiz := &Quiz{}
	err := row.Scan(&quiz.QuizID, &quiz.HeritageID, &quiz.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NotFound("Quiz not found for this heritage")
		}
		return nil, errors.InternalServerError("Database error")
	}

	return quiz, nil
}

func (r *MySQLRepository) GetQuestionsByQuizID(ctx context.Context, quizID int) ([]Question, error) {
	query := `SELECT question_id, quiz_id, question_text, option1, option2, option3, option4, correct_answer, created_at 
			  FROM questions WHERE quiz_id = ? ORDER BY question_id`
	
	rows, err := r.db.QueryContext(ctx, query, quizID)
	if err != nil {
		return nil, errors.InternalServerError("Database error")
	}
	defer rows.Close()

	var questions []Question
	for rows.Next() {
		var q Question
		err := rows.Scan(&q.QuestionID, &q.QuizID, &q.QuestionText, 
			&q.Option1, &q.Option2, &q.Option3, &q.Option4, 
			&q.CorrectAnswer, &q.CreatedAt)
		if err != nil {
			return nil, errors.InternalServerError("Database error")
		}
		questions = append(questions, q)
	}

	return questions, nil
}

func (r *MySQLRepository) UpdateUserVisitScore(ctx context.Context, userID, heritageID, score int) error {
	query := `UPDATE user_visits 
			  SET quiz_completed = TRUE, quiz_score = ? 
			  WHERE user_id = ? AND heritage_id = ?`
	
	_, err := r.db.ExecContext(ctx, query, score, userID, heritageID)
	if err != nil {
		return errors.InternalServerError("Failed to update quiz score")
	}

	return nil
}

func (r *MySQLRepository) GetBadgeByCondition(ctx context.Context, conditionType string) (*Badge, error) {
	query := `SELECT badge_id, name, description, image_url, condition_type, created_at 
			  FROM badges WHERE condition_type = ? LIMIT 1`
	
	row := r.db.QueryRowContext(ctx, query, conditionType)

	badge := &Badge{}
	err := row.Scan(&badge.BadgeID, &badge.Name, &badge.Description, 
		&badge.ImageURL, &badge.ConditionType, &badge.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.InternalServerError("Database error")
	}

	return badge, nil
}

func (r *MySQLRepository) CreateUserBadge(ctx context.Context, userID, badgeID int) error {
	query := `INSERT INTO user_badges (user_id, badge_id) VALUES (?, ?)`
	_, err := r.db.ExecContext(ctx, query, userID, badgeID)
	if err != nil {
		return errors.InternalServerError("Failed to award badge")
	}

	return nil
}

func (r *MySQLRepository) CheckUserBadgeExists(ctx context.Context, userID, badgeID int) (bool, error) {
	query := `SELECT COUNT(*) FROM user_badges WHERE user_id = ? AND badge_id = ?`
	row := r.db.QueryRowContext(ctx, query, userID, badgeID)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, errors.InternalServerError("Database error")
	}

	return count > 0, nil
}