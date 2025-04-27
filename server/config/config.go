package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db-name"`
	Config   string `mapstructure:"config"`
}

func (m *MySQLConfig) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		m.Username, m.Password, m.Host, m.Port, m.DBName, m.Config)
}

type SystemConfig struct {
	Port string `mapstructure:"port"`
}

type UploadConfig struct {
	UploadDir      string `mapstructure:"upload_dir"`
	AvatarDir      string `mapstructure:"avatar_dir"`
	SongDir        string `mapstructure:"song_dir"`
	SongPicDir     string `mapstructure:"song_pic_dir"`
	SongLrcDir     string `mapstructure:"song_lrc_dir"`
	SingerPicDir   string `mapstructure:"singer_pic_dir"`
	SongListPicDir string `mapstructure:"song_list_pic_dir"`
	BannerPicDir   string `mapstructure:"banner_pic_dir"`
}

type ServerConfig struct {
	MySqlConfig  MySQLConfig  `mapstructure:"mysql"`
	SystemConfig SystemConfig `mapstructure:"system"`
	UploadConfig UploadConfig `mapstructure:"upload"`
}

func InitConfig() (*ServerConfig, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}

	var config ServerConfig
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件发生变化:", e.Name)
		if err := v.Unmarshal(&config); err != nil {
			fmt.Printf("重新加载配置文件失败: %v\n", err)
		}
	})

	return &config, nil
}
