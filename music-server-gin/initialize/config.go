package initialize

import (
	"fmt"
	"path/filepath"

	"music-server-gin/config"
	"music-server-gin/global"
	"music-server-gin/util"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

func ReadConfig() config.ServerConfig {
	v := viper.New()

	v.SetConfigFile("config.yaml")

	if err := v.ReadInConfig(); err != nil {
		fmt.Println("config read failed", err)
		util.Logger.Fatal(err)
	}
	serverConfig := config.ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		fmt.Println("config unmarshal failed", err)
		util.Logger.Fatal(err)
	}

	// 设置默认上传目录
	if serverConfig.UploadConfig.UploadDir == "" {
		serverConfig.UploadConfig.UploadDir = "uploads"
	}

	// 设置各类文件的上传目录
	serverConfig.UploadConfig.AvatarDir = filepath.Join(serverConfig.UploadConfig.UploadDir, "avatar")
	serverConfig.UploadConfig.SongDir = filepath.Join(serverConfig.UploadConfig.UploadDir, "song")
	serverConfig.UploadConfig.SongPicDir = filepath.Join(serverConfig.UploadConfig.UploadDir, "song_pic")
	serverConfig.UploadConfig.SongLrcDir = filepath.Join(serverConfig.UploadConfig.UploadDir, "song_lrc")
	serverConfig.UploadConfig.SingerPicDir = filepath.Join(serverConfig.UploadConfig.UploadDir, "singer_pic")
	serverConfig.UploadConfig.SongListPicDir = filepath.Join(serverConfig.UploadConfig.UploadDir, "song_list_pic")
	serverConfig.UploadConfig.BannerPicDir = filepath.Join(serverConfig.UploadConfig.UploadDir, "banner_pic")

	// 初始化上传目录
	global.UploadConfig = serverConfig.UploadConfig

	return serverConfig
}
