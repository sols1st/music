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

func (c *CollectRouter) AddCollect(ctx *gin.Context) {
	var collect model.Collect
	if err := ctx.ShouldBindJSON(&collect); err != nil {
		ctx.JSON(400, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	if err := c.service.AddCollect(&collect); err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": "添加失败"})
		return
	}
	ctx.JSON(200, gin.H{"code": 200, "msg": "添加成功"})
}

func (c *CollectRouter) DeleteCollect(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	if err := c.service.DeleteCollect(uint(id)); err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": "删除失败"})
		return
	}
	ctx.JSON(200, gin.H{"code": 200, "msg": "删除成功"})
}

func (c *CollectRouter) CollectOfUserId(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Query("userId"))
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	collects, err := c.service.CollectOfUserId(uint(userId))
	if err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": "获取失败"})
		return
	}
	ctx.JSON(200, gin.H{"code": 200, "data": collects})
}

func (c *CollectRouter) IsCollect(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Query("userId"))
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	songId, err := strconv.Atoi(ctx.Query("songId"))
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	collects, err := c.service.CollectOfUserId(uint(userId))
	if err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": "获取失败"})
		return
	}
	for _, collect := range collects {
		if collect.SongID == uint(songId) {
			ctx.JSON(200, gin.H{"code": 200, "data": true})
			return
		}
	}
	ctx.JSON(200, gin.H{"code": 200, "data": false})
}

func (s *CollectRouter) AllCollect(c *gin.Context) {
	collects, err := service.CollectServiceApp.AllCollect()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, collects)
}

func (s *CollectRouter) CollectOfSongId(c *gin.Context) {
	songId, err := strconv.Atoi(c.Query("songId"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid songId"})
		return
	}
	collects, err := service.CollectServiceApp.CollectOfSongId(uint(songId))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, collects)
}
