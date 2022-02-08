package server

import (
	"TrainingProgram/middleware"
	"TrainingProgram/api"
	"github.com/gin-gonic/gin"
	"os"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.CurrentUser())

	RegisterRouter(r)
	return r
}

func RegisterRouter(r *gin.Engine) {
	g := r.Group("/api/v1")

	g.POST("/auth/login")

	// 需要登录保护的
	g.Use(middleware.AuthRequired())
	{
		// 成员管理
		g.POST("/member/create", api.CreateMember)
		g.GET("/member", api.GetAMember)
		g.GET("/member/list", api.GetMemberList)
		g.POST("/member/update", api.UpdateMember)
		g.POST("/member/delete", api.DeleteMember)

		// 登录
		g.POST("/auth/logout")
		g.GET("/auth/whoami")

		// 排课
		g.POST("/course/create")
		g.GET("/course/get")

		g.POST("/teacher/bind_course")
		g.POST("/teacher/unbind_course")
		g.GET("/teacher/get_course")
		g.POST("/course/schedule")

		// 抢课
		g.POST("/student/book_course")
		g.GET("/student/course")
	}
}
