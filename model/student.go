package model

import (
	. "TrainingProgram/resource"
	"gorm.io/gorm"
	"strconv"
)

type StudentCourse struct {
	gorm.Model
	StudentID  uint   `gorm:"column:student_id;"`
	CoursesID  uint   `gorm:"column:course_id"`
	CourseName string `gorm:"column:course_name"`
	TeacherID  uint   `gorm:"column:teacher_id"`
}

func GetStudentCourse(StudentID string) (*[]StudentCourse, ErrNo) {
	var student Member
	ID, _ := strconv.Atoi(StudentID)

	//查找是否有对应学生
	result := DB.First(&student, ID)
	if result.Error != nil {
		return nil, StudentNotExisted
	}

	//查找到学生之后得出课表
	var Courses []StudentCourse
	result2 := DB.Where("StudentID = ?", ID).Find(&Courses)
	if result2.Error != nil {
		return nil, StudentHasNoCourse
	} else {
		return &Courses, StudentHasCourse
	}

}
