package dto

type MonthlyReportResponse struct {
	Id         int      `json:"key"`
	Title      string   `json:"title"`
	Tags       []string `json:"tags"`
	CreateTime string   `json:"create_time"`
	URL        string   `json:"url"`
}
