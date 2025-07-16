package badge

import (
	"context"
	"database/sql"

	"hackathon-dsm-server/internal/pkg/utils/errors"
)

type Repository interface {
	GetAllBadges(ctx context.Context) ([]BadgeWithHeritage, error)
	GetAllBadgesCount(ctx context.Context) (int, error)
	GetBadgeByID(ctx context.Context, badgeID int) (*BadgeWithHeritage, error)
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

func (r *MySQLRepository) GetAllBadges(ctx context.Context) ([]BadgeWithHeritage, error) {
	query := `SELECT b.badge_id, b.heritage_id, b.name, b.description, b.image_url, h.name as heritage_name, b.created_at
			  FROM badges b 
			  JOIN heritage h ON b.heritage_id = h.heritage_id
			  ORDER BY b.created_at DESC`
	
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.InternalServerError("Database error")
	}
	defer rows.Close()

	var badges []BadgeWithHeritage
	for rows.Next() {
		var badge BadgeWithHeritage
		err := rows.Scan(&badge.BadgeID, &badge.HeritageID, &badge.Name, &badge.Description, 
			&badge.ImageURL, &badge.HeritageName, &badge.CreatedAt)
		if err != nil {
			return nil, errors.InternalServerError("Database error")
		}
		badges = append(badges, badge)
	}

	return badges, nil
}

func (r *MySQLRepository) GetAllBadgesCount(ctx context.Context) (int, error) {
	query := `SELECT COUNT(*) FROM badges`
	row := r.db.QueryRowContext(ctx, query)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, errors.InternalServerError("Database error")
	}

	return count, nil
}

func (r *MySQLRepository) GetBadgeByID(ctx context.Context, badgeID int) (*BadgeWithHeritage, error) {
	query := `SELECT b.badge_id, b.heritage_id, b.name, b.description, b.image_url, h.name as heritage_name, b.created_at
			  FROM badges b 
			  JOIN heritage h ON b.heritage_id = h.heritage_id
			  WHERE b.badge_id = ?`
	
	row := r.db.QueryRowContext(ctx, query, badgeID)

	badge := &BadgeWithHeritage{}
	err := row.Scan(&badge.BadgeID, &badge.HeritageID, &badge.Name, &badge.Description, 
		&badge.ImageURL, &badge.HeritageName, &badge.CreatedAt)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NotFound("Badge not found")
		}
		return nil, errors.InternalServerError("Database error")
	}

	return badge, nil
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