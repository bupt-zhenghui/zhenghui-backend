package db_dal

import "zhenghui-backend/biz/model/dao"

func SearchLeetcodeList() ([]dao.Leetcode, error) {
	var codeList []dao.Leetcode
	err := dao.DB.Order("question_id").Find(&codeList).Error
	return codeList, err
}
