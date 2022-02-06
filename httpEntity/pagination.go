package httpEntity

type Pagination struct {
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	Sort      string `json:"sort"`
	TotalData int64  `json:"total_data"`
}
