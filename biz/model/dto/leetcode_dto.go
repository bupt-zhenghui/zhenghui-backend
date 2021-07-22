package dto

import "zhenghui-backend/biz/model/dao"

type AcDataDTO struct {
	Date     string `gorm:"column:date" json:"date"`
	AcNumber int64  `gorm:"column:new_ac_number" json:"ac_number"`
}

type LeetcodeStatisticsDTO struct {
	Statistics    dao.LeetcodeStatistics `json:"statistics"`
	HistoryAcData []AcDataDTO            `json:"history_ac_data"`
}
