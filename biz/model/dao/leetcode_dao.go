package dao

import "time"

type Leetcode struct {
	QuestionId int64  `gorm:"column:question_id" json:"key"`
	SpId       string `gorm:"column:sp_id" json:"sp_id"`
	Title      string `gorm:"column:title" json:"title"`
	Status     string `gorm:"column:status" json:"status"`
	Difficulty int64  `gorm:"column:difficulty" json:"difficulty"`
	URL        string `gorm:"column:url" json:"url"`
}

type LeetcodeStatistics struct {
	Id              int64     `gorm:"column:id" json:"id"`
	Date            string    `gorm:"column:date" json:"date"`
	TotalNumber     int64     `gorm:"column:total_number" json:"total_number"`
	EasyNumber      int64     `gorm:"column:easy_number" json:"easy_number"`
	MediumNumber    int64     `gorm:"column:medium_number" json:"medium_number"`
	HardNumber      int64     `gorm:"column:hard_number" json:"hard_number"`
	AcNumber        int64     `gorm:"column:ac_number" json:"ac_number"`
	AcEasyNumber    int64     `gorm:"column:ac_easy_number" json:"ac_easy_number"`
	AcMediumNumber  int64     `gorm:"column:ac_medium_number" json:"ac_medium_number"`
	AcHardNumber    int64     `gorm:"column:ac_hard_number" json:"ac_hard_number"`
	NotAcNumber     int64     `gorm:"column:not_ac_number" json:"not_ac_number"`
	UntouchedNumber int64     `gorm:"column:untouched_number" json:"untouched_number"`
	NewAcNumber     int64     `gorm:"column:new_ac_number" json:"new_ac_number"`
	CreateTime      time.Time `gorm:"column:create_time;default:null" json:"create_time"`
	UpdateTime      time.Time `gorm:"column:update_time;default:null" json:"update_time"`
}
