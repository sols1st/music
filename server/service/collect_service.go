package service

import (
	"errors"
	"music-server-gin/global"
	"music-server-gin/model"
	"time"
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

// IsCollection 检查是否已收藏
func (s *CollectService) IsCollection(userId uint, typeStr string, songId uint) (bool, error) {
	var collect model.Collect
	err := global.DB.Where("user_id = ? AND type = ? AND song_id = ?", userId, typeStr, songId).First(&collect).Error
	if err != nil {
		return false, nil
	}
	return true, nil
}

// AddCollection 添加收藏
func (s *CollectService) AddCollection(collect *model.Collect) error {
	// 验证必填字段
	if collect.UserID == 0 {
		return errors.New("用户ID不能为空")
	}
	if collect.Type == 0 {
		return errors.New("收藏类型不能为空")
	}
	if collect.Type == 0 && collect.SongID == 0 {
		return errors.New("歌曲ID不能为空")
	}
	if collect.Type == 1 && collect.SongListID == 0 {
		return errors.New("歌单ID不能为空")
	}

	// 检查是否已收藏
	exists, err := s.IsCollection(collect.UserID, string(collect.Type), collect.SongID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("已经收藏过了")
	}

	// 设置创建时间
	collect.CreateTime = time.Now()

	// 使用事务
	tx := global.DB.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(collect).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// DeleteCollection 删除收藏
func (s *CollectService) DeleteCollection(userId uint, typeStr string, songId uint) error {
	// 使用事务
	tx := global.DB.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Where("user_id = ? AND type = ? AND song_id = ?", userId, typeStr, songId).Delete(&model.Collect{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// CollectionOfUser 获取用户的收藏
func (s *CollectService) CollectionOfUser(userId uint) ([]model.Collect, error) {
	var collects []model.Collect
	err := global.DB.Where("user_id = ?", userId).Find(&collects).Error
	if err != nil {
		return nil, err
	}
	return collects, nil
}
