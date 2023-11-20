package controller

import (
	"github.com/gin-gonic/gin"
	"wgin/common/request"
	"wgin/common/response"
	"wgin/service"
)

// Login 进行入参校验，并调用 UserService 和 JwtService 服务，颁发 Token
func Login(c *gin.Context) {
	var form request.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
	}

	if err, user := service.UserService.Login(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		token, err, _ := service.JwtService.CreateToken(service.AppGuardName, user)
		if err != nil {
			response.BusinessFail(c, err.Error())
			return
		}
		response.Success(c, token)
	}
}
