package dao

type Leetcode struct {
	QuestionId int64  `gorm:"column:question_id" json:"key"`
	SpId       string `gorm:"column:sp_id" json:"sp_id"`
	Title      string `gorm:"column:title" json:"title"`
	Status     string `gorm:"column:status" json:"status"`
	Difficulty int64  `gorm:"column:difficulty" json:"difficulty"`
	URL        string `gorm:"column:url" json:"url"`
}
