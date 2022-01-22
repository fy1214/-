package TrainingProgram

import (
	"TrainingProgram/types"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 注册路由
	types.RegisterRouter(r)
	// 启动并监听端口
	r.Run(":8080")
}
