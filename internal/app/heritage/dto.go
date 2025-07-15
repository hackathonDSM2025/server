package heritage

type SearchRequest struct {
	Keyword string `form:"keyword" binding:"required"`
}

type ScanRequest struct {
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

type ScanResponse struct {
	Success bool     `json:"success"`
	Data    ScanData `json:"data"`
}

type ScanData struct {
	HeritageID   int    `json:"heritageId"`
	Name         string `json:"name"`
	ImageURL     string `json:"imageUrl"`
	Description  string `json:"description"`
	Story        string `json:"story"`
	IsFirstVisit bool   `json:"isFirstVisit"`
}