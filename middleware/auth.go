package middleware

import (
	"TrainingProgram/model"
	"TrainingProgram/resource"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("user_id")
		if uid != nil {
			member, err := model.GetMember(uid)
			if err == nil {
				c.Set("user", &member)
			}
		}
		c.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if member, _ := c.Get("user"); member != nil {
			if _, ok := member.(*resource.Member); ok {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusOK, model.LoginResponse{
			Code: model.LoginRequired,
			Data: struct{ UserID string }{UserID: ""},
		})
		c.Abort()
	}
}
