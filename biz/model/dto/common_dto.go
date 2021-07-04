package dto

import "time"

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