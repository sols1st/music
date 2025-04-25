package router

import (
	"music-server-gin/model"
	"music-server-gin/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BannerRouter struct {
	service *service.BannerService
}

var BannerRouterApp = &BannerRouter{
	service: service.BannerServiceApp,
}

func (s *BannerRouter) AddBanner(ctx *gin.Context) {
	var banner model.Banner
	if err := ctx.ShouldBindJSON(&banner); err != nil {
		ctx.JSON(400, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	if err := s.service.AddBanner(&banner); err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": "添加失败"})
		return
	}
	ctx.JSON(200, gin.H{"code": 200, "msg": "添加成功"})
}

func (s *BannerRouter) UpdateBannerMsg(ctx *gin.Context) {
	var banner model.Banner
	if err := ctx.ShouldBindJSON(&banner); err != nil {
		ctx.JSON(400, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	if err := s.service.UpdateBannerMsg(&banner); err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": "更新失败"})
		return
	}
	ctx.JSON(200, gin.H{"code": 200, "msg": "更新成功"})
}

func (s *BannerRouter) UpdateBannerPic(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.PostForm("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	file, err := ctx.FormFile("pic")
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "msg": "文件上传失败"})
		return
	}
	fileData, err := file.Open()
	if err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": "文件读取失败"})
		return
	}
	defer fileData.Close()

	fileBytes := make([]byte, file.Size)
	if _, err := fileData.Read(fileBytes); err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": "文件读取失败"})
		return
	}

	if err := s.service.UpdateBannerPic(uint(id), string(fileBytes)); err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": "更新失败"})
		return
	}
	ctx.JSON(200, gin.H{"code": 200, "msg": "更新成功"})
}

func (s *BannerRouter) DeleteBanner(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	if err := s.service.DeleteBanner(uint(id)); err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": "删除失败"})
		return
	}
	ctx.JSON(200, gin.H{"code": 200, "msg": "删除成功"})
}

func (s *BannerRouter) AllBanner(ctx *gin.Context) {
	banners, err := s.service.AllBanner()
	if err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": "获取失败"})
		return
	}
	ctx.JSON(200, gin.H{"code": 200, "data": banners})
}
