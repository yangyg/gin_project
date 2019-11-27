package middleware

import (
	"fmt"
	"gin_project/model"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//AuthCheckRole 权限检查中间件
func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		//根据上下文获取载荷claims 从claims获得role
		claims := c.MustGet("claims").(*jwt.StandardClaims)
		username := claims.Audience
		e := model.Casbin()
		//检查权限
		fmt.Println("user:", username, c.Request.URL.Path, c.Request.Method)
		res, err := e.Enforce(username, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"msg":    "错误消息" + err.Error(),
			})
			c.Abort()
			return
		}
		if res {
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg":    "很抱歉您没有此权限",
			})
			c.Abort()
			return
		}
	}
}
