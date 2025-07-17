package user

import (
	"context"
)

type Service interface {
	GetUserProfile(ctx context.Context, userID int) (*UserProfileData, error)
	GetMyVisits(ctx context.Context, userID int) (*MyVisitsData, error)
	GetMyVisitsCount(ctx context.Context, userID int) (int, error)
	GetMyReviews(ctx context.Context, userID int) (*MyReviewsData, error)
	GetMyReviewsCount(ctx context.Context, userID int) (int, error)
	GetMyBadges(ctx context.Context, userID int) (*MyBadgesData, error)
	GetMyBadgesCount(ctx context.Context, userID int) (int, error)
}

type UserService struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserProfile(ctx context.Context, userID int) (*UserProfileData, error) {
	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &UserProfileData{
		UserID:    user.UserID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *UserService) GetMyVisits(ctx context.Context, userID int) (*MyVisitsData, error) {
	visitCount, err := s.repo.GetVisitedCount(ctx, userID)
	if err != nil {
		return nil, err
	}

	visits, err := s.repo.GetVisitedPlaces(ctx, userID)
	if err != nil {
		return nil, err
	}

	var visitsDTO []VisitedPlaceDTO
	for _, visit := range visits {
		visitsDTO = append(visitsDTO, VisitedPlaceDTO{
			Name:      visit.Name,
			VisitedAt: visit.VisitedAt.Format("2006-01-02"),
			Completed: visit.Completed,
		})
	}

	return &MyVisitsData{
		VisitCount: visitCount,
		Visits:     visitsDTO,
	}, nil
}

func (s *UserService) GetMyVisitsCount(ctx context.Context, userID int) (int, error) {
	return s.repo.GetVisitedCount(ctx, userID)
}

func (s *UserService) GetMyReviews(ctx context.Context, userID int) (*MyReviewsData, error) {
	reviewCount, err := s.repo.GetUserReviewCount(ctx, userID)
	if err != nil {
		return nil, err
	}

	reviews, err := s.repo.GetUserReviews(ctx, userID)
	if err != nil {
		return nil, err
	}

	var reviewsDTO []ReviewDTO
	for _, review := range reviews {
		reviewsDTO = append(reviewsDTO, ReviewDTO{
			ReviewID:         review.ReviewID,
			HeritageID:       review.HeritageID,
			HeritageName:     review.HeritageName,
			HeritageImageURL: review.HeritageImageURL,
			Rating:           review.Rating,
			ReviewText:       review.ReviewText,
			CreatedAt:        review.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:        review.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &MyReviewsData{
		ReviewCount: reviewCount,
		Reviews:     reviewsDTO,
	}, nil
}

func (s *UserService) GetMyReviewsCount(ctx context.Context, userID int) (int, error) {
	return s.repo.GetUserReviewCount(ctx, userID)
}

func (s *UserService) GetMyBadges(ctx context.Context, userID int) (*MyBadgesData, error) {
	badgeCount, err := s.repo.GetUserBadgeCount(ctx, userID)
	if err != nil {
		return nil, err
	}

	badges, err := s.repo.GetUserBadges(ctx, userID)
	if err != nil {
		return nil, err
	}

	var badgesDTO []BadgeDTO
	for _, badge := range badges {
		badgesDTO = append(badgesDTO, BadgeDTO{
			Name:             badge.Name,
			ImageURL:         badge.ImageURL,
			EarnedAt:         badge.EarnedAt.Format("2006-01-02"),
			HeritageName:     badge.HeritageName,
			HeritageImageURL: badge.HeritageImageURL,
		})
	}

	return &MyBadgesData{
		BadgeCount: badgeCount,
		Badges:     badgesDTO,
	}, nil
}

func (s *UserService) GetMyBadgesCount(ctx context.Context, userID int) (int, error) {
	return s.repo.GetUserBadgeCount(ctx, userID)
}
