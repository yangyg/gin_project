package middleware

import (
	"fmt"
	"gin_project/config"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LoggerToFile 日志记录到文件
func LoggerToFile() gin.HandlerFunc {

	logFilePath := config.LOG_FILE_PATH
	logFileName := config.LOG_FILE_NAME

	//日志文件
	fileName := path.Join(logFilePath, logFileName)

	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//实例化
	logger := logrus.New()

	//设置输出
	logger.Out = src

	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006/01/02 15:04:05",
	})

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqURI := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// reference
		referer := c.Request.Referer()

		clientUserAgent := c.Request.UserAgent()

		entry := logger.WithFields(logrus.Fields{
			"statusCode": statusCode,
			"latency":    latencyTime, // time to process
			"clientIP":   clientIP,
			"method":     reqMethod,
			"path":       reqURI,
			"referer":    referer,
			"userAgent":  clientUserAgent,
		})

		msg := fmt.Sprintf("| %3d | %13v | %15s | %s | %s | %s | %s |", statusCode, latencyTime, clientIP, reqMethod, reqURI, referer, clientUserAgent)
		if statusCode > 499 {
			entry.Error(msg)
		} else if statusCode > 399 {
			entry.Warn(msg)
		} else {
			entry.Info(msg)
		}

		// 日志格式
		// logger.Infof("| %3d | %13v | %15s | %s | %s |",
		// 	statusCode,
		// 	latencyTime,
		// 	clientIP,
		// 	reqMethod,
		// 	reqURI,
		// )
	}
}

// LoggerToMongo 日志记录到 MongoDB
func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// LoggerToES 日志记录到 ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// LoggerToMQ 日志记录到 MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
