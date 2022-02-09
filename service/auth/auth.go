package auth

import (
	"TrainingProgram/model"
	"github.com/gin-contrib/sessions"
)

func Login(session sessions.Session, username, password string) (struct{ UserID string }, model.ErrNo) {
	member, err := model.LoginMember(username, password)
	if err != nil {
		return struct{ UserID string }{}, model.WrongPassword
	}

	session.Set("user_id", member.UserID)
	_ = session.Save()

	return struct{ UserID string }{UserID: member.UserID}, model.OK
}

func WhoAmI(session sessions.Session) (model.TMember, model.ErrNo) {
	uid := session.Get("user_id")
	if uid == nil {
		return model.TMember{}, model.LoginRequired
	}

	member, err := model.GetMember(uid)
	if err != nil {
		return model.TMember{}, model.UnknownError // user should exist
	}

	return member, model.OK
}
