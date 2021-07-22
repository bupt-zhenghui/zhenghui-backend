package dto

import (
	"encoding/json"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"net/http"
	"time"
	"zhenghui-backend/biz/consts"
)

type Location struct {
	IP       string `json:"ip"`
	Province string `json:"pro"`
	City     string `json:"city"`
	ISP      string `json:"addr"`
}

type AccessHistory struct {
	IP             string    `gorm:"column:ip" json:"ip"`
	Count          int64     `gorm:"column:count" json:"count"`
	ISP            string    `gorm:"column:isp" json:"isp"`
	LastAccessTime time.Time `gorm:"column:last_access_time" json:"last_access_time"`
}

type AccessStatistics struct {
	TotalNumber         int64 `gorm:"total_number" json:"total_number"`
	RecentWeekNumber    int64 `gorm:"recent_week_number" json:"recent_week_number"`
	RecentMonthNumber   int64 `gorm:"recent_month_number" json:"recent_month_number"`
	TotalIPNumber       int64 `gorm:"total_ip_number" json:"total_ip_number"`
	RecentMonthIPNumber int64 `gorm:"recent_month_ip_number" json:"recent_month_ip_number"`
}

type AccessStatisticsMid struct {
	DailyAccessList []DailyAccess `json:"daily_access_list"`
	PageAccessList  []PageAccess  `json:"page_access_list"`
}

type DailyAccess struct {
	Date   string `json:"date"`
	Number int64  `json:"number"`
}

type PageAccess struct {
	Page   string `json:"page"`
	Number int64  `json:"number"`
}

// fetch client location and ISP based on client ip address
func ParseIP(ip string) (Location, error) {
	var response Location
	resp, err := http.Get(consts.IP2LocationSite + ip)
	if err != nil {
		return response, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	bodyStr := mahonia.NewDecoder("gbk").ConvertString(string(body))
	err = json.Unmarshal([]byte(string(bodyStr)), &response) // 将string 格式转成json格式
	return response, err
}
