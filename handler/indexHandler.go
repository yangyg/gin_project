package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexHandler index
func IndexHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "index")
}
