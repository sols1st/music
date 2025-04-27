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

type ConsumerService struct{}

var ConsumerServiceApp = new(ConsumerService)

func (s *ConsumerService) AddUser(consumer *model.Consumer) error {
	if err := global.DB.Where("username = ?", consumer.Username).First(&model.Consumer{}).Error; err == nil {
		return errors.New("用户名已存在")
	}

	if err := global.DB.Where("email = ?", consumer.Email).First(&model.Consumer{}).Error; err == nil {
		return errors.New("邮箱不允许重复")
	}

	if consumer.PhoneNum == "" {
		consumer.PhoneNum = ""
	}
	if consumer.Email == "" {
		consumer.Email = ""
	}

	consumer.Avatar = "/img/avatar/user.jpg"
	consumer.Password = s.encryptPassword(consumer.Password)
	consumer.CreateTime = time.Now()
	consumer.UpdateTime = time.Now()

	tx := global.DB.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(consumer).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (s *ConsumerService) DeleteUser(id uint) error {
	tx := global.DB.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Delete(&model.Consumer{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (s *ConsumerService) AllUser() ([]model.Consumer, error) {
	var consumers []model.Consumer
	err := global.DB.Find(&consumers).Error
	return consumers, err
}

func (s *ConsumerService) UserOfId(id uint) (model.Consumer, error) {
	var consumer model.Consumer
	err := global.DB.First(&consumer, id).Error
	return consumer, err
}

func (s *ConsumerService) UpdateUserMsg(consumer *model.Consumer) error {
	consumer.UpdateTime = time.Now()

	tx := global.DB.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Model(consumer).Updates(consumer).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (s *ConsumerService) UpdateUserAvatar(id uint, avatarPath string) error {
	tx := global.DB.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Model(&model.Consumer{}).Where("id = ?", id).Update("avatar", avatarPath).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (s *ConsumerService) ConsumerOfName(name string) ([]model.Consumer, error) {
	var consumers []model.Consumer
	err := global.DB.Where("username = ?", name).Find(&consumers).Error
	return consumers, err
}

func (s *ConsumerService) LoginStatus(username, password string) (model.Consumer, error) {
	var consumer model.Consumer
	err := global.DB.Where("username = ? AND password = ?", username, s.encryptPassword(password)).First(&consumer).Error
	if err != nil {
		return model.Consumer{}, errors.New("用户名或密码错误")
	}
	return consumer, nil
}

func (s *ConsumerService) LoginEmailStatus(email, password string) (model.Consumer, error) {
	var consumer model.Consumer
	err := global.DB.Where("email = ?", email).First(&consumer).Error
	if err != nil {
		return model.Consumer{}, errors.New("邮箱不存在")
	}

	if consumer.Password != s.encryptPassword(password) {
		return model.Consumer{}, errors.New("密码错误")
	}

	return consumer, nil
}

func (s *ConsumerService) FindByEmail(email string) (model.Consumer, error) {
	var consumer model.Consumer
	err := global.DB.Where("email = ?", email).First(&consumer).Error
	if err != nil {
		return model.Consumer{}, err
	}
	return consumer, nil
}

func (s *ConsumerService) ResetPassword(email, password string) error {
	return global.DB.Model(&model.Consumer{}).Where("email = ?", email).Update("password", password).Error
}

func (s *ConsumerService) SendVerificationCode(email string) error {
	// TODO: 实现发送验证码的逻辑
	return nil
}

func (s *ConsumerService) UpdatePassword(id uint, oldPassword, newPassword string) error {
	if !s.VerifyPassword(id, oldPassword) {
		return errors.New("密码输入错误")
	}

	encryptedPassword := s.encryptPassword(newPassword)

	tx := global.DB.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Model(&model.Consumer{}).Where("id = ?", id).Update("password", encryptedPassword).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (s *ConsumerService) UpdatePassword01(id uint, newPassword string) error {
	encryptedPassword := s.encryptPassword(newPassword)

	tx := global.DB.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Model(&model.Consumer{}).Where("id = ?", id).Update("password", encryptedPassword).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (s *ConsumerService) VerifyPassword(id uint, password string) bool {
	var consumer model.Consumer
	if err := global.DB.First(&consumer, id).Error; err != nil {
		return false
	}
	return consumer.Password == s.encryptPassword(password)
}

func (s *ConsumerService) encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(SALT + password))
	return hex.EncodeToString(h.Sum(nil))
}
