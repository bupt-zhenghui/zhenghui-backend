package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"zhenghui-backend/biz/consts"
)

func UploadFile(c *gin.Context, header *multipart.FileHeader, fileType consts.FileType) error {
	filePath := fileType.GetPath() + header.Filename
	err := c.SaveUploadedFile(header, filePath)
	if err != nil {
		log.Println("SaveUploadedFile error, err = ", err)
	}
	return err
}
