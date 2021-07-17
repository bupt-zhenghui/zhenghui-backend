package dto

import (
	"time"
	"zhenghui-backend/biz/consts"
)

type ProjectDTO struct {
	Id         int64              `gorm:"column:id" json:"id"`
	Type       consts.ProjectType `gorm:"column:type" json:"type"`
	Name       string             `gorm:"column:name" json:"name"`
	Tag        []string           `gorm:"column:tag" json:"tag"`
	MdURL      string             `gorm:"column:md_url" json:"md_url"`
	GithubURL  string             `gorm:"column:github_url" json:"github_url"`
	ProjectURL string             `gorm:"column:project_url" json:"project_url"`
	CreateTime time.Time          `gorm:"column:create_time;default:null" json:"create_time"`
	UpdateTime time.Time          `gorm:"column:update_time;default:null" json:"update_time"`
}
