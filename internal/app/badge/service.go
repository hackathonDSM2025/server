package badge

import (
	"context"
)

type Service interface {
	GetUserBadges(ctx context.Context, userID int) (*BadgeListData, error)
}

type BadgeService struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &BadgeService{repo: repo}
}

func (s *BadgeService) GetUserBadges(ctx context.Context, userID int) (*BadgeListData, error) {
	badgeCount, err := s.repo.GetBadgeCount(ctx, userID)
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
			Name:         badge.Name,
			ImageURL:     badge.ImageURL,
			EarnedAt:     badge.EarnedAt.Format("2006-01-02"),
			HeritageName: badge.HeritageName,
		})
	}

	return &BadgeListData{
		BadgeCount: badgeCount,
		Badges:     badgesDTO,
	}, nil
}