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

// AddComment 添加评论
func (c *CommentRouter) AddComment(ctx *gin.Context) {
	// 定义请求结构体
	type CommentRequest struct {
		Content    string `json:"content"`
		SongID     *uint  `json:"songId"`
		SongListID *uint  `json:"songListId"`
		NowType    byte   `json:"nowType"`
		UserID     *uint  `json:"userId"`
	}

	var req CommentRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, BadRequest("参数错误: "+err.Error()))
		return
	}

	// 验证必填字段
	if req.Content == "" {
		ctx.JSON(400, BadRequest("评论内容不能为空"))
		return
	}

	// 构建评论对象
	comment := &model.Comment{
		Content: req.Content,
		Type:    req.NowType,
	}

	// 根据类型设置相应的ID
	if req.NowType == 0 {
		comment.SongID = *req.SongID
		comment.SongListID = 0
	} else {
		comment.SongListID = *req.SongListID
		comment.SongID = 0
	}

	// 从上下文中获取用户ID
	if req.UserID == nil {
		ctx.JSON(401, Unauthorized("用户未登录"))
		return
	}
	comment.UserID = *req.UserID

	if err := c.service.AddComment(comment); err != nil {
		ctx.JSON(500, Error("添加失败: "+err.Error()))
		return
	}

	ctx.JSON(200, Success("添加成功"))
}

// DeleteComment 删除评论
func (c *CommentRouter) DeleteComment(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(400, BadRequest("参数错误: "+err.Error()))
		return
	}

	// 验证用户权限
	userId, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(401, Unauthorized("用户未登录"))
		return
	}

	// 获取评论信息
	comment, err := c.service.CommentOfId(uint(id))
	if err != nil {
		ctx.JSON(500, Error("获取评论失败: "+err.Error()))
		return
	}

	// 验证是否是评论作者
	if comment.UserID != userId.(uint) {
		ctx.JSON(403, Forbidden("无权删除此评论"))
		return
	}

	if err := c.service.DeleteComment(uint(id)); err != nil {
		ctx.JSON(500, Error("删除失败: "+err.Error()))
		return
	}
	ctx.JSON(200, Success("删除成功"))
}

// CommentOfSongId 获取歌曲评论
func (c *CommentRouter) CommentOfSongId(ctx *gin.Context) {
	songId, err := strconv.Atoi(ctx.Query("songId"))
	if err != nil {
		ctx.JSON(400, BadRequest("参数错误: "+err.Error()))
		return
	}
	comments, err := c.service.CommentOfSongId(uint(songId))
	if err != nil {
		ctx.JSON(500, Error("获取失败: "+err.Error()))
		return
	}
	ctx.JSON(200, Success(comments))
}

// CommentOfSongListId 获取歌单评论
func (c *CommentRouter) CommentOfSongListId(ctx *gin.Context) {
	songListId, err := strconv.Atoi(ctx.Query("songListId"))
	if err != nil {
		ctx.JSON(400, BadRequest("参数错误: "+err.Error()))
		return
	}
	comments, err := c.service.CommentOfSongListId(uint(songListId))
	if err != nil {
		ctx.JSON(500, Error("获取失败: "+err.Error()))
		return
	}
	ctx.JSON(200, Success(comments))
}

// UpdateCommentMsg 更新评论
func (c *CommentRouter) UpdateCommentMsg(ctx *gin.Context) {
	var comment model.Comment
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(400, BadRequest("参数错误: "+err.Error()))
		return
	}

	// 验证用户权限
	userId, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(401, Unauthorized("用户未登录"))
		return
	}

	// 获取原评论信息
	oldComment, err := c.service.CommentOfId(comment.ID)
	if err != nil {
		ctx.JSON(500, Error("获取评论失败: "+err.Error()))
		return
	}

	// 验证是否是评论作者
	if oldComment.UserID != userId.(uint) {
		ctx.JSON(403, Forbidden("无权修改此评论"))
		return
	}

	if err := c.service.UpdateCommentMsg(&comment); err != nil {
		ctx.JSON(500, Error("更新失败: "+err.Error()))
		return
	}
	ctx.JSON(200, Success("更新成功"))
}

// AllComment 获取所有评论
func (s *CommentRouter) AllComment(c *gin.Context) {
	comments, err := service.CommentServiceApp.AllComment()
	if err != nil {
		c.JSON(500, Error("获取失败: "+err.Error()))
		return
	}
	c.JSON(200, Success(comments))
}
