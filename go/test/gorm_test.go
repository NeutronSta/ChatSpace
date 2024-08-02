package test

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

/*
func TestGormTest(t *testing.T) {
	dsn := "root:168168hyS@tcp(114.116.204.137:3306)/chatspace?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	err = db.AutoMigrate(&modules.Message{})
	if err != nil {
		t.Fatal(err)
	}

	/*var users []modules.User
	db.Find(&users)

	for v := range users {
		fmt.Printf("User: %v \n", v)
	}

*/
