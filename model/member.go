package model

import (
	. "TrainingProgram/resource"
	"golang.org/x/crypto/bcrypt"
)

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
)

func GetMember(ID interface{}) (TMember, error) {
	var member TMember
	result := DB.First(&member, ID)
	return member, result.Error
}

func (member *TMember) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	member.PasswordDigest = string(bytes)
	return nil
}

//Member增删改查
//创建Member
func CreateAMember(member *Member) (err error) {
	err = DB.Create(&member).Error
	if err != nil {
		return err
	}
	return
}

//修改Member
func UpdateAMember(member *Member, colume string, value string) {
	DB.Model(&member).UpdateColumn(colume, value)
	return
}

//删除Member
func DeleteAMember(member *Member) {
	DB.Model(&member).UpdateColumn("deleted", true)
	return
}

//单个查询Member
func GetAMember(id uint64) (member *Member, err error) {
	err = DB.First(&member, id).Error
	if err != nil {
		return nil, err
	}
	return
}

//批量查询Member
func GetMemberList(limit int, offset int) (memberList []*Member, err error) {
	err = DB.Where("deleted = ?", false).Limit(limit).Offset(offset).Find(&memberList).Error
	if err != nil {
		return nil, err
	}
	return
}
