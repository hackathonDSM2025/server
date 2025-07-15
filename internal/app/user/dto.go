package user

type UserProfileResponse struct {
	Success bool            `json:"success"`
	Data    UserProfileData `json:"data"`
}

type UserProfileData struct {
	Nickname      string            `json:"nickname"`
	VisitedCount  int               `json:"visitedCount"`
	BadgeCount    int               `json:"badgeCount"`
	VisitedPlaces []VisitedPlaceDTO `json:"visitedPlaces"`
	Badges        []BadgeDTO        `json:"badges"`
}

type VisitedPlaceDTO struct {
	Name      string `json:"name"`
	VisitedAt string `json:"visitedAt"`
	Completed bool   `json:"completed"`
}

type BadgeDTO struct {
	Name     string `json:"name"`
	ImageURL string `json:"imageUrl"`
	EarnedAt string `json:"earnedAt"`
}