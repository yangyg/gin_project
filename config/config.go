package config

import (
	"gin_project/model"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Secret  jwt secret
const Secret = "gin project"

// OneDayOfHours for jwt expires time
const OneDayOfHours = 60 * 60 * 24

func init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 读取翻译文件
	if err := LoadLocales("config/locales/zh-cn.yaml"); err != nil {
		log.Panicln("翻译文件加载失败", err)
	}

	// 连接数据库
	model.InitDb(os.Getenv("MYSQL_DSN"))
}
