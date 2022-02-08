package model

import "TrainingProgram/resource"

//执行数据迁移
func migration() {
	// 自动迁移模式
	//成员表迁移
	_ = DB.AutoMigrate(&resource.Member{})
}
