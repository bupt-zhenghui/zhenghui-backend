package service

import (
	"log"
	"strings"
	"zhenghui-backend/biz/model/dto"
	"zhenghui-backend/biz/utils"
)

func SearchMonthlyReport() ([]dto.MonthlyReportResponse, error) {
	var response []dto.MonthlyReportResponse
	fileList, err := utils.GetAllFile("./file/monthly_report")
	if err != nil {
		log.Fatal("utils.GetAllFile error, err = ", err)
		return response, err
	}
	for idx, file := range fileList {
		fileFull := strings.Split(file, "/")[len(strings.Split(file, "/"))-1]
		fileNameTags := strings.Split(fileFull, ".")[0]
		fileName, tags := strings.Split(fileNameTags, ";")[0], strings.Split(fileNameTags, ";")[1]
		report := dto.MonthlyReportResponse{
			Id:         idx,
			Title:      fileName,
			Tags:       strings.Split(tags, ","),
			CreateTime: utils.GetFileModTime(file),
			URL:        fileFull,
		}
		response = append(response, report)
	}
	return response, nil
}
