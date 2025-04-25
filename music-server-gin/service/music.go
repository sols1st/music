package service

import (
	"errors"

	"github.com/MephistoSolsist/mysql-practice/global"
	"github.com/MephistoSolsist/mysql-practice/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MusicService struct{}

func (*MusicService) GetMusicList(list *[]model.Music) error {
	err := global.DB.Raw("select music_name,music.music_id,album_name,album.album_id,artist.artist_id,artist_name from music,album,artist where ifnull(music.artist_id,0) = artist.artist_id and IFNULL(music.album_id,0)=album.album_id").Order("music_id ASC").Scan(&list)
	if err != nil {
		return err.Error
	}
	return nil
}

func (*MusicService) GetMusicById(id int) (model.Music, error) {
	var m model.Music
	err := global.DB.Raw("select music_name,music.music_id,album_name,album.album_id,artist.artist_id,artist_name from music,album,artist where ifnull(music.artist_id,0) = artist.artist_id and IFNULL(music.album_id,0)=album.album_id and music_id = ?", id).Scan(&m)
	if err != nil {
		return m, err.Error
	}
	return m, nil
}

func (*MusicService) DeleteMusic(id int) error {
	m := model.MusicDB{MusicId: id}
	err := global.DB.Delete(&m)
	if err != nil {
		return err.Error
	}
	return nil
}

func (*MusicService) UploadMusic(m *model.Music) error {
	var a model.Artist
	var al model.Album
	err := global.DB.Where("artist_name = ?", m.ArtistName).First(&a).Error
	a.ArtistName = m.ArtistName
	if errors.Is(err, gorm.ErrRecordNotFound) {
		global.DB.Create(&a)
	}
	err = global.DB.Where("artist_name = ?", m.ArtistName).First(&a).Error
	if err != nil {
		return err
	}
	err = global.DB.Where("album_name = ?", m.AlbumName).First(&al).Error
	al.AlbumName = m.AlbumName
	if errors.Is(err, gorm.ErrRecordNotFound) {
		global.DB.Create(&al)
	}
	err = global.DB.Where("album_name = ?", m.AlbumName).First(&al).Error
	if err != nil {
		return err
	}
	var music model.MusicDB
	{
		music.AlbumId = al.AlbumId
		music.ArtistId = a.ArtistId
		music.MusicName = m.MusicName
	}
	err = global.DB.Create(&music).Error
	if err != nil {
		return err
	}
	return nil
}

func (*MusicService) UpdateMusic(m *model.Music) error {
	var a model.Artist
	var al model.Album
	err := global.DB.Where("artist_name = ?", m.ArtistName).First(&a).Error
	a.ArtistName = m.ArtistName
	if errors.Is(err, gorm.ErrRecordNotFound) {
		global.DB.Create(&a)
	}
	err = global.DB.Where("artist_name = ?", m.ArtistName).First(&a).Error
	if err != nil {
		return err
	}
	err = global.DB.Where("album_name = ?", m.AlbumName).First(&al).Error
	al.AlbumName = m.AlbumName
	if errors.Is(err, gorm.ErrRecordNotFound) {
		global.DB.Create(&al)
	}
	err = global.DB.Where("album_name = ?", m.AlbumName).First(&al).Error
	if err != nil {
		return err
	}
	var music model.MusicDB
	{
		music.MusicId = m.MusicId
		music.AlbumId = al.AlbumId
		music.ArtistId = a.ArtistId
		music.MusicName = m.MusicName
	}
	global.DB.Save(music)
	return nil
}

var MusicServiceApp = new(MusicService)
