package service

import (
	"log"
	"time"
	"zhenghui-backend/biz/consts"
	"zhenghui-backend/biz/dal/db_dal"
	"zhenghui-backend/biz/model/dao"
	"zhenghui-backend/biz/model/dto"
)

func SearchLeetcode() ([]dao.Leetcode, error) {
	return db_dal.SearchLeetcodeList()
}

func SearchLeetcodeStatistics() (dto.LeetcodeStatisticsDTO, error) {
	var response dto.LeetcodeStatisticsDTO
	var err error
	response.Statistics, err = db_dal.SearchLeetcodeStatistics()
	if err != nil {
		log.Println("db_dal.SearchLeetcodeStatistics error, err = ", err)
		return response, err
	}
	response.HistoryAcData, err = db_dal.SearchAcNumberPerDay()
	if err != nil {
		log.Println("db_dal.SearchAcNumberPerDay error, err = ", err)
	}
	return response, err
}

func GenerateLeetcodeStatistics() (dao.LeetcodeStatistics, error) {
	var statistics dao.LeetcodeStatistics
	var err error
	statistics.Date = time.Now().Format(consts.ShortDate)
	statistics.TotalNumber, err = db_dal.QueryLeetcodeCount("question_id > 0")
	if err != nil {
		log.Println("QueryLeetcodeCount error, err = ", err)
		return statistics, err
	}
	statistics.EasyNumber, err = db_dal.QueryLeetcodeCount("difficulty = 1")
	if err != nil {
		log.Println("QueryLeetcodeCount error, err = ", err)
		return statistics, err
	}
	statistics.MediumNumber, err = db_dal.QueryLeetcodeCount("difficulty = 2")
	if err != nil {
		log.Println("QueryLeetcodeCount error, err = ", err)
		return statistics, err
	}
	statistics.HardNumber, err = db_dal.QueryLeetcodeCount("difficulty = 3")
	if err != nil {
		log.Println("QueryLeetcodeCount error, err = ", err)
		return statistics, err
	}
	statistics.AcNumber, err = db_dal.QueryLeetcodeCount("status = 'ac'")
	if err != nil {
		log.Println("QueryLeetcodeCount error, err = ", err)
		return statistics, err
	}
	statistics.AcEasyNumber, err = db_dal.QueryLeetcodeCount("status = 'ac' and difficulty = 1")
	if err != nil {
		log.Println("QueryLeetcodeCount error, err = ", err)
		return statistics, err
	}
	statistics.AcMediumNumber, err = db_dal.QueryLeetcodeCount("status = 'ac' and difficulty = 2")
	if err != nil {
		log.Println("QueryLeetcodeCount error, err = ", err)
		return statistics, err
	}
	statistics.AcHardNumber, err = db_dal.QueryLeetcodeCount("status = 'ac' and difficulty = 3")
	if err != nil {
		log.Println("QueryLeetcodeCount error, err = ", err)
		return statistics, err
	}
	statistics.NotAcNumber, err = db_dal.QueryLeetcodeCount("status = 'notac'")
	if err != nil {
		log.Println("QueryLeetcodeCount error, err = ", err)
		return statistics, err
	}
	statistics.UntouchedNumber, err = db_dal.QueryLeetcodeCount("status is null")
	if err != nil {
		log.Println("QueryLeetcodeCount error, err = ", err)
		return statistics, err
	}
	statistics.NewAcNumber, err = db_dal.CalculateNewAcNumber(statistics.AcNumber)
	if err != nil {
		log.Println("CalculateNewAcNumber error, err = ", err)
	}
	return statistics, err
}

func UpdateLeetcodeStatistics() error {
	statistics, err := GenerateLeetcodeStatistics()
	if err != nil {
		log.Println("GenerateLeetcodeStatistics failed, err = ", err)
		return err
	}
	curDate := time.Now().Format(consts.ShortDate)
	isNewDate, err := db_dal.CheckNewDate(curDate)
	if err != nil {
		log.Println("db_dal.CheckNewDate failed, err = ", err)
		return err
	}
	if isNewDate {
		err = db_dal.CreateLeetcodeStatistics(statistics)
		if err != nil {
			log.Println("db_dal.CreateLeetcodeStatistics failed, err = ", err)
		}
	} else {
		err = db_dal.UpdateLeetcodeStatistics(statistics)
		if err != nil {
			log.Println("db_dal.UpdateLeetcodeStatistics failed, err = ", err)
		}
	}
	return err
}
