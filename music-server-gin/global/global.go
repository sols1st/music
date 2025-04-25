package global

import (
	"music-server-gin/config"

	"github.com/jinzhu/gorm"
)

var (
	DB           *gorm.DB
	CONFIG       config.ServerConfig
	UploadConfig config.UploadConfig
)
