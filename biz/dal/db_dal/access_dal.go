package db_dal

import (
	"fmt"
	"zhenghui-backend/biz/model/dao"
	"zhenghui-backend/biz/model/dto"
)

func InsertAccessData(accessData dao.AccessData) error {
	return dao.DB.Create(&accessData).Error
}

// 所有ip的最新访问记录表
func SearchAccessHistory() ([]dto.AccessHistory, error) {
	var response []dto.AccessHistory
	// select ip, count(ip) as count, max(create_time) as last_access_time, max(ISP) as isp from access_data group by ip order by last_access_time desc
	err := dao.DB.Table("access_data").Select("ip, count(ip) as count, max(create_time) as last_access_time, max(ISP) as isp").
		Order("last_access_time desc").Group("ip").Find(&response).Error
	return response, err
}

// 访问记录统计数据
func SearchAccessStatistics() (dto.AccessStatistics, error) {
	var response dto.AccessStatistics
	dbPrefix := dao.DB.Table("access_data")
	// 1. 查询浏览量总数
	if err := dbPrefix.Select("count(*) as total_number").Find(&response.TotalNumber).Error; err != nil {
		return response, err
	}
	fmt.Println(response)
	// 2. 查询本周浏览量
	if err := dbPrefix.Select("count(*) as recent_week_number").Where("DATE_FORMAT(create_time, '%Y%m%u') = DATE_FORMAT(now(),'%Y%m%u')").
		Find(&response.RecentWeekNumber).Error; err != nil {
		return response, err
	}
	fmt.Println(response)
	// 3. 查询本月浏览量
	if err := dbPrefix.Select("count(*) as recent_30_day_number").Where("DATE_FORMAT(create_time, '%Y%m') = DATE_FORMAT(NOW(),'%Y%m')").
		Find(&response.RecentMonthNumber).Error; err != nil {
		return response, err
	}
	fmt.Println(response)
	// 4. 查询总浏览人数（相同ip只记1次）
	if err := dbPrefix.Select("count(DISTINCT ip)").Find(&response.TotalIPNumber).Error; err != nil {
		return response, err
	}
	fmt.Println(response)
	err := dbPrefix.Select("count(DISTINCT ip)").Where("DATE_FORMAT(create_time, '%Y%m') = DATE_FORMAT(NOW(),'%Y%m')").
		Find(&response.RecentMonthIPNumber).Error
	return response, err
}

// select substring(create_time, 1, 10), count(0) from access_data group by substring(create_time, 1, 10)
