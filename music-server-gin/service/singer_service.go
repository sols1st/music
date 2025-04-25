package service

import (
	"music-server-gin/global"
	"music-server-gin/model"
)

type SingerService struct{}

func (s *SingerService) AddSinger(singer *model.Singer) error {
	return global.DB.Create(singer).Error
}

func (s *SingerService) UpdateSingerMsg(singer *model.Singer) error {
	return global.DB.Model(singer).Updates(singer).Error
}

func (s *SingerService) UpdateSingerPic(id uint, pic []byte) error {
	return global.DB.Model(&model.Singer{}).Where("id = ?", id).Update("pic", pic).Error
}

func (s *SingerService) DeleteSinger(id uint) error {
	return global.DB.Delete(&model.Singer{}, id).Error
}

func (s *SingerService) AllSinger() ([]model.Singer, error) {
	var singers []model.Singer
	err := global.DB.Find(&singers).Error
	return singers, err
}

func (s *SingerService) SingerOfId(id uint) (model.Singer, error) {
	var singer model.Singer
	err := global.DB.First(&singer, id).Error
	return singer, err
}

func (s *SingerService) SingerOfName(name string) ([]model.Singer, error) {
	var singers []model.Singer
	err := global.DB.Where("name = ?", name).Find(&singers).Error
	return singers, err
}

func (s *SingerService) SingerOfSex(sex string) ([]model.Singer, error) {
	var singers []model.Singer
	err := global.DB.Where("sex = ?", sex).Find(&singers).Error
	return singers, err
}

var SingerServiceApp = new(SingerService)
