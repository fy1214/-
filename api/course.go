package api

import (
	"TrainingProgram/model"
	"TrainingProgram/service/course"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateCourseRequest(c *gin.Context) {
	var request model.CreateCourseRequest
	err := c.ShouldBind(&request)
	if err != nil {
		//c.String(http.StatusPreconditionFailed, "添加失败", nil)
	} else {
		course, err := course.CreateCourse(request.Name, request.Cap)
		c.JSON(http.StatusOK, model.CreateCourseResponse{
			Code: err,
			Data: struct{ CourseID string }{CourseID: strconv.Itoa(int(course.ID))},
		})
	}
}

func GetCourseRequest(c *gin.Context) {
	var request model.GetCourseRequest
	err := c.ShouldBind(&request)
	if err != nil {

	} else {
		t, err := course.GetCourse(request.CourseID)
		c.JSON(http.StatusOK, model.GetCourseResponse{
			Code: err,
			Data: t,
		})
	}
}

func GetAllCourseRequest(c *gin.Context) {
	var request model.GetTeacherCourseRequest
	err := c.ShouldBind(&request)
	if err != nil {

	} else {
		list, err := course.GetAllCourse(request.TeacherID)
		c.JSON(http.StatusOK, model.GetTeacherCourseResponse{
			Code: err,
			Data: struct{ CourseList []*model.TCourse }{CourseList: list},
		})
	}
}

func BindCourseRequest(c *gin.Context) {
	var request model.BindCourseRequest
	err := c.ShouldBind(&request)
	if err != nil {

	} else {
		err := course.BindCourse(request.CourseID, request.TeacherID)
		c.JSON(http.StatusOK, model.BindCourseResponse{
			Code: err,
		})
	}
}

func UnBindCourseRequest(c *gin.Context) {
	var request model.UnbindCourseRequest
	err := c.ShouldBind(&request)
	if err != nil {

	} else {
		err := course.UnBindCourse(request.CourseID, request.TeacherID)
		c.JSON(http.StatusOK, model.UnbindCourseResponse{
			Code: err,
		})
	}
}

func GetScheduleCourse(c *gin.Context) {
	var request model.ScheduleCourseRequest
	err := c.ShouldBind(&request)
	if err != nil {

	} else {
		result, err := course.GetScheduleCourseService(request.TeacherCourseRelationShip)
		c.JSON(http.StatusOK, model.ScheduleCourseResponse{
			Code: err,
			Data: result,
		})
	}
}
