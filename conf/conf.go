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
	godotenv.Load()

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// mysql初始化
	model.Database(os.Getenv("MYSQL_DSN"))

	// redis初始化
	cache.Redis()
}
