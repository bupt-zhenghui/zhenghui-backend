package dto

type MonthlyReportResponse struct {
	Id         int      `json:"key"`
	Title      string   `json:"title"`
	Tags       []string `json:"tags"`
	UpdateTime string   `json:"update_time"`
	URL        string   `json:"url"`
}
