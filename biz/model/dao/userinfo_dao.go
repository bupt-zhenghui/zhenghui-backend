package dao

import (
	"time"
	"zhenghui-backend/biz/consts"
)

type IInfo interface {
}

type CommonBase struct {
	Id         int64     `gorm:"column:id" json:"id"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

type UserInfo struct {
	CommonBase
	Name       string            `gorm:"column:name" json:"name"`
	Gender     consts.GenderType `gorm:"column:gender" json:"gender"`
	Age        int8              `gorm:"column:age" json:"age"`
	TelNumber  string            `gorm:"column:tel-number" json:"tel-number"`
	Email      string            `gorm:"column:email" json:"email"`
	Company    string            `gorm:"column:company" json:"company"`
	University string            `gorm:"column:university" json:"university"`
}

type ConsumerInfo struct {
	CommonBase
}
