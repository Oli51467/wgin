package service

import (
	"errors"
	"strconv"
	"wgin/common/request"
	"wgin/global"
	"wgin/model"
	"wgin/util"
)

type userService struct {
}

var UserService = new(userService)

// Register 用户注册服务
func (userService *userService) Register(params request.Register) (err error, user model.User) {
	var userDB = global.App.Database.Where("mobile = ?", params.Mobile).Select("id").First(&model.User{})
	if userDB.RowsAffected != 0 {
		err = errors.New("手机号已存在")
		return
	}
	user = model.User{Name: params.Name, Mobile: params.Mobile, Password: util.MakeBcrypt([]byte(params.Password))}
	err = global.App.Database.Create(&user).Error
	return
}

// Login 用户登录服务
func (userService *userService) Login(params request.Login) (err error, user *model.User) {
	err = global.App.Database.Where("mobile = ?", params.Mobile).First(&user).Error
	if err != nil {
		err = errors.New("用户不存在")
	}
	if !util.CheckBcrypt([]byte(params.Password), user.Password) {
		err = errors.New("用户名或密码错误")
	}
	return
}

// GetUserInfo 获取用户信息
func (userService *userService) GetUserInfo(id string) (err error, user model.User) {
	uid, err := strconv.Atoi(id)
	err = global.App.Database.First(&user, uid).Error
	if err != nil {
		err = errors.New("用户不存在")
	}
	return
}
