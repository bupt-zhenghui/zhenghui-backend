package db_dal

import "zhenghui-backend/biz/model/dao"

func SearchProjectList() ([]dao.Project, error) {
	var response []dao.Project
	err := dao.DB.Find(&response).Error
	return response, err
}

func InsertProject(project dao.Project) error {
	return dao.DB.Create(&project).Error
}

func UpdateProject(project dao.Project) error {
	return dao.DB.Updates(&project).Where("id = ?", project.Id).Error
}
