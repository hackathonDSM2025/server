package badge

type AllBadgesResponse struct {
	Success bool          `json:"success"`
	Data    AllBadgesData `json:"data"`
}

type AllBadgesData struct {
	BadgeCount int          `json:"badgeCount"`
	Badges     []BadgeBasic `json:"badges"`
}

type BadgeBasic struct {
	BadgeID      int    `json:"badgeId"`
	Name         string `json:"name"`
	ImageURL     string `json:"imageUrl"`
	HeritageName string `json:"heritageName"`
	Description  string `json:"description"`
}

type BadgeDetailResponse struct {
	Success bool            `json:"success"`
	Data    BadgeDetailData `json:"data"`
}

type BadgeDetailData struct {
	BadgeID      int    `json:"badgeId"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	ImageURL     string `json:"imageUrl"`
	HeritageName string `json:"heritageName"`
	CreatedAt    string `json:"createdAt"`
}