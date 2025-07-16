package user

import (
	"context"
	"database/sql"

	"hackathon-dsm-server/internal/pkg/utils/errors"
)

type Repository interface {
	GetUserByID(ctx context.Context, userID int) (*User, error)
	GetVisitedPlaces(ctx context.Context, userID int) ([]VisitedPlace, error)
	GetVisitedCount(ctx context.Context, userID int) (int, error)
	GetUserReviewCount(ctx context.Context, userID int) (int, error)
	GetUserReviews(ctx context.Context, userID int) ([]UserReview, error)
	GetUserBadgeCount(ctx context.Context, userID int) (int, error)
	GetUserBadges(ctx context.Context, userID int) ([]UserBadge, error)
}

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) Repository {
	return &MySQLRepository{db: db}
}

func (r *MySQLRepository) GetUserByID(ctx context.Context, userID int) (*User, error) {
	query := `SELECT user_id, username, created_at FROM users WHERE user_id = ?`
	row := r.db.QueryRowContext(ctx, query, userID)

	user := &User{}
	err := row.Scan(&user.UserID, &user.Username, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NotFound("User not found")
		}
		return nil, errors.InternalServerError("Database error")
	}

	return user, nil
}

func (r *MySQLRepository) GetVisitedPlaces(ctx context.Context, userID int) ([]VisitedPlace, error) {
	query := `SELECT h.name, uv.visited_at, uv.quiz_completed 
			  FROM user_visits uv 
			  JOIN heritage h ON uv.heritage_id = h.heritage_id 
			  WHERE uv.user_id = ? 
			  ORDER BY uv.visited_at DESC`
	
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, errors.InternalServerError("Database error")
	}
	defer rows.Close()

	var places []VisitedPlace
	for rows.Next() {
		var place VisitedPlace
		err := rows.Scan(&place.Name, &place.VisitedAt, &place.Completed)
		if err != nil {
			return nil, errors.InternalServerError("Database error")
		}
		places = append(places, place)
	}

	return places, nil
}

func (r *MySQLRepository) GetVisitedCount(ctx context.Context, userID int) (int, error) {
	query := `SELECT COUNT(*) FROM user_visits WHERE user_id = ?`
	row := r.db.QueryRowContext(ctx, query, userID)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, errors.InternalServerError("Database error")
	}

	return count, nil
}

func (r *MySQLRepository) GetUserReviewCount(ctx context.Context, userID int) (int, error) {
	query := `SELECT COUNT(*) FROM heritage_reviews WHERE user_id = ?`
	row := r.db.QueryRowContext(ctx, query, userID)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, errors.InternalServerError("Database error")
	}

	return count, nil
}

func (r *MySQLRepository) GetUserReviews(ctx context.Context, userID int) ([]UserReview, error) {
	query := `SELECT hr.review_id, hr.heritage_id, h.name as heritage_name, hr.rating, hr.review_text, hr.created_at, hr.updated_at
			  FROM heritage_reviews hr
			  JOIN heritage h ON hr.heritage_id = h.heritage_id
			  WHERE hr.user_id = ?
			  ORDER BY hr.created_at DESC`
	
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, errors.InternalServerError("Database error")
	}
	defer rows.Close()

	var reviews []UserReview
	for rows.Next() {
		var review UserReview
		err := rows.Scan(&review.ReviewID, &review.HeritageID, &review.HeritageName, 
			&review.Rating, &review.ReviewText, &review.CreatedAt, &review.UpdatedAt)
		if err != nil {
			return nil, errors.InternalServerError("Database error")
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}

func (r *MySQLRepository) GetUserBadgeCount(ctx context.Context, userID int) (int, error) {
	query := `SELECT COUNT(*) FROM user_badges WHERE user_id = ?`
	row := r.db.QueryRowContext(ctx, query, userID)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, errors.InternalServerError("Database error")
	}

	return count, nil
}

func (r *MySQLRepository) GetUserBadges(ctx context.Context, userID int) ([]UserBadge, error) {
	query := `SELECT b.name, b.image_url, ub.earned_at, h.name as heritage_name
			  FROM user_badges ub
			  JOIN badges b ON ub.badge_id = b.badge_id
			  JOIN heritage h ON b.heritage_id = h.heritage_id
			  WHERE ub.user_id = ?
			  ORDER BY ub.earned_at DESC`
	
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, errors.InternalServerError("Database error")
	}
	defer rows.Close()

	var badges []UserBadge
	for rows.Next() {
		var badge UserBadge
		err := rows.Scan(&badge.Name, &badge.ImageURL, &badge.EarnedAt, &badge.HeritageName)
		if err != nil {
			return nil, errors.InternalServerError("Database error")
		}
		badges = append(badges, badge)
	}

	return badges, nil
}