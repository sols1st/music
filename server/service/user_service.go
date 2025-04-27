package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"music-server-gin/global"
	"music-server-gin/model"
	"time"
)

const (
	SALT = "music" // 密码加密盐值
)

type UserService struct{}

var UserServiceApp = new(UserService)

func (s *UserService) AddUser(user *model.User) error {
	if err := global.DB.Where("username = ?", user.Username).First(&model.User{}).Error; err == nil {
		return errors.New("用户名已存在")
	}

	if err := global.DB.Where("email = ?", user.Email).First(&model.User{}).Error; err == nil {
		return errors.New("邮箱已存在")
	}

	if user.PhoneNum == "" {
		user.PhoneNum = ""
	}
	if user.Email == "" {
		user.Email = ""
	}

	user.Avatar = "/img/avatar/user.jpg"
	user.Password = s.encryptPassword(user.Password)
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()

	tx := global.DB.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (s *UserService) DeleteUser(id uint) error {
	tx := global.DB.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Delete(&model.User{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (s *UserService) AllUser() ([]model.User, error) {
	var users []model.User
	err := global.DB.Find(&users).Error
	return users, err
}

func (s *UserService) UserOfId(id uint) (model.User, error) {
	var user model.User
	err := global.DB.First(&user, id).Error
	return user, err
}

func (s *UserService) UpdateUserMsg(user *model.User) error {
	user.UpdateTime = time.Now()

	tx := global.DB.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Model(user).Updates(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (s *UserService) UpdateUserAvatar(id uint, avatarPath string) error {
	tx := global.DB.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Model(&model.User{}).Where("id = ?", id).Update("avatar", avatarPath).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (s *UserService) UserOfName(name string) ([]model.User, error) {
	var users []model.User
	err := global.DB.Where("username = ?", name).Find(&users).Error
	return users, err
}

func (s *UserService) LoginStatus(username, password string) (model.User, error) {
	var user model.User
	err := global.DB.Where("username = ? AND password = ?", username, s.encryptPassword(password)).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (s *UserService) LoginEmailStatus(email, password string) (model.User, error) {
	var user model.User
	err := global.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return model.User{}, err
	}

	if user.Password != s.encryptPassword(password) {
		return model.User{}, errors.New("密码错误")
	}

	return user, nil
}

func (s *UserService) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := global.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (s *UserService) ResetPassword(email, password string) error {
	return global.DB.Model(&model.User{}).Where("email = ?", email).Update("password", password).Error
}

func (s *UserService) SendVerificationCode(email string) error {
	// TODO: 实现发送验证码的逻辑
	return nil
}

func (s *UserService) UpdatePassword(id uint, oldPassword, newPassword string) error {
	if !s.VerifyPassword(id, oldPassword) {
		return errors.New("密码输入错误")
	}

	encryptedPassword := s.encryptPassword(newPassword)

	tx := global.DB.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Model(&model.User{}).Where("id = ?", id).Update("password", encryptedPassword).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (s *UserService) UpdatePassword01(id uint, newPassword string) error {
	encryptedPassword := s.encryptPassword(newPassword)

	tx := global.DB.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Model(&model.User{}).Where("id = ?", id).Update("password", encryptedPassword).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (s *UserService) VerifyPassword(id uint, password string) bool {
	var user model.User
	if err := global.DB.First(&user, id).Error; err != nil {
		return false
	}
	return user.Password == s.encryptPassword(password)
}

func (s *UserService) encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(SALT + password))
	return hex.EncodeToString(h.Sum(nil))
}
