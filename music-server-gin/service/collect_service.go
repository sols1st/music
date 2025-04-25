package service

import (
	"music-server-gin/global"
	"music-server-gin/model"
)

type CollectService struct{}

var CollectServiceApp = new(CollectService)

func (s *CollectService) AddCollect(collect *model.Collect) error {
	return global.DB.Create(collect).Error
}

func (s *CollectService) DeleteCollect(id uint) error {
	return global.DB.Delete(&model.Collect{}, id).Error
}

func (s *CollectService) AllCollect() ([]model.Collect, error) {
	var collects []model.Collect
	err := global.DB.Find(&collects).Error
	return collects, err
}

func (s *CollectService) CollectOfUserId(userId uint) ([]model.Collect, error) {
	var collects []model.Collect
	err := global.DB.Where("user_id = ?", userId).Find(&collects).Error
	return collects, err
}

func (s *CollectService) CollectOfSongId(songId uint) ([]model.Collect, error) {
	var collects []model.Collect
	err := global.DB.Where("song_id = ?", songId).Find(&collects).Error
	return collects, err
} 