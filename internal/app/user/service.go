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
	visitedCount, err := s.repo.GetVisitedCount(ctx, userID)
	if err != nil {
		return nil, err
	}

	visitedPlaces, err := s.repo.GetVisitedPlaces(ctx, userID)
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

	return &UserProfileData{
		VisitedCount:  visitedCount,
		VisitedPlaces: visitedPlacesDTO,
	}, nil
}
