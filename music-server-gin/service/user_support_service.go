package service

import (
	"music-server-gin/global"
	"music-server-gin/model"
)

type UserSupportService struct{}

var UserSupportServiceApp = new(UserSupportService)

func (s *UserSupportService) AddUserSupport(userSupport *model.UserSupport) error {
	return global.DB.Create(userSupport).Error
}

func (s *UserSupportService) DeleteUserSupport(id uint) error {
	return global.DB.Delete(&model.UserSupport{}, id).Error
}

func (s *UserSupportService) AllUserSupport() ([]model.UserSupport, error) {
	var userSupports []model.UserSupport
	err := global.DB.Find(&userSupports).Error
	return userSupports, err
}

func (s *UserSupportService) UserSupportOfUserId(userId uint) ([]model.UserSupport, error) {
	var userSupports []model.UserSupport
	err := global.DB.Where("user_id = ?", userId).Find(&userSupports).Error
	return userSupports, err
}

func (s *UserSupportService) UserSupportOfCommentId(commentId uint) ([]model.UserSupport, error) {
	var userSupports []model.UserSupport
	err := global.DB.Where("comment_id = ?", commentId).Find(&userSupports).Error
	return userSupports, err
}

func (s *UserSupportService) UpdateUserSupportMsg(userSupport *model.UserSupport) error {
	return global.DB.Save(userSupport).Error
}
