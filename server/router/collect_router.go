package router

import (
	"music-server-gin/model"
	"music-server-gin/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CollectRouter struct {
	service *service.CollectService
}

var CollectRouterApp = &CollectRouter{
	service: service.CollectServiceApp,
}

// IsCollection 检查是否已收藏
func (c *CollectRouter) IsCollection(ctx *gin.Context) {
	var request struct {
		UserID uint   `json:"userId"`
		Type   string `json:"type"`
		SongID uint   `json:"songId"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, BadRequest("参数错误: "+err.Error()))
		return
	}

	isCollect, err := c.service.IsCollection(request.UserID, request.Type, request.SongID)
	if err != nil {
		ctx.JSON(500, Error("检查失败: "+err.Error()))
		return
	}

	ctx.JSON(200, Success(isCollect))
}

// AddCollection 添加收藏
func (c *CollectRouter) AddCollection(ctx *gin.Context) {
	var collect model.Collect
	if err := ctx.ShouldBindJSON(&collect); err != nil {
		ctx.JSON(400, BadRequest("参数错误: "+err.Error()))
		return
	}

	if err := c.service.AddCollection(&collect); err != nil {
		ctx.JSON(500, Error("添加失败: "+err.Error()))
		return
	}
	ctx.JSON(200, Success("添加成功"))
}

// DeleteCollection 删除收藏
func (c *CollectRouter) DeleteCollection(ctx *gin.Context) {
	var request struct {
		UserID uint   `json:"userId"`
		Type   string `json:"type"`
		SongID uint   `json:"songId"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, BadRequest("参数错误: "+err.Error()))
		return
	}

	if err := c.service.DeleteCollection(request.UserID, request.Type, request.SongID); err != nil {
		ctx.JSON(500, Error("删除失败: "+err.Error()))
		return
	}
	ctx.JSON(200, Success("删除成功"))
}

// CollectionOfUser 获取用户的收藏
func (c *CollectRouter) CollectionOfUser(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Query("userId"))
	if err != nil {
		ctx.JSON(400, BadRequest("参数错误: "+err.Error()))
		return
	}

	collects, err := c.service.CollectionOfUser(uint(userId))
	if err != nil {
		ctx.JSON(500, Error("获取失败: "+err.Error()))
		return
	}
	ctx.JSON(200, Success(collects))
}

func (c *CollectRouter) CollectOfUserId(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Query("userId"))
	if err != nil {
		ctx.JSON(400, BadRequest("参数错误"))
		return
	}
	collects, err := c.service.CollectOfUserId(uint(userId))
	if err != nil {
		ctx.JSON(500, Error("获取失败"))
		return
	}
	ctx.JSON(200, Success(collects))
}

func (c *CollectRouter) IsCollect(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Query("userId"))
	if err != nil {
		ctx.JSON(400, BadRequest("参数错误"))
		return
	}
	songId, err := strconv.Atoi(ctx.Query("songId"))
	if err != nil {
		ctx.JSON(400, BadRequest("参数错误"))
		return
	}
	collects, err := c.service.CollectOfUserId(uint(userId))
	if err != nil {
		ctx.JSON(500, Error("获取失败"))
		return
	}
	for _, collect := range collects {
		if collect.SongID == uint(songId) {
			ctx.JSON(200, Success(true))
			return
		}
	}
	ctx.JSON(200, Success(false))
}

func (s *CollectRouter) AllCollect(c *gin.Context) {
	collects, err := service.CollectServiceApp.AllCollect()
	if err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	c.JSON(200, Success(collects))
}

func (s *CollectRouter) CollectOfSongId(c *gin.Context) {
	songId, err := strconv.Atoi(c.Query("songId"))
	if err != nil {
		c.JSON(400, BadRequest("无效的歌曲ID"))
		return
	}
	collects, err := service.CollectServiceApp.CollectOfSongId(uint(songId))
	if err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	c.JSON(200, Success(collects))
}
