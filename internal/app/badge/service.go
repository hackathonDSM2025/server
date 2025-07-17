package badge

import (
	"context"
)

type Service interface {
	GetAllBadges(ctx context.Context) (*AllBadgesData, error)
	GetBadgeByID(ctx context.Context, badgeID int) (*BadgeDetailData, error)
}

type BadgeService struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &BadgeService{repo: repo}
}

func (s *BadgeService) GetAllBadges(ctx context.Context) (*AllBadgesData, error) {
	badgeCount, err := s.repo.GetAllBadgesCount(ctx)
	if err != nil {
		return nil, err
	}

	badges, err := s.repo.GetAllBadges(ctx)
	if err != nil {
		return nil, err
	}

	var badgesDTO []BadgeBasic
	for _, badge := range badges {
		badgesDTO = append(badgesDTO, BadgeBasic{
			BadgeID:           badge.BadgeID,
			Name:              badge.Name,
			ImageURL:          badge.ImageURL,
			HeritageName:      badge.HeritageName,
			HeritageImageURL:  badge.HeritageImageURL,
			Description:       badge.Description,
		})
	}

	return &AllBadgesData{
		BadgeCount: badgeCount,
		Badges:     badgesDTO,
	}, nil
}

func (s *BadgeService) GetBadgeByID(ctx context.Context, badgeID int) (*BadgeDetailData, error) {
	badge, err := s.repo.GetBadgeByID(ctx, badgeID)
	if err != nil {
		return nil, err
	}

	return &BadgeDetailData{
		BadgeID:           badge.BadgeID,
		Name:              badge.Name,
		Description:       badge.Description,
		ImageURL:          badge.ImageURL,
		HeritageName:      badge.HeritageName,
		HeritageImageURL:  badge.HeritageImageURL,
		CreatedAt:         badge.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}