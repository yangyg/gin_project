package service

import (
	"fmt"
	"gin_project/config"
	"gin_project/model"
	"gin_project/serializer"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// UserLoginService  登录
type UserLoginService struct {
	Username string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// Login 用户登录函数
func (service *UserLoginService) Login() (model.User, *serializer.Response) {
	var user model.User

	if err := model.Db.Where("username = ?", service.Username).First(&user).Error; err != nil {
		return user, &serializer.Response{
			Status: 40001,
			Msg:    config.T("Login.Invalid"),
		}
	}

	if user.CheckPassword(service.Password) == false {
		return user, &serializer.Response{
			Status: 40001,
			Msg:    config.T("Login.Invalid"),
		}
	}

	return user, nil
}

// GenerateJWT  jwt generate
func (service UserLoginService) GenerateJWT(user model.User) (string, error) {
	expiresTime := time.Now().Unix() + int64(config.OneDayOfHours)
	ID := fmt.Sprint(user.ID)
	claims := jwt.StandardClaims{
		Audience:  user.Username,     // 受众
		ExpiresAt: expiresTime,       // 失效时间
		Id:        ID,                // 编号
		IssuedAt:  time.Now().Unix(), // 签发时间
		Issuer:    "gin project",     // 签发人
		NotBefore: time.Now().Unix(), // 生效时间
		Subject:   "login",           // 主题
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 通过密码和保留字段加密
	var jwtSecret = []byte(config.Secret + user.PasswordDigest)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}
