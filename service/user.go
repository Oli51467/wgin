package service

import (
	"errors"
	"wgin/common/request"
	"wgin/global"
	"wgin/model"
	"wgin/util"
)

type userService struct {
}

var UserService = new(userService)

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
