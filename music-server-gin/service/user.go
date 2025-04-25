package service

import (
	"errors"

	"github.com/MephistoSolsist/mysql-practice/global"
	"github.com/MephistoSolsist/mysql-practice/model"
	"github.com/MephistoSolsist/mysql-practice/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserService struct{}

func (*UserService) Register(user *model.User) error {
	var u model.User
	if !errors.Is(global.DB.Where("user_id = ?", user.UserId).First(&u).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已被注册")
	}
	user.Password = util.BcryptHash(user.Password)
	err := global.DB.Create(user).Error
	return err
}

func (*UserService) Login(user *model.User) (model.User, error) {
	var u model.User
	err := global.DB.Where("user_id = ?", user.UserId).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.User{}, errors.New("用户不存在")
	}
	if err == nil {
		if ok := util.BcryptCheck(user.Password, u.Password); !ok {
			return model.User{}, errors.New("密码错误")
		}
	}
	return u, err
}

func (*UserService) Delete(user *model.User) error {
	err := global.DB.Delete(user).Error
	return err
}

func (*UserService) ChangePassword(user *model.User) (err error) {
	var u model.User
	err = global.DB.Where("user_id=?", u.UserId).First(&u).Error
	if err != nil {
		return
	}
	u.Password = util.BcryptHash(user.Password)
	err = global.DB.Save(&u).Error
	return
}

var UserServiceApp = new(UserService)
