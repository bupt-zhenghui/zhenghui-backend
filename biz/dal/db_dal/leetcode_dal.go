package db_dal

import (
	"zhenghui-backend/biz/model/dao"
	"zhenghui-backend/biz/model/dto"
)

func SearchLeetcodeList() ([]dao.Leetcode, error) {
	var codeList []dao.Leetcode
	err := dao.DB.Order("question_id").Find(&codeList).Error
	return codeList, err
}

func SearchLeetcodeStatistics() (dao.LeetcodeStatistics, error) {
	var response dao.LeetcodeStatistics
	err := dao.DB.Table("leetcode_statistics").Order("date desc").Limit(1).Find(&response).Error
	return response, err
}

func QueryLeetcodeCount(condition string) (int64, error) {
	var count int64
	err := dao.DB.Table("leetcode").Select("count(*)").Where(condition).Find(&count).Error
	return count, err
}

func CalculateNewAcNumber(curNum int64) (int64, error) {
	var lastNumber int64
	//lastDate := time.Now().AddDate(0, 0, -1).Format(consts.ShortDate)
	err := dao.DB.Table("leetcode_statistics").Select("ac_number").Order("date desc").Find(&lastNumber).Limit(1).Error
	return curNum - lastNumber, err
}

func CreateLeetcodeStatistics(statistics dao.LeetcodeStatistics) error {
	return dao.DB.Create(&statistics).Error
}

func UpdateLeetcodeStatistics(statistics dao.LeetcodeStatistics) error {
	return dao.DB.Where("date = ?", statistics.Date).Updates(statistics).Error
}

func CheckNewDate(date string) (bool, error) {
	var count int64
	response := false
	err := dao.DB.Table("leetcode_statistics").Select("count(*)").Where("date = ?", date).Find(&count).Error
	if count == 0 {
		response = true
	}
	return response, err
}

func SearchAcNumberPerDay() ([]dto.AcDataDTO, error) {
	var response []dto.AcDataDTO
	err := dao.DB.Table("leetcode_statistics").Where("create_time >= DATE_SUB(CURDATE(), INTERVAL 30 DAY)").
		Select("date, new_ac_number").Find(&response).Error
	for idx, item := range response {
		response[idx].Date = item.Date[4:6] + "-" + item.Date[6:]
	}
	return response, err
}
