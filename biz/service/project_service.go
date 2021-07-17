package service

import (
	"log"
	"zhenghui-backend/biz/dal/db_dal"
	"zhenghui-backend/biz/model/dao"
)

func SearchProjectList() ([]dao.Project, error) {
	response, err := db_dal.SearchProjectList()
	if err != nil {
		log.Println("SearchProjectList error, err = ", err)
	}
	return response, err
}

func UpdateProject(project dao.Project) error {
	return db_dal.UpdateProject(project)
}

func InsertProject(project dao.Project) error {
	return db_dal.InsertProject(project)
}

func DeleteProject(id int64) error {
	return db_dal.DeleteProject(id)
}
