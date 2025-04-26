package router

import (
	"music-server-gin/model"
	"music-server-gin/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ListSongRouter struct {
	service *service.ListSongService
}

var ListSongRouterApp = &ListSongRouter{
	service: service.ListSongServiceApp,
}

func (l *ListSongRouter) AddListSong(ctx *gin.Context) {
	var listSong model.ListSong
	if err := ctx.ShouldBindJSON(&listSong); err != nil {
		ctx.JSON(400, BadRequest("参数错误"))
		return
	}
	if err := l.service.AddListSong(&listSong); err != nil {
		ctx.JSON(500, Error("添加失败"))
		return
	}
	ctx.JSON(200, Success("添加成功"))
}

func (l *ListSongRouter) DeleteListSong(ctx *gin.Context) {
	songId, err := strconv.Atoi(ctx.Query("songId"))
	if err != nil {
		ctx.JSON(400, BadRequest("参数错误"))
		return
	}
	if err := l.service.DeleteListSong(uint(songId)); err != nil {
		ctx.JSON(500, Error("删除失败"))
		return
	}
	ctx.JSON(200, Success("删除成功"))
}

func (l *ListSongRouter) ListSongOfSongId(ctx *gin.Context) {
	songListId, err := strconv.Atoi(ctx.Query("songListId"))
	if err != nil {
		ctx.JSON(400, BadRequest("参数错误"))
		return
	}
	listSongs, err := l.service.ListSongOfSongListId(uint(songListId))
	if err != nil {
		ctx.JSON(500, Error("获取失败"))
		return
	}
	ctx.JSON(200, Success(listSongs))
}

func (l *ListSongRouter) UpdateListSongMsg(ctx *gin.Context) {
	var listSong model.ListSong
	if err := ctx.ShouldBindJSON(&listSong); err != nil {
		ctx.JSON(400, BadRequest("参数错误"))
		return
	}
	if err := l.service.UpdateListSongMsg(&listSong); err != nil {
		ctx.JSON(500, Error("更新失败"))
		return
	}
	ctx.JSON(200, Success("更新成功"))
}

func (l *ListSongRouter) GetExcle(ctx *gin.Context) {
	// TODO: 实现导出 Excel 的功能
	ctx.JSON(200, Success("导出成功"))
}

func (s *ListSongRouter) AllListSong(c *gin.Context) {
	listSongs, err := service.ListSongServiceApp.AllListSong()
	if err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	c.JSON(200, Success(listSongs))
}

func (s *ListSongRouter) ListSongOfSongListId(c *gin.Context) {
	songListId, err := strconv.Atoi(c.Query("songListId"))
	if err != nil {
		c.JSON(400, BadRequest("无效的歌单ID"))
		return
	}
	listSongs, err := service.ListSongServiceApp.ListSongOfSongListId(uint(songListId))
	if err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	c.JSON(200, Success(listSongs))
}
