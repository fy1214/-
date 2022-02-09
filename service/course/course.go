package course

import (
	"TrainingProgram/model"
	"strconv"
)

func CreateCourse(name string, cap int) (*model.Course, model.ErrNo) {
	course := model.Course{
		Name:      name,
		Cap:       cap,
		TeacherID: 0,
	}
	return course.CreateCourse()
}

func GetCourse(CourseID string) (model.TCourse, model.ErrNo) {
	course, err := model.GetOneCourse(CourseID)
	if err != model.OK {
		return model.TCourse{}, err
	}
	var teacherID string
	if course.TeacherID != 0 {
		teacherID = strconv.Itoa(int(course.ID))
	}
	return model.TCourse{
		CourseID:  CourseID,
		Name:      course.Name,
		TeacherID: teacherID,
	}, err
}

func GetAllCourse(TeacherID string) ([]*model.TCourse, model.ErrNo) {
	courseList, err := model.GetAllCourse(TeacherID)
	if err == model.CourseNotExisted {
		return nil, err
	} else {
		t_courses := make([]*model.TCourse, len(*courseList))
		for i, v := range *courseList {
			t_courses[i] = &model.TCourse{
				CourseID:  strconv.Itoa(int(v.ID)),
				Name:      v.Name,
				TeacherID: strconv.Itoa(int(v.TeacherID)),
			}
		}
		return t_courses, err
	}
}

func BindCourse(CourseID, TeacherID string) model.ErrNo {
	return model.BindCourse(CourseID, TeacherID)
}

func UnBindCourse(CourseID, TeacherID string) model.ErrNo {
	return model.UnBindCourse(CourseID, TeacherID)
}

func GetScheduleCourseService(relationShip map[string][]string) (map[string]string, model.ErrNo) {
	name_map := make(map[string]int)
	t_index, c_index := 0, len(relationShip)

	for k, v := range relationShip {
		name_map[k] = t_index
		t_index++
		for _, s := range v {
			_, ok := name_map[s]
			if !ok {
				name_map[s] = c_index
				c_index++
			}
		}
	}

	spot := len(name_map)
	namespace := make([]string, spot)
	for k, v := range name_map {
		namespace[v] = k
	}

	graph := make([][]int, spot)
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]int, spot)
	}

	for k, v := range relationShip {
		ti := name_map[k]
		for _, s := range v {
			si := name_map[s]
			graph[ti][si] = 1
			graph[si][ti] = 1
		}
	}

	Len := len(graph)
	used := make([]bool, Len)
	match := make([]int, Len)
	for i := 0; i < len(match); i++ {
		match[i] = -1
	}

	for i := 0; i < Len; i++ {
		if match[i] == -1 {
			MapSearch(&graph, used, &match, i)
		}
	}

	result := make(map[string]string)
	for i := 0; i < len(relationShip); i++ {
		if match[i] != -1 {
			result[namespace[i]] = namespace[match[i]]
		}
	}
	return result, model.OK
}

func MapSearch(graph *[][]int, used []bool, match *[]int, x int) bool {
	tGraph, tMatch := *graph, *match
	for i := 0; i < len(tGraph); i++ {
		if tGraph[x][i] == 1 {
			if !used[i] {
				used[i] = true
				if tMatch[i] == -1 || MapSearch(graph, used, match, tMatch[i]) {
					tMatch[i] = x
					tMatch[x] = i
					return true
				}
			}
		}
	}
	return false
}
