package model

import (
	. "TrainingProgram/resource"
	"TrainingProgram/util"
	"strconv"
)

type StudentCourse struct {
	StudentID  uint   `gorm:"column:student_id;primary_key" sql:"type:INT(10) UNSIGNED NOT NULL"`
	CoursesID  uint   `gorm:"column:course_id;primary_key" sql:"type:INT(10) UNSIGNED NOT NULL"`
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

// AddStudentCourseInBatches 批量插入学生课表数据
func AddStudentCourseInBatches(studentIds []string, courseId string) ErrNo {
	// 根据courseId查询课程名称和老师id
	course, errNo := GetOneCourse(courseId)
	if errNo != OK {
		return errNo
	}

	n := len(studentIds)
	courseStudents := make([]StudentCourse, n)
	for i := 0; i < len(studentIds); i++ {
		studentIdInt, _ := strconv.Atoi(studentIds[i])
		courseStudents[i] = StudentCourse{
			StudentID:  uint(studentIdInt),
			CoursesID:  course.ID,
			CourseName: course.Name,
			TeacherID:  course.TeacherID,
		}
	}
	if err := DB.CreateInBatches(courseStudents, n).Error; err != nil {
		util.Log().Error("AddStudentCourseInBatches error: %s\n", err.Error())
		return UnknownError
	} else {
		return OK
	}
}
