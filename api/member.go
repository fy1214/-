package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "TrainingProgram/model"
	. "TrainingProgram/resource"
	"regexp"
	"strconv"
)

func CreateUser(c *gin.Context) {

	curUser, _ := c.Get("member")

	//不是管理员，没有操作权限，返回PermDenied状态码
	if curUser.(Member).UserType != Admin {
		c.JSON(http.StatusOK, ResponseMeta{Code: PermDenied})
		return
	}

	//从请求中拿出数据
	var member Member
	err := c.BindJSON(&member)
	if err != nil {
		return
	}

	//判断参数是否合法
	//用户昵称，不小于 4 位，不超过 20 位（字节）
	matchNickname, _ := regexp.MatchString("^.{4,20}$", member.Nickname)
	if !matchNickname {
		c.JSON(http.StatusOK, ResponseMeta{Code: ParamInvalid})
		return
	}

	//用户名，支持大小写，不小于 8 位 不超过 20 位（字节）
	matchUsername, _ := regexp.MatchString("^[a-zA-Z]{8,20}$", member.Username)
	if !matchUsername {
		c.JSON(http.StatusOK, ResponseMeta{Code: ParamInvalid})
		return
	}

	//密码，同时包括大小写、数字，不少于 8 位 不超过 20 位（字节）
	//matchPassword, _ := regexp.MatchString("^(?<![0-9a-z]+$)(?<![0-9A-Z]+$)(?<![a-zA-Z]+$)[0-9A-Za-z]{8,20}$", member.Password)
	//if !matchPassword {
	//	c.JSON(http.StatusOK, model.ResponseMeta{Code: model.ParamInvalid})
	//	return
	//}
	matchPassword1, _ := regexp.MatchString("^[0-9A-Za-z]{8,20}$", member.Password)
	if !matchPassword1 {
		c.JSON(http.StatusOK, ResponseMeta{Code: ParamInvalid})
		return
	}
	matchPassword2, _ := regexp.MatchString("^[0-9A-Z]+$", member.Password)
	if matchPassword2 {
		c.JSON(http.StatusOK, ResponseMeta{Code: ParamInvalid})
		return
	}
	matchPassword3, _ := regexp.MatchString("^[a-zA-Z]+$", member.Password)
	if matchPassword3 {
		c.JSON(http.StatusOK, ResponseMeta{Code: ParamInvalid})
		return
	}
	matchPassword4, _ := regexp.MatchString("^[0-9a-z]+$", member.Password)
	if matchPassword4 {
		c.JSON(http.StatusOK, ResponseMeta{Code: ParamInvalid})
		return
	}

	//用户类型，枚举值1,2,3
	if !(member.UserType == Admin || member.UserType == Student || member.UserType == Teacher) {
		c.JSON(http.StatusOK, ResponseMeta{Code: ParamInvalid})
		return
	}

	//密码加密
	SetPassword(&member, member.Password)

	//数据添加至数据库
	err = CreateAMember(&member)
	if err != nil {
		c.JSON(http.StatusOK, ResponseMeta{Code: UnknownError})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Code": OK,
		"Data": struct{ UserID string }{UserID: strconv.FormatUint(member.UserID, 10)},
	})
	return
}

func UpdateUser(c *gin.Context) {

	//获取当前用户
	curUser, _ := c.Get("member")

	//从请求中拿出数据
	var member1 UpdateMemberRequest
	c.BindJSON(&member1)
	id, _ := strconv.ParseUint(member1.UserID, 10, 64)
	nickname := member1.Nickname

	//判断新昵称是否合法
	matchName, _ := regexp.MatchString("^.{4,20}$", nickname)
	if !matchName {
		c.JSON(http.StatusOK, ResponseMeta{Code: ParamInvalid})
		return
	}

	//如果修改的不是自己的昵称，且不是管理员
	if curUser.(Member).UserType != Admin && curUser.(Member).UserID != id {
		c.JSON(http.StatusOK, ResponseMeta{Code: PermDenied})
		return
	}

	//查找需要修改的用户
	member2, err := GetAMember(id)
	if err != nil {
		c.JSON(http.StatusOK, ResponseMeta{Code: UserNotExisted})
		return
	}
	if member2.Deleted {
		c.JSON(http.StatusOK, ResponseMeta{Code: UserHasDeleted})
		return
	}

	//修改昵称
	UpdateAMember(member2, "name", nickname)
	c.JSON(http.StatusOK, ResponseMeta{Code: OK})
	return
}

func DeleteUser(c *gin.Context) {

	curUser, _ := c.Get("member")

	//不是管理员，没有操作权限，返回PermDenied状态码
	if curUser.(Member).UserType != Admin {
		c.JSON(http.StatusOK, ResponseMeta{Code: PermDenied})
		return
	}

	//从请求中拿出数据
	var member1 DeleteMemberRequest
	c.BindJSON(&member1)
	id, _ := strconv.ParseUint(member1.UserID, 10, 64)

	//查找并删除数据
	member2, err := GetAMember(id)
	if err != nil {
		c.JSON(http.StatusOK, ResponseMeta{Code: UserNotExisted})
		return
	}
	DeleteAMember(member2)
	c.JSON(http.StatusOK, ResponseMeta{Code: OK})
	return

	//如果删除的用户是当前用户，强制退出，LoginRequired ？

}

func GetAUser(c *gin.Context) {

	curUser, _ := c.Get("member")

	//不是管理员，没有操作权限，返回PermDenied状态码
	if curUser.(Member).UserType != Admin {
		c.JSON(http.StatusOK, ResponseMeta{Code: PermDenied})
		return
	}

	//从请求中拿出数据
	var member1 GetMemberRequest
	c.BindJSON(&member1)
	id, _ := strconv.ParseUint(member1.UserID, 10, 64)

	//查询该id的成员
	member2, err := GetAMember(id)
	//用户不存在
	if err != nil {
		c.JSON(http.StatusOK, ResponseMeta{Code: UserNotExisted})
		return
	}
	//用户已删除
	if member2.Deleted {
		c.JSON(http.StatusOK, ResponseMeta{Code: UserHasDeleted})
		return
	}
	//返回
	c.JSON(http.StatusOK, gin.H{
		"code": OK,
		"Data": TMember{
			UserID:   strconv.FormatUint(member2.UserID, 10),
			Nickname: member2.Nickname,
			Username: member2.Username,
			UserType: member2.UserType,
		},
	})
	return
}

func GetUserList(c *gin.Context) {

	curUser, _ := c.Get("member")

	//不是管理员，没有操作权限，返回PermDenied状态码
	if curUser.(Member).UserType != Admin {
		c.JSON(http.StatusOK, ResponseMeta{Code: PermDenied})
		return
	}

	//从请求中拿出数据
	var listRequest GetMemberListRequest
	c.BindJSON(&listRequest)
	offset := listRequest.Offset
	limit := listRequest.Limit

	memberList, err := GetMemberList(limit, offset)
	if err != nil {
		c.JSON(http.StatusOK, ResponseMeta{Code: UnknownError})
		return
	}
	c.JSON(http.StatusOK, memberList)
	return
}
