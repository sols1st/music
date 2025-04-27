package service

import (
	"music-server-gin/global"
	"music-server-gin/model"
)

type AdminService struct{}

var AdminServiceApp = new(AdminService)

func (s *AdminService) VerifyPassword(username, password string) bool {
	var admin model.Admin
	err := global.DB.Where("name = ? AND password = ?", username, password).First(&admin).Error
	return err == nil
}
