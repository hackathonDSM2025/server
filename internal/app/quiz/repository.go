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
	GetBadgeByHeritageID(ctx context.Context, heritageID int) (*BadgeWithHeritage, error)
	CreateUserBadge(ctx context.Context, userID, badgeID int) error
	CheckUserBadgeExists(ctx context.Context, userID, badgeID int) (bool, error)
	GetUserVisitStatus(ctx context.Context, userID, heritageID int) (*UserVisitStatus, error)
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
	// First check if user_visits record exists
	var visitID int
	checkQuery := `SELECT visit_id FROM user_visits WHERE user_id = ? AND heritage_id = ?`
	err := r.db.QueryRowContext(ctx, checkQuery, userID, heritageID).Scan(&visitID)
	
	if err == sql.ErrNoRows {
		// No visit record exists, create one
		insertQuery := `INSERT INTO user_visits (user_id, heritage_id, quiz_completed, quiz_score) 
						VALUES (?, ?, TRUE, ?)`
		_, err := r.db.ExecContext(ctx, insertQuery, userID, heritageID, score)
		if err != nil {
			return errors.InternalServerError("Failed to create visit record with quiz score")
		}
	} else if err != nil {
		return errors.InternalServerError("Database error while checking visit record")
	} else {
		// Visit record exists, update it
		updateQuery := `UPDATE user_visits 
						SET quiz_completed = TRUE, quiz_score = ? 
						WHERE user_id = ? AND heritage_id = ?`
		_, err := r.db.ExecContext(ctx, updateQuery, score, userID, heritageID)
		if err != nil {
			return errors.InternalServerError("Failed to update quiz score")
		}
	}

	return nil
}

func (r *MySQLRepository) GetBadgeByHeritageID(ctx context.Context, heritageID int) (*BadgeWithHeritage, error) {
	query := `SELECT b.badge_id, b.name, b.description, b.image_url, h.name as heritage_name, h.image_url as heritage_image_url, b.created_at 
			  FROM badges b 
			  JOIN heritage h ON b.heritage_id = h.heritage_id
			  WHERE b.heritage_id = ? LIMIT 1`
	
	row := r.db.QueryRowContext(ctx, query, heritageID)

	badge := &BadgeWithHeritage{}
	err := row.Scan(&badge.BadgeID, &badge.Name, &badge.Description, 
		&badge.ImageURL, &badge.HeritageName, &badge.HeritageImageURL, &badge.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.InternalServerError("GetBadgeByHeritageID DB error: " + err.Error())
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

func (r *MySQLRepository) GetUserVisitStatus(ctx context.Context, userID, heritageID int) (*UserVisitStatus, error) {
	query := `SELECT quiz_completed, quiz_score FROM user_visits WHERE user_id = ? AND heritage_id = ?`
	row := r.db.QueryRowContext(ctx, query, userID, heritageID)

	var quizCompleted bool
	var quizScore int
	err := row.Scan(&quizCompleted, &quizScore)
	if err != nil {
		if err == sql.ErrNoRows {
			return &UserVisitStatus{
				Visited:       false,
				QuizCompleted: false,
				QuizScore:     0,
			}, nil
		}
		return nil, errors.InternalServerError("Database error")
	}

	return &UserVisitStatus{
		Visited:       true,
		QuizCompleted: quizCompleted,
		QuizScore:     quizScore,
	}, nil
}