package model

import "golang.org/x/crypto/bcrypt"

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
