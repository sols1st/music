package service

import (
	"music-server-gin/global"
	"music-server-gin/model"
)

type SongListService struct{}

func (s *SongListService) AddSongList(songList *model.SongList) error {
	return global.DB.Create(songList).Error
}

func (s *SongListService) UpdateSongListMsg(songList *model.SongList) error {
	return global.DB.Model(songList).Updates(songList).Error
}

func (s *SongListService) UpdateSongListPic(id uint, pic []byte) error {
	return global.DB.Model(&model.SongList{}).Where("id = ?", id).Update("pic", pic).Error
}

func (s *SongListService) DeleteSongList(id uint) error {
	return global.DB.Delete(&model.SongList{}, id).Error
}

func (s *SongListService) AllSongList() ([]model.SongList, error) {
	var songLists []model.SongList
	err := global.DB.Find(&songLists).Error
	return songLists, err
}

func (s *SongListService) SongListOfId(id uint) (model.SongList, error) {
	var songList model.SongList
	err := global.DB.First(&songList, id).Error
	return songList, err
}

func (s *SongListService) SongListOfTitle(title string) ([]model.SongList, error) {
	var songLists []model.SongList
	err := global.DB.Where("title = ?", title).Find(&songLists).Error
	return songLists, err
}

func (s *SongListService) SongListOfStyle(style string) ([]model.SongList, error) {
	var songLists []model.SongList
	err := global.DB.Where("style = ?", style).Find(&songLists).Error
	return songLists, err
}

var SongListServiceApp = new(SongListService)
