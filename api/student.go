package api

import (
	"TrainingProgram/model"
	"TrainingProgram/service/student"
	"TrainingProgram/util"
	"github.com/gin-gonic/gin"
)

// BookCourseInit 抢课前的初始化函数，模拟一次抢课活动
func BookCourseInit(c *gin.Context) {
	res := student.BookCourseInit()
	c.JSON(200, res)
}

// BookCourse 抢课处理函数
func BookCourse(c *gin.Context) {
	var service student.BookCourseService
	if err := c.ShouldBind(&service); err == nil {
		res := service.BookCourse()
		c.JSON(200, res)
	} else {
		util.Log().Info("BookCourse error: %s", err.Error())
		c.JSON(200, model.BookCourseResponse{
			Code: 255,
		})
	}
}
