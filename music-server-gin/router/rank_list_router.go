package router

import (
	"music-server-gin/model"
	"music-server-gin/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RankListRouter struct {
	service *service.RankListService
}

var RankListRouterApp = &RankListRouter{
	service: service.RankListServiceApp,
}

func (r *RankListRouter) AddRank(ctx *gin.Context) {
	var rank model.RankList
	if err := ctx.ShouldBindJSON(&rank); err != nil {
		ctx.JSON(400, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	if err := r.service.AddRank(&rank); err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": "添加失败"})
		return
	}
	ctx.JSON(200, gin.H{"code": 200, "msg": "添加成功"})
}

func (r *RankListRouter) RankOfSongListId(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	rank, err := r.service.RankOfSongListId(uint(id))
	if err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": "获取失败"})
		return
	}
	ctx.JSON(200, gin.H{"code": 200, "data": rank})
}

func (r *RankListRouter) RankOfConsumerId(ctx *gin.Context) {
	songListId, err := strconv.Atoi(ctx.Query("songListId"))
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	consumerId, err := strconv.Atoi(ctx.Query("consumerId"))
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	rank, err := r.service.GetUserRank(uint(songListId), uint(consumerId))
	if err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": "获取失败"})
		return
	}
	ctx.JSON(200, gin.H{"code": 200, "data": rank})
}

func (r *RankListRouter) GetAverageScore(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	avg, err := r.service.GetAverageScore(uint(id))
	if err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": "获取失败"})
		return
	}
	ctx.JSON(200, gin.H{"code": 200, "data": gin.H{"avg": avg}})
}
