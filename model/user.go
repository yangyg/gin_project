package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	gorm.Model
	Username       string `json:"username" form:"username"`
	Email          string `json:"email" form:"email" binding:"email" gorm:"unique;not null"`
	PasswordDigest string
	Status         string
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
)

// func init() {
// 	if !Db.HasTable(&User{}) {
// 		Db.CreateTable(&User{})
// 	}
// }

// Save user
func (user User) Save() (uint, error) {

	result := Db.Create(&user)

	if result.Error != nil {
		// log.Panicln("user insert error", result.Error)
	}

	return user.ID, result.Error
}

// QueryByEmail  查询用户
func (user *User) QueryByEmail() User {
	Db.Where("email = ?", user.Email).First(user)
	return *user
}

// QueryByUsername ...查询用户
func (user User) QueryByUsername() User {
	Db.Where("username = ?", user.Username).First(&user)
	return user
}

// QueryByID  query user by id查询用户
func (user User) QueryByID() User {
	Db.First(&user)
	return user
}

// UserList user list
func UserList() []User {
	users := make([]User, 0)
	Db.Find(&users)
	return users
}

// UserDelete 删除
func (user User) UserDelete() error {
	result := Db.Delete(&user)
	return result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
