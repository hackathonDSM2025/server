package user

type UserProfileResponse struct {
	Success bool            `json:"success"`
	Data    UserProfileData `json:"data"`
}

type UserProfileData struct {
	VisitedCount  int               `json:"visitedCount"`
	VisitedPlaces []VisitedPlaceDTO `json:"visitedPlaces"`
}

type VisitedPlaceDTO struct {
	Name      string `json:"name"`
	VisitedAt string `json:"visitedAt"`
	Completed bool   `json:"completed"`
}

