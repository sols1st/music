package service

import (
	"music-server-gin/model"

	"github.com/jinzhu/gorm"
)

type ConfigService struct {
	db *gorm.DB
}

func NewConfigService(db *gorm.DB) *ConfigService {
	return &ConfigService{db: db}
}

func (s *ConfigService) GetConfig() (*model.Config, error) {
	var config model.Config
	err := s.db.First(&config).Error
	if err == gorm.ErrRecordNotFound {
		// 如果配置不存在，创建默认配置
		config = model.Config{
			MaxLoginAttempts: 5,
			LockoutDuration:  30, // 分钟
		}
		err = s.db.Create(&config).Error
	}
	return &config, err
}

func (s *ConfigService) UpdateConfig(config *model.Config) error {
	return s.db.Save(config).Error
}
