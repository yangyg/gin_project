package service

import (
	"gin_project/model"
)

// UserRegisterService register user
type UserRegisterService struct {
	Email           string `form:"email" json:"email" binding:"required"`
	Username        string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// Valid 验证表单
func (service *UserRegisterService) Valid() *model.Result {
	if service.PasswordConfirm != service.Password {
		return &model.Result{
			Code:    40001,
			Message: "两次输入的密码不相同",
		}
	}

	count := 0
	model.Db.Model(&model.User{}).Where("username = ?", service.Username).Count(&count)
	if count > 0 {
		return &model.Result{
			Code:    40001,
			Message: "用户名已经注册",
		}
	}

	count = 0
	model.Db.Model(&model.User{}).Where("email = ?", service.Email).Count(&count)
	if count > 0 {
		return &model.Result{
			Code:    40001,
			Message: "邮箱已经注册",
		}
	}

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() (model.User, *model.Result) {
	user := model.User{
		Username: service.Username,
		Email:    service.Email,
		Status:   model.Active,
	}

	// 表单验证
	if err := service.Valid(); err != nil {
		return user, err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return user, &model.Result{
			Code:    40002,
			Message: "密码加密失败",
		}
	}

	// 创建用户
	if err := model.Db.Create(&user).Error; err != nil {
		return user, &model.Result{
			Code:    40002,
			Message: "注册失败",
		}
	}

	return user, nil
}
