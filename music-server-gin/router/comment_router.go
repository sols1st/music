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
		ctx.JSON(400, BadRequest("参数错误"))
		return
	}
	if err := c.service.AddComment(&comment); err != nil {
		ctx.JSON(500, Error("添加失败"))
		return
	}
	ctx.JSON(200, Success("添加成功"))
}

func (c *CommentRouter) DeleteComment(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(400, BadRequest("参数错误"))
		return
	}
	if err := c.service.DeleteComment(uint(id)); err != nil {
		ctx.JSON(500, Error("删除失败"))
		return
	}
	ctx.JSON(200, Success("删除成功"))
}

func (c *CommentRouter) CommentOfSongId(ctx *gin.Context) {
	songId, err := strconv.Atoi(ctx.Query("songId"))
	if err != nil {
		ctx.JSON(400, BadRequest("参数错误"))
		return
	}
	comments, err := c.service.CommentOfSongId(uint(songId))
	if err != nil {
		ctx.JSON(500, Error("获取失败"))
		return
	}
	ctx.JSON(200, Success(comments))
}

func (c *CommentRouter) CommentOfSongListId(ctx *gin.Context) {
	songListId, err := strconv.Atoi(ctx.Query("songListId"))
	if err != nil {
		ctx.JSON(400, BadRequest("参数错误"))
		return
	}
	comments, err := c.service.CommentOfSongListId(uint(songListId))
	if err != nil {
		ctx.JSON(500, Error("获取失败"))
		return
	}
	ctx.JSON(200, Success(comments))
}

func (c *CommentRouter) UpdateCommentMsg(ctx *gin.Context) {
	var comment model.Comment
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(400, BadRequest("参数错误"))
		return
	}
	if err := c.service.UpdateCommentMsg(&comment); err != nil {
		ctx.JSON(500, Error("更新失败"))
		return
	}
	ctx.JSON(200, Success("更新成功"))
}

func (s *CommentRouter) AllComment(c *gin.Context) {
	comments, err := service.CommentServiceApp.AllComment()
	if err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	c.JSON(200, Success(comments))
}
