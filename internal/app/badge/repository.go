package badge

import (
	"context"
	"database/sql"

	"hackathon-dsm-server/internal/pkg/utils/errors"
)

type Repository interface {
	GetUserBadges(ctx context.Context, userID int) ([]UserBadgeInfo, error)
	GetBadgeCount(ctx context.Context, userID int) (int, error)
	GetBadgeByHeritageID(ctx context.Context, heritageID int) (*Badge, error)
	CheckUserBadgeExists(ctx context.Context, userID, badgeID int) (bool, error)
	CreateUserBadge(ctx context.Context, userID, badgeID int) error
}

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) Repository {
	return &MySQLRepository{db: db}
}

func (r *MySQLRepository) GetUserBadges(ctx context.Context, userID int) ([]UserBadgeInfo, error) {
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

	var badges []UserBadgeInfo
	for rows.Next() {
		var badge UserBadgeInfo
		err := rows.Scan(&badge.Name, &badge.ImageURL, &badge.EarnedAt, &badge.HeritageName)
		if err != nil {
			return nil, errors.InternalServerError("Database error")
		}
		badges = append(badges, badge)
	}

	return badges, nil
}

func (r *MySQLRepository) GetBadgeCount(ctx context.Context, userID int) (int, error) {
	query := `SELECT COUNT(*) FROM user_badges WHERE user_id = ?`
	row := r.db.QueryRowContext(ctx, query, userID)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, errors.InternalServerError("Database error")
	}

	return count, nil
}

func (r *MySQLRepository) GetBadgeByHeritageID(ctx context.Context, heritageID int) (*Badge, error) {
	query := `SELECT badge_id, heritage_id, name, description, image_url, created_at FROM badges WHERE heritage_id = ?`
	row := r.db.QueryRowContext(ctx, query, heritageID)

	badge := &Badge{}
	err := row.Scan(&badge.BadgeID, &badge.HeritageID, &badge.Name, &badge.Description, &badge.ImageURL, &badge.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NotFound("Badge not found")
		}
		return nil, errors.InternalServerError("Database error")
	}

	return badge, nil
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

func (r *MySQLRepository) CreateUserBadge(ctx context.Context, userID, badgeID int) error {
	query := `INSERT INTO user_badges (user_id, badge_id) VALUES (?, ?)`
	_, err := r.db.ExecContext(ctx, query, userID, badgeID)
	if err != nil {
		return errors.InternalServerError("Failed to create user badge")
	}

	return nil
}