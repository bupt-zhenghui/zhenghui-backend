package db_dal

import (
	"zhenghui-backend/biz/model/dao"
)

func SearchUserInfo(userId int64) (dao.UserInfo, error) {
	var user dao.UserInfo
	err := dao.DB.Where("id = ?", userId).First(&user).Error
	return user, err
}
