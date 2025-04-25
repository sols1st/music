package service

import (
	"music-server-gin/model"
	"time"

	"github.com/jinzhu/gorm"
)

type ResetPasswordService struct {
	db *gorm.DB
}

func NewResetPasswordService(db *gorm.DB) *ResetPasswordService {
	return &ResetPasswordService{db: db}
}

func (s *ResetPasswordService) CreateResetRequest(request *model.ResetPasswordRequest) error {
	return s.db.Create(request).Error
}

func (s *ResetPasswordService) GetResetRequestByToken(token string) (*model.ResetPasswordRequest, error) {
	var request model.ResetPasswordRequest
	err := s.db.Where("token = ? AND expires_at > ?", token, time.Now()).First(&request).Error
	return &request, err
}

func (s *ResetPasswordService) DeleteResetRequest(id uint) error {
	return s.db.Delete(&model.ResetPasswordRequest{}, id).Error
}

func (s *ResetPasswordService) DeleteExpiredRequests() error {
	return s.db.Where("expires_at < ?", time.Now()).Delete(&model.ResetPasswordRequest{}).Error
}
