package resource

import . "TrainingProgram/model"

// Member数据表
type Member struct {
	UserID   uint64   `gorm:"primaryKey;autoIncrement"`          //ID 主键
	Nickname string   `gorm:"<-;type:varchar(20)"`               //昵称
	Username string   `gorm:"<-:create;unique;type:varchar(20)"` //用户名
	Password string   `gorm:"<-:create;type:varchar(20)"`        //密码
	UserType UserType `gorm:"<-:create"`                         //用户类型：1.管理员；2.学生；3.老师
	Deleted  bool     `gorm:"<-;default:false"`                  //是否删除
}
