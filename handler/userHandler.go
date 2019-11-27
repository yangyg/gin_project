package handler

import (
	"gin_project/model"
	"gin_project/serializer"
	"gin_project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// UserSave handler
func UserSave(context *gin.Context) {
	username := context.Param("name")
	context.String(http.StatusOK, username+"用户已经保存")
}

// Users 列表
func Users(context *gin.Context) {
	result := serializer.UserList()
	context.JSON(http.StatusOK, result)
}

// RandomUsers ...
func RandomUsers(context *gin.Context) {
	result := serializer.RandomUsers()
	context.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

// UserRegister 注册
func UserRegister(context *gin.Context) {
	var service service.UserRegisterService
	result := &model.Result{
		Code:    200,
		Message: "注册成功",
		Data:    nil,
	}
	if err := context.ShouldBind(&service); err == nil {
		if user, err := service.Register(); err != nil {
			result = err
			context.JSON(http.StatusInternalServerError, result)
		} else {
			res := serializer.BuildUserResponse(user)
			context.JSON(http.StatusOK, res)
		}
	} else {
		result.Code = 500
		result.Message = err.Error()
		result.Data = err
		context.JSON(http.StatusInternalServerError, result)
	}
}

// UserLogin 登录
func UserLogin(ctx *gin.Context) {
	// 获取用户
	service := &service.UserLoginService{}

	result := &model.Result{
		Code:    200,
		Message: "登录成功",
		Data:    nil,
	}
	if e := ctx.ShouldBind(&service); e != nil {
		result.Message = "数据绑定失败"
		result.Code = http.StatusUnauthorized
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"result": result,
		})
		return
	}

	if user, err := service.Login(); err != nil {
		ctx.JSON(200, err)
	} else {
		// 设置jwt
		token, err := service.GenerateJWT(user)
		if err != nil {
			result.Data = err
			ctx.JSON(http.StatusOK, gin.H{
				"result": result,
			})
		} else {
			result.Data = "Bearer " + token
			ctx.JSON(http.StatusOK, gin.H{
				"result": result,
			})
		}
	}

}

// UserDelete 删除
func UserDelete(context *gin.Context) {
	sid := context.PostForm("id")
	iid, _ := strconv.Atoi(sid)
	id := uint(iid)
	user := model.User{Model: gorm.Model{ID: id}}
	u := user.QueryByID()

	err := u.UserDelete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"result": err,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"result": "delete",
		})
	}
}
