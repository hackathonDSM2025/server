package heritage

import (
	"context"
	"database/sql"

	"hackathon-dsm-server/internal/pkg/utils/errors"
)

type Repository interface {
	SearchHeritage(ctx context.Context, keyword string) (*Heritage, error)
	GetHeritageByQRCode(ctx context.Context, qrCode string) (*Heritage, error)
	GetUserVisit(ctx context.Context, userID, heritageID int) (*UserVisit, error)
	CreateUserVisit(ctx context.Context, userID, heritageID int) error
	GetBadgeByHeritageID(ctx context.Context, heritageID int) (*Badge, error)
	CheckUserBadgeExists(ctx context.Context, userID, badgeID int) (bool, error)
	CreateUserBadge(ctx context.Context, userID, badgeID int) error
	CreateHeritageReview(ctx context.Context, userID, heritageID int, rating int, reviewText string) error
	GetHeritageReview(ctx context.Context, userID, heritageID int) (*HeritageReview, error)
	UpdateHeritageReview(ctx context.Context, userID, heritageID int, rating int, reviewText string) error
}

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) Repository {
	return &MySQLRepository{db: db}
}

func (r *MySQLRepository) SearchHeritage(ctx context.Context, keyword string) (*Heritage, error) {
	query := `SELECT heritage_id, name, description, image_url, qr_code, story_text, latitude, longitude, created_at 
			  FROM heritage 
			  WHERE name LIKE ? 
			  ORDER BY CASE WHEN name = ? THEN 1 ELSE 2 END, name 
			  LIMIT 1`

	searchPattern := "%" + keyword + "%"

	row := r.db.QueryRowContext(ctx, query, searchPattern, keyword)

	heritage := &Heritage{}
	err := row.Scan(&heritage.HeritageID, &heritage.Name, &heritage.Description,
		&heritage.ImageURL, &heritage.QRCode, &heritage.StoryText,
		&heritage.Latitude, &heritage.Longitude, &heritage.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NotFound("Heritage not found")
		}
		return nil, errors.InternalServerError("Database error")
	}

	return heritage, nil
}

func (r *MySQLRepository) GetHeritageByQRCode(ctx context.Context, qrCode string) (*Heritage, error) {
	query := `SELECT heritage_id, name, description, image_url, qr_code, story_text, latitude, longitude, created_at 
			  FROM heritage 
			  WHERE qr_code = ?`

	row := r.db.QueryRowContext(ctx, query, qrCode)

	heritage := &Heritage{}
	err := row.Scan(&heritage.HeritageID, &heritage.Name, &heritage.Description,
		&heritage.ImageURL, &heritage.QRCode, &heritage.StoryText,
		&heritage.Latitude, &heritage.Longitude, &heritage.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NotFound("Invalid QR code")
		}
		return nil, errors.InternalServerError("Database error")
	}

	return heritage, nil
}

func (r *MySQLRepository) GetUserVisit(ctx context.Context, userID, heritageID int) (*UserVisit, error) {
	query := `SELECT visit_id, user_id, heritage_id, quiz_completed, quiz_score, visited_at 
			  FROM user_visits 
			  WHERE user_id = ? AND heritage_id = ?`

	row := r.db.QueryRowContext(ctx, query, userID, heritageID)

	visit := &UserVisit{}
	err := row.Scan(&visit.VisitID, &visit.UserID, &visit.HeritageID,
		&visit.QuizCompleted, &visit.QuizScore, &visit.VisitedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.InternalServerError("Database error")
	}

	return visit, nil
}

func (r *MySQLRepository) CreateUserVisit(ctx context.Context, userID, heritageID int) error {
	query := `INSERT INTO user_visits (user_id, heritage_id, quiz_completed, quiz_score) 
			  VALUES (?, ?, FALSE, 0)
			  ON DUPLICATE KEY UPDATE visited_at = CURRENT_TIMESTAMP`

	_, err := r.db.ExecContext(ctx, query, userID, heritageID)
	if err != nil {
		return errors.InternalServerError("Failed to create visit record")
	}

	return nil
}

func (r *MySQLRepository) GetBadgeByHeritageID(ctx context.Context, heritageID int) (*Badge, error) {
	query := `SELECT badge_id, name, description, image_url, created_at 
			  FROM badges WHERE heritage_id = ? LIMIT 1`

	row := r.db.QueryRowContext(ctx, query, heritageID)

	badge := &Badge{}
	err := row.Scan(&badge.BadgeID, &badge.Name, &badge.Description,
		&badge.ImageURL, &badge.CreatedAt)
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

func (r *MySQLRepository) CreateHeritageReview(ctx context.Context, userID, heritageID int, rating int, reviewText string) error {
	query := `INSERT INTO heritage_reviews (user_id, heritage_id, rating, review_text) 
			  VALUES (?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, userID, heritageID, rating, reviewText)
	if err != nil {
		return errors.InternalServerError("CreateHeritageReview DB error: " + err.Error())
	}

	return nil
}

func (r *MySQLRepository) GetHeritageReview(ctx context.Context, userID, heritageID int) (*HeritageReview, error) {
	query := `SELECT review_id, user_id, heritage_id, rating, review_text, created_at, updated_at 
			  FROM heritage_reviews 
			  WHERE user_id = ? AND heritage_id = ?`

	row := r.db.QueryRowContext(ctx, query, userID, heritageID)

	review := &HeritageReview{}
	err := row.Scan(&review.ReviewID, &review.UserID, &review.HeritageID,
		&review.Rating, &review.ReviewText, &review.CreatedAt, &review.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.InternalServerError("Database error")
	}

	return review, nil
}

func (r *MySQLRepository) UpdateHeritageReview(ctx context.Context, userID, heritageID int, rating int, reviewText string) error {
	query := `UPDATE heritage_reviews 
			  SET rating = ?, review_text = ?, updated_at = CURRENT_TIMESTAMP 
			  WHERE user_id = ? AND heritage_id = ?`

	_, err := r.db.ExecContext(ctx, query, rating, reviewText, userID, heritageID)
	if err != nil {
		return errors.InternalServerError("Failed to update review")
	}

	return nil
}
