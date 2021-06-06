package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"zhenghui-backend/biz/utils"
)

var DB *gorm.DB

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		log.Fatalf("Init DB error: %v dsn: %v", err, dsn)
	}
	return db, err
}

func ConnectDB() {
	dsn, err := utils.GenerateMySQLDSN()
	if err != nil {
		log.Fatal("Invalid MySQL Configuration, please check file: ./conf/config.yaml")
	}
	DB, _ = InitDB(dsn)
}
