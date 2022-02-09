package conf

import (
	"TrainingProgram/cache"
	"TrainingProgram/model"
	"TrainingProgram/util"
	"github.com/joho/godotenv"
	"os"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	_ = godotenv.Load()
	if os.Getenv("ACTIVE_ENV") == "DEV" {
		_ = godotenv.Load(".env.dev")
	} else if os.Getenv("ACTIVE_ENV") == "PROD" {
		_ = godotenv.Load(".env.prod")
	}

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// mysql初始化
	model.Database(os.Getenv("MYSQL_DSN"))

	// redis初始化
	cache.Redis()
}
