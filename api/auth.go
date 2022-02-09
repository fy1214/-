package api

import (
	"TrainingProgram/model"
	"TrainingProgram/service/auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var loginRequest model.LoginRequest
	if err := c.ShouldBind(&loginRequest); err != nil {
		c.JSON(http.StatusOK, model.LoginResponse{
			Code: model.ParamInvalid,
		})
		return
	}

	result, err := auth.Login(sessions.Default(c), loginRequest.Username, loginRequest.Password)
	c.JSON(http.StatusOK, model.LoginResponse{
		Code: err,
		Data: result,
	})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user_id")
	_ = session.Save()

	c.JSON(http.StatusOK, model.LogoutResponse{
		Code: model.OK,
	})
}

func WhoAmI(c *gin.Context) {
	result, err := auth.WhoAmI(sessions.Default(c))
	c.JSON(http.StatusOK, model.WhoAmIResponse{
		Code: err,
		Data: result,
	})
}
