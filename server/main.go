package main

import (
	"log"
	"music-server-gin/config"
	"music-server-gin/global"
	"music-server-gin/router"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 确保所有必要的目录存在
// func ensureDirectories(config config.UploadConfig) error {
// 	dirs := []string{
// 		config.UploadDir,
// 		config.AvatarDir,
// 		config.SongDir,
// 		config.SongPicDir,
// 		config.SongLrcDir,
// 		config.SingerPicDir,
// 		config.SongListPicDir,
// 		config.BannerPicDir,
// 	}

// 	for _, dir := range dirs {
// 		if dir != "" {
// 			if err := os.MkdirAll(dir, 0755); err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	return nil
// }

func main() {
	// 初始化配置
	serverConfig, err := config.InitConfig()
	if err != nil {
		log.Fatalf("初始化配置失败: %v", err)
	}
	global.CONFIG = *serverConfig
	global.UploadConfig = serverConfig.UploadConfig

	// 确保上传目录存在
	// if err := ensureDirectories(serverConfig.UploadConfig); err != nil {
	// 	log.Fatalf("创建上传目录失败: %v", err)
	// }

	// 初始化数据库连接
	dsn := global.CONFIG.MySqlConfig.Dsn()
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	global.DB = db

	// log.Printf("配置信息: %+v", serverConfig)
	// log.Printf("上传目录: %s", serverConfig.UploadConfig.UploadDir)
	// log.Printf("Banner目录: %s", serverConfig.UploadConfig.BannerPicDir)

	// 创建路由
	r := router.SetupRouter()
	// 启动服务器
	r.Run(":" + global.CONFIG.SystemConfig.Port)
}
