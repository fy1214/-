package main

import (
	"TrainingProgram/conf"
	"TrainingProgram/server"
	"TrainingProgram/service/student"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 初始化抢课相关
	student.Init()
	student.BookCourseInit()

	// 装载路由
	r := server.NewRouter()
	r.Run(":8080")
}
