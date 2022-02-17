package auth

import (
	"TrainingProgram/model"
	. "TrainingProgram/resource"
	"github.com/gin-contrib/sessions"
	"strconv"
)

func Login(session sessions.Session, username, password string) (struct{ UserID string }, model.ErrNo) {
	member, err := model.LoginMember(username, password)
	if err != nil {
		return struct{ UserID string }{}, model.WrongPassword
	}

	session.Set("user_id", member.UserID)
	_ = session.Save()

	return struct{ UserID string }{UserID: strconv.FormatUint(member.UserID, 10)}, model.OK
}

func WhoAmI(session sessions.Session) (model.TMember, model.ErrNo) {
	uid := session.Get("user_id")
	if uid == nil {
		return model.TMember{}, model.LoginRequired
	}

	member, err := model.GetMember(uid)
	if err != nil || member.Deleted {
		return model.TMember{}, model.UnknownError // user should exist
	}

	return convertMemberToTMember(member), model.OK
}

func convertMemberToTMember(member Member) model.TMember {
	return model.TMember{
		UserID:         strconv.FormatUint(member.UserID, 10),
		Nickname:       member.Nickname,
		Username:       member.Username,
		// PasswordDigest: member.Password,
		UserType:       member.UserType,
	}
}
