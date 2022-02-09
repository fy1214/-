package model

import (
	. "TrainingProgram/resource"
	"gorm.io/gorm"
	"strconv"
)

type Data struct {
	CourseList []TCourse
}

type StudentCourse struct {
	gorm.Model
	StudentID      uint     `gorm:"index:idx_studentID;comment:'选课学生'"`
	StudentCourses []Course `gorm:""`
}

func GetStudentCourse(StudentID string) (Data, ErrNo) {
	var student Member
	ID, _ := strconv.Atoi(StudentID)
	result := DB.First(&student, ID)

}
