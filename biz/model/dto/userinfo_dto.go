package dto

type UserInfoResponse struct {
	Id         int64  `gorm:"column:id" json:"id"`
	Name       string `gorm:"column:name" json:"name"`
	TelNumber  string `gorm:"column:tel-number" json:"tel-number"`
	Email      string `gorm:"column:email" json:"email"`
	Company    string `gorm:"column:company" json:"company"`
	University string `gorm:"column:university" json:"university"`
}
