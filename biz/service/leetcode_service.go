package service

import (
	"zhenghui-backend/biz/dal/db_dal"
	"zhenghui-backend/biz/model/dao"
)

func SearchLeetcode() ([]dao.Leetcode, error) {
	return db_dal.SearchLeetcodeList()
}
