package router

import (
	"music-server-gin/model"
	"music-server-gin/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentRouter struct {
	service *service.CommentService
}

var CommentRouterApp = &CommentRouter{
	service: service.CommentServiceApp,
}

func (c *CommentRouter) AddComment(ctx *gin.Context) {
	var comment model.Comment
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(400, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	if err := c.service.AddComment(&comment); err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": "添加失败"})
		return
	}
	ctx.JSON(200, gin.H{"code": 200, "msg": "添加成功"})
}

func (c *CommentRouter) DeleteComment(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	if err := c.service.DeleteComment(uint(id)); err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": "删除失败"})
		return
	}
	ctx.JSON(200, gin.H{"code": 200, "msg": "删除成功"})
}

func (c *CommentRouter) CommentOfSongId(ctx *gin.Context) {
	songId, err := strconv.Atoi(ctx.Query("songId"))
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	comments, err := c.service.CommentOfSongId(uint(songId))
	if err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": "获取失败"})
		return
	}
	ctx.JSON(200, gin.H{"code": 200, "data": comments})
}

func (c *CommentRouter) CommentOfSongListId(ctx *gin.Context) {
	songListId, err := strconv.Atoi(ctx.Query("songListId"))
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	comments, err := c.service.CommentOfSongListId(uint(songListId))
	if err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": "获取失败"})
		return
	}
	ctx.JSON(200, gin.H{"code": 200, "data": comments})
}

func (c *CommentRouter) UpdateCommentMsg(ctx *gin.Context) {
	var comment model.Comment
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(400, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	if err := c.service.UpdateCommentMsg(&comment); err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": "更新失败"})
		return
	}
	ctx.JSON(200, gin.H{"code": 200, "msg": "更新成功"})
}

func (s *CommentRouter) AllComment(c *gin.Context) {
	comments, err := service.CommentServiceApp.AllComment()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, comments)
}
