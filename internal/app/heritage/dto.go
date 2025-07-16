package heritage

type SearchRequest struct {
	Keyword string `form:"keyword" binding:"required"`
}

type CreateVisitRequest struct {
	QRCode string `json:"qrCode" binding:"required"`
}

type SearchResponse struct {
	Success bool       `json:"success"`
	Data    SearchData `json:"data"`
}

type SearchData struct {
	Name        string  `json:"name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	ImageURL    string  `json:"imageUrl"`
	Description string  `json:"description"`
}

type CreateVisitResponse struct {
	Success bool            `json:"success"`
	Data    CreateVisitData `json:"data"`
}

type CreateVisitData struct {
	HeritageID   int    `json:"heritageId"`
	Name         string `json:"name"`
	ImageURL     string `json:"imageUrl"`
	Description  string `json:"description"`
	Story        string `json:"story"`
	IsFirstVisit bool   `json:"isFirstVisit"`
}

type ReviewRequest struct {
	Rating     int    `json:"rating" binding:"required,min=1,max=5"`
	ReviewText string `json:"reviewText" binding:"required"`
}

type GetMyReviewResponse struct {
	Success bool        `json:"success"`
	Data    *ReviewData `json:"data"`
}

type ReviewData struct {
	ReviewID   int    `json:"reviewId"`
	Rating     int    `json:"rating"`
	ReviewText string `json:"reviewText"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

type CreateReviewResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type UpdateReviewResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}