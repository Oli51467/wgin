package controller

import (
	"github.com/gin-gonic/gin"
	"wgin/common/request"
	"wgin/common/response"
	"wgin/service"
)

// Register 用户注册
func Register(c *gin.Context) {
	var form request.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	// 在controller捕获service抛出的异常
	if err, user := service.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}
