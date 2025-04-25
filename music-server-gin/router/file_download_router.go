package router

import (
	"io"
	"music-server-gin/global"
	"music-server-gin/service"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FileDownloadRouter struct {
	service *service.FileDownloadService
}

var FileDownloadRouterApp = &FileDownloadRouter{
	service: service.FileDownloadServiceApp,
}

func (f *FileDownloadRouter) DownloadFile(ctx *gin.Context) {
	fileName := ctx.Param("fileName")
	if fileName == "" {
		ctx.JSON(400, gin.H{"code": 400, "msg": "文件名不能为空"})
		return
	}

	filePath := filepath.Join(global.UploadConfig.UploadDir, fileName)
	file, err := os.Open(filePath)
	if err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": "文件不存在"})
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": "获取文件信息失败"})
		return
	}

	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", "attachment; filename="+fileName)
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

	_, err = io.Copy(ctx.Writer, file)
	if err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": "文件下载失败"})
		return
	}
}
