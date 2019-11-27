package routers

import (
	"gin_project/handler"
	"gin_project/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter 路由
func SetupRouter() *gin.Engine {

	router := gin.Default()
	// logrus
	router.Use(middleware.LoggerToFile())
	// log request body
	router.Use(middleware.RequestLogger())
	// cors 跨域设置
	router.Use(middleware.Cors())

	// index
	index := router.Group("/")
	{
		index.Any("", handler.IndexHandler)
		index.POST("/register", handler.UserRegister)
		index.POST("/login", handler.UserLogin)
	}
	// auth routers
	v1 := router.Group("/api/v1", middleware.Auth())
	v1.Use(middleware.AuthCheckRole())
	{
		userRouter := v1.Group("/users")
		{
			userRouter.GET("", handler.Users)
			userRouter.GET("/random", handler.RandomUsers)
			userRouter.DELETE("/", handler.UserDelete)
			// userRouter.GET("/:name", handler.UserSave)
		}
		auth := v1.Group("/auth")
		{
			auth.POST("/policy", handler.AddPolicy)
			auth.DELETE("/policy", handler.RemovePolicy)
			auth.POST("/role", handler.Role)
			auth.DELETE("/role", handler.Role)
		}
		subjectRouter := v1.Group("/subjects")
		{
			subjectRouter.GET("", handler.Subjects)
			subjectRouter.POST("", handler.SubjectCreate)
			subjectRouter.PATCH("/:id", handler.SubjectUpdate)
			subjectRouter.DELETE("/:id", handler.SubjectDelete)
		}
	}

	return router
}
