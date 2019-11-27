package middleware

import (
	"gin_project/config"
	"gin_project/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Auth 权限认证
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {

		result := model.Result{
			Code:    http.StatusUnauthorized,
			Message: "无法认证，重新登录",
			Data:    nil,
		}
		auth := context.Request.Header.Get("Authorization")

		if len(auth) == 0 {
			context.Abort()
			context.JSON(http.StatusUnauthorized, result)
			return
		}

		auth = strings.Fields(auth)[1]
		// 校验token
		claims, err := parseToken(auth)
		if err != nil {
			context.Abort()
			result.Message = "token error " + err.Error()
			context.JSON(http.StatusUnauthorized, result)
		} else {
			println("token 正确")
		}
		context.Set("claims", claims)

		context.Next()
	}
}

func parseToken(token string) (*jwt.StandardClaims, error) {

	// 分割出来载体
	payload := strings.Split(token, ".")
	bytes, e := jwt.DecodeSegment(payload[1])

	if e != nil {
		println(e.Error())
	}
	content := ""
	for i := 0; i < len(bytes); i++ {
		content += string(bytes[i])
	}
	split := strings.Split(content, ",")
	id := strings.SplitAfter(split[2], ":")
	i := strings.Split(id[1], "\"")
	// i = strings.Split(i[1], "\"")

	ID, err := strconv.Atoi(i[1])
	if err != nil {
		println(err.Error())
	}

	user := model.User{}
	user.ID = uint(ID)
	u := model.User.QueryByID(user)
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{},
		func(token *jwt.Token) (i interface{}, e error) {
			return []byte(config.Secret + u.PasswordDigest), nil
		})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}
