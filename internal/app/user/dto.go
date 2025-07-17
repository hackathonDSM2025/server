package user

type UserProfileResponse struct {
	Success bool            `json:"success"`
	Data    UserProfileData `json:"data"`
}

type UserProfileData struct {
	UserID    int    `json:"userId"`
	Username  string `json:"username"`
	CreatedAt string `json:"createdAt"`
}

type VisitedPlaceDTO struct {
	Name      string `json:"name"`
	VisitedAt string `json:"visitedAt"`
	Completed bool   `json:"completed"`
}

type MyVisitsResponse struct {
	Success bool         `json:"success"`
	Data    MyVisitsData `json:"data"`
}

type MyVisitsData struct {
	VisitCount int               `json:"visitCount"`
	Visits     []VisitedPlaceDTO `json:"visits"`
}

type MyVisitsCountResponse struct {
	Success bool              `json:"success"`
	Data    MyVisitsCountData `json:"data"`
}

type MyVisitsCountData struct {
	VisitCount int `json:"visitCount"`
}

type MyReviewsResponse struct {
	Success bool          `json:"success"`
	Data    MyReviewsData `json:"data"`
}

type MyReviewsData struct {
	ReviewCount int         `json:"reviewCount"`
	Reviews     []ReviewDTO `json:"reviews"`
}

type MyReviewsCountResponse struct {
	Success bool               `json:"success"`
	Data    MyReviewsCountData `json:"data"`
}

type MyReviewsCountData struct {
	ReviewCount int `json:"reviewCount"`
}

type ReviewDTO struct {
	ReviewID         int    `json:"reviewId"`
	HeritageID       int    `json:"heritageId"`
	HeritageName     string `json:"heritageName"`
	HeritageImageURL string `json:"heritageImageUrl"`
	Rating           int    `json:"rating"`
	ReviewText       string `json:"reviewText"`
	CreatedAt        string `json:"createdAt"`
	UpdatedAt        string `json:"updatedAt"`
}

type MyBadgesResponse struct {
	Success bool         `json:"success"`
	Data    MyBadgesData `json:"data"`
}

type MyBadgesData struct {
	BadgeCount int       `json:"badgeCount"`
	Badges     []BadgeDTO `json:"badges"`
}

type MyBadgesCountResponse struct {
	Success bool              `json:"success"`
	Data    MyBadgesCountData `json:"data"`
}

type MyBadgesCountData struct {
	BadgeCount int `json:"badgeCount"`
}

type BadgeDTO struct {
	Name             string `json:"name"`
	ImageURL         string `json:"imageUrl"`
	EarnedAt         string `json:"earnedAt"`
	HeritageName     string `json:"heritageName"`
	HeritageImageURL string `json:"heritageImageUrl"`
}

