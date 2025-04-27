package router

import (
	"music-server-gin/model"
	"music-server-gin/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SongListRouter struct{}

func (s *SongListRouter) AddSongList(c *gin.Context) {
	var songList model.SongList
	if err := c.ShouldBindJSON(&songList); err != nil {
		c.JSON(400, BadRequest(err.Error()))
		return
	}
	if err := service.SongListServiceApp.AddSongList(&songList); err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	c.JSON(200, Success(songList))
}

func (s *SongListRouter) DeleteSongList(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, BadRequest("无效的ID"))
		return
	}
	if err := service.SongListServiceApp.DeleteSongList(uint(id)); err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	c.JSON(200, Success("删除成功"))
}

func (s *SongListRouter) AllSongList(c *gin.Context) {
	songLists, err := service.SongListServiceApp.AllSongList()
	if err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	c.JSON(200, Success(songLists))
}

func (s *SongListRouter) SongListOfTitle(c *gin.Context) {
	title := c.Query("title")
	if title == "" {
		c.JSON(400, BadRequest("标题不能为空"))
		return
	}
	songLists, err := service.SongListServiceApp.SongListOfTitle(title)
	if err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	c.JSON(200, Success(songLists))
}

func (s *SongListRouter) SongListOfStyle(c *gin.Context) {
	style := c.Query("style")
	if style == "" {
		c.JSON(400, BadRequest("风格不能为空"))
		return
	}
	songLists, err := service.SongListServiceApp.SongListOfStyle(style)
	if err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	c.JSON(200, Success(songLists))
}

func (s *SongListRouter) UpdateSongListMsg(c *gin.Context) {
	var songList model.SongList
	if err := c.ShouldBindJSON(&songList); err != nil {
		c.JSON(400, BadRequest(err.Error()))
		return
	}
	if err := service.SongListServiceApp.UpdateSongListMsg(&songList); err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	c.JSON(200, Success("更新成功"))
}

func (s *SongListRouter) UpdateSongListPic(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(400, BadRequest("无效的ID"))
		return
	}
	file, err := c.FormFile("pic")
	if err != nil {
		c.JSON(400, BadRequest(err.Error()))
		return
	}
	fileData, err := file.Open()
	if err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	defer fileData.Close()

	fileBytes := make([]byte, file.Size)
	if _, err := fileData.Read(fileBytes); err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}

	if err := service.SongListServiceApp.UpdateSongListPic(uint(id), fileBytes); err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	c.JSON(200, Success("更新成功"))
}

var SongListRouterApp = new(SongListRouter)
