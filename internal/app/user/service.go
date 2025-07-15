package user

import (
	"context"
)

type Service interface {
	GetUserProfile(ctx context.Context, userID int) (*UserProfileData, error)
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

	visitedCount, err := s.repo.GetVisitedCount(ctx, userID)
	if err != nil {
		return nil, err
	}

	badgeCount, err := s.repo.GetBadgeCount(ctx, userID)
	if err != nil {
		return nil, err
	}

	visitedPlaces, err := s.repo.GetVisitedPlaces(ctx, userID)
	if err != nil {
		return nil, err
	}

	badges, err := s.repo.GetUserBadges(ctx, userID)
	if err != nil {
		return nil, err
	}

	var visitedPlacesDTO []VisitedPlaceDTO
	for _, place := range visitedPlaces {
		visitedPlacesDTO = append(visitedPlacesDTO, VisitedPlaceDTO{
			Name:      place.Name,
			VisitedAt: place.VisitedAt.Format("2006-01-02"),
			Completed: place.Completed,
		})
	}

	var badgesDTO []BadgeDTO
	for _, badge := range badges {
		badgesDTO = append(badgesDTO, BadgeDTO{
			Name:     badge.Name,
			ImageURL: badge.ImageURL,
			EarnedAt: badge.EarnedAt.Format("2006-01-02"),
		})
	}

	return &UserProfileData{
		Nickname:      user.Nickname,
		VisitedCount:  visitedCount,
		BadgeCount:    badgeCount,
		VisitedPlaces: visitedPlacesDTO,
		Badges:        badgesDTO,
	}, nil
}