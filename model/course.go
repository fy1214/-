package model

import (
	"gorm.io/gorm"
	"strconv"
)

type Course struct {
	gorm.Model
	//CourseID  string
	Name      string `gorm:"type:varchar(20);not null;comment:'课程名称'"`
	Cap       int    `gorm:"comment:'课程容量'"`
	TeacherID uint   `gorm:"index:idx_teacherID;comment:'任课老师'"`
}

func (c *Course) CreateCourse() (*Course, ErrNo) {
	result := DB.Where("name = ?", c.Name).First(&c)
	if result.Error == nil {
		return nil, UnknownError
	}
	result = DB.Create(c)
	if result.Error != nil {
		return nil, UnknownError
	}
	var course Course
	result = DB.First(&course)
	if result.Error != nil {
		return nil, UnknownError
	} else {
		return c, OK
	}
}

func GetOneCourse(CourseID string) (*Course, ErrNo) {
	var course Course
	ID, _ := strconv.Atoi(CourseID)
	result := DB.First(&course, ID)
	if result.Error != nil {
		return nil, CourseNotExisted
	} else {
		return &course, OK
	}
}

func GetAllCourse(TeacherID string) (*[]Course, ErrNo) {
	tID, _ := strconv.Atoi(TeacherID)
	var courseList []Course
	DB.Where("teacher_id = ?", tID).Find(&courseList)
	if len(courseList) == 0 {
		return nil, CourseNotExisted
	} else {
		return &courseList, OK
	}
}

// GetAllCourses 查询所有的课程
func GetAllCourses() ([]Course, ErrNo) {
	var courseList []Course
	DB.Find(courseList)
	return courseList, OK
}

func BindCourse(CourseID, TeacherID string) ErrNo {
	cID, _ := strconv.Atoi(CourseID)
	var course Course
	result := DB.First(&course, cID)
	if result.Error != nil {
		return CourseNotExisted
	}
	if course.TeacherID != 0 {
		return CourseHasBound
	}
	tID, _ := strconv.Atoi(TeacherID)
	result = DB.Model(&course).Update("teacher_id", tID)
	if result.Error != nil {
		return UnknownError
	} else {
		return OK
	}
}

func UnBindCourse(CourseID, TeacherID string) ErrNo {
	cID, _ := strconv.Atoi(CourseID)
	var course Course
	result := DB.First(&course, cID)
	if result.Error != nil {
		return CourseNotExisted
	}
	tID, _ := strconv.Atoi(TeacherID)
	if int(course.TeacherID) != tID {
		return CourseNotBind
	}
	DB.Model(&course).Update("teacher_id", 0)
	return OK
}
