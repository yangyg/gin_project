package handler

import (
	"gin_project/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Subjects index
func Subjects(context *gin.Context) {
	subjects, err := model.SubjectList()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"result": nil,
			"err":    err,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"result": subjects,
		})
	}
}

// SubjectCreate ...
func SubjectCreate(context *gin.Context) {
	var subject model.Subject
	if err := context.ShouldBind(&subject); err != nil {
		context.String(http.StatusBadRequest, "输入的数据不合法")
		log.Panicln("err ->", err.Error())
	}
	_, err := subject.SubjectCreate()
	log.Println(err)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"result": err,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"result": subject,
		})
	}
}

// SubjectUpdate ...
func SubjectUpdate(context *gin.Context) {

}

// SubjectDelete ...
func SubjectDelete(context *gin.Context) {

}
