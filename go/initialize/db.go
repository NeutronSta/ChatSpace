package initialize

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB = Init()

func Init() *gorm.DB {
	//dsn := "root:168168hyS@tcp(127.0.0.1:3306)/chatspace?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:168168hyS@tcp(114.116.204.137:3306)/chatspace?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm Init Error: ", err)
	}
	return db
}
