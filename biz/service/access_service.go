package service

import (
	"log"
	"zhenghui-backend/biz/consts"
	"zhenghui-backend/biz/dal/db_dal"
	"zhenghui-backend/biz/model/dao"
	"zhenghui-backend/biz/model/dto"
)

func InsertAccessData(ip string, accessPage consts.AccessPage) error {
	accessDataRequest, err := dto.ParseIP(ip)
	if err != nil {
		log.Println("utils.ParseIP error, err = ", err)
		return err
	}
	accessData := dao.AccessData{
		IP:       ip,
		ISP:      accessDataRequest.ISP,
		Page:     accessPage.ToString(),
		Province: accessDataRequest.Province,
		City:     accessDataRequest.City,
	}
	return db_dal.InsertAccessData(accessData)
}

func SearchAccessHistory() ([]dto.AccessHistory, error) {
	return db_dal.SearchAccessHistory()
}

func SearchAccessStatistics() (dto.AccessStatistics, error) {
	return db_dal.SearchAccessStatistics()
}

func SearchAccessStatisticsMid() (dto.AccessStatisticsMid, error) {
	var response dto.AccessStatisticsMid
	var err error

	dailyAccessList, err := db_dal.SearchAccessNumberPerDay()
	if err != nil {
		log.Println("SearchAccessNumberPerDay error, err = ", err)
		return response, err
	}
	var count int64
	for i := 0; i < len(dailyAccessList); i++ {
		dailyAccessList[i].Number += count
		count = dailyAccessList[i].Number
	}
	response.DailyAccessList = dailyAccessList
	response.PageAccessList, err = db_dal.SearchPageAccessNumber()
	if err != nil {
		log.Println("SearchPageAccessNumber error, err = ", err)
	}
	return response, err
}
