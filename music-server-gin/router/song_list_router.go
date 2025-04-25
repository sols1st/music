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
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := service.SongListServiceApp.AddSongList(&songList); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, songList)
}

func (s *SongListRouter) DeleteSongList(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	if err := service.SongListServiceApp.DeleteSongList(uint(id)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "delete success"})
}

func (s *SongListRouter) AllSongList(c *gin.Context) {
	songLists, err := service.SongListServiceApp.AllSongList()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, songLists)
}

func (s *SongListRouter) SongListOfTitle(c *gin.Context) {
	title := c.Query("title")
	if title == "" {
		c.JSON(400, gin.H{"error": "title is required"})
		return
	}
	songLists, err := service.SongListServiceApp.SongListOfTitle(title)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, songLists)
}

func (s *SongListRouter) SongListOfStyle(c *gin.Context) {
	style := c.Query("style")
	if style == "" {
		c.JSON(400, gin.H{"error": "style is required"})
		return
	}
	songLists, err := service.SongListServiceApp.SongListOfStyle(style)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, songLists)
}

func (s *SongListRouter) UpdateSongListMsg(c *gin.Context) {
	var songList model.SongList
	if err := c.ShouldBindJSON(&songList); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := service.SongListServiceApp.UpdateSongListMsg(&songList); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "update success"})
}

func (s *SongListRouter) UpdateSongListPic(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	file, err := c.FormFile("pic")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fileData, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer fileData.Close()

	fileBytes := make([]byte, file.Size)
	if _, err := fileData.Read(fileBytes); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := service.SongListServiceApp.UpdateSongListPic(uint(id), fileBytes); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "update success"})
}

var SongListRouterApp = new(SongListRouter)
