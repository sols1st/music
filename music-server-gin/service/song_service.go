package service

import (
	"music-server-gin/global"
	"music-server-gin/model"
)

type SongService struct{}

func (s *SongService) AddSong(song *model.Song) error {
	return global.DB.Create(song).Error
}

func (s *SongService) UpdateSongMsg(song *model.Song) error {
	return global.DB.Model(song).Updates(song).Error
}

func (s *SongService) UpdateSongPic(id uint, pic []byte) error {
	return global.DB.Model(&model.Song{}).Where("id = ?", id).Update("pic", pic).Error
}

func (s *SongService) UpdateSongUrl(id uint, url []byte) error {
	return global.DB.Model(&model.Song{}).Where("id = ?", id).Update("url", url).Error
}

func (s *SongService) UpdateSongLrc(id uint, lrc []byte) error {
	return global.DB.Model(&model.Song{}).Where("id = ?", id).Update("lyric", lrc).Error
}

func (s *SongService) DeleteSong(id uint) error {
	return global.DB.Delete(&model.Song{}, id).Error
}

func (s *SongService) AllSong() ([]model.Song, error) {
	var songs []model.Song
	err := global.DB.Find(&songs).Error
	return songs, err
}

func (s *SongService) SongOfId(id uint) (model.Song, error) {
	var song model.Song
	err := global.DB.First(&song, id).Error
	return song, err
}

func (s *SongService) SongOfSingerId(singerId uint) ([]model.Song, error) {
	var songs []model.Song
	err := global.DB.Where("singer_id = ?", singerId).Find(&songs).Error
	return songs, err
}

func (s *SongService) SongOfSingerName(name string) ([]model.Song, error) {
	var songs []model.Song
	err := global.DB.Joins("JOIN singers ON songs.singer_id = singers.id").
		Where("singers.name = ?", name).
		Find(&songs).Error
	return songs, err
}

var SongServiceApp = new(SongService)
