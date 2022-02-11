package student

import (
	cache "TrainingProgram/cache"
	"TrainingProgram/model"
	"TrainingProgram/util"
	cron "github.com/robfig/cron/v3"
	"strconv"
	"strings"
)

type BookCourseService struct {
	StudentId string `form:"student_id" json:"student_id" binding:"required"`
	CourseID  string `form:"course_id" json:"course_id" binding:"required"`
}

var SaveUserIdChan chan string

func Init() {
	SaveUserIdChan = make(chan string, 1024)
	// 定时or课抢满了时保存抢到课的学生id的协程
	go func() {
		for true {
			// 需要进行保存的课程id
			courseId, ok := <-SaveUserIdChan
			if ok {
				util.Log().Info("[BookCourse] save course %s's students to db\n", courseId)
				// 获得抢到该课程的所有学生id
				userIds, _ := cache.RedisClient.LRange(courseId+":user_id", 0, -1).Result()
				// 删除redis key
				cache.RedisClient.Del(courseId + ":user_id")
				// 保存id到数据库
				model.AddStudentCourseInBatches(userIds, courseId)
			}
		}
	}()

	// 定时保存
	crontab := cron.New()
	task := func() {
		keys, err := cache.RedisClient.Keys("*:user_id").Result()
		if err != nil {
			return
		}
		if len(keys) == 0 {
			return
		}
		util.Log().Info("begin to save student id\n")
		for _, key := range keys {
			courseId := strings.Split(key, ":")[0]
			SaveUserIdChan <- courseId
		}
	}
	// 添加定时任务, * * * * * 是 crontab,表示每分钟执行一次
	crontab.AddFunc("* * * * *", task)
	// 启动定时器
	crontab.Start()
}

// BookCourseInit 抢课前的初始化函数，模拟一次抢课活动
func BookCourseInit() model.BookCourseResponse {
	// 每次启动服务时清空上一次留在Redis中的数据
	cache.RedisClient.FlushAll()
	bookCourses, errNo := model.GetAllCourses()
	// 如果数据库为空，则不执行下列初始化
	if errNo != model.OK || len(bookCourses) == 0 {
		return model.BookCourseResponse{Code: model.OK}
	}
	for _, course := range bookCourses {
		courseId := strconv.Itoa(int(course.ID))
		if course.Cap <= 0 {
			continue
		}
		// 保存为redis中的hash类型，key为课程id，value为课程容量
		cache.RedisClient.HSet("book_course", courseId, course.Cap)
		// 记录课程是否抢满了，redis key："book_course:finish"，hash key：课程id
		cache.RedisClient.HSet("book_course:finish", courseId, 0)
	}
	return model.BookCourseResponse{Code: model.OK}
}

// BookCourse 抢课函数
// TODO: 过滤重复的抢课请求
func (service *BookCourseService) BookCourse() model.BookCourseResponse {
	userId := service.StudentId
	courseId := service.CourseID
	_, err := cache.RedisClient.HGet("book_course", courseId).Result()
	// 课程id找不到
	if err != nil {
		// 先去数据库里找
		course, errNo := model.GetOneCourse(courseId)
		// 数据库里有且容量大于0
		if errNo == model.OK && course.Cap > 0 {
			cache.RedisClient.HSet("book_course", courseId, course.Cap)
			cache.RedisClient.HSet("book_course:finish", courseId, 0)
		} else {
			// 数据库里也没有
			return model.BookCourseResponse{Code: model.CourseNotExisted}
		}
	}

	// 如果课程已经抢满了，直接返回
	res, err := cache.RedisClient.HGet("book_course:finish", courseId).Result()
	if err != nil {
		return model.BookCourseResponse{Code: model.UnknownError}
	}
	if intRes, _ := strconv.Atoi(res); intRes == 1 {
		return model.BookCourseResponse{Code: model.CourseNotAvailable}
	}

	leftCap, err := cache.RedisClient.HIncrBy("book_course", courseId, -1).Result()
	if err != nil {
		return model.BookCourseResponse{Code: model.CourseNotAvailable}
	}
	// 抢课成功
	if leftCap >= 0 {
		// 记录抢到课的学生id，redis list key："{courseId}:user_id"
		cache.RedisClient.LPush(courseId+":user_id", userId)
		util.Log().Info("[BookCourse] user %s book course %s\n", userId, courseId)
		// 课程抢满了
		if leftCap == 0 {
			// 记录课程抢满了
			_, err := cache.RedisClient.HSet("book_course:finish", courseId, 1).Result()
			// 通知go routine将抢到课程的学生id保存到数据库
			SaveUserIdChan <- courseId
			if err != nil {
				return model.BookCourseResponse{Code: model.UnknownError}
			}
		}
		return model.BookCourseResponse{Code: model.OK}
	} else {
		// 高并发导致的课程超卖
		return model.BookCourseResponse{Code: model.CourseNotAvailable}
	}
}
