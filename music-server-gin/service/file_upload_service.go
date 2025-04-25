package service

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"music-server-gin/global"
)

type FileUploadService interface {
	UploadAvatar(file *multipart.FileHeader) (string, error)
	UploadSong(file *multipart.FileHeader) (string, error)
	UploadSongPic(file *multipart.FileHeader) (string, error)
	UploadSongLrc(file *multipart.FileHeader) (string, error)
	UploadSingerPic(file *multipart.FileHeader) (string, error)
	UploadSongListPic(file *multipart.FileHeader) (string, error)
	UploadBannerPic(file *multipart.FileHeader) (string, error)
	GetFileUrl(filename string) string
}

type fileUploadService struct{}

func NewFileUploadService() FileUploadService {
	return &fileUploadService{}
}

func (s *fileUploadService) saveFile(file *multipart.FileHeader, dir string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// 创建目标目录
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	// 生成文件名
	ext := filepath.Ext(file.Filename)
	filename := time.Now().Format("20060102150405") + ext
	dstPath := filepath.Join(dir, filename)

	// 创建目标文件
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// 复制文件内容
	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return filename, nil
}

func (s *fileUploadService) UploadAvatar(file *multipart.FileHeader) (string, error) {
	return s.saveFile(file, global.UploadConfig.AvatarDir)
}

func (s *fileUploadService) UploadSong(file *multipart.FileHeader) (string, error) {
	return s.saveFile(file, global.UploadConfig.SongDir)
}

func (s *fileUploadService) UploadSongPic(file *multipart.FileHeader) (string, error) {
	return s.saveFile(file, global.UploadConfig.SongPicDir)
}

func (s *fileUploadService) UploadSongLrc(file *multipart.FileHeader) (string, error) {
	return s.saveFile(file, global.UploadConfig.SongLrcDir)
}

func (s *fileUploadService) UploadSingerPic(file *multipart.FileHeader) (string, error) {
	return s.saveFile(file, global.UploadConfig.SingerPicDir)
}

func (s *fileUploadService) UploadSongListPic(file *multipart.FileHeader) (string, error) {
	return s.saveFile(file, global.UploadConfig.SongListPicDir)
}

func (s *fileUploadService) UploadBannerPic(file *multipart.FileHeader) (string, error) {
	return s.saveFile(file, global.UploadConfig.BannerPicDir)
}

func (s *fileUploadService) GetFileUrl(filename string) string {
	if filename == "" {
		return ""
	}
	return "/uploads/" + filename
}
