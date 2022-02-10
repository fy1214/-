package student

import (
	"TrainingProgram/model"
)

type Data struct {
	CourseList []model.TCourse
}

// GetStudentCourse 函数获取最终的输出
func GetStudentCourse(StudentID string) (Data, model.ErrNo) {
	studentCourseList, err := model.GetStudentCourse(StudentID)
	if err == model.StudentNotExisted || err == model.StudentHasNoCourse {
		return struct{ CourseList []model.TCourse }{CourseList: nil}, err
	} else {
		//将course类型转换成TCourse类型以返回
		return struct{ CourseList []model.TCourse }{CourseList: courseToTCourse(studentCourseList)}, err
	}
}

// courseTOTCourse 函数把[]studentcourse转换成[]Tcourse
// TODO: 转换函数
func courseToTCourse(studentCourseList *[]model.StudentCourse) []model.TCourse {
	return nil
}
