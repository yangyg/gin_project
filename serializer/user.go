package serializer

import (
	"gin_project/model"
	"time"

	"github.com/jinzhu/gorm"
)

// User serializer
type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// UserResponse 单个用户序列化
type UserResponse struct {
	Response
	Data User `json:"data"`
}

// UserListResponse 列表序列化
type UserListResponse struct {
	Response
	Data []User `json:"data"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		Username:  user.Username,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) UserResponse {
	return UserResponse{
		Data: BuildUser(user),
	}
}

// UserList 用户列表
func UserList() UserListResponse {
	users := make([]User, 0)
	model.Db.Find(&users)
	return UserListResponse{
		Data: users,
	}
}

// RandomUsers 随机列表
func RandomUsers() (users []User) {
	model.Db.Order(gorm.Expr("rand()")).Limit(3).Find(&users)
	return
}
