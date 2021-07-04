package dao

import "time"

type AccessData struct {
	Id         int64     `gorm:"column:id" json:"id"`
	IP         string    `gorm:"column:ip" json:"ip"`
	Page       string    `gorm:"column:page" json:"page"`
	ISP        string    `gorm:"column:ISP" json:"isp"`
	Province   string    `gorm:"column:province" json:"province"`
	City       string    `gorm:"column:city" json:"city"`
	CreateTime time.Time `gorm:"column:create_time;default:null" json:"create_time"`
}
