package heritage

import (
	"context"

	"hackathon-dsm-server/internal/pkg/utils/errors"
)

type Service interface {
	SearchHeritage(ctx context.Context, keyword string) (*SearchData, error)
	ScanQRCode(ctx context.Context, userID int, qrCode string) (*ScanData, error)
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

func (s *HeritageService) ScanQRCode(ctx context.Context, userID int, qrCode string) (*ScanData, error) {
	heritage, err := s.repo.GetHeritageByQRCode(ctx, qrCode)
	if err != nil {
		return nil, err
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

	return &ScanData{
		HeritageID:   heritage.HeritageID,
		Name:         heritage.Name,
		ImageURL:     heritage.ImageURL,
		Description:  heritage.Description,
		Story:        heritage.StoryText,
		IsFirstVisit: isFirstVisit,
	}, nil
}