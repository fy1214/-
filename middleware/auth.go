package middleware

import (
	"TrainingProgram/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("user_id")
		if uid != nil {
			member, err := model.GetMember(uid)
			if err == nil {
				c.Set("member", &member)
			}
		}
		c.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if member, _ := c.Get("member"); member != nil {
			if _, ok := member.(*model.TMember); ok {
				c.Next()
				return
			}
		}

		c.JSON(200, model.LoginResponse{
			Code: 6,
			Data: struct{ UserID string }{UserID: ""},
		})
		c.Abort()
	}
}
