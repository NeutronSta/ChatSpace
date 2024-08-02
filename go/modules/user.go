package modules

import (
	"ChatSpace/initialize"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Name     string `gorm:"column:name;type:varchar(100);index" json:"name" form:"name"`
	Password []byte `gorm:"column:password;type:BLOB" json:"password" form:"password"`
	State    bool   `gorm:"column:state;type:tinyint(1)" json:"connect_state" form:"connect_state"`
}

// SetPassword 设置密码
func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = []byte(hashedPassword)
	return nil
}

// ChangePassword 修改密码
func ChangePassword(id uint, password string) error {
	u := User{}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if err := initialize.DB.Where("id = ?", id).First(&u).Error; err != nil {
		return err
	}
	u.Password = []byte(hashedPassword)
	if err := initialize.DB.Save(&u).Error; err != nil {
		return err
	}
	return nil
}

// CheckPassword 验证密码
func CheckPassword(id uint, password string) bool {
	var user User
	if err := initialize.DB.Table(`users`).Where("id = ?", id).First(&user).Error; err != nil {
		// 处理数据库查询错误
		return false
	}

	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		// 密码验证失败，返回 false
		return false
	}

	// 密码验证成功，返回 true
	return true
}

// CheckUserExists 检查用户名是否已存在的示例函数
func CheckUserExists(id uint) bool {
	var user User
	result := initialize.DB.Table(`users`).Where("id = ?", id).First(&user)
	if result.Error == nil {
		return true // 用户名已存在
	}
	return false // 用户名不存在
}

// UserId 获取用户id
func UserId(name string) uint {
	var user User
	result := initialize.DB.Table(`users`).Where("name = ?", name).First(&user)
	if result.Error == nil {
		return user.ID // 用户名已存在
	} else {
		return 0 // 用户名不存在
	}
}

// UserName 获取用户name
func UserName(id uint) string {
	var user User
	result := initialize.DB.Table(`users`).Where("id = ?", id).First(&user)
	if result.Error == nil {
		return user.Name // 用户名存在
	}
	return ""
}

// CreateUser 创建用户的示例函数
func CreateUser(name, password string) bool {
	var user User
	user.Name = name
	user.Password = []byte(password)
	err := user.SetPassword(password)
	if err != nil {
		log.Println(err)
	}

	result := initialize.DB.Table(`users`).Create(&user)
	if result.Error != nil {
		//log.Println("creat failed")
		return false
	} else {
		//log.Println("register successful")
		return true
	}
}
