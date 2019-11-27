package handler

import (
	"encoding/json"
	"gin_project/model"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddPolicy add policy
func AddPolicy(ctx *gin.Context) {
	casbinModel := model.CasbinModel{}
	if err := ctx.ShouldBind(&casbinModel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	ok, err := casbinModel.AddPolicy()
	if ok && err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"res": casbinModel,
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
}

// RemovePolicy delete policy
func RemovePolicy(ctx *gin.Context) {
	casbinModel := model.CasbinModel{}
	if err := ctx.ShouldBind(&casbinModel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	}
	ok, err := casbinModel.RemovePolicy()
	if ok && err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"res": casbinModel,
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
}

// Role role for user
func Role(ctx *gin.Context) {

	body, _ := ioutil.ReadAll(ctx.Request.Body)

	var mapResult map[string]string
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal(body, &mapResult); err != nil {
		log.Fatal(err)
		ctx.Abort()
	}
	if ctx.Request.Method == "POST" {
		model.AddRoleForUser(mapResult["user_name"], mapResult["role_name"])
	} else if ctx.Request.Method == "DELETE" {
		model.DeleteRoleForUser(mapResult["user_name"], mapResult["role_name"])
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
