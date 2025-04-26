package service

import (
	"errors"
	"music-server-gin/global"
	"music-server-gin/model"
	"time"
)

type ConsumerService struct{}

var ConsumerServiceApp = new(ConsumerService)

func (s *ConsumerService) AddConsumer(consumer *model.Consumer) error {
	if err := global.DB.Where("username = ?", consumer.Username).First(&model.Consumer{}).Error; err == nil {
		return errors.New("用户名已存在")
	}

	consumer.Avatar = "img/avatorImages/user.jpg"
	consumer.CreateTime = time.Now()
	consumer.UpdateTime = time.Now()
	return global.DB.Create(consumer).Error
}

func (s *ConsumerService) DeleteConsumer(id uint) error {
	return global.DB.Delete(&model.Consumer{}, id).Error
}

func (s *ConsumerService) AllConsumer() ([]model.Consumer, error) {
	var consumers []model.Consumer
	err := global.DB.Find(&consumers).Error
	return consumers, err
}

func (s *ConsumerService) ConsumerOfId(id uint) (model.Consumer, error) {
	var consumer model.Consumer
	err := global.DB.First(&consumer, id).Error
	return consumer, err
}

func (s *ConsumerService) UpdateConsumerMsg(consumer *model.Consumer) error {
	return global.DB.Model(consumer).Updates(consumer).Error
}

func (s *ConsumerService) UpdateConsumerAvatar(id uint, file interface{}) error {
	return global.DB.Model(&model.Consumer{}).Where("id = ?", id).Update("avatar", file).Error
}

func (s *ConsumerService) ConsumerOfName(name string) ([]model.Consumer, error) {
	var consumers []model.Consumer
	err := global.DB.Where("username = ?", name).Find(&consumers).Error
	return consumers, err
}

func (s *ConsumerService) VerifyPassword(username, password string) (model.Consumer, error) {
	var consumer model.Consumer
	err := global.DB.Where("username = ? AND password = ?", username, password).First(&consumer).Error
	return consumer, err
}

func (s *ConsumerService) VerifyEmailPassword(email, password string) (model.Consumer, error) {
	var consumer model.Consumer
	err := global.DB.Where("email = ? AND password = ?", email, password).First(&consumer).Error
	return consumer, err
}

func (s *ConsumerService) ResetPassword(email, password string) error {
	return global.DB.Model(&model.Consumer{}).Where("email = ?", email).Update("password", password).Error
}

func (s *ConsumerService) SendVerificationCode(email string) error {
	// TODO: 实现发送验证码的逻辑
	return nil
}

func (s *ConsumerService) UpdatePassword(consumer *model.Consumer) error {
	return global.DB.Model(consumer).Update("password", consumer.Password).Error
}
