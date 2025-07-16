package badge

type BadgeListResponse struct {
	Success bool         `json:"success"`
	Data    BadgeListData `json:"data"`
}

type BadgeListData struct {
	BadgeCount int        `json:"badgeCount"`
	Badges     []BadgeDTO `json:"badges"`
}

type BadgeDTO struct {
	Name         string `json:"name"`
	ImageURL     string `json:"imageUrl"`
	EarnedAt     string `json:"earnedAt"`
	HeritageName string `json:"heritageName"`
}