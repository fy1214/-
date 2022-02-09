package server

import (
	"TrainingProgram/api"
	"TrainingProgram/middleware"
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

	g.POST("/auth/login", api.Login)

	// 需要登录保护的
	g.Use(middleware.AuthRequired())
	{
		// 成员管理
		g.POST("/member/create", api.CreateUser)
		g.GET("/member", api.GetAUser)
		g.GET("/member/list", api.GetUserList)
		g.POST("/member/update", api.UpdateUser)
		g.POST("/member/delete", api.DeleteUser)

		// 登录
		g.POST("/auth/logout", api.Logout)
		g.GET("/auth/whoami", api.WhoAmI)

		// 排课
		g.POST("/course/create", api.CreateCourseRequest)
		g.GET("/course/get", api.GetCourseRequest)

		g.POST("/teacher/bind_course", api.BindCourseRequest)
		g.POST("/teacher/unbind_course", api.UnBindCourseRequest)
		g.GET("/teacher/get_course", api.GetAllCourseRequest)
		g.POST("/course/schedule", api.GetScheduleCourse)

		// 抢课
		g.POST("/student/book_course")
		g.GET("/student/course")
	}
}
