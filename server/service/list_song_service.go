package service

import (
	"music-server-gin/global"
	"music-server-gin/model"
)

type ListSongService struct{}

func (s *ListSongService) AddListSong(listSong *model.ListSong) error {
	return global.DB.Create(listSong).Error
}

func (s *ListSongService) UpdateListSongMsg(listSong *model.ListSong) error {
	return global.DB.Model(listSong).Updates(listSong).Error
}

func (s *ListSongService) DeleteListSong(id uint) error {
	return global.DB.Delete(&model.ListSong{}, id).Error
}

func (s *ListSongService) AllListSong() ([]model.ListSong, error) {
	var listSongs []model.ListSong
	err := global.DB.Find(&listSongs).Error
	return listSongs, err
}

func (s *ListSongService) ListSongOfSongId(songId uint) ([]model.ListSong, error) {
	var listSongs []model.ListSong
	err := global.DB.Where("song_id = ?", songId).Find(&listSongs).Error
	return listSongs, err
}

func (s *ListSongService) ListSongOfSongListId(songListId uint) ([]model.ListSong, error) {
	var listSongs []model.ListSong
	err := global.DB.Where("song_list_id = ?", songListId).Find(&listSongs).Error
	return listSongs, err
}

var ListSongServiceApp = new(ListSongService)
