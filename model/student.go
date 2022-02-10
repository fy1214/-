package model

import (
	. "TrainingProgram/resource"
	"gorm.io/gorm"
	"strconv"
)

type StudentCourse struct {
	gorm.Model
	StudentID uint   `gorm:"index:idx_studentID;comment:'选课学生'"`
	CoursesID string `gorm:"comment:'学生课程'"`
}

func GetStudentCourse(StudentID string) (*[]StudentCourse, ErrNo) {
	var Courses []StudentCourse
	var student Member
	ID, _ := strconv.Atoi(StudentID)
	//查找是否有对应学生
	result := DB.First(&student, ID)
	if result.Error != nil {
		return nil, StudentNotExisted
	}
	//查找到学生之后得出课表
	result2 := DB.Where("StudentID = ?", ID).Find(&Courses)
	if result2.Error != nil {
		return nil, StudentHasNoCourse
	} else {
		return &Courses, StudentHasCourse
	}

}
