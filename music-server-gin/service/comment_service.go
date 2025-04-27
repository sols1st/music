package service

import (
	"music-server-gin/global"
	"music-server-gin/model"
	"time"
)

type CommentService struct{}

func (s *CommentService) AddComment(comment *model.Comment) error {
	// 添加createTime
	comment.CreateTime = time.Now()
	return global.DB.Create(comment).Error
}

func (s *CommentService) UpdateCommentMsg(comment *model.Comment) error {
	return global.DB.Model(comment).Updates(comment).Error
}

func (s *CommentService) DeleteComment(id uint) error {
	return global.DB.Delete(&model.Comment{}, id).Error
}

func (s *CommentService) AllComment() ([]model.Comment, error) {
	var comments []model.Comment
	err := global.DB.Find(&comments).Error
	return comments, err
}

func (s *CommentService) CommentOfId(id uint) (model.Comment, error) {
	var comment model.Comment
	err := global.DB.First(&comment, id).Error
	return comment, err
}

func (s *CommentService) CommentOfSongId(songId uint) ([]model.Comment, error) {
	var comments []model.Comment
	err := global.DB.Where("song_id = ?", songId).Find(&comments).Error
	return comments, err
}

func (s *CommentService) CommentOfSongListId(songListId uint) ([]model.Comment, error) {
	var comments []model.Comment
	err := global.DB.Where("song_list_id = ?", songListId).Find(&comments).Error
	return comments, err
}

var CommentServiceApp = new(CommentService)
