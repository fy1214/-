package util

import (
	"TrainingProgram/model"
	. "TrainingProgram/resource"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // only set seed once
}

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func ConvertMemberToTMember(member Member) model.TMember {
	return model.TMember{
		UserID:         strconv.FormatUint(member.UserID, 10),
		Nickname:       member.Nickname,
		Username:       member.Username,
		PasswordDigest: member.Password,
		UserType:       member.UserType,
	}
}
