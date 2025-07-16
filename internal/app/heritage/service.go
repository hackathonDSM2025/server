package heritage

import (
	"context"

	"hackathon-dsm-server/internal/pkg/utils/errors"
)

type Service interface {
	SearchHeritage(ctx context.Context, keyword string) (*SearchData, error)
	CreateVisit(ctx context.Context, userID, heritageID int, qrCode string) (*CreateVisitData, error)
	CreateOnlyReview(ctx context.Context, userID, heritageID int, rating int, reviewText string) error
	UpdateOnlyReview(ctx context.Context, userID, heritageID int, rating int, reviewText string) error
	GetMyReview(ctx context.Context, userID, heritageID int) (*ReviewData, error)
}

type HeritageService struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &HeritageService{repo: repo}
}

func (s *HeritageService) SearchHeritage(ctx context.Context, keyword string) (*SearchData, error) {
	heritage, err := s.repo.SearchHeritage(ctx, keyword)
	if err != nil {
		return nil, err
	}

	return &SearchData{
		Name:        heritage.Name,
		Latitude:    heritage.Latitude,
		Longitude:   heritage.Longitude,
		ImageURL:    heritage.ImageURL,
		Description: heritage.Description,
	}, nil
}

func (s *HeritageService) CreateVisit(ctx context.Context, userID, heritageID int, qrCode string) (*CreateVisitData, error) {
	heritage, err := s.repo.GetHeritageByQRCode(ctx, qrCode)
	if err != nil {
		return nil, err
	}

	if heritage.HeritageID != heritageID {
		return nil, errors.BadRequest("QR code does not match heritage ID")
	}

	visit, err := s.repo.GetUserVisit(ctx, userID, heritage.HeritageID)
	if err != nil {
		return nil, err
	}

	isFirstVisit := visit == nil

	if isFirstVisit {
		if err := s.repo.CreateUserVisit(ctx, userID, heritage.HeritageID); err != nil {
			return nil, errors.InternalServerError("Failed to record visit")
		}

		// Award "first_visit" badge if user doesn't have it yet
		badge, err := s.repo.GetBadgeByCondition(ctx, "first_visit")
		if err == nil && badge != nil {
			exists, err := s.repo.CheckUserBadgeExists(ctx, userID, badge.BadgeID)
			if err == nil && !exists {
				s.repo.CreateUserBadge(ctx, userID, badge.BadgeID)
			}
		}
	}

	return &CreateVisitData{
		HeritageID:   heritage.HeritageID,
		Name:         heritage.Name,
		ImageURL:     heritage.ImageURL,
		Description:  heritage.Description,
		Story:        heritage.StoryText,
		IsFirstVisit: isFirstVisit,
	}, nil
}

func (s *HeritageService) CreateOnlyReview(ctx context.Context, userID, heritageID int, rating int, reviewText string) error {
	existingReview, err := s.repo.GetHeritageReview(ctx, userID, heritageID)
	if err != nil {
		return err
	}

	if existingReview != nil {
		return errors.BadRequest("Review already exists for this heritage")
	}

	return s.repo.CreateHeritageReview(ctx, userID, heritageID, rating, reviewText)
}

func (s *HeritageService) UpdateOnlyReview(ctx context.Context, userID, heritageID int, rating int, reviewText string) error {
	existingReview, err := s.repo.GetHeritageReview(ctx, userID, heritageID)
	if err != nil {
		return err
	}

	if existingReview == nil {
		return errors.NotFound("Review not found for this heritage")
	}

	return s.repo.UpdateHeritageReview(ctx, userID, heritageID, rating, reviewText)
}

func (s *HeritageService) GetMyReview(ctx context.Context, userID, heritageID int) (*ReviewData, error) {
	review, err := s.repo.GetHeritageReview(ctx, userID, heritageID)
	if err != nil {
		return nil, err
	}

	if review == nil {
		return nil, nil
	}

	return &ReviewData{
		ReviewID:   review.ReviewID,
		Rating:     review.Rating,
		ReviewText: review.ReviewText,
		CreatedAt:  review.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  review.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}