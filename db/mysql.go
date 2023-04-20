package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

//-------------------------连接数据库-------------------------------------
var DB *gorm.DB

func InitDB() *gorm.DB {

	dsn := "root:root@tcp(10.117.54.37:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"
	//配置日志对象
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
		},
	)
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	//迁移模型类，将模型类转换成SQL
	DB.AutoMigrate(
		&Teacher{},
		&Class{},
		&Student{},
		&Course{},
		&Admin{},
	)
	return DB

}
