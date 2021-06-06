package service

import (
	"log"
	"zhenghui-backend/biz/dal/db_dal"
	"zhenghui-backend/biz/model/dto"
)

func SearchUserInfoById(userId int64) (dto.UserInfoResponse, error) {
	var response dto.UserInfoResponse
	user, err := db_dal.SearchUserInfo(userId)
	if err != nil {
		log.Println("db_dal.SearchUserInfo failed to execute, error: ", err)
		return response, err
	}
	response = dto.UserInfoResponse{
		Id:         user.Id,
		Name:       user.Name,
		TelNumber:  user.TelNumber,
		Email:      user.Email,
		Company:    user.Company,
		University: user.University,
	}
	return response, err
}
