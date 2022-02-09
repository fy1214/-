package resource

// Member 成员数据表
type Member struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`          //ID 主键
	Name     string `gorm:"<-;type:varchar(20)"`               //昵称
	UserName string `gorm:"<-:create;unique;type:varchar(20)"` //用户名
	Password string `gorm:"<-:create;type:varchar(20)"`        //密码
	UserType uint8  `gorm:"<-:create"`                         //用户类型：1.管理员；2.学生；3.老师
	Deleted  bool   `gorm:"<-;default:false"`                  //是否删除
	CourseID uint   `gorm:"<-:create"`                         //外键，关联老师的课程表
}
